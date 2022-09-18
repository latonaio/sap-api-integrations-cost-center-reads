package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	ProductionOrder struct {
		ManufacturingOrder             string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    string      `json:"quantity"`
		MfgOrderItemActualDeviationQty string      `json:"picked_quantity"`
		Price                          string      `json:"price"`
	  Batch                          string      `json:"batch"`
	} `json:"document"`
	ManufacturingOrder struct {
		ManufacturingOrder             string      `json:"document_no"`
		StatusCode                     string      `json:"status"`
		DeliverTo                      string      `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    string      `json:"quantity"`
		MfgOrderItemActualDeviationQty string      `json:"completed_quantity"`
	    PlannedStartDate               string      `json:"planned_start_date"`
	    MfgOrderItemPlndDeliveryDate   string      `json:"planned_validated_date"`
	    ActualStartDate                string      `json:"actual_start_date"`
	    MfgOrderItemActualDeliveryDate string      `json:"actual_validated_date"`
	    Batch                          string      `json:"batch"`
		BillOfOperations struct {
			OrderIntBillOfOperationsItem string      `json:"work_no"`
			OpPlannedTotalQuantity       string      `json:"quantity"`
			OpTotalConfirmedYieldQty     string      `json:"completed_quantity"`
			ErroredQuantity              string      `json:"errored_quantity"`
			Component                    string      `json:"component"`
			MaterialCompOriginalQuantity string      `json:"planned_component_quantity"`
			OpErlstSchedldExecStrtDte    string      `json:"planned_start_date"`
			OpErlstSchedldExecStrtTme    string      `json:"planned_start_time"`
			OpErlstSchedldExecEndDte     string      `json:"planned_validated_date"`
			OpErlstSchedldExecEndTme     string      `json:"planned_validated_time"`
			OpActualExecutionStartDate   string      `json:"actual_start_date"`
			OpActualExecutionStartTime   string      `json:"actual_start_time"`
			OpActualExecutionEndDate     string      `json:"actual_validated_date"`
			OpActualExecutionEndTime     string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                      string      `json:"api_schema"`
	MaterialCode                   string      `json:"material_code"`
	ProductionPlant                string      `json:"plant/supplier"`
	Stock                          string      `json:"stock"`
	ManufacturingOrderType         string      `json:"document_type"`
	ManufacturingOrderNo           string      `json:"document_no"`
	MfgOrderItemPlndDeliveryDate   string      `json:"planned_date"`
	MfgOrderItemActualDeliveryDate string      `json:"validated_date"`
	Deleted                        bool        `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	CostCenter    struct {
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
		CostCenterText                 struct {
			Language              string `json:"Language"`
			ValidityEndDate       string `json:"ValidityEndDate"`
			ValidityStartDate     string `json:"ValidityStartDate"`
			CostCenterName        string `json:"CostCenterName"`
			CostCenterDescription string `json:"CostCenterDescription"`
		} `json:"CostCenterText"`
	} `json:"CostCenter"`
	APISchema      string   `json:"api_schema"`
	Accepter       []string `json:"accepter"`
	CostCenterCode string   `json:"cost_center_code"`
	Deleted        bool     `json:"deleted"`
}
