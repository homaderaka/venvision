package config

type ListenConfig struct {
	SocketFile string `env:"SOCKET_FILE" env-default:"app.sock"`
	Type       string `env:"LISTEN_TYPE" env-default:"port"`
	BindIP     string `env:"BIND_IP" env-default:"0.0.0.0"`
	Port       string `env:"PORT" env-default:"8000"`
}
