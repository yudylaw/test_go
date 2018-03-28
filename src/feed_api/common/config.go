package common

type Config struct {
	Server struct {
		Bind string `toml:"bind"`
	}
	Log struct {
		Level       string `toml:"level"`
		AccessLog   string `toml:"access_log"`
		BusinessLog string `toml:"business_log"`
		ErrorLog    string `toml:"error_log"`
		StatLog     string `toml:"stat_log"`
	}
}
