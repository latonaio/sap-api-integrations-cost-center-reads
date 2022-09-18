package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-cost-center-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetCostCenter(controllingArea, costCenter, language, costCenterName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(controllingArea, costCenter)
				wg.Done()
			}()
		case "CostCenterName":
			func() {
				c.CostCenterName(language, costCenterName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(controllingArea, costCenter string) {
	headerData, err := c.callCostCenterSrvAPIRequirementHeader("A_CostCenter", controllingArea, costCenter)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	textData, err := c.callToText(headerData[0].ToText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(textData)
}

func (c *SAPAPICaller) callCostCenterSrvAPIRequirementHeader(api, controllingArea, costCenter string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_COSTCENTER_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, controllingArea, costCenter)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToText(url string) ([]sap_api_output_formatter.ToText, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CostCenterName(language, costCenterName string) {
	data, err := c.callCostCenterSrvAPIRequirementCostCenterName("A_CostCenterText", language, costCenterName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callCostCenterSrvAPIRequirementCostCenterName(api, language, costCenterName string) ([]sap_api_output_formatter.Text, error) {
	url := strings.Join([]string{c.baseURL, "API_COSTCENTER_SRV", api}, "/")

	param := c.getQueryWithCostCenterName(map[string]string{}, language, costCenterName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, controllingArea, costCenter string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("ControllingArea eq '%s' and CostCenter eq '%s'", controllingArea, costCenter)
	return params
}

func (c *SAPAPICaller) getQueryWithCostCenterName(params map[string]string, language, costCenterName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Language eq '%s' and substringof('%s', CostCenterName)", language, costCenterName)
	return params
}
