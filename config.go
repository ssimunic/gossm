package gossm

import (
	"encoding/json"
)

type Config struct {
	Servers  Servers  `json:"servers"`
	Settings Settings `json:"settings"`
}

// New returns pointer to Config which is created from provided JSON data.
// Guarantees to be validated.
func NewConfig(jsonData []byte) *Config {
	config := &Config{}
	err := json.Unmarshal(jsonData, config)
	if err != nil {
		panic("error parsing json configuration data")
	}
	config.validate()
	return config
}
