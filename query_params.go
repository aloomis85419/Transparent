package main

import "time"

//Filing query params

var filingYears = []int{
	1999,
}

var filingPeriod = []string{
	"first_quarter", "second_quarter", "third_quarter", "fourth_quarter", "mid_year", "year_end",
}

var filingTypes []FilingType
var governmentEntities []GovernmentEntity
var lobbyActivityIssues []LobbyingActivityIssue
var countries []Country
var states []State
var lobbyistPrefixes []LobbyistPrefix
var lobbyistSuffixes []LobbyistSuffix
var contributionItemTypes []ContributionItemType

var filingQueryParams = struct {
	AffiliatedOrganizationCountry         string
	AffiliatedOrganizationListedIndicator string
	AffiliatedOrganizationName            string
	ClientCountry                         string
	ClientID                              string
	ClientName                            string
	ClientPPBCountry                      string
	ClientPPBState                        string
	ClientState                           string
	FilingAmountReportedMax               string
	FilingAmountReportedMin               string
	FilingDtPostedAfter                   string
	FilingDtPostedBefore                  string
	FilingPeriod                          string
	FilingSpecificLobbyingIssues          string
	FilingType                            string
	FilingUUID                            string
	FilingYear                            string
	ForeignEntityCountry                  string
	ForeignEntityListedIndicator          string
	ForeignEntityName                     string
	ForeignEntityOwnershipPercentageMax   string
	ForeignEntityOwnershipPercentageMin   string
	ForeignEntityPPBCountry               string
	LobbyistConvictionDateRangeAfter      string
	LobbyistConvictionDateRangeBefore     string
	LobbyistConvictionDisclosure          string
	LobbyistConvictionDisclosureIndicator string
	LobbyistCoveredPosition               string
	LobbyistCoveredPositionIndicator      string
	LobbyistID                            string
	LobbyistName                          string
	Ordering                              string
	Page                                  string
	PageSize                              string
	RegistrantCountry                     string
	RegistrantID                          string
	RegistrantName                        string
	RegistrantPPBCountry                  string
}{"AffiliatedOrganizationCountry", "AffiliatedOrganizationListedIndicator", "AffiliatedOrganizationName",
	"ClientCountry", "ClientID", "ClientName", "ClientPPBCountry", "ClientPPBState", "ClientState",
	"FilingAmountReportedMax", "FilingAmountReportedMin", "FilingDtPostedAfter", "FilingDtPostedBefore", "" +
		"FilingPeriod", "FilingSpecificLobbyingIssues", "FilingType", "FilingUUID", "FilingYear", "ForeignEntityCountry", "ForeignEntityListedIndicator",
	"ForeignEntityName", "ForeignEntityOwnershipPercentageMax", "ForeignEntityOwnershipPercentageMin", "ForeignEntityPPBCountry",
	"LobbyistConvictionDateRangeAfter", "LobbyistConvictionDateRangeBefore",
	"LobbyistConvictionDisclosure", "LobbyistConvictionDisclosureIndicator", "LobbyistCoveredPosition", "LobbyistCoveredPositionIndicator", "LobbyistID",
	"LobbyistName", "Ordering", "Page", "PageSize", "RegistrantCountry", "RegistrantID", "RegistrantName", "RegistrantPPBCountry",
}

var contributionQueryParams = struct {
	ContributionAmountMax   string
	ContributionAmountMin   string
	ContributionContributor string
	ContributionDateAfter   string
	ContributionDateBefore  string
	ContributionHonoree     string
	ContributionPayee       string
	ContributionType        string
	FilingDtPostedAfter     string
	FilingDtPostedBefore    string
	FilingPeriod            string
	FilingType              string
	FilingUUID              string
	FilingYear              string
	LobbyistExclude         string
	LobbyistID              string
	LobbyistName            string
	Ordering                string
	Page                    string
	PageSize                string
	RegistrantID            string
	RegistrantName          string
}{
	"ContributionAmountMax", "ContributionAmountMin", "ContributionContributor", "ContributionDateAfter", "ContributionDateBefore",
	"ContributionHonoree", "ContributionPayee", "ContributionType", "FilingDtPostedAfter", "FilingDtPostedBefore", "FilingPeriod",
	"FilingType", "FilingUUID", "FilingYear", "LobbyistExclude", "LobbyistID", "LobbyistName", "Ordering", "Page", "PageSize",
	"RegistrantID", "RegistrantName",
}

var registrantQueryParams = struct {
	Country         string
	DtUpdatedAfter  string
	DtUpdatedBefore string
	ID              string
	Ordering        string
	Page            string
	PageSize        string
	PPBCountry      string
	RegistrantName  string
	State           string
}{
	"Country", "DtUpdatedAfter", "DtUpdatedBefore", "ID", "Ordering",
	"Page", "PageSize", "PPBCountry", "RegistrantName", "State",
}

var clientQueryParams = struct {
	ClientCountry    string
	ClientName       string
	ClientPPBCountry string
	ClientPPBState   string
	ClientState      string
	ID               string
	Ordering         string
	Page             string
	PageSize         string
	RegistrantID     string
	RegistrantName   string
}{
	"ClientCountry", "ClientName", "ClientPPBCountry", "ClientPPBState", "ClientState",
	"ID", "Ordering", "Page", "PageSize", "RegistrantID", "RegistrantName",
}

var lobbyistQueryParams = struct {
	ID             string
	LobbyistName   string
	Ordering       string
	Page           string
	PageSize       string
	RegistrantID   string
	RegistrantName string
}{
	"ID", "LobbyistName", "Ordering", "Page", "PageSize", "RegistrantID", "RegistrantName",
}

func loadDynamicConstants() {
	populateFilingYears()
	filingTypes = client.listFilingTypes()
	countries = client.listCountries()
	governmentEntities = client.listGovernmentEntities()
	lobbyActivityIssues = client.listLobbyingActivityGeneralIssues()
	states = client.listStates()
	lobbyistPrefixes = client.listLobbyistPrefixes()
	lobbyistSuffixes = client.listLobbyistSuffixes()
	contributionItemTypes = client.listContributionItemTypes()
}

func populateFilingYears() {
	for i, j := filingYears[0], 0; i <= time.Now().Year(); i, j = i+1, j+1 {
		filingYears[j] = i
	}
}
