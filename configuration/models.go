package configuration

type Agent struct {
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
	}
	Server struct {
		Port    uint16 `yaml:"port"`
		Host    string `yaml:"host"`
		Gateway string `yaml:"gateway"`
	}

	RPC struct {
		Agent          Agent
		Agentctl       Agentctl
		Agentrunner    Agentexecuter
	}
	Warmup struct {
		Concurrency       int `yaml:"concurrency"`
		MigrationInterval int `yaml:"migration_interval"`
	}

	Nra struct {
		Concurrency int `yaml:"concurrency"`
	}

	Redis struct {
		Host         string `yaml:"host"`
		Port         uint16    `yaml:"port"`
		DB           int    `yaml:"db"`
		Password     string `yaml:"password"`
		Key          string `yaml:"key"`
		DialTimeout  int32  `yaml:"dial_timeout"`
		ReadTimeout  int32  `yaml:"read_timeout"`
		WriteTimeout int32  `yaml:"write_timeout"`
		PoolSize     int    `yaml:"pool_size"`
		PoolTimeout  int32  `yaml:"pool_timeout"`
	}
	MQ struct {
		Host                     string `yaml:"host"`
		Port                     uint16    `yaml:"port"`
		Topic                    string `yaml:"topic"`
		WarmUpTopic              string `yaml:"warmup_topic"`
		NraTopic                 string `yaml:"nra_topic"`
		OperationTopic           string `yaml:"operation_topic"`
		CommandTopic             string `yaml:"command_topic"`
		Username                 string `yaml:"Username"`
		Password                 string `yaml:"Password"`
	}

	AgentGateway struct{
		Host			string `yaml:"host"`
		Port            uint16 `yaml:"port"`
		Gateway			string `yaml:"gateway"`
		Storage			string `yaml:"storage"`
	}

	Agentd struct {
		Master			AgentMaster	`yaml:"master"`
		Worker			AgentWorker `yaml:"worker"`
	}

}
