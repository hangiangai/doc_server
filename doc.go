package doc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// 服务器配置
type Server struct {
	Addr      string // 地址
	Port      string // 端口
	ApiServer string // 测试接口访问地址
}

// 存储每个文件的注释内容
type Doc struct {
	Name  string
	Count int
	Param []map[string][]string
}

// 响应值
type Response struct {
	Docs   []Doc
	IpAddr string
	Port   string
}

var (
	docs              []Doc                      //文档数组
	server            Server                     //服务信息
	defaultConfigPath = "doc/public/config.json" //默认配置文件
	defaultStaticPath = "doc/public"             //默认静态目录
	ipAddr            = getIpv4Addr()            //部署主机局域网地址
	configInfo         map[string]interface{}
)

func checkError(err error, code int) {
	switch code {

	}
	if err != nil {
		log.Printf("Error: %s", err)
	}
}

// 获取本机ip地址
func getIpv4Addr() string {
	var ipv4 string
	netInterfaces, err := net.Interfaces()
	checkError(err, 200)
	for _, v := range netInterfaces {
		// 判断ip地址有效
		if v.Flags&net.FlagUp != 0 {
			addrs, _ := v.Addrs()
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ipv4 = ipnet.IP.String()
					}
				}
			}
		}
	}
	return ipv4
}

// 静态文件服务
func startStaticFileServer(_cp string) {
	fs := http.FileServer(http.Dir(_cp))
	http.Handle("/doc/", http.StripPrefix("/doc/", fs))
}

// 开启服务
func startServer(s Server) {
	checkError(http.ListenAndServe(s.Addr+":"+s.Port, nil), 200)
}

//注册路由
func registerRouter() {
	http.HandleFunc("/doc/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		bytes, err := json.Marshal(Response{
			Docs:   docs,
			IpAddr: ipAddr,
			Port:   server.Port,
		})
		checkError(err, 200)
		_, _ = w.Write(bytes)
	})
}

// 处理
func toDoc() {

}

func getMapVal(val map[string]interface{}, key string) string {
	var result string
	if v, ok := val[key]; ok {
		result = v.(string)
	}
	return result
}

// 更新配置文件
func updateConfig(config []byte) {
	checkError(ioutil.WriteFile("doc/public/config.json", config, 0644), 200)
}

//读取配置文件
func ReadConfig(_f string) {
	cfg, err := ioutil.ReadFile(_f)
	checkError(err, 200)
	config := make(map[string]interface{})
	configInfo = config
	checkError(json.Unmarshal(cfg, &config), 200)
	// 读取配置文件配置
	server = Server{
		Addr:      getMapVal(config, "addr"),
		Port:      getMapVal(config, "port"),
		ApiServer: getMapVal(config, "apiServer"),
	}
	// 读取需要处理的文件
	if files, ok := config["files"]; ok {
		if v, ok := files.(map[string]interface{}); ok {
			for k, file := range v { // 读取并解析文件
				doc_ := readFile(file.(string))
				docs = append(docs, Doc{
					Name:  k,
					Param: doc_,
					Count: len(doc_),
				})
			}
		}
	}
}

func InitConfig(s Server) {
	if s.Port != "" {
		server.Port = s.Port
		configInfo["port"] = server.Port
	}
	if s.Addr != "" {
		server.Addr = ipAddr
		configInfo["addr"] = ipAddr
	}
	if s.ApiServer != "" {
		server.ApiServer = s.ApiServer
		configInfo["apiServer"] = server.ApiServer
	}
	data, _ := json.Marshal(configInfo)
	updateConfig(data)
}

func InitAndRun(s Server) {
	// 读取配置
	ReadConfig(defaultConfigPath)
	// 初始化配置
	InitConfig(s)
	// 注册和更新数据
	registerRouter()
	// 开启静态文件服务
	startStaticFileServer(defaultStaticPath)
	// 开启服务
	startServer(s)
}
