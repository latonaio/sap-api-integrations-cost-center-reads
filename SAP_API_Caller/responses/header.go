package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
			ToText                         struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Text"`
		} `json:"results"`
	} `json:"d"`
}
