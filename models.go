package main

import "time"

// Filing represents the structure for a filing record
type Filing struct {
	URL                        string                   `json:"url"`
	FilingUUID                 string                   `json:"filing_uuid"`
	FilingType                 string                   `json:"filing_type"`
	FilingTypeDisplay          string                   `json:"filing_type_display"`
	FilingYear                 int                      `json:"filing_year"`
	FilingPeriod               string                   `json:"filing_period"`
	FilingPeriodDisplay        string                   `json:"filing_period_display"`
	FilingDocumentURL          string                   `json:"filing_document_url"`
	FilingDocumentContentType  string                   `json:"filing_document_content_type"`
	Income                     string                   `json:"income"`
	Expenses                   string                   `json:"expenses"`
	ExpensesMethod             string                   `json:"expenses_method"`
	ExpensesMethodDisplay      string                   `json:"expenses_method_display"`
	PostedByName               string                   `json:"posted_by_name"`
	DtPosted                   time.Time                `json:"dt_posted"`
	TerminationDate            string                   `json:"termination_date"`
	RegistrantCountry          string                   `json:"registrant_country"`
	RegistrantPPBCountry       string                   `json:"registrant_ppb_country"`
	RegistrantAddress1         string                   `json:"registrant_address_1"`
	RegistrantAddress2         string                   `json:"registrant_address_2"`
	RegistrantDifferentAddress bool                     `json:"registrant_different_address"`
	RegistrantCity             string                   `json:"registrant_city"`
	RegistrantState            string                   `json:"registrant_state"`
	RegistrantZIP              string                   `json:"registrant_zip"`
	Registrant                 Registrant               `json:"registrant"`
	Client                     Client                   `json:"client"`
	LobbyingActivities         []LobbyingActivity       `json:"lobbying_activities"`
	ConvictionDisclosures      []ConvictionDisclosure   `json:"conviction_disclosures"`
	ForeignEntities            []ForeignEntity          `json:"foreign_entities"`
	AffiliatedOrganizations    []AffiliatedOrganization `json:"affiliated_organizations"`
}

// Registrant represents the structure for a registrant
type Registrant struct {
	ID                int       `json:"id"`
	URL               string    `json:"url"`
	HouseRegistrantID int       `json:"house_registrant_id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Address1          string    `json:"address_1"`
	Address2          string    `json:"address_2"`
	Address3          string    `json:"address_3"`
	Address4          string    `json:"address_4"`
	City              string    `json:"city"`
	State             string    `json:"state"`
	StateDisplay      string    `json:"state_display"`
	ZIP               string    `json:"zip"`
	Country           string    `json:"country"`
	CountryDisplay    string    `json:"country_display"`
	PPBCountry        string    `json:"ppb_country"`
	PPBCountryDisplay string    `json:"ppb_country_display"`
	ContactName       string    `json:"contact_name"`
	ContactTelephone  string    `json:"contact_telephone"`
	DtUpdated         time.Time `json:"dt_updated"`
}

// Client represents the structure for a client
type Client struct {
	ID                     int        `json:"id"`
	URL                    string     `json:"url"`
	ClientID               string     `json:"client_id"`
	Name                   string     `json:"name"`
	GeneralDescription     string     `json:"general_description"`
	ClientGovernmentEntity bool       `json:"client_government_entity"`
	ClientSelfSelect       bool       `json:"client_self_select"`
	State                  string     `json:"state"`
	StateDisplay           string     `json:"state_display"`
	Country                string     `json:"country"`
	CountryDisplay         string     `json:"country_display"`
	PPBState               string     `json:"ppb_state"`
	PPBStateDisplay        string     `json:"ppb_state_display"`
	PPBCountry             string     `json:"ppb_country"`
	PPBCountryDisplay      string     `json:"ppb_country_display"`
	EffectiveDate          time.Time  `json:"effective_date"`
	Registrant             Registrant `json:"registrant"`
}

// LobbyingActivity represents the structure for lobbying activities
type LobbyingActivity struct {
	GeneralIssueCode        string      `json:"general_issue_code"`
	GeneralIssueCodeDisplay string      `json:"general_issue_code_display"`
	Description             string      `json:"description"`
	ForeignEntityIssues     string      `json:"foreign_entity_issues"`
	Lobbyists               []Lobbyist  `json:"lobbyists"`
	GovernmentEntities      []GovEntity `json:"government_entities"`
}

// Lobbyist represents the structure for a lobbyist involved in lobbying activities
type Lobbyist struct {
	ID            int        `json:"id"`
	Prefix        string     `json:"prefix"`
	PrefixDisplay string     `json:"prefix_display"`
	FirstName     string     `json:"first_name"`
	Nickname      string     `json:"nickname"`
	MiddleName    string     `json:"middle_name"`
	LastName      string     `json:"last_name"`
	Suffix        string     `json:"suffix"`
	SuffixDisplay string     `json:"suffix_display"`
	Registrant    Registrant `json:"registrant"`
}

// Person represents an individual with prefixes, suffixes, and name details
type Person struct {
	ID            int    `json:"id"`
	Prefix        string `json:"prefix"`
	PrefixDisplay string `json:"prefix_display"`
	FirstName     string `json:"first_name"`
	Nickname      string `json:"nickname"`
	MiddleName    string `json:"middle_name"`
	LastName      string `json:"last_name"`
	Suffix        string `json:"suffix"`
	SuffixDisplay string `json:"suffix_display"`
}

// GovEntity represents a government entity involved in lobbying
type GovEntity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ConvictionDisclosure represents a disclosed conviction for a lobbyist
type ConvictionDisclosure struct {
	Lobbyist    Person `json:"lobbyist"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

// ForeignEntity represents a foreign entity involved in lobbying
type ForeignEntity struct {
	Name                string `json:"name"`
	Contribution        string `json:"contribution"`
	OwnershipPercentage string `json:"ownership_percentage"`
	Address             string `json:"address"`
	City                string `json:"city"`
	State               string `json:"state"`
	StateDisplay        string `json:"state_display"`
	Country             string `json:"country"`
	CountryDisplay      string `json:"country_display"`
	PPBCity             string `json:"ppb_city"`
	PPBState            string `json:"ppb_state"`
	PPBStateDisplay     string `json:"ppb_state_display"`
	PPBCountry          string `json:"ppb_country"`
	PPBCountryDisplay   string `json:"ppb_country_display"`
}

// AffiliatedOrganization represents an affiliated organization involved in lobbying
type AffiliatedOrganization struct {
	Name              string `json:"name"`
	URL               string `json:"url"`
	Address1          string `json:"address_1"`
	Address2          string `json:"address_2"`
	Address3          string `json:"address_3"`
	Address4          string `json:"address_4"`
	City              string `json:"city"`
	State             string `json:"state"`
	StateDisplay      string `json:"state_display"`
	ZIP               string `json:"zip"`
	Country           string `json:"country"`
	CountryDisplay    string `json:"country_display"`
	PPBCity           string `json:"ppb_city"`
	PPBState          string `json:"ppb_state"`
	PPBStateDisplay   string `json:"ppb_state_display"`
	PPBCountry        string `json:"ppb_country"`
	PPBCountryDisplay string `json:"ppb_country_display"`
}

type ContributionReport struct {
	URL                       string             `json:"url"`
	FilingUUID                string             `json:"filing_uuid"`
	FilingType                string             `json:"filing_type"`
	FilingTypeDisplay         string             `json:"filing_type_display"`
	FilingYear                int                `json:"filing_year"`
	FilingPeriod              string             `json:"filing_period"`
	FilingPeriodDisplay       string             `json:"filing_period_display"`
	FilingDocumentURL         string             `json:"filing_document_url"`
	FilingDocumentContentType string             `json:"filing_document_content_type"`
	FilerType                 string             `json:"filer_type"`
	FilerTypeDisplay          string             `json:"filer_type_display"`
	DtPosted                  time.Time          `json:"dt_posted"`
	ContactName               string             `json:"contact_name"`
	Comments                  string             `json:"comments"`
	Address1                  string             `json:"address_1"`
	Address2                  string             `json:"address_2"`
	City                      string             `json:"city"`
	State                     string             `json:"state"`
	StateDisplay              string             `json:"state_display"`
	ZIP                       string             `json:"zip"`
	Country                   string             `json:"country"`
	CountryDisplay            string             `json:"country_display"`
	Registrant                Registrant         `json:"registrant"`
	Lobbyist                  Lobbyist           `json:"lobbyist"`
	NoContributions           bool               `json:"no_contributions"`
	PACs                      []string           `json:"pacs"`
	ContributionItems         []ContributionItem `json:"contribution_items"`
}

// ContributionItem represents a single contribution record
type ContributionItem struct {
	ContributionType        string    `json:"contribution_type"`
	ContributionTypeDisplay string    `json:"contribution_type_display"`
	ContributorName         string    `json:"contributor_name"`
	PayeeName               string    `json:"payee_name"`
	HonoreeName             string    `json:"honoree_name"`
	Amount                  string    `json:"amount"`
	Date                    time.Time `json:"date"`
}

// FilingType represents the structure for a filing type
type FilingType struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// FilingTypes represents a list of multiple filing types
type FilingTypes []FilingType

// LobbyingIssue represents the structure for a lobbying activity issue
type LobbyingIssue struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// LobbyingIssues represents a list of lobbying activity issues
type LobbyingIssues []LobbyingIssue

// GovernmentEntity represents the structure for a government entity
type GovernmentEntity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GovernmentEntities represents a list of government entities
type GovernmentEntities []GovernmentEntity

// Country represents the structure for a country entity
type Country struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// Countries represents a list of country entities
type Countries []Country

// State represents the structure for a U.S. state or territory
type State struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// States represents a list of states and territories
type States []State

// LobbyistPrefix represents the structure for a lobbyist prefix
type LobbyistPrefix struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// LobbyistPrefixes represents a list of lobbyist prefixes
type LobbyistPrefixes []LobbyistPrefix

// LobbyistSuffix represents the structure for a lobbyist suffix
type LobbyistSuffix struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// LobbyistSuffixes represents a list of lobbyist suffixes
type LobbyistSuffixes []LobbyistSuffix

// ContributionItemType represents the structure for a contribution item type
type ContributionItemType struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// ContributionItemTypes represents a list of contribution item types
type ContributionItemTypes []ContributionItemType
