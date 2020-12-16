package model

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
	1100 = urgent task
	1101 = command shell
	 */
	ExeType	int	`json:"exe_type"`
	/*
	Execute id from frontend
	 */
	ExeId	int	`json:"exe_id"`
}
type AgentExeRequest struct {
	AgentIp	string	`json:"agent_ip"`
	/*
		Execute type
		1100 = urgent task
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
	ProdileIp	string	`json:"prodile_ip"`
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
