package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/abkhan/ddvco/internal/myvelo"
	"github.com/abkhan/ddvco/pkg/velocloud"
)

const tok = "ey"

var linkId = flag.Int("l", 0, "ID of the edge device")
var edgeID = flag.Int("e", 1, "ID of the edge device")
var enterpriseId = flag.Int("enterpriseId", 10, "ID of the VeloCloud Enterprise")
var host = flag.String("h", "", "VeloCloud api host")
var token = flag.String("t", tok, "VeloCloud api access token")
var ssl = flag.Bool("ssl", false, "enable ssl")
var summ = flag.Bool("v", false, "print initial option summary")

var adds = flag.String("f", "na", "comma separated additional fields")
var waitm = flag.Int("w", 5, "minutes between poll")

func main() {
	flag.Parse()

	if *summ {

		log.Printf("Poll for EdgeId: %d", *edgeID)
		if *linkId != 0 {
			log.Printf("Print for LinkID: %d", *linkId)
		}
		log.Printf("Additional Fields: %s", *adds)
		log.Printf("Poll Every %dmin", *waitm)
	}

	vco, err := velocloud.NewTokenClient(host, token, ssl)
	if err != nil {
		log.Fatalf("error creating token client: %v", err)
	}
	vco.SetEnterprise(*enterpriseId)
	fl := fieldList()

	for {

		flist, e := myvelo.GetEdgeLinkMetricFields(vco, *edgeID, time.Now().Add(time.Duration(-1)*time.Hour), time.Now(), fl)
		if e != nil {
			log.Fatalf("error getting link metrics: %v", e)
		}
		vcprint(fl, flist)
		time.Sleep(time.Minute * time.Duration(*waitm))
	}
}

func fieldList() []string {
	fl := []string{"bytesRx", "bytesTx", "bpsOfBestPathRx", "bpsOfBestPathTx"}
	if *adds != "na" {
		adfl := strings.Split(*adds, ",")
		fl = append(fl, adfl...)
	}
	return fl
}

func vcprint(fl []string, data map[int]map[string]interface{}) {

	for link, datamap := range data {

		if *linkId != 0 && *linkId != link {
			continue
		}

		line := fmt.Sprintf("Link: %d >> ", link)

		for _, f := range fl {
			if fv, there := datamap[f]; there {
				line += fmt.Sprintf("%s: %.0f, ", f, fv)
			}
		}

		log.Print(line)
	}

}
