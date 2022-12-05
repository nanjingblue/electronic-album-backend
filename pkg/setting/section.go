package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type DatabaseSettings struct {
	DBType        string
	Username      string
	Password      string
	Host          string
	DBName        string
	TablePrefix   string
	Charset       string
	ParseTime     bool
	MaxIdleConns  int
	MaxOpenConns  int
	SqliteDB      string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

type JwtSettingS struct {
	Key string
}

type OSSSettingS struct {
	END_POINT         string
	ACCESS_KEY_ID     string
	ACCESS_KEY_SECRET string
	BUCKET            string
	DOMAIN string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return err
}
