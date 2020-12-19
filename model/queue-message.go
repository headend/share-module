package model

import (
	"encoding/json"
)


type AgentExeSingleRequest struct {
	AgentId	int64	`json:"agent_id"`
	/*
		Execute type
		1100 = urgent task (default)
		1101 = command shell
	*/
	ExeType	int	`json:"exe_type"`
	/*
		Execute id from frontend
	*/
	ExeId	int	`json:"exe_id"`
	/*
		Data transmit through frontend --> backend --> agent --> backend --> frontend
	*/
	TunnelData	map[string]string	`json:"tunnel_data"`
}

type AgentCTLQueueRequest struct {
	AgentCtlRequest	
	ControlType int	`json:"control_type"`
	EventTime int64	`json:"event_time"`
}

type AgentEXEQueueRequest struct {
	AgentExeSingleRequest
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
