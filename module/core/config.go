package core

// Configuration ==> Configuration which save the global key value
type Configuration struct {
	props map[string]string
}

func (configuration *Configuration) get(k string) string {
	return configuration.get(k)
}

// Config ==>  global configuration variable
var Config = &Configuration{}

func init() {
	loadProperties(Config)
	loadYMAL(Config)
}

func loadProperties(Config *Configuration) {}

func loadYMAL(Config *Configuration) {}
