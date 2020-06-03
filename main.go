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
)

//type Params struct {
//	InsuranceType uint `json:"insurance_type"`
//	Status uint `json:"status"`
//	RegistrationMark string `json:"registration_mark"`
//	DateFrom time.Time `json:"date_from"`
//	DateTo time.Time `json:"date_to"`
//	InsurancePolicyNo string `json:"insurance_policy_no"`
//	ChassisNumber string `json:"chassis_number"`
//}
//
//var (
//
//	parA []Params
//	jsonBuffer []byte
//)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}




var rstaPost = http.HandlerFunc(func (response http.ResponseWriter, request *http.Request)  {
	response.Header().Set("content-type", "application/json")
	queryValues := request.URL.Query()
	_, _ = fmt.Fprintf(response, "hello, %s!\n", queryValues.Get("registrationMark"))

	insuranceType := queryValues.Get("insuranceType")
	status := queryValues.Get("status")
	registrationMark := queryValues.Get("registrationMark")
	dateFrom := queryValues.Get("dateFrom")
	dateTo := queryValues.Get("dateTo")
	insurancePolicyNo := queryValues.Get("insurancePolicyNo")
	chassisNumber := queryValues.Get("chassisNumber")

	var username  = "Douglas.Chilungu"
	var passwd = "aplusgeneral@2019"
	client := &http.Client{}

	var newA []string

	newA = append(newA, insuranceType)
	newA = append(newA, status)
	newA = append(newA, registrationMark)
	newA = append(newA, dateFrom)
	newA = append(newA, dateTo)
	newA = append(newA, insurancePolicyNo)
	newA = append(newA, chassisNumber)

	log.Println("New A", newA)

	nP, _ := json.Marshal( &newA)

	url := "https://zampointzidb.eservices.gov.zm/ZIDB/ReceiveInsurancePolicies"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(nP))

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
	router.HandleFunc("/rtsa", rstaPost).Methods("GET")

	port := os.Getenv("PORT")

	log.Println(port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}


