package warmup

import "encoding/json"

type WarmupMessage struct {
	EventTime	int64 `json:"event_time"`
	Data	[]WarmupElement `json:"data"`
	/*
	event: warmup from event connect/disconect from agent (default)
	interval: intervel system check status agent
	*/
	WupType	string `json:"wup_type"`
}

type WarmupElement struct {
	IP	string	`json:"ip"`
	Status	bool	`json:"status"`
	Port	string `json:"port"`
	Cpu	int	`json:"cpu"`
	Ram int `json:"ram"`
	Disk int `json:"disk"`
}

func (WM *WarmupMessage) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(WM)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (WM *WarmupMessage) LoadFromJsonString(JsonString string) (err error) {
	err = json.Unmarshal([]byte(JsonString), WM)
	if err != nil {
		return err
	}
	return
}

func (WE *WarmupElement) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(WE)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (WE *WarmupElement) LoadFromJsonString(JsonString string) (err error) {
	err = json.Unmarshal([]byte(JsonString), WE)
	if err != nil {
		return err
	}
	return
}
