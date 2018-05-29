package config

type LogConfig struct {
	logPath  string
	logLevel string
}

func (dc LogConfig) LogPath() string {
	return dc.logPath
}

func (dc LogConfig) LogLevel() string {
	return dc.logLevel
}
