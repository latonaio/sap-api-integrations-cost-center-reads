package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-cost-center-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			ControllingArea:                data.ControllingArea,
			CostCenter:                     data.CostCenter,
			ValidityEndDate:                data.ValidityEndDate,
			ValidityStartDate:              data.ValidityStartDate,
			CompanyCode:                    data.CompanyCode,
			BusinessArea:                   data.BusinessArea,
			CostCtrResponsiblePersonName:   data.CostCtrResponsiblePersonName,
			CostCtrResponsibleUser:         data.CostCtrResponsibleUser,
			CostCenterCurrency:             data.CostCenterCurrency,
			ProfitCenter:                   data.ProfitCenter,
			Department:                     data.Department,
			FunctionalArea:                 data.FunctionalArea,
			Country:                        data.Country,
			Region:                         data.Region,
			CityName:                       data.CityName,
			CostCenterStandardHierArea:     data.CostCenterStandardHierArea,
			CostCenterCategory:             data.CostCenterCategory,
			IsBlkdForPrimaryCostsPosting:   data.IsBlkdForPrimaryCostsPosting,
			IsBlkdForSecondaryCostsPosting: data.IsBlkdForSecondaryCostsPosting,
			IsBlockedForRevenuePosting:     data.IsBlockedForRevenuePosting,
			IsBlockedForCommitmentPosting:  data.IsBlockedForCommitmentPosting,
			IsBlockedForPlanPrimaryCosts:   data.IsBlockedForPlanPrimaryCosts,
			IsBlockedForPlanSecondaryCosts: data.IsBlockedForPlanSecondaryCosts,
			IsBlockedForPlanRevenues:       data.IsBlockedForPlanRevenues,
			ConsumptionQtyIsRecorded:       data.ConsumptionQtyIsRecorded,
			Language:                       data.Language,
			CostCenterCreationDate:         data.CostCenterCreationDate,
			ToText:                         data.ToText.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToText(raw []byte, l *logger.Logger) ([]Text, error) {
	pm := &responses.Text{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Text. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	text := make([]Text, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		text = append(text, Text{
			CostCenter:            data.CostCenter,
			ControllingArea:       data.ControllingArea,
			Language:              data.Language,
			ValidityEndDate:       data.ValidityEndDate,
			ValidityStartDate:     data.ValidityStartDate,
			CostCenterName:        data.CostCenterName,
			CostCenterDescription: data.CostCenterDescription,
		})
	}

	return text, nil
}

func ConvertToToText(raw []byte, l *logger.Logger) ([]ToText, error) {
	pm := &responses.ToText{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToText. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toText := make([]ToText, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toText = append(toText, ToText{
			CostCenter:            data.CostCenter,
			ControllingArea:       data.ControllingArea,
			Language:              data.Language,
			ValidityEndDate:       data.ValidityEndDate,
			ValidityStartDate:     data.ValidityStartDate,
			CostCenterName:        data.CostCenterName,
			CostCenterDescription: data.CostCenterDescription,
		})
	}

	return toText, nil
}
