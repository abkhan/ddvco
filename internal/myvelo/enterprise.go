package myvelo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abkhan/ddvco/pkg/velocloud"
)

// GetEnterprises ...
func GetEnterprises(c *velocloud.Client) ([]string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/network/getNetworkEnterprises", c.HostURL), nil)

	es := []string{}
	if err != nil {
		fmt.Println(err.Error())
		return es, err
	}

	// Send the request
	res, err := c.DoRequest(req)
	if err != nil {
		fmt.Println(err.Error())
		return es, err
	}

	// Unmarschal
	var list []map[string]interface{}
	err = json.Unmarshal(res, &list)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return es, err
	}

	for _, v := range list {
		es = append(es, v["name"].(string))
	}

	return es, nil
}
