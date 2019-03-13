package config

type config struct {
	DB db
}

func newConfig() *config {
	return &config{
		DB: newDB(),
	}
}

var shared *config

func init() {
	shared = newConfig()
}

func Get() *config {
	return shared
}
