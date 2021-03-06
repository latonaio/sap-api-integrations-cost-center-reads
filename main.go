package main

import (
	sap_api_caller "sap-api-integrations-cost-center-reads/SAP_API_Caller"
	"sap-api-integrations-cost-center-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Cost_Center_Cost_Center_Name_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "CostCenterName",
		}
	}

	caller.AsyncGetCostCenter(
		inoutSDC.CostCenter.ControllingArea,
		inoutSDC.CostCenter.CostCenter,
		inoutSDC.CostCenter.CostCenterText.Language,
		inoutSDC.CostCenter.CostCenterText.CostCenterName,
		accepter,
	)
}
