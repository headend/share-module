package register

import "encoding/json"

type Register struct {
	IP		string	`json:"ip"`
	Port	string `json:"port"`
	Cpu		int	`json:"cpu"`
	Ram		int `json:"ram"`
	Disk	int `json:"disk"`
	Uuid	string	`json:"uuid"`
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

type RegisterReply struct {
	id		int32	`json:"id"`
	IP		string	`json:"ip"`
	Port	string `json:"port"`
	Uuid	string	`json:"uuid"`
	EventTime int64 `json:"event_time"`
}

func (NRAR *RegisterReply) LoadFromJsonString(JsonString string) (err error) {
	err = json.Unmarshal([]byte(JsonString), NRAR)
	if err != nil {
		return err
	}
	return
}

func (NRAR *RegisterReply) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(NRAR)
	if err != nil {
		return "", err
	}
	return string(b), nil
}