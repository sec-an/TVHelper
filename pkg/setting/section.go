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

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
