package register

import "encoding/json"

type Register struct {
	IP		string	`json:"ip"`
	Port	string `json:"port"`
	Cpu		int	`json:"cpu"`
	Ram		int `json:"ram"`
	Disk	int `json:"disk"`
	EventTime int64 `json:"event_time"`
}

func (NRA *Register) LoadFromJsonString(JsonString string) (err error) {
	err = json.Unmarshal([]byte(JsonString), NRA)
	if err != nil {
		return err
	}
	return
}

func (NRA *Register) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(NRA)
	if err != nil {
		return "", err
	}
	return string(b), nil
}