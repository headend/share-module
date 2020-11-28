package communication

//WarmUpMessage để sync xuống MQ
type WarmUpMessage struct {
	// Slice of agent info
	// required: true
	// min length: 1
	Time   int64          `json:"time"`
	// Slice of agent status info
	// required: true
	// min length: 1
	Agents []AgentStatus `json:"data"`
}
//AgentStatus để làm slice của json
type AgentStatus struct {
	// UUID gen from agent private key
	// required: true
	// min length: 1
	UUID     string `json:"UUID"`
	// Status
	// required: true
	// min: 0
	Status bool   `json:"status"`
	// Slice of string version
	// required: true
	// min length: 1
	Version []string `json:"version"`
	// 1-true allow connect, 0-false not allow connect
	// required: true
	// min length: false
	IsEnabled bool `json:"is_enabled"`
}

