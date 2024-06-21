package config

import "time"

var Config configStruct

// 应用信息
type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
}

// JWT 配置
type jwt struct {
	Secret string        `json:"secret"`
	Issuer string        `json:"issuer"`
	Expire time.Duration `json:"expire"`
}

// 日志信息
type log struct {
	Path       string     `yaml:"path"`
	Level      string     `yaml:"level"`
	FilePrefix string     `yaml:"filePrefix"`
	FileFormat string     `yaml:"fileFormat"`
	OutFormat  string     `yaml:"outFormat"`
	LumberJack lumberJack `yaml:"lumberJack"`
}

// 日志切割
type lumberJack struct {
	MaxSize    int  `yaml:"maxSize"`    //单文件最大容量(单位MB)
	MaxBackups int  `yaml:"maxBackups"` // 保留旧文件的最大数量
	MaxAge     int  `yaml:"maxAge"`     // 旧文件最多保存几天
	Compress   bool `yaml:"compress"`   // 是否压缩/归档旧文件
}

// MySQL配置
type mysql struct {
	Host                      string `json:"host"` // Host
	Port                      int    `json:"port"` // 端口
	User                      string `json:"user"`
	Password                  string `json:"password"`
	Database                  string `json:"database"`
	TablePrefix               string `json:"tablePrefix"`
	Charset                   string `json:"charset"`                   // 要支持完整的UTF-8编码,需设置成: utf8mb4
	ParseTime                 bool   `json:"parseTime"`                 // 解析time.Time类型
	TimeZone                  string `json:"timeZone"`                  // 时区,若设置 Asia/Shanghai,需写成: Asia%2fShanghai,默认Local
	DefaultStringSize         int    `json:"defaultStringSize"`         // string 类型字段的默认长度
	DisableDatetimePrecision  bool   `json:"disableDatetimePrecision"`  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	SkipInitializeWithVersion bool   `json:"skipInitializeWithVersion"` // 根据当前 MySQL 版本自动配置
	AutoMigrate               bool   `json:"autoMigrate"`               // 开启时，每次服务启动都会根据实体创建/更新表结构
	SlowSql                   string `json:"slowSql"`                   // 慢sql时间。单位毫秒
	LogLevel                  string `json:"logLevel"`                  // error、info、warn
	IgnoreRecordNotFoundError bool   `json:"ignoreRecordNotFoundError"` // 是否忽略ErrRecordNotFound(未查到记录错误)
}

// Redis 配置
type redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type configStruct struct {
	App   app   `yaml:"app"`
	Log   log   `yaml:"log"`
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
	Jwt   jwt   `yaml:"jwt"`
}
