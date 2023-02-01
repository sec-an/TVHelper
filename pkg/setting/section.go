package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type LogSettingS struct {
	Path       string
	FilePrefix string
	Encoder    string
	Output     string
	LumberJack LumberJack
}

type LumberJack struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type SpiderSettingS struct {
	DouBanClientTimeout time.Duration
	AListClientTimeout  time.Duration
	ParserClientTimeout time.Duration
}

type MysqlSettingS struct {
	Running  bool
	Host     string
	Port     uint16
	Username string
	Password string
	Database string
}

type RedisSettingS struct {
	Running      bool
	Host         string
	Port         uint16
	Auth         string
	Database     int
	IdleTimeout  time.Duration
	PoolSize     int
	SubCacheTime time.Duration
}

type AListSettingS struct {
	Host  string
	Token string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
