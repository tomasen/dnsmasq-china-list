package chinaDomainList

import (
	"bufio"
	"log"
	"net"
	"os"
	"regexp"

	"golang.org/x/net/publicsuffix"
)

var (
	// ']: reply d.dropbox.com is '
	logEntry = regexp.MustCompile(`\]\:\ reply\ ([a-zA-z0-9\-\.]+) is`)
	chinaNS  = regexp.MustCompile(`(qq.com|dnspod|360safe|sina|dnsv5|taobao)`)
)

// ReadDNSMasqLogfile analyz dnsmasq log file
func ReadDNSMasqLogfile(f string) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b := logEntry.FindAllStringSubmatch(scanner.Text(), -1)
		if len(b) == 1 && len(b[0]) == 2 {
			checkDomain(b[0][1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkDomain(domain string) {
	if isIgnored(domain) {
		return
	}

	tldPlusOne, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		log.Fatal(err)
	}

	// put it in ignores list to avoid double check
	addToIgnoreList(tldPlusOne)

	check(domain, tldPlusOne)

}

func check(domain string, tldPlusOne string) bool {
	// check ns record
	if len(tldPlusOne) == 0 {
		tldPlusOne, _ = publicsuffix.EffectiveTLDPlusOne(domain)
	}

	nss, err := net.LookupNS(tldPlusOne)
	if err != nil {
		log.Println("LookupNS failed", tldPlusOne, err)
	}

	for _, v := range nss {
		if chinaNS.MatchString(v.Host) {
			addToChinaList(tldPlusOne)
			return true
		}
	}
	if len(nss) > 0 {
		log.Println(nss[0].Host)
	}

	return false
}
