package rajaongkir

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ShippingDetails struct {
	OriginSubID     int
	Weight          int
	DestSubID       int
	ShipName        string
	ShipServiceName string
}

func GetCost(shipping ShippingDetails) {

	url := "https://pro.rajaongkir.com/api/cost"
	payloadStr := fmt.Sprintf("origin=%v&originType=subdistrict&destination=%v&destinationType=subdistrict&weight=%v&courier=%v", shipping.OriginSubID, shipping.DestSubID, shipping.Weight, shipping.ShipName)
	payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("key", "your-api-key")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
