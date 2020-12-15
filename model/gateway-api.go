package model

type AgentCtlRequest struct {
	AgentIp	string	`json:"agent_ip"`
	ActionType	int	`json:"action_type"`
	ActionId	int	`json:"action_id"`
	TunnelData	map[string]string	`json:"tunnel_data"`
}

type AgentCtlResponse struct {
	ReturnCode	string `json:"return_code"`
	ReturnData	AgentCtlDataResponse	`json:"return_data"`
	ReturnMessage	string	`json:"return_message"`
	TunnelData	map[string]string	`json:"tunnel_data"`
}

type AgentCtlDataResponse struct {
	ActionId	int	`json:"action_id"`
	ActionType	int	`json:"action_type"`
}

