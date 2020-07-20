package doc

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

type Response struct {
	Docs      []Doc
	IpAddr    string
	Port      string
	ApiServer string
}

type FileInfo struct {
	Sum string
	Len int
}

var (
	defaultStaticDir = "doc/public" //默认静态目录
	updateQueue      = NewQueue(10) //保存更新的信息
	docInfo          = make(map[string]FileInfo)
	docs             []Doc
)

func getIpv4Addr() string {
	var ipv4 string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
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

func staticFileServer(dir string) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/doc/", http.StripPrefix("/doc/", fs))
}

//注册路由
func registerRouter(docs []Doc, port string, apiServer string) {
	http.HandleFunc("/doc/v2/updated", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		bytes, err := json.Marshal(map[string]interface{}{
			"updated": updateQueue.All(),
		})
		if err != nil {
			log.Fatal(err)
		}
		_, _ = w.Write(bytes)
	})

	http.HandleFunc("/doc/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		bytes, err := json.Marshal(Response{
			Docs:      docs,
			IpAddr:    getIpv4Addr(),
			Port:      port,
			ApiServer: apiServer,
		})
		if err != nil {
			log.Fatal(err)
		}
		_, _ = w.Write(bytes)
	})
}

func startServer(addr string, port string) {
	log.Printf("server: %s:%s", addr, port)
	log.Printf("website: %s:%s/doc", addr, port)
	if err := http.ListenAndServe(addr+":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func key(filepath string) string {
	filenameDecode := url.QueryEscape(filepath)
	randNumber := strconv.FormatInt(rand.Int63(), 10)
	key := fmt.Sprintf("%x", md5.Sum([]byte(randNumber+filenameDecode)))
	return key
}

func InitAndRun(apiServer string) {
	// 初始化配置
	c := newConfig()
	c.initConfig()
	// 文件监听
	fWatcher := newFWatcher(c.Files)
	fWatcher.run(&updateQueue)
	// 处理文档
	docs = make([]Doc, len(c.Files))
	for k, v := range c.Files {
		docs[k] = toDoc(v)
	}
	// 注册路由
	registerRouter(docs, c.Port, c.ApiServer)
	// 开启文件服务
	staticFileServer(defaultStaticDir)
	// 开启服务
	startServer(c.Addr, c.Port)
}
