package model

import (
	"encoding/json"
)

type ProfileForAgentdElement struct {
	MonitorId	int64	`json:"monitor_id"`
	ProfileId	int64	`json:"profile_id"`
	AgentId		int64	`json:"agent_id"`
	Status		int64	`json:"status"`
	VideoStatus	bool	`json:"video_status"`
	AudioStatus bool	`json:"audio_status"`
	MulticastIP string	`json:"multicast_ip"`
}

type MonitorInputForAgent struct {
	MonitorType	int	`json:"monitor_type"`
	ProfileList []ProfileForAgentdElement `json:"profile_list"`
}

func (PFA *ProfileForAgentdElement) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(PFA)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (PFA *ProfileForAgentdElement) LoadFromJsonString(JsonString string) (err error) {
	err = json.Unmarshal([]byte(JsonString), PFA)
	if err != nil {
		return err
	}
	return
}

func (MIFA *MonitorInputForAgent) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(MIFA)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (MIFA *MonitorInputForAgent) LoadFromJsonString(JsonString string) (err error) {
	err = json.Unmarshal([]byte(JsonString), MIFA)
	if err != nil {
		return err
	}
	return
}

type ProfileChangeStatus struct {
	MonitorType		int	`json:"monitor_type"`
	MonitorID		int64 `json:"monitor_id"`
	ProfileId		int64	`json:"profile_id"`
	AgentId			int64	`json:"agent_id"`
	OldStatus		int64	`json:"old_status"`
	NewStatus		int64	`json:"new_status"`
	OldSignalStatus	bool	`json:"old_signal_status"`
	NewSignalStatus	bool	`json:"new_signal_status"`
	OldVideoStatus	bool	`json:"old_video_status"`
	NewVideoStatus	bool	`json:"new_video_status"`
	OldAudioStatus	bool	`json:"old_audio_status"`
	NewAudioStatus	bool	`json:"new_audio_status"`
	EventTime int64	`json:"event_time"`
}

func (PCS *ProfileChangeStatus) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(PCS)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (PCS *ProfileChangeStatus) LoadFromJsonString(JsonString string) (err error) {
	err = json.Unmarshal([]byte(JsonString), PCS)
	if err != nil {
		return err
	}
	return
}

type Register struct {
	Id		int32	`json:"id"`
	Uuid	int32	`json:"uuid"`
}

