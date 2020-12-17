package model

import "encoding/json"

type AgentCtlRequest struct {
	/*
	Agent control ip
	 */
	AgentIp	string	`json:"agent_ip"`
	/*
		Control ID from frontend
	*/
	ControlId	int	`json:"control_id"`
	/*
	Data transmit through frontend --> backend --> agent --> backend --> frontend
	 */
	TunnelData	map[string]string	`json:"tunnel_data"`
}

type AgentCtlResponse struct {
	/*
	Return code
		1 - Success
		0 - Fail
	 */
	ReturnCode	int `json:"return_code"`
	/*
	Data
	 */
	ReturnData	AgentCtlDataResponse	`json:"return_data"`
	/*
	Message success or fail
	 */
	ReturnMessage	string	`json:"return_message"`
	/*
		Data transmit through frontend --> backend --> agent --> backend --> frontend
	*/
	TunnelData	map[string]string	`json:"tunnel_data"`
}

type AgentCtlDataResponse struct {
	/*
		Control ID from frontend
	*/
	ControlId	int	`json:"control_id"`
}

type AgentExeDataResponse struct {
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
		Profile IP need check
	*/
	ListProfileId	[]string	`json:"list_profile_id"`
	/*
		Chose agent to run check (default all if null)
	*/
	ListAgentToRun	[]string	`json:"list_agent_to_run"`
}
type AgentExeRequest struct {
	AgentIp	string	`json:"agent_ip"`
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
	Profile IP need check
	 */
	ListProfileId	[]int	`json:"list_profile_id"`
	/*
		Chose agent to run check (default all if null)
	*/
	ListAgentToRun	[]string	`json:"list_agent_to_run"`
	/*
		Data transmit through frontend --> backend --> agent --> backend --> frontend
	*/
	TunnelData	map[string]string	`json:"tunnel_data"`
}
type AgentExeResponse struct {
	/*
		Return code
			1 - Success
			0 - Fail
	*/
	ReturnCode	int `json:"return_code"`
	/*
		Data
	*/
	ReturnData	AgentExeDataResponse	`json:"return_data"`
	/*
		Message success or fail
	*/
	ReturnMessage	string	`json:"return_message"`
	/*
		Data transmit through frontend --> backend --> agent --> backend --> frontend
	*/
	TunnelData	map[string]string	`json:"tunnel_data"`
}

func (actlr *AgentCtlRequest) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(actlr)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (exer *AgentExeRequest) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(exer)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (actlrsp *AgentCtlResponse) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(actlrsp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (exersp *AgentExeResponse) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(exersp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}