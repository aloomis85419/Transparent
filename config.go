package main

type Config struct {
	BaseUrl         string
	AuthInfo        AuthInfo
	EndpointInfo    map[string]EndpointInfo
	DefaultPageSize int
}

type EndpointInfo struct {
	Name string
	Url  string
}

type AuthInfo struct {
	ApiToken string
}

var ENDPOINTS = map[string]EndpointInfo{
	"filings":                {Name: "filings", Url: "filings/"},
	"contributions":          {Name: "contributions", Url: "contributions/"},
	"registrants":            {Name: "registrants", Url: "registrants/"},
	"clients":                {Name: "clients", Url: "clients/"},
	"lobbyists":              {Name: "lobbyists", Url: "lobbyists/"},
	"filingtypes":            {Name: "constants/filing/filingtypes", Url: "constants/filing/filingtypes/"},
	"lobbyingactivityissues": {Name: "constants/filing/lobbyingactivityissues", Url: "constants/filing/lobbyingactivityissues/"},
	"governmententities":     {Name: "constants/filing/governmententities", Url: "constants/filing/governmententities/"},
	"countries":              {Name: "constants/general/countries", Url: "constants/general/countries/"},
	"states":                 {Name: "constants/general/states", Url: "constants/general/states/"},
	"lobbyistPrefixes":       {Name: "constants/lobbyist/prefixes", Url: "constants/lobbyist/prefixes/"},
	"lobbyistSuffixes":       {Name: "constants/lobbyist/suffixes", Url: "constants/lobbyist/suffixes/"},
	"contributionItemTypes":  {Name: "constants/contribution/itemtypes/", Url: "constants/contribution/itemtypes/"},
}

func Configure(apiToken string) Config {
	cfg := Config{}
	authInfo := AuthInfo{}
	authInfo.ApiToken = apiToken
	cfg.BaseUrl = "https://lda.senate.gov/api/v1/"
	cfg.EndpointInfo = ENDPOINTS
	cfg.AuthInfo = authInfo
	cfg.DefaultPageSize = 100
	return cfg
}
