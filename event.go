package main

import (
	"encoding/json"
)

type Event struct {
	userIdentity      Identity
	EventTime         string          `json:"eventTime"`
	EventSource       string          `json:"eventSource"`
	AwsRegion         string          `json:"awsRegion"`
	EventName         string          `json:"eventName"`
	SourceIPAddress   string          `json:"sourceIPAddress"`
	UserAgent         string          `json:"userAgent"`
	RequestParameters json.RawMessage `json:"requestParameters"`
	ResponseElements  json.RawMessage `json:"responseElements"`
	EventType         string          `json:"eventType"`
	EventID           string          `json:"eventID"`
}

type Identity struct {
	TypeIdentity string `json:"type"`
	PrincipalID  string `json:"principalId"`
	Arn          string `json:"arn"`
	AccountID    string `json:"accountId"`
	AccessKeyID  string `json:"accessKeyId"`
	UserName     string `json:"userName"`
}

func NewEvent() *Event {
	return &Event{}
}
