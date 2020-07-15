package models

import (
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

	ParA       []Params
	jsonBuffer []byte
)