package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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
	resp := lda.get(lda.Cfg.EndpointInfo["filings"])
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	CheckErr(err, "", false)
	var filings []Filing
	if err := json.Unmarshal(body, &filings); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return filings
}

func (lda *LDAClient) retrieveFiling(id interface{}) Filing {
	resp := lda.getById(lda.Cfg.EndpointInfo["filings"], id)
	defer resp.Body.Close()
	body := []byte{}
	resp.Body.Read(body)
	filing := Filing{}
	json.Unmarshal(body, filing)
	return filing
}

func (lda *LDAClient) listContributionReports() []ContributionReport {
	resp := lda.get(lda.Cfg.EndpointInfo["contributions"])
	body := []byte{}
	resp.Body.Read(body)
	contributions := []ContributionReport{}
	json.Unmarshal(body, contributions)
	return contributions
}

func (lda *LDAClient) retrieveContributionReport(id interface{}) ContributionReport {
	resp := lda.getById(lda.Cfg.EndpointInfo["contributions"], id)
	body := []byte{}
	resp.Body.Read(body)
	contribution := ContributionReport{}
	json.Unmarshal(body, contribution)
	return contribution
}

func (lda *LDAClient) listRegistrants() *http.Response {
	resp := lda.get(lda.Cfg.EndpointInfo["registrants"])
	body := []byte{}
	resp.Body.Read(body)
	registrants := []Registrant{}
	json.Unmarshal(body, registrants)
	return resp
}

func (lda *LDAClient) retrieveRegistrant(id interface{}) Registrant {
	resp := lda.getById(lda.Cfg.EndpointInfo["registrants"], id)
	body := []byte{}
	resp.Body.Read(body)
	registrant := Registrant{}
	json.Unmarshal(body, registrant)
	return registrant
}

func (lda *LDAClient) listClients() []Client {
	resp := lda.get(lda.Cfg.EndpointInfo["clients"])
	body := []byte{}
	resp.Body.Read(body)
	clients := []Client{}
	json.Unmarshal(body, clients)
	return clients
}

func (lda *LDAClient) retrieveClient(id interface{}) Client {
	resp := lda.getById(lda.Cfg.EndpointInfo["clients"], id)
	body := []byte{}
	resp.Body.Read(body)
	client := Client{}
	json.Unmarshal(body, client)
	return client
}

func (lda *LDAClient) listLobbyists() []Lobbyist {
	resp := lda.get(lda.Cfg.EndpointInfo["lobbyists"])
	body := []byte{}
	resp.Body.Read(body)
	lobbyists := []Lobbyist{}
	json.Unmarshal(body, lobbyists)
	return lobbyists
}

func (lda *LDAClient) retrieveLobbyist(id interface{}) Lobbyist {
	resp := lda.getById(lda.Cfg.EndpointInfo["lobbyists"], id)
	body := []byte{}
	resp.Body.Read(body)
	lobbyist := Lobbyist{}
	json.Unmarshal(body, lobbyist)
	return lobbyist
}

func (lda *LDAClient) listFilingTypes() []FilingType {
	resp := lda.get(lda.Cfg.EndpointInfo["filingtypes"])
	body := []byte{}
	resp.Body.Read(body)
	filingTypes := []FilingType{}
	json.Unmarshal(body, filingTypes)
	return filingTypes
}

func (lda *LDAClient) listLobbyingActivityGeneralIssues() []LobbyingActivityIssue {
	resp := lda.get(lda.Cfg.EndpointInfo["lobbyingactivityissues"])
	body := []byte{}
	resp.Body.Read(body)
	lobbyingActivityIssues := []LobbyingActivityIssue{}
	json.Unmarshal(body, lobbyingActivityIssues)
	return lobbyingActivityIssues
}

func (lda *LDAClient) listGovernmentEntities() []GovernmentEntity {
	resp := lda.get(lda.Cfg.EndpointInfo["governmententities"])
	body := []byte{}
	resp.Body.Read(body)
	governmentEntities := []GovernmentEntity{}
	json.Unmarshal(body, governmentEntities)
	return governmentEntities
}

func (lda *LDAClient) listCountries() []Country {
	resp := lda.get(lda.Cfg.EndpointInfo["countries"])
	body := []byte{}
	resp.Body.Read(body)
	countries := []Country{}
	json.Unmarshal(body, countries)
	return countries
}

func (lda *LDAClient) listStates() []State {
	resp := lda.get(lda.Cfg.EndpointInfo["states"])
	body := []byte{}
	resp.Body.Read(body)
	states := []State{}
	json.Unmarshal(body, states)
	return states
}

func (lda *LDAClient) listLobbyistPrefixes() []LobbyistPrefix {
	resp := lda.get(lda.Cfg.EndpointInfo["lobbyistPrefixes"])
	body := []byte{}
	resp.Body.Read(body)
	lobbyistPrefixes := []LobbyistPrefix{}
	json.Unmarshal(body, lobbyistPrefixes)
	return lobbyistPrefixes
}

func (lda *LDAClient) listLobbyistSuffixes() []LobbyistSuffix {
	resp := lda.get(lda.Cfg.EndpointInfo["lobbyistSuffixes"])
	body := []byte{}
	resp.Body.Read(body)
	lobbyistSuffixes := []LobbyistSuffix{}
	json.Unmarshal(body, lobbyistSuffixes)
	return lobbyistSuffixes
}

func (lda *LDAClient) listContributionItemTypes() []ContributionItemType {
	resp := lda.get(lda.Cfg.EndpointInfo["contributionItemTypes"])
	body := []byte{}
	resp.Body.Read(body)
	contributionItemTypes := []ContributionItemType{}
	json.Unmarshal(body, contributionItemTypes)
	return contributionItemTypes
}
