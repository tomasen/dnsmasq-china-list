package main

import (
	"flag"

	"github.com/tomasen/dnsmasq-china-list/src"
)

func main() {

	logfile := flag.String("logfile", "/var/log/dnsmasq.log", "dnsmasq query log file to be analyzed")
	confDir := flag.String("confdir", "./", "current dnsmasq conf dir")
	output := flag.String("output", "accelerated-domains2.conf", "output config filename")
	server := flag.String("server", "119.29.29.29", "dns source server")

	flag.Parse()

	chinaDomainList.ReadConfDir(*confDir)
	chinaDomainList.ReadDNSMasqLogfile(*logfile)
	chinaDomainList.WriteChinaConf(*output, *server)
}
