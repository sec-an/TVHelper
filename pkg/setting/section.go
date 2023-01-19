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
	ParserClientTimeout time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
