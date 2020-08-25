package mailer

type Config struct {
	Server   string `config:"server"`
	Hostname string `config:"hostname"`
	From     string `config:"from"`
}
