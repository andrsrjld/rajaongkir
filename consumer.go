package rajaongkir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetCost(request CheckCostRequest) (response CheckCostResponse, err error) {
	url := URLCheckOngkir

	courier := SeparateCourierList(request.Courier)
	strPayload := fmt.Sprintf("origin=%s&originType=%s&destination=%s&destinationType=%s&weight=%d&courier=%s",
		request.Origin, request.OriginType, request.Destination, request.DestinationType, request.Weight, courier)

	payload := strings.NewReader(strPayload)
	RORequest, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println("RAJAONGKIR [GET COST] NEW REQUEST ERROR :", err)
		return
	}

	RORequest.Header.Add("key", APIKey)
	RORequest.Header.Add("content-type", "application/x-www-form-urlencoded")
	ROResponse, err := http.DefaultClient.Do(RORequest)
	if err != nil {
		log.Println("RAJAONGKIR [GET COST] RESPONSE ERROR :", err)
		return
	}

	defer ROResponse.Body.Close()
	body, err := ioutil.ReadAll(ROResponse.Body)
	if err != nil {
		log.Println("RAJAONGKIR [GET COST] READ RESPONSE BODY ERROR :", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("RAJAONGKIR [GET COST] UNMARSHAL RESPONSE BODY ERROR :", err)
		return
	}

	return
}

func CheckAwb(request WaybillRequest) (response WaybillResponse, err error) {
	url := URLCheckWaybill

	strPayload := fmt.Sprintf("waybill=%s&courier=%s", request.NomorResi, request.Courier)

	payload := strings.NewReader(strPayload)
	RORequest, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println("RAJAONGKIR [CHECK AWB] NEW REQUEST ERROR :", err)
		return
	}

	RORequest.Header.Add("key", APIKey)
	RORequest.Header.Add("content-type", "application/x-www-form-urlencoded")
	ROReponse, err := http.DefaultClient.Do(RORequest)
	if err != nil {
		log.Println("RAJAONGKIR [CHECK AWB] RESPONSE ERROR :", err)
		return
	}

	defer ROReponse.Body.Close()
	body, err := ioutil.ReadAll(ROReponse.Body)
	if err != nil {
		log.Println("RAJAONGKIR [CHECK AWB] READ RESPONSE BODY ERROR :", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("RAJAONGKIR [CHECK AWB] UNMARSHAL RESPONSE BODY ERROR :", err)
		return
	}

	return
}
