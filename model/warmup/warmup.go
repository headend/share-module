package warmup

type WarmupMessage struct {
	EventTime	int64 `json:"event_time"`
	Data	[]WarmupElement `json:"data"`
}

type WarmupElement struct {
	IP	string	`json:"ip"`
	Port	string `json:"port"`
	Cpu	int	`json:"cpu"`
	Ram int `json:"ram"`
	Disk int `json:"disk"`
}