package apiserver

type Config struct {
	Host string `config:"host"`
	Port int    `config:"port"`
}
