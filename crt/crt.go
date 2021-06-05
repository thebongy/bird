package crt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Crt struct {
	URL string
}

/*
   {
       "issuer_ca_id": 179890,
       "issuer_name": "C=BE, O=GlobalSign nv-sa, CN=GlobalSign Atlas R3 DV TLS CA 2020",
       "common_name": "f.cloud.github.com",
       "name_value": "f.cloud.github.com",
       "id": 4514434585,
       "entry_timestamp": "2021-05-12T18:55:40.063",
       "not_before": "2021-05-12T18:55:36",
       "not_after": "2022-06-13T18:55:35",
       "serial_number": "015c7639616e88e6fcd9feff14b9192a"
   },
*/

type CrtRecord struct {
	CommonName string `json:"common_name,omitempty"`
}

func removeDuplicateValues(stringSlice []CrtRecord) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range stringSlice {
		if _, value := keys[entry.CommonName]; !value {
			keys[entry.CommonName] = true
			list = append(list, entry.CommonName)
		}
	}
	return list
}

// New returns a new Brute object
func New(url string) *Crt {
	return &Crt{
		URL: url,
	}
}

func (c *Crt) Parse() []string {
	domain := strings.TrimPrefix(strings.TrimPrefix(c.URL, "http://"), "https://")
	res, _ := http.Get(fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain))
	var data []CrtRecord
	json.NewDecoder(res.Body).Decode(&data)

	return removeDuplicateValues(data)
}
