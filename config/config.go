package config

import (
	"github.com/go-ini/ini"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var AppSetting *App

type App struct {
	JwtSecret             string `yaml:"jwt_secret"`
	JwtRefreshSecret      string `yaml:"jjwt_refresh_secret"`
	RuntimeRootPath       string `yaml:"runtime_root_path"`
	CasbinModelPath       string `yaml:"casbin_model_path"`
	LogSavePath           string `yaml:"log_save_path"`
	LogSaveName           string `yaml:"log_save_name"`
	LogFileExt            string `yaml:"log_file_ext"`
	FileCachePath         string `yaml:"file_cache_path"`
	SuperUsername         string `yaml:"super_username"`
	DingWebHook           string `yaml:"ding_web_hook"`
	DingSecret            string `yaml:"ding_secret"`
	MailHost              string `yaml:"mail_host"`
	MailAddress           string `yaml:"mail_address"`
	MailPassword          string `yaml:"mail_password"`
	RuleEngineDingWebHook string `yaml:"rule_engine_ding_web_hook"`
}

var MysqlSetting *MySqlConfig

type MySqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

var MongoDBSetting *MongoDBConfig

type MongoDBConfig struct {
	Uri      string `yaml:"uri"`
	Database string `yaml:"database"`
}

var RedisSetting *RedisConfig

type RedisConfig struct {
	Addr              string `yaml:"addr"`
	Password          string `yaml:"password"`
	DBIndex           int    `yaml:"db_index"`
	MaxIdle           int    `yaml:"max_idle"`
	MaxActive         int    `yaml:"max_active"`
	IdleTimeoutSecond int    `yaml:"idle_timeout"`
}

var ServerSetting *Server

type Server struct {
	RunMode            string `yaml:"run_mode"`
	HttpPort           int    `yaml:"http_port"`
	ReadTimeoutSecond  int64  `yaml:"read_timeout"`
	WriteTimeoutSecond int64  `yaml:"write_timeout"`
}

var cfg *ini.File

var GlobalConf = &Config{}

type Config struct {
	App     App           `yaml:"app"`
	Server  Server        `yaml:"server"`
	Mysql   MySqlConfig   `yaml:"mysql"`
	MongoDB MongoDBConfig `yaml:"mongodb"`
	Redis   RedisConfig   `yaml:"redis"`
}

func MustSetupYaml() {
	cfgPath := "config/app.yaml"
	if f, err := os.Open(cfgPath); err != nil {
		panic(err)
	} else {
		if yaml.NewDecoder(f).Decode(GlobalConf) != nil {
			panic(err)
		}
	}
	if GlobalConf.App.JwtSecret == "" {
		panic("配置文件初始化失败")
	}

	AppSetting = &GlobalConf.App
	ServerSetting = &GlobalConf.Server
	MysqlSetting = &GlobalConf.Mysql
	MongoDBSetting = &GlobalConf.MongoDB
	RedisSetting = &GlobalConf.Redis
}

func SetupIni() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("mysql", MysqlSetting)
	mapTo("mongodb", MongoDBSetting)
	mapTo("redis", RedisSetting)

	log.Println(MysqlSetting.Database)
	GlobalConf.App = *AppSetting
	GlobalConf.Server = *ServerSetting
	GlobalConf.Mysql = *MysqlSetting
	GlobalConf.MongoDB = *MongoDBSetting
	GlobalConf.Redis = *RedisSetting
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
