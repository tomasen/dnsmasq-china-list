package chinaDomainList

import (
	"log"
	"os"
)

// WriteChinaConf write dnsmasq config file
func WriteChinaConf(conf string) {
	log.Println("writing dnsmasq file:", conf)
	f, err := os.Create(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for k, v := range chndomains {
		if v {
			// server=/csi.gstatic.com/114.114.114.114
			f.WriteString("server=/" + k + "/223.5.5.5\n")
		}
	}
}
