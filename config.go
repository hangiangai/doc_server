package api_doc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	baseKeys = map[string][]string{
		"@":       {"@"},
		"@title":  {"title"},
		"@url":    {"url"},
		"@header": {"header"},
		"@method": {"method"},
		"@param":  {"param"},
		"@return": {"return"},
		"@hint":   {"hint"},
	}
)

type Config struct {
	Addr        string
	Port        string
	Files       []string
	ApiServ     string
	MatchKeys   []map[string]string
	defaultPath string
	defaultPort string
	defaultLogs string
}

func NewConfig(path string) *Config {
	c := &Config{
		defaultPath: "config.json",
		defaultLogs: "",
	}
	if path != "" {
		c.defaultPath = path
	}
	c.readConfigFile()
	c.collectMatchKey()
	c.WriteMatchFunc()
	return c
}

func (c *Config) readConfigFile() {
	cfg, err := ioutil.ReadFile(c.defaultPath)
	checkError(err, true)
	if err := json.Unmarshal(cfg, c); err != nil {
		panic(err)
	}

	fmt.Println(c)
}

func (c *Config) collectMatchKey() map[string][]string {
	for _, key := range c.MatchKeys {
		for k, v := range key {
			baseKeys[k] = append(baseKeys[k], v[1:])
		}
	}
	return baseKeys
}

func (c *Config) files() []string {
	return c.Files
}

func (c *Config) handlerConfigUpdated() map[string]int {
	toMap := make(map[string]int)
	return toMap
}

func (c *Config) WriteMatchFunc() {
	// 写入匹配函数
	var matchStr string
	matchStr += "\npackage doc"
	matchStr += "\n\nfunc match(note *note, key string, value []string){"
	matchStr += "\n\tswitch key {"
	matchStr += "\n\tcase \"" + strings.Join(baseKeys["@title"], "\",\"") + "\":"
	matchStr += "\n\t\tnote.Title = value[0]"
	matchStr += "\n\tcase \"" + strings.Join(baseKeys["@url"], "\",\"") + "\":"
	matchStr += "\n\t\tnote.Url = value[0]"
	matchStr += "\n\tcase \"" + strings.Join(baseKeys["@header"], "\",\"") + "\":"
	matchStr += "\n\t\tnote.Header = value[0]"
	matchStr += "\n\tcase \"" + strings.Join(baseKeys["@method"], "\",\"") + "\":"
	matchStr += "\n\t\tnote.Method = value[0]"
	matchStr += "\n\tcase \"" + strings.Join(baseKeys["@param"], "\",\"") + "\":"
	matchStr += "\n\t\tnote.Params = append(note.Params, value)"
	matchStr += "\n\t} \n}"
	if err := ioutil.WriteFile("match.go", []byte(matchStr), 0777); err != nil {
		log.Println(err)
	}
}
