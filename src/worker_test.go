package chinaDomainList

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_LogEntry(t *testing.T) {
	pwd, _ := os.Getwd()
	ReadDNSMasqLogfile(filepath.Join(pwd, "_test.log"))
}

func Test_IsChina(t *testing.T) {
	if !check("shd.xd.com", "") {
		t.Fail()
	}
}
