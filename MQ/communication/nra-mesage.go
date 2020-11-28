package communication

import (
	"encoding/json"
)

//NRAMessageRequest để request NRA
type NRAMessageRequest struct {
	// Event time
	// required: false
	// min: 1
	Time int64 `json:"time"`
	// Slice of agent info
	// required: true
	// min length: 1
	Agent NewRegisterAgentRequest `json:"data"`
}

//NewRegisterAgentRequest để khởi tạo struct thông tin của agent gửi sang nra
type NewRegisterAgentRequest struct {
	// UUID gen from private key, use it instead of ID
	// required: true
	// min length: 1
	UUID string `json:"uuid"`
	// Agent's IP
	// required: true
	// min length: 7
	IP string `json:"ip"`
	// Agent public key
	// required: true
	// min length: 1
	PublicKey string `json:"public_key"`
	Info      string `json:"Info"`
}

type NRAAuthentication struct {
	UUID       chan string
	IP         string
	PublicKey  string
	Info       string
	UpdateInfo bool
}

//NRAMessageResponse nhận kết quả trả về của bên nhận đăng ký
type NRAMessageResponse struct {
	// Event time
	// required: true
	// min: 1
	Time uint `json:"time"`
	// Slice of agent info
	// required: true
	// min length: 1
	Agent NewRegisterAgentResponse `json:"data"`
}

//NewRegisterAgentResponse là kết quả trả về thông tin nhận được
type NewRegisterAgentResponse struct {
	// UUID gen from private key, use it instead of ID
	// required: true
	// min length: 1
	UUID string `json:"uuid"`
	// Agent's IP
	// required: true
	// min length: 1
	IP string `json:"ip"`
	// Backend public key
	// required: true
	// min length: 7
	PublicKey string `json:"public_key"`
}

//ToString convert NRAMessageRequest to json string
func (Object *NRAMessageRequest) ToString() (string, error) {
	jsonbin, err := json.Marshal(Object)
	return string(jsonbin), err
}

//ToString convert NRAMessageResponse to json string
func (Object *NRAMessageResponse) ToString() (string, error) {
	jsonbin, err := json.Marshal(Object)
	return string(jsonbin), err
}

//ConvertToStructResponese convert data to struct NRAMessageResponse
func ConvertToStructResponese(data []byte) (*NRAMessageResponse, error) {
	ReturnStruct := &NRAMessageResponse{}
	err := json.Unmarshal(data, ReturnStruct)
	return ReturnStruct, err
}

//ConvertToStructRequest convert data to struct NRAMessageRequest
func ConvertToStructRequest(data []byte) (*NRAMessageRequest, error) {
	ReturnStruct := &NRAMessageRequest{}
	err := json.Unmarshal(data, ReturnStruct)
	return ReturnStruct, err
}

//AgentdAuthen sẽ gửi lên những data này
type AgentdAuthen struct {
	State         State
	UUID          string
	IP            string
	PublicKey     string
	MasterVersion string
	Info          string
	InfoDeny      string
	TransactionID string
	GroupCommand  []string
	KeyCode       string
	ProcessName   string
}

type State uint8

const (
	RequestNRA = iota + 1
	AcceptedNRA
	DenyNRA
	RequestInfoOnceMore
	SendInfo
)
