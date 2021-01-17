package model

import "encoding/json"

type WorkerUpdateSignal struct {
	FileName	string	`json:"fileName"`
	FilePath	string	`json:"filePath"`
	FileSizeInByte	int64	`json:"fileSizeInByte"`
	Md5	string	`json:"md5"`
}

func (wu *WorkerUpdateSignal) GetJsonString() (JsonString string, err error) {
	b, err := json.Marshal(wu)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (wu *WorkerUpdateSignal) LoadFromJsonString(JsonString string) (err error) {
	err = json.Unmarshal([]byte(JsonString), wu)
	if err != nil {
		return err
	}
	return
}
