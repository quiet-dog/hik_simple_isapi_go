package hik

import "encoding/xml"

type GateWayConnect struct {
	HikConfig HikConfig
	IsConnect bool
}

type UserCheckRes struct {
	XMLNAME           xml.Name `xml:"userCheck"`
	StatusValue       int      `xml:"statusValue"`
	StatusString      string   `xml:"statusString"`
	IsDefaultPassword bool     `xml:"isDefaultPassword"`
	IsRiskPassword    bool     `xml:"isRiskPassword"`
	IsActivated       bool     `xml:"isActivated"`
	ResidualValidity  int      `xml:"residualValidity"`
	LockStatus        string   `xml:"lockStatus"`
	UnlockTime        int      `xml:"unlockTime"`
	RetryLoginTime    int      `xml:"retryLoginTime"`
}

type ResponseStatusXML struct {
	XMLNAME       xml.Name `xml:"ResponseStatus"`
	RequestURL    string   `xml:"requestURL" json:"requestURL"`
	StatusCode    int      `xml:"statusCode" json:"statusCode"`
	StatusString  string   `xml:"statusString" json:"statusString"`
	SubStatusCode string   `xml:"subStatusCode" json:"subStatusCode"`
	ErrorCode     int      `xml:"errorCode" json:"errorCode"`
	ErrorMsg      string   `xml:"errorMsg" json:"errorMsg"`
}

type ErrorMsg struct {
	StatusCode    int    `json:"statusCode"`
	StatusString  string `json:"statusString"`
	SubStatusCode string `json:"subStatusCode"`
	ErrorCode     int    `json:"errorCode"`
	ErrorMsg      string `json:"errorMsg"`
}

type Msg struct {
	Ip   string
	Data Result
	Type string
	Msg  string
}

const (
	SUCCESS = "success"
	ERROR   = "error"
)
