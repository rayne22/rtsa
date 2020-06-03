package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Params struct {
	InsuranceType uint `json:"insuranceType"`
	Status uint `json:"status"`
	RegistrationMark string `json:"registrationMark"`
	DateFrom time.Time`json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
	InsurancePolicyNo string `json:"insurancePolicyNo"`
	ChassisNumber string `json:"chassisNumber"`
}


var (

	parA []Params
	jsonBuffer []byte
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}




var rstaPost = http.HandlerFunc(func (response http.ResponseWriter, request *http.Request)  {

	response.Header().Set("content-type", "application/json")
	var messages []Params
	var username  = "Douglas.Chilungu"
	var passwd = "aplusgeneral@2019"
	client := &http.Client{}
	_ = json.NewDecoder(request.Body).Decode(&messages)
	var ppl []interface{}
	for _, p := range messages {
		ppl = append(ppl, p)
	}

	_ = json.NewEncoder(response).Encode(ppl)


	jsonReq, err := json.Marshal(ppl)


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




func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/rtsa", rstaPost).Methods("Post")

	port := os.Getenv("PORT")

	log.Println(port)
	log.Fatal(http.ListenAndServe(":8060" , router))
}


