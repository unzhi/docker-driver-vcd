package rancher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatchRancherInstallScript(t *testing.T) {
	in := `#!/bin/sh
STRICT_VERIFY="true"
CATTLE_ROLE_NONE=true
CATTLE_SERVER=https://rancher.example.com
`
	out := patchRancherInstallScript(in)
	assert.Contains(t, out, `STRICT_VERIFY="false"`)
	assert.Contains(t, out, "CATTLE_ROLE_NONE=false")
	assert.NotContains(t, out, "CATTLE_ROLE_NONE=true")
}
