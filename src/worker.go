package chinaDomainList

import (
	"bufio"
	"log"
	"net"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/publicsuffix"
)

var (
	// ']: reply d.dropbox.com is '
	logEntry = regexp.MustCompile(`\]\:\ reply\ ([a-zA-z0-9\-\.]+) is`)
	chinaNS  = regexp.MustCompile(`(qq\.com|dnspod|360safe|sina|\.dnsv[0-9]+\.com|dnspai|51dns|xincache|yunjiasu|xundns|baidu|lecloud|5173|tudoudns|letvlb|qingcdn|xinhuanet|youku|yodao|duowanns|sogou|alidns|kingsoft|aliyun|xunlei|alipay|ourdvs|taobao|uc\.cn|hichina|iqiyi|chinacache|ccgslb|\.cn\.|nease|aoyou365|sohu)`)
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
	domain = strings.ToLower(domain)

	if !isDomainName(domain) {
		return
	}

	if isIgnored(domain) {
		return
	}

	tldPlusOne, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		log.Println(err)
		tldPlusOne = domain
	}

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
		addToIgnoreList(tldPlusOne, false)
		addToIgnoreList(domain, false)
		return false
	}

	for _, v := range nss {
		if chinaNS.MatchString(v.Host) {
			addToChinaList(tldPlusOne)
			return true
		}

		// check if ns record is belong to china domain
		ns := strings.TrimSuffix(strings.TrimSpace(v.Host), ".")
		ns, err = publicsuffix.EffectiveTLDPlusOne(domain)
		if err == nil && isChina(ns) {
			addToChinaList(tldPlusOne)
			return true
		}
	}

	// put it in ignores list to avoid double check
	addToIgnoreList(tldPlusOne, false)

	if len(nss) > 0 {
		log.Println("out-china ns server:", nss[0].Host)
	}

	return false
}

func isDomainName(s string) bool {
	// See RFC 1035, RFC 3696.
	if len(s) == 0 {
		return false
	}
	if len(s) > 255 {
		return false
	}

	last := byte('.')
	ok := false // Ok once we've seen a letter.
	partlen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		default:
			return false
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_':
			ok = true
			partlen++
		case '0' <= c && c <= '9':
			// fine
			partlen++
		case c == '-':
			// Byte before dash cannot be dot.
			if last == '.' {
				return false
			}
			partlen++
		case c == '.':
			// Byte before dot cannot be dot, dash.
			if last == '.' || last == '-' {
				return false
			}
			if partlen > 63 || partlen == 0 {
				return false
			}
			partlen = 0
		}
		last = c
	}
	if last == '-' || partlen > 63 {
		return false
	}

	return ok
}
