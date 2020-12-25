package static_config

const (
	GatewayHost = "0.0.0.0"
	GatewayPort = 8000
	GatewayStorage = "asset/"
	GatewayUser = "user"
	GatewayPassword = "pass"
	GatewayTransportProtocol = "websocket"
	MasterHost = "127.0.0.1"
	MasterPort = 45656
	MasterStorage = "asset/"
	MasterUser = "user"
	MasterPassword = "pass"
	MasterTransportProtocol = "websocket"
	InstallationPath = "/opt/iptv"
	LogPath = "/opt/iptv/log"
	BinaryPath = "/opt/iptv/sbin"
	ConfigFilePath = "/opt/iptv/application.yml"
	VersionFilePath = "/opt/iptv/version"
)

// signal declare
const (
	// start agent
	StartAgentd = 10000
	// stop agent
	StopAgentd = 10001
	// Connect agent
	ConnectAgentd = 10030
	// Disconnect agent
	DisconnectAgentd = 10031
	// Start Worker
	StartWorker = 20000
	// Stop worker
	StopWorker = 20001
	//Update Worker
	UpdateWorker = 20010
	StopMonitorSignal = 30001
	StartMonitorSignal = 30002
	StartMonitorVideo = 300011
	StopMonitorVideo = 30012
	StartMonitorAudio = 30021
	StopMonitorAudio = 30022
)

// Monitor type declare
const (
	Signal = 70001
	Video = 70002
	Audio = 70003
	BlackScreen = 70004
	RunThread = 7005
)

//Executer type
const (
	UrgentTask = 1100
	CommandShell = 1101
)

