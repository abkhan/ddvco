package myvelo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/abkhan/ddvco/pkg/velocloud"
)

func GetEdgeLinkMetricFields(c *velocloud.Client, eid int, from, to time.Time, fields []string) (map[int]map[string]interface{}, error) {

	retv := make(map[int]map[string]interface{})

	linklist, err := GetEdgeLinkMetrics(c, eid, from, to)
	if err != nil {
		return nil, err
	}

	for _, data := range linklist {
		lid := data.(map[string]interface{})["linkId"]

		linkmap := make(map[string]interface{})
		for _, field := range fields {
			fd := data.(map[string]interface{})[field]
			linkmap[field] = fd
		}

		retv[int(lid.(float64))] = linkmap
	}
	return retv, nil
}

// GetEdgeLinkMetrics return link metrics for an edge
// eid is edgeId
func GetEdgeLinkMetrics(c *velocloud.Client, eid int, from, to time.Time) ([]interface{}, error) {

	req := GetEdgeLinkMetricsRequest{
		EnterpriseId: c.EnterpriseId,
		EdgeId:       eid,
		IntervalStr: IntervalStr{
			Start: from.Format(time.RFC3339),
			End:   to.Format(time.RFC3339),
		},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(req)
	if err != nil {
		log.Fatal(err)
	}

	httpreq, err := http.NewRequest("POST", fmt.Sprintf("%s/metrics/getEdgeLinkMetrics", c.HostURL), &buf)

	es := []interface{}{}
	if err != nil {
		fmt.Println(err.Error())
		return es, err
	}

	// Send the request
	res, err := c.DoRequest(httpreq)
	if err != nil {
		fmt.Println(err.Error())
		return es, err
	}

	// Unmarschal
	err = json.Unmarshal(res, &es)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return es, err
	}

	return es, nil
}
