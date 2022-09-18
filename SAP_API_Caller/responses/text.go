package responses

type Text struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			CostCenter            string `json:"CostCenter"`
			ControllingArea       string `json:"ControllingArea"`
			Language              string `json:"Language"`
			ValidityEndDate       string `json:"ValidityEndDate"`
			ValidityStartDate     string `json:"ValidityStartDate"`
			CostCenterName        string `json:"CostCenterName"`
			CostCenterDescription string `json:"CostCenterDescription"`
		} `json:"results"`
	} `json:"d"`
}
