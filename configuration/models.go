package configuration

type Agent struct {
	Host    string `yaml:"host"`
	Port    uint16 `yaml:"port"`
	Gateway string `yaml:"gateway"`
}

type Logging struct {
	Host    string `yaml:"host"`
	Port    uint16 `yaml:"port"`
	Gateway string `yaml:"gateway"`
}
type AgentMonitor struct {
	Host    string `yaml:"host"`
	Port    uint16 `yaml:"port"`
	Gateway string `yaml:"gateway"`
}

type AgentMaster struct {
	Host             string `yaml:"host"`
	Port             uint16 `yaml:"port"`
	WorkingDir       string `yaml:"WorkingDir"`
	IsDebugEnable    bool   `yaml:"IsDebugEnable"`
	PathFileLog      string `yaml:"PathFileLog"`
	RetryConnectTime uint16 `yaml:"RetryConnectTime"`
}

type AgentWorker struct{
	MaxLifeTime		int `yaml:"maxLifeTime"`
	MinThread		int `yaml:"minThread"`
	MaxThread		int `yaml:"maxThread"`
}

type Agentctl struct {
	Host    string `yaml:"host"`
	Port    uint16 `yaml:"port"`
	Gateway string `yaml:"gateway"`
}

type Agentexecuter struct {
	Host    string `yaml:"host"`
	Port    uint16 `yaml:"port"`
	Gateway string `yaml:"gateway"`
}




type Conf struct {
	ConfigureFile string
	DB            struct {
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}	`yaml:"db"`
	Telegram	struct{
		APIBaseURL		string	`yaml:"api_base_url"`
		APISendMessage	string	`yaml:"api_send_message"`
		Bot				struct{
			TokenEnv	string	`yaml:"token_env"`
			ChatID		string	`yaml:"chat_id"`
		}	`yaml:"bot"`
	}	`yaml:"telegram"`
	Server struct {
		Port    uint16 `yaml:"port"`
		Host    string `yaml:"host"`
		Gateway string `yaml:"gateway"`
		RequestTimeout	int	`yaml:"request_timeout"`
	}	`yaml:"server"`

	RPC struct {
		Agent          Agent	`yaml:"agent"`
		Agentctl       Agentctl	`yaml:"agentctl"`
		Agentrunner    Agentexecuter	`yaml:"agentexe"`
		AgentMonitor	AgentMonitor `yaml:"agent_monitor"`
		Logging			Logging		`yaml:"logging"`
	}	`yaml:"rpc"`
	Warmup struct {
		Concurrency       int `yaml:"concurrency"`
		MigrationInterval int `yaml:"migration_interval"`
	}

	Nra struct {
		Concurrency int `yaml:"concurrency"`
	}	`yaml:"nra"`

	Redis struct {
		Host         string `yaml:"host"`
		Port         uint16	`yaml:"port"`
		DB           int    `yaml:"db"`
		Password     string `yaml:"password"`
		Key          string `yaml:"key"`
		DialTimeout  int32  `yaml:"dial_timeout"`
		ReadTimeout  int32  `yaml:"read_timeout"`
		WriteTimeout int32  `yaml:"write_timeout"`
		PoolSize     int    `yaml:"pool_size"`
		PoolTimeout  int32  `yaml:"pool_timeout"`
	}	`yaml:"redis"`
	MQ struct {
		Host                     string `yaml:"host"`
		Port                     uint16	`yaml:"port"`
		Topic                    string `yaml:"topic"`
		WarmUpTopic              string `yaml:"warmup_topic"`
		NraTopic                 string `yaml:"nra_topic"`
		OperationTopic           string `yaml:"operation_topic"`
		CommandTopic             string `yaml:"command_topic"`
		MonitorLogsTopic			string `yaml:"monitor_logs_topic"`
		Username                 string `yaml:"Username"`
		Password                 string `yaml:"Password"`
	}	`yaml:"mq"`

	AgentGateway struct{
		Host			string `yaml:"host"`
		Port            uint16 `yaml:"port"`
		Gateway			string `yaml:"gateway"`
		Storage			string `yaml:"storage"`
	} `yaml:"agent_gateway"`

	Agentd struct {
		Master			AgentMaster	`yaml:"master"`
		Worker			AgentWorker `yaml:"worker"`
	} `yaml:"agentd"`

}
