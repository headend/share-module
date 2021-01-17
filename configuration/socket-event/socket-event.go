package socket_event

// topic declare
const (
	// event to receive notice
	ThongBao  = "notice"
	// event to receive message
	TinNhan  = "message"
	// event to recieve error
	Loi = "error"
	// event to receive connection message
	KetNoi = "connection"
	// event to handle disconnect event
	NgatKetNoi = "disconnection"
	// event recieve file
	NhanFile = "receivefile"
	// event recieve log
	NhanLog = "log"
	// event recieve update
	CapNhat = "update-request"
	// event update response
	KetQuaCapNhat = "update-response"
	// Nhom chung
	NhomChung = "nhomchung"
	// Nhom source goc
	NhomOrigin = "nhomorigin"
	// Nhom source sau encode
	NhomSauEncode = "nhomsauencode"
	// Nhom source ha tang
	NhomHaTang = "nhomhatang"
	// Thuc thi lenh
	ThucThiLenh = "execute-request"
	// Ket qua thuc thi lenh
	KetQuaThucThiLenh = "execute-response"
	// Dieu khien stop/ start
	DieuKhien = "control"
	UpdateWorker = "worker-update"
	MonitorResponse = "monitor-response"
	ProfileRequest = "profile-monitor-request"
	ProfileResponse = "profile-monitor-response"
	PingPing = "ping"
	PongPong = "pong"
)



