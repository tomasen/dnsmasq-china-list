package chinaDomainList

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/publicsuffix"
)

var serverEntry = regexp.MustCompile(`^server=\/([a-zA-z0-9\-\.]+)\/[0-9\.]+$`)

var (
	ignores    = make(map[string]bool)
	chndomains = make(map[string]bool)
)

// ReadConfDir put domain in ignore list
func ReadConfDir(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if !f.IsDir() {
			if filepath.Ext(f.Name()) == ".conf" {
				addDnsmasqConfToIgnoreList(filepath.Join(dir, f.Name()))
			}
		}
	}
}

func addDnsmasqConfToIgnoreList(f string) {
	log.Println("loading dnsmasq conf:", f)

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b := serverEntry.FindAllStringSubmatch(scanner.Text(), -1)
		if len(b) == 1 && len(b[0]) == 2 {
			addToIgnoreList(b[0][1], true)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func addToIgnoreList(domain string, ischina bool) {
	ignores[domain] = ischina
}

func isIgnored(domain string) bool {
	if _, ok := ignores[domain]; ok {
		return true
	}

	d, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		log.Println(err)
		return true
	}

	if _, ok := ignores[d]; ok {
		return true
	}

	for k := range ignores {
		if strings.HasSuffix(domain, k) {
			return true
		}
	}

	return false
}

func addToChinaList(domain string) {
	log.Println("found new china-domain:", domain)
	chndomains[domain] = true
	addToIgnoreList(domain, true)
}

func isChina(domain string) bool {
	if val, ok := ignores[domain]; ok {
		return val
	}

	d, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		log.Println(err)
		return false
	}

	if val, ok := ignores[d]; ok {
		return val
	}

	for k, v := range ignores {
		if strings.HasSuffix(domain, k) {
			return v
		}
	}

	return false
}
