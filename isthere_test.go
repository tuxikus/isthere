package isthere

import (
	"os/exec"
	"strings"
	"testing"
)

func TestIsThere(t *testing.T) {
	// bash
	whichCmd, err := exec.Command("which", "bash").Output()
	if err != nil {
	}

	whichCmdStr := string(whichCmd)
	whichCmdStr = strings.ReplaceAll(whichCmdStr, "\n", "")
	whichCmdStr = strings.TrimSpace(whichCmdStr)
	isThereCmd, _ := IsThere("bash")

	if whichCmdStr != isThereCmd {
		t.Errorf("which: %s, isthere: %s", whichCmdStr, isThereCmd)
	}
}
