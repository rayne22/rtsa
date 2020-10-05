package models

import (
	"time"
)

type Params struct {
	InsuranceType uint `json:"insuranceType"`
	Status uint `json:"status"`
	DateFrom time.Time`json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
	InsurancePolicyNo string `json:"insurancePolicyNo"`
	ChassisNumber string `json:"chassisNumber"`
	Quarter uint `json:"quarter"`
	InsuranceCompany string `json:"insuranceCompany"`
	InsuranceCompanyId uint `json:"insuranceCompanyId"`
	RegistrationMark string `json:"registrationMark"`
}



var (

	ParA       []Params
	jsonBuffer []byte
)