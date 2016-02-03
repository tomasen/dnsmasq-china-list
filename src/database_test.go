package chinaDomainList

import (
	"log"
	"os"
	"testing"
)

func Test_AddDNSMasqEntry(t *testing.T) {
	pwd, _ := os.Getwd()
	ReadConfDir(pwd)

	if len(ignores) != 15 {
		log.Println(ignores)
		t.Fail()
	}
}
