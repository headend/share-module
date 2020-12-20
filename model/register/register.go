package register

type Register struct {
	IP		string	`json:"ip"`
	Port	string `json:"port"`
	Cpu		int	`json:"cpu"`
	Ram		int `json:"ram"`
	Disk	int `json:"disk"`
	EventTime int64 `json:"event_time"`
}
