package sap_api_output_formatter

type CostCenter struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	CostCenter         string `json:"cost_center_code"`
	Deleted            bool   `json:"deleted"`
}

type Header struct {
	ControllingArea                string `json:"ControllingArea"`
	CostCenter                     string `json:"CostCenter"`
	ValidityEndDate                string `json:"ValidityEndDate"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	CompanyCode                    string `json:"CompanyCode"`
	BusinessArea                   string `json:"BusinessArea"`
	CostCtrResponsiblePersonName   string `json:"CostCtrResponsiblePersonName"`
	CostCtrResponsibleUser         string `json:"CostCtrResponsibleUser"`
	CostCenterCurrency             string `json:"CostCenterCurrency"`
	ProfitCenter                   string `json:"ProfitCenter"`
	Department                     string `json:"Department"`
	FunctionalArea                 string `json:"FunctionalArea"`
	Country                        string `json:"Country"`
	Region                         string `json:"Region"`
	CityName                       string `json:"CityName"`
	CostCenterStandardHierArea     string `json:"CostCenterStandardHierArea"`
	CostCenterCategory             string `json:"CostCenterCategory"`
	IsBlkdForPrimaryCostsPosting   string `json:"IsBlkdForPrimaryCostsPosting"`
	IsBlkdForSecondaryCostsPosting string `json:"IsBlkdForSecondaryCostsPosting"`
	IsBlockedForRevenuePosting     string `json:"IsBlockedForRevenuePosting"`
	IsBlockedForCommitmentPosting  string `json:"IsBlockedForCommitmentPosting"`
	IsBlockedForPlanPrimaryCosts   string `json:"IsBlockedForPlanPrimaryCosts"`
	IsBlockedForPlanSecondaryCosts string `json:"IsBlockedForPlanSecondaryCosts"`
	IsBlockedForPlanRevenues       string `json:"IsBlockedForPlanRevenues"`
	ConsumptionQtyIsRecorded       string `json:"ConsumptionQtyIsRecorded"`
	Language                       string `json:"Language"`
	CostCenterCreationDate         string `json:"CostCenterCreationDate"`
	ToText                         string `json:"to_Text"`
}

type Text struct {
	CostCenter            string `json:"CostCenter"`
	ControllingArea       string `json:"ControllingArea"`
	Language              string `json:"Language"`
	ValidityEndDate       string `json:"ValidityEndDate"`
	ValidityStartDate     string `json:"ValidityStartDate"`
	CostCenterName        string `json:"CostCenterName"`
	CostCenterDescription string `json:"CostCenterDescription"`
}

type ToText struct {
	CostCenter            string `json:"CostCenter"`
	ControllingArea       string `json:"ControllingArea"`
	Language              string `json:"Language"`
	ValidityEndDate       string `json:"ValidityEndDate"`
	ValidityStartDate     string `json:"ValidityStartDate"`
	CostCenterName        string `json:"CostCenterName"`
	CostCenterDescription string `json:"CostCenterDescription"`
}
