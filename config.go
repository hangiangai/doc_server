package doc

import (
	"encoding/json"
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

type config struct {
	Addr        string
	Port        string
	Files       []string
	ApiServ     string
	MatchKeys   []map[string]string
	defaultPath string
	defaultPort string
	defaultLogs string
}

func NewConfig() *config {
	c := &config{
		defaultPath: "doc/public/config.json",
		defaultLogs: "",
	}
	c.readConfigFile();
	c.collectMatchKey();
	return c
}

func (c *config) readConfigFile() {
	cfg, err := ioutil.ReadFile(c.defaultPath)
	checkError(err, true)
	if err := json.Unmarshal(cfg, c); err != nil {
		panic(err)
	}
}

func (c *config) collectMatchKey() map[string][]string {
	for _, key := range c.MatchKeys {
		for k, v := range key {
			baseKeys[k] = append(baseKeys[k], v[1:])
		}
	}
	return baseKeys
}

func (c *config) files() []string {
	return c.Files
}

func (c *config) handlerConfigUpdated() map[string]int {
	toMap := make(map[string]int)
	return toMap
}



func (c *config) WriteMatchFunc() {
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
	if err := ioutil.WriteFile("doc/match.go", []byte(matchStr), 0777); err != nil {
		log.Println(err)
	}
}
