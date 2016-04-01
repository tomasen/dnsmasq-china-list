package chinaDomainList

import (
	"log"
	"os"
)

// WriteChinaConf write dnsmasq config file
func WriteChinaConf(conf, server string) {
	log.Println("writing dnsmasq file:", conf)
	f, err := os.Create(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for k, v := range chndomains {
		if v {
			f.WriteString("server=/" + k + "/" + server + "\n")
		}
	}
}
