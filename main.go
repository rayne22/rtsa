package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Params struct {
	InsuranceType uint `json:"insurance_type"`
	Status uint `json:"status"`
	RegistrationMark string `json:"registration_mark"`
	DateFrom time.Time `json:"date_from"`
	DateTo time.Time `json:"date_to"`
	InsurancePolicyNo string `json:"insurance_policy_no"`
	ChassisNumber string `json:"chassis_number"`
}

var (
	par1 []Params
	jsonBuffer []byte
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func rstaPost (w http.ResponseWriter, r *http.Request) {
	var username  = "Douglas.Chilungu"
	var passwd = "aplusgeneral@2019"

	client := &http.Client{}
	var jsonStr = `[{
   "insuranceType": 0,
   "status": 0,
   "registrationMark": "test1",
   "dateFrom": "2020-06-02T08:39:58.579Z",
   "dateTo": "2020-06-02T08:39:58.579Z",
   "insurancePolicyNo": "test1",
   "chassisNumber": "test1"
 }]`

	//var para Params
	//
	//byts := json.Unmarshal([]byte(jsonStr), &para)

	req, err := http.NewRequest("POST", "https://zampointzidb.eservices.gov.zm/ZIDB/ReceiveInsurancePolicies", bytes.NewBuffer([]byte(jsonStr)))
	req.SetBasicAuth(username, passwd)
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))


}


//var rstaPost = http.HandlerFunc(func (response http.ResponseWriter, request *http.Request)  {
//	response.Header().Set("content-type", "application/json")
//	_ = json.NewDecoder(request.Body).Decode(&par1)
//
//
//	var username  = "Douglas.Chilungu"
//	var passwd = "aplusgeneral@2019"
//	client := &http.Client{}
//
//
//	payload := par1
//
//	nP, _ := json.Marshal( &payload)
//
//	log.Println(bytes.NewBuffer(nP))
//
//
//	url := "https://zampointzidb.eservices.gov.zm/ZIDB/ReceiveInsurancePolicies"
//	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(nP))
//
//	log.Println(request.Body)
//
//	req.SetBasicAuth(username, passwd)
//	req.Header.Set("content-type", "application/json")
//
//	res, err := client.Do(req)
//	if err != nil{
//				log.Fatal(err)
//			}
//	defer res.Body.Close()
//	body, _ := ioutil.ReadAll(res.Body)
//
//	fmt.Println(string(body))
//})












func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/rtsa", rstaPost).Methods("POST")
	log.Fatal(http.ListenAndServe(":8060", router))
}


