package conf

type Config struct {
	ConsulAddr           string `json:"consul_addr"`
	ConsulHealth         bool   `json:"consul_health"`
	ConsulHealthInterval string `json:"consul_health_interval"`
	ConsulHealthTimeout  string `json:"consul_health_timeout"`
	ConsulHealthPort     string `json:"consul_health_port"`
	LogLevel             string `json:"log_level"`
	LogFile              string `json:"log_file"`
	HttpAddr             string `json:"http_addr"`
	ServiceName          string `json:"service_name"`
	HttpReadOverTime     int    `json:"http_read_over_time"`
	HttpWriteOverTime    int    `json:"http_write_over_time"`
	IsProduction         bool   `json:"is_production"`
}
