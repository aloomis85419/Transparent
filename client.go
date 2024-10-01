package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

var cfg = Configure("57a8172387f671b61ff90cdd614f9a4bf82c071d")

var client = newLDAClient(cfg, http.Client{})

type LDAClient struct {
	Cfg    Config
	Client http.Client
}

func newLDAClient(cfg Config, client http.Client) LDAClient {
	ldaClient := LDAClient{
		Cfg:    cfg,
		Client: client,
	}
	return ldaClient
}

func (lda *LDAClient) addHeaders(req *http.Request) {
	req.Header.Set("Authorization", lda.Cfg.AuthInfo.ApiToken)
	req.Header.Set("Content-Type", "application/json")
}

func (lda *LDAClient) get(endpointInfo EndpointInfo) *http.Response {
	jsonData := []byte{}
	req, err := http.NewRequest("GET", lda.Cfg.BaseUrl+endpointInfo.Url, bytes.NewBuffer(jsonData))
	lda.addHeaders(req)
	resp, err := lda.Client.Do(req)
	CheckErr(err, fmt.Sprintf("GET request to %s returned an error", endpointInfo.Url), false)
	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Sprintf("Error: GET request to %s returned an with status code: %s, message: %s", resp.StatusCode, resp.Status))
	}
	return resp
}

func (lda *LDAClient) getWQuery(endpointInfo EndpointInfo, queryParams map[string]string, page int) *http.Response {
	jsonData := []byte{}
	//set page size to default
	queryParams["page_size"] = strconv.Itoa(lda.Cfg.DefaultPageSize)
	queryParams["page"] = strconv.Itoa(page)
	req, err := http.NewRequest("GET", lda.Cfg.BaseUrl+endpointInfo.Url, bytes.NewBuffer(jsonData))
	CheckErr(err, "creating request", true)
	addQueryParams(queryParams, req)
	lda.addHeaders(req)
	resp, err := lda.Client.Do(req)
	CheckErr(err, fmt.Sprintf("GET request to %s returned an error", endpointInfo.Url), false)
	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Sprintf("Error: GET request to %s returned an with status code: %s, message: %s", resp.StatusCode, resp.Status))
	}
	return resp
}

func (lda *LDAClient) getById(endpointInfo EndpointInfo, id interface{}) *http.Response {
	jsonData := []byte{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/%s", lda.Cfg.BaseUrl, endpointInfo.Url, id), bytes.NewBuffer(jsonData))
	lda.addHeaders(req)
	resp, err := lda.Client.Do(req)
	CheckErr(err, fmt.Sprintf("GET request to %s returned an error", endpointInfo.Url), false)
	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Sprintf("Error: GET request to %s returned an with status code: %s, message: %s", resp.StatusCode, resp.Status))
	}
	return resp
}

func (lda *LDAClient) listFilings() []Filing {
	resp := lda.getWQuery(lda.Cfg.EndpointInfo["filings"], map[string]string{}, 2)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	var paginatedFilingResponse PaginatedFilingResponse
	err = json.Unmarshal(body, &paginatedFilingResponse)
	CheckErr(err, "unmarshalling json", true)
	return paginatedFilingResponse.Results
}

func (lda *LDAClient) retrieveFiling(id interface{}) Filing {
	resp := lda.getById(lda.Cfg.EndpointInfo["filings"], id)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	filing := Filing{}
	if err := json.Unmarshal(body, &filing); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return filing
}

func (lda *LDAClient) listContributionReports() []ContributionReport {
	resp := lda.get(lda.Cfg.EndpointInfo["contributions"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	var paginatedContributionReportResponse PaginatedContributionReportResponse
	err = json.Unmarshal(body, &paginatedContributionReportResponse)
	CheckErr(err, "unmarshalling json", true)
	return paginatedContributionReportResponse.Results
}

func (lda *LDAClient) retrieveContributionReport(id interface{}) ContributionReport {
	resp := lda.getById(lda.Cfg.EndpointInfo["contributions"], id)
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	contribution := ContributionReport{}
	paginatedRegistrantResponse := []PaginatedRegistrantResponse{}
	err = json.Unmarshal(body, &paginatedRegistrantResponse)
	CheckErr(err, "unmarshalling json", true)
	return contribution
}

func (lda *LDAClient) listRegistrants() []Registrant {
	resp := lda.get(lda.Cfg.EndpointInfo["registrants"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	var paginatedRegistrantResponse PaginatedRegistrantResponse
	err = json.Unmarshal(body, &paginatedRegistrantResponse)
	CheckErr(err, "unmarshalling json", true)
	return paginatedRegistrantResponse.Results
}

func (lda *LDAClient) retrieveRegistrant(id interface{}) Registrant {
	resp := lda.getById(lda.Cfg.EndpointInfo["registrants"], id)
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	registrant := Registrant{}
	if err := json.Unmarshal(body, &registrant); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return registrant
}

func (lda *LDAClient) listClients() []Client {
	resp := lda.get(lda.Cfg.EndpointInfo["clients"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	var paginatedClientResponse PaginatedClientResponse
	err = json.Unmarshal(body, &paginatedClientResponse)
	CheckErr(err, "unmarshalling json", true)
	return paginatedClientResponse.Results
}

func (lda *LDAClient) retrieveClient(id interface{}) Client {
	resp := lda.getById(lda.Cfg.EndpointInfo["clients"], id)
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	client := Client{}
	if err := json.Unmarshal(body, &client); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return client
}

func (lda *LDAClient) listLobbyists() []Lobbyist {
	resp := lda.get(lda.Cfg.EndpointInfo["lobbyists"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	var paginatedLobbyistResponse PaginatedLobbyistResponse
	err = json.Unmarshal(body, &paginatedLobbyistResponse)
	CheckErr(err, "unmarshalling json", true)
	return paginatedLobbyistResponse.Results
}

func (lda *LDAClient) retrieveLobbyist(id interface{}) Lobbyist {
	resp := lda.getById(lda.Cfg.EndpointInfo["lobbyists"], id)
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	lobbyist := Lobbyist{}
	if err := json.Unmarshal(body, &lobbyist); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return lobbyist
}

func (lda *LDAClient) listFilingTypes() []FilingType {
	resp := lda.get(lda.Cfg.EndpointInfo["filingtypes"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	filingTypes := []FilingType{}
	if err := json.Unmarshal(body, &filingTypes); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return filingTypes
}

func (lda *LDAClient) listLobbyingActivityGeneralIssues() []LobbyingActivityIssue {
	resp := lda.get(lda.Cfg.EndpointInfo["lobbyingactivityissues"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	lobbyingActivityIssues := []LobbyingActivityIssue{}
	if err := json.Unmarshal(body, &lobbyingActivityIssues); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return lobbyingActivityIssues
}

func (lda *LDAClient) listGovernmentEntities() []GovernmentEntity {
	resp := lda.get(lda.Cfg.EndpointInfo["governmententities"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	governmentEntities := []GovernmentEntity{}
	if err := json.Unmarshal(body, &governmentEntities); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return governmentEntities
}

func (lda *LDAClient) listCountries() []Country {
	resp := lda.get(lda.Cfg.EndpointInfo["countries"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	countries := []Country{}
	if err := json.Unmarshal(body, &countries); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return countries
}

func (lda *LDAClient) listStates() []State {
	resp := lda.get(lda.Cfg.EndpointInfo["states"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	states := []State{}
	if err := json.Unmarshal(body, &states); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return states
}

func (lda *LDAClient) listLobbyistPrefixes() []LobbyistPrefix {
	resp := lda.get(lda.Cfg.EndpointInfo["lobbyistPrefixes"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	lobbyistPrefixes := []LobbyistPrefix{}
	if err := json.Unmarshal(body, &lobbyistPrefixes); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return lobbyistPrefixes
}

func (lda *LDAClient) listLobbyistSuffixes() []LobbyistSuffix {
	resp := lda.get(lda.Cfg.EndpointInfo["lobbyistSuffixes"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	lobbyistSuffixes := []LobbyistSuffix{}
	if err := json.Unmarshal(body, &lobbyistSuffixes); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return lobbyistSuffixes
}

func (lda *LDAClient) listContributionItemTypes() []ContributionItemType {
	resp := lda.get(lda.Cfg.EndpointInfo["contributionItemTypes"])
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	contributionItemTypes := []ContributionItemType{}
	if err := json.Unmarshal(body, &contributionItemTypes); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return contributionItemTypes
}

func addQueryParams(queryParams map[string]string, r *http.Request) {
	query := r.URL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	r.URL.RawQuery = query.Encode()
}
