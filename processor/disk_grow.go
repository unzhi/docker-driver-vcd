package processor

// guestCustomizationGrowRootDisk — vCD меняет SizeMb диска, но Ubuntu LVM-шаблон
// оставляет LV ~8G. Скрипт выполняется в guest customization до bootstrap RKE2.
const guestCustomizationGrowRootDisk = `
# expand root FS after vCD disk resize (Ubuntu LVM cloud template)
if command -v growpart >/dev/null 2>&1; then
  growpart /dev/sda 3 2>/dev/null || growpart /dev/sda 2 2>/dev/null || true
fi
PV=$(pvs --noheadings -o pv_name 2>/dev/null | awk 'NF{print $1; exit}')
if [ -n "$PV" ]; then
  pvresize "$PV" 2>/dev/null || true
fi
LV=$(findmnt -n -o SOURCE / 2>/dev/null)
if [ -n "$LV" ]; then
  lvextend -l +100%FREE "$LV" 2>/dev/null || lvextend -l +100%FREE /dev/ubuntu-vg/ubuntu-lv 2>/dev/null || true
  resize2fs "$LV" 2>/dev/null || xfs_growfs / 2>/dev/null || true
fi
`
