package model

type AgentCTLRequest struct {
	AgentCtlRequest	
	ControlType int	`json:"control_type"`
	EventTime int64	`json:"event_time"`
}

type AgentEXERequest struct {
	AgentExeRequest
	ExeType int	`json:"exe_type"`
	EventTime int64	`json:"event_time"`
}


