package rancher

import (
	"strings"

	"github.com/docker/machine/libmachine/log"
	"gopkg.in/yaml.v2"
)

type CloudInitRancher struct {
	Runcmd     []string `yaml:"runcmd"`
	WriteFiles []struct {
		Content     string `yaml:"content"`
		Encoding    string `yaml:"encoding"`
		Path        string `yaml:"path"`
		Permissions string `yaml:"permissions"`
	} `yaml:"write_files"`
}

func GetCloudInitRancher(s string) string {
	out := CloudInitRancher{}
	err := yaml.Unmarshal([]byte(s), &out)
	if err != nil {
		log.Debugf("Unmarshal: %v", err)
	}

	for _, entry := range out.WriteFiles {
		return patchRancherInstallScript(entry.Content)
	}

	return ""
}

// patchRancherInstallScript правит bootstrap install.sh перед записью на ВМ.
// vCD driver доставляет скрипт через guest customization (до SSH docker-machine).
// ROLE_NONE=true блокирует регистрацию ролей; strict verify ломается с --no-cacerts.
func patchRancherInstallScript(content string) string {
	content = strings.ReplaceAll(content, "CATTLE_ROLE_NONE=true", "CATTLE_ROLE_NONE=false")
	content = strings.ReplaceAll(content, `STRICT_VERIFY="true"`, `STRICT_VERIFY="false"`)
	return content
}
