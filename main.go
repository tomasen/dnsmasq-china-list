package main

import (
	"flag"
	"log"
	"os"

	"github.com/tomasen/dnsmasq-china-list/src"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println("get pwd error:", err)
	}
	var logfile = flag.String("logfile", "/var/log/dnsmasq.log", "dnsmasq query log file to be analyzed")
	var confDir = flag.String("confdir", pwd, "current dnsmasq conf dir")
	var output = flag.String("output", "accelerated-domains2.conf", "output config filename")

	flag.Parse()

	chinaDomainList.ReadConfDir(*confDir)
	chinaDomainList.ReadDNSMasqLogfile(*logfile)
	chinaDomainList.WriteChinaConf(*output)
}
