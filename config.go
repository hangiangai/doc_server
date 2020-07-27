package doc

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Addr       string
	Port       string
	Files      []string
	ApiServer  string
	configPath string
	staticDir  string
}

func newConfig() *Config {
	return &Config{
		Port:       "8888",
		configPath: "doc/public/config.json",
		staticDir:  "doc/public",
	}
}

func (c *Config) initConfig() {
	cfg, err := ioutil.ReadFile(c.configPath)
	checkError(err, true)
	config := make(map[string]interface{})
	checkError(json.Unmarshal(cfg, &config), true)
	c.Addr = config["addr"].(string)
	c.Port = config["port"].(string)
	c.ApiServer = config["apiServer"].(string)
	if files, ok := config["files"]; ok {
		if fp, ok := files.([]interface{}); ok {
			for _, v := range fp {
				c.Files = append(c.Files, v.(string))
			}
		}
	}
	c.Files = append(c.Files, c.configPath)
}
