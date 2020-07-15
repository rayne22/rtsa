package goldenlotus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rtsa/models"
)


var RtsaPostGoldenLotus = http.HandlerFunc(func (response http.ResponseWriter, request *http.Request)  {

	response.Header().Set("content-type", "application/json")

	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	response.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	//var messages []Params

	var username  = "ronald.mwenda"
	var passwd = "g0lden2020"
	client := &http.Client{}
	//_ = json.NewDecoder(request.Body).Decode(&messages)
	//var ppl []interface{}
	//for _, p := range messages {
	//	ppl = append(ppl, p)
	//}

	reqBody, _ := ioutil.ReadAll(request.Body)
	var article models.Params
	_ = json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	models.ParA = append(models.ParA, article)

	_ = json.NewEncoder(response).Encode(models.ParA)


	jsonReq, err := json.Marshal(models.ParA)


	url := "https://zampointzidb.eservices.gov.zm/ZIDB/ReceiveInsurancePolicies"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))

	log.Println(request.Body)

	req.SetBasicAuth(username, passwd)
	req.Header.Set("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

})