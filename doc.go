package doc

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Response struct {
	Docs      []Doc
	IpAddr    string
	Port      string
	ApiServer string
}

var (
	defaultStaticDir = "doc/public" //默认静态目录
	updateQueue      = NewQueue(10) //保存更新的信息
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

func saveContent() {

}


func startServer(addr string, port string) {
	log.Printf("server: %s:%s", addr, port)
	log.Printf("website: %s:%s/doc", addr, port)
	if err := http.ListenAndServe(addr+":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type ClientMsg struct {
	data     Queue
	create   time.Time
	identity string
	status   int
}

var (
	conns       = make([]*websocket.Conn, 0)
	updatedData = NewQueue(10)
)

//注册路由
func registerRouter(docs []Doc, port string, apiServer string) {
	http.HandleFunc("/doc/v2/updated", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		bytes, err := json.Marshal(map[string]interface{}{
			"updated": updatedData.data,
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

/*
@param: type:int:传输类型: 1-注册 2-chat 3-doc
@param: data:string:传输信息
@param: state:int:传输动作
@param: identity:string:身份信息
*/
type TransmissionData struct {
	Type     int
	Data     string
	Status   int
	Identity string
}

var (
	clients     = make(map[string]*websocket.Conn, 0)
	chatStorage = make([]string, 0)
)

func DataDistribution(t TransmissionData) {
	t.Status = 2 // 2-服务端发送 1-客户端发送
	b, _ := json.Marshal(t)
	for _, v := range clients {
		_, _ = v.Write(b)
	}
}

func Websocket() {
	http.Handle("/updated", websocket.Handler(func(conn *websocket.Conn) {
		var reply []byte
		for {
			// 接收数据
			if err := websocket.Message.Receive(conn, &reply); err == nil {
				var info TransmissionData
				_ = json.Unmarshal(reply, &info)
				switch info.Type {
				case 1: // 注册
					clients[info.Identity] = conn
				case 2: // 聊天信息
					chatStorage = append(chatStorage, info.Data)
					DataDistribution(info)
				case 3: // 文档相关通知
					log.Println("文档请求")
				}
			}
		}
	}))
}

// 数据交换
func exchange() {

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

	// 更新信息传输通道
	var p = make(chan bool)
	// 开启websocket
	Websocket()

	// 当有文件修改 进行数据分发
	go func(updatedNotify *chan bool) {
		for range *updatedNotify {
			toJson, _ := json.Marshal(TransmissionData{
				Type: 3,
				Status: 1,
			})
			for k, c := range clients {
				n, err := c.Write(toJson)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("%s:%d\n", k, n);
			}
		}
	}(&p)

	// 文件监听
	fWatcher := newFWatcher(c.Files)
	fWatcher.run(&p)

	// 处理文档
	docs = make([]Doc, len(c.Files))
	for k, v := range c.Files {
		docs[k] = toDoc(v)
	}

	registerRouter(docs, c.Port, c.ApiServer)
	staticFileServer(defaultStaticDir)
	startServer(c.Addr, c.Port)
}
