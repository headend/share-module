package model

import (
	"encoding/json"
)

type AgentCTLQueueRequest struct {
	AgentCtlRequest	
	ControlType int	`json:"control_type"`
	EventTime int64	`json:"event_time"`
}

type AgentEXEQueueRequest struct {
	AgentExeRequest
	ExeType int	`json:"exe_type"`
	EventTime int64	`json:"event_time"`
}

func (actlr *AgentCTLQueueRequest) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(actlr)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (exer *AgentEXEQueueRequest) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(exer)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
