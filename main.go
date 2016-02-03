package main

import (
	"flag"

	"github.com/tomasen/dnsmasq-china-list/src"
)

func main() {

	var logfile = flag.String("logfile", "/var/log/dnsmasq.log", "dnsmasq query log file to be analyzed")
	var confDir = flag.String("confdir", "./", "current dnsmasq conf dir")
	var output = flag.String("output", "accelerated-domains2.conf", "output config filename")

	flag.Parse()

	chinaDomainList.ReadConfDir(*confDir)
	chinaDomainList.ReadDNSMasqLogfile(*logfile)
	chinaDomainList.WriteChinaConf(*output)
}
