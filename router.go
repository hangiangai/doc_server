package api_doc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"time"
)

// 存放路由
type methodTree struct {
	method string
	routes map[string]HandlersChain
}

type methodTrees []methodTree

func (m *methodTrees) getValue(method, relativePath string) HandlersChain {
	for _, v := range *m {
		if v.method == method {
			if hs, ok := v.routes[relativePath]; ok {
				return hs
			}
		}
	}
	return nil
}

func (m *methodTrees) addRoute(method, relativePath string, handlers HandlersChain) {
	var isExist bool
	for _, v := range *m {
		if v.method == method {
			v.routes[relativePath] = handlers
			isExist = true
			break
		}
	}
	if !isExist {
		*m = append(*m, methodTree{
			method: method,
			routes: map[string]HandlersChain{
				relativePath: handlers,
			},
		})
	}
}

type Context struct {
	Request *http.Request
	Write   http.ResponseWriter
}

func (ctx *Context) FormValue(k string) string {
	return ctx.Request.FormValue(k)
}

func (ctx *Context) SetHeader(k, v string) {
	ctx.Write.Header().Set(k, v)
}

func (ctx *Context) Json(v interface{}) (int, error) {
	wCon, err := json.Marshal(v)
	if err != nil {
		return 0, err
	}
	n, err := ctx.Write.Write(wCon)
	return n, err
}

type HttpHandlerFunc func(ctx *Context)
type HandlersChain []HttpHandlerFunc

type IRoutes interface {
	Use(...HttpHandlerFunc) IRoutes
	Handler(string, string, ...HttpHandlerFunc) IRoutes
	GET(string, ...HttpHandlerFunc) IRoutes
	POST(string, ...HttpHandlerFunc) IRoutes
	DELETE(string, ...HttpHandlerFunc) IRoutes
	PUT(string, ...HttpHandlerFunc) IRoutes
	OPTION(string, ...HttpHandlerFunc) IRoutes
}

type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
}

func (group *RouterGroup) BasePath() string {
	return group.basePath
}

func (group *RouterGroup) Group(relativePath string, handlers ...HttpHandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
		engine:   group.engine,
	}
}

func (group *RouterGroup) Handler(httpMethod, relativePath string, handlers ...HttpHandlerFunc) IRoutes {
	if matches, err := regexp.MatchString("^[A-Z]+$", httpMethod); !matches || err != nil {
		panic("http method " + httpMethod + " is not valid")
	}
	return group.handler(httpMethod, relativePath, handlers)
}

func (group *RouterGroup) GET(relativePath string, handlers ...HttpHandlerFunc) IRoutes {
	return group.handler(http.MethodGet, relativePath, handlers)
}

func (group *RouterGroup) POST(relativePath string, handlers ...HttpHandlerFunc) IRoutes {
	return group.handler(http.MethodPost, relativePath, handlers)
}

func (group *RouterGroup) DELETE(relativePath string, handlers ...HttpHandlerFunc) IRoutes {
	return group.handler(http.MethodDelete, relativePath, handlers)
}

func (group *RouterGroup) PUT(relativePath string, handlers ...HttpHandlerFunc) IRoutes {
	return group.handler(http.MethodPut, relativePath, handlers)
}

func (group *RouterGroup) OPTION(relativePath string, handlers ...HttpHandlerFunc) IRoutes {
	return group.handler(http.MethodOptions, relativePath, handlers)
}

func (group *RouterGroup) Use(middleware ...HttpHandlerFunc) IRoutes {
	group.Handlers = append(group.Handlers, middleware...)
	return group
}

func (group *RouterGroup) handler(HttpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.engine.addRoute(HttpMethod, absolutePath, handlers)
	return group
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}

func (group *RouterGroup) endChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	if relativePath == "" {
		return group.basePath
	}
	finalPath := path.Join(group.basePath, relativePath)
	appendSlash := group.endChar(relativePath) == '/' && group.endChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}

var DefaultWriter io.Writer = os.Stdout

func debugPrint(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	_, _ = fmt.Fprintf(DefaultWriter, "[debug] "+format, values...)
}

func defaultLogFormatter(time string, method string, path string) string {
	return fmt.Sprintf("%c[0;0;32m%s %c[0;0;32m[%s] %s %c[0m\n",
		0x1B, time, 0x1B, method, path, 0x1B)
}

func logger() HttpHandlerFunc {
	return func(ctx *Context) {
		nowTime := time.Now().Format("2006/01/02-15:04:05")
		method := ctx.Request.Method
		p := ctx.Request.URL.Path
		_, _ = fmt.Fprint(DefaultWriter, defaultLogFormatter(nowTime, method, p))
	}
}

/*

 */
type Engine struct {
	RouterGroup
	methodTrees methodTrees
	pool        sync.Pool
}

func NewEngine() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: make(HandlersChain, 0),
			basePath: "/",
		},
		methodTrees: make(methodTrees, 0, 5),
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return &Context{}
	}
	return engine
}

func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
	engine.methodTrees.addRoute(method, path, handlers) // 添加路由
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := engine.pool.Get().(*Context)
	c.Request = request
	c.Write = writer
	_ = c.Request.ParseForm()
	engine.handleHttpRequest(c)
}

func (engine *Engine) Run(addr string) (err error) {
	debugPrint("Listening and serving HTTP on %s\n", addr)
	err = http.ListenAndServe(addr, engine)
	return
}

func (engine *Engine) handleHttpRequest(ctx *Context) {
	httpMethod := ctx.Request.Method
	rPath := ctx.Request.URL.Path
	handlers := engine.methodTrees.getValue(httpMethod, rPath)
	if handlers != nil {
		for _, h := range handlers {
			h(ctx)
		}
	}
}

type Value struct {
	data   interface{}
	key    string
	expire time.Time
}

type Store struct {
	values map[string]Value
	sync   sync.Locker
	add    chan int
}

func NewStore() *Store {
	store := &Store{
		values: make(map[string]Value, 0),
		add:    make(chan int),
	}
	return store
}

func insertionSort(arr []int) {
	var arrLen = len(arr)
	for i := 1; i < arrLen; i++ {
		curr := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > curr {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = curr
	}
}

func (s *Store) Run() {
	t := time.NewTicker(time.Second)
	go func(s *Store) {
		for {
			select {
			case <-t.C:
				for k, v := range s.values {
					if v.expire.Before(time.Now()) {
						s.Delete(k)
					}
				}
			case <-s.add:
				insertionSort([]int{})
			}
		}
	}(s)
}

func (s *Store) Set(k string, v interface{}, expire time.Duration) {
	s.values[k] = Value{
		data:   v,
		key:    k,
		expire: time.Now().Add(expire),
	}
}

func (s *Store) Delete(k string) {
	if _, ok := s.values[k]; ok {
		delete(s.values, k)
	}
}

func (s *Store) GetValue(k string) interface{} {
	return s.values[k].data
}

func (s *Store) GetString(k string) string {
	return s.values[k].data.(string)
}

func (s *Store) GetFloat64(k string) float64 {
	return s.values[k].data.(float64)
}

func (s *Store) GetFloat32(k string) float32 {
	return s.values[k].data.(float32)
}

func (s *Store) GetInt(k string) int {
	return s.values[k].data.(int)
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second    // 60 读取客户端响应
	pingPeriod     = (pongWait * 9) / 10 // 向客户端写入心跳包
	maxMessageSize = 512
)

type exchangeBody struct {
	Type     int
	Identity string
	Data     interface{}
}

// 管理客户
type Hub struct {
	clients    map[*client]bool
	broadcast  chan exchangeBody
	register   chan *client
	unregister chan *client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan exchangeBody),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

type client struct {
	conn *websocket.Conn
	send chan exchangeBody
	hub  *Hub
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *client) readPump() {
	defer func() {
		c.hub.unregister <- c
		_ = c.conn.Close()
	}()
	// 设置最大读取字节个数
	c.conn.SetReadLimit(maxMessageSize)
	// 设置最大读取时间
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// 心跳包处理函数
	c.conn.SetPongHandler(func(appData string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		// 会形成阻
		_, message, err := c.conn.ReadMessage()
		// 客户端失去链接
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Panicf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		var body = exchangeBody{}
		_ = json.Unmarshal(message, &body)
		fmt.Println(body)
	}
}

func (c *client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err)
				return
			}
			sb, _ := json.Marshal(message)
			_, _ = w.Write(sb)

			n := len(c.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(newline)
				//z_, _ = w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			//写入心跳包
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func initServer(port string, apiServer string) {

	en := NewEngine()

	en.Use(func(ctx *Context) {
		//允许访问所有域
		ctx.SetHeader("Access-Control-Allow-Origin", "*")
		//header的类型
		ctx.SetHeader("Access-Control-Allow-Headers", "Content-Type")
	})

	en.Use(logger())

	en.GET("/updated", func(ctx *Context) {
		conn, err := upgrade.Upgrade(ctx.Write, ctx.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &client{
			hub:  hub,
			conn: conn,
			send: make(chan exchangeBody, 256),
		}
		client.hub.register <- client
		go client.readPump()
		go client.writePump()

	})
	// 文档更新路由
	en.GET("/doc/v2/updated", func(ctx *Context) {
		_, _ = ctx.Json(map[string]interface{}{
			"updated": store.GetValue("updated"),
		})
	})
	// 文档添加路由
	en.GET("/doc/v2/added", func(ctx *Context) {
		//identity := ctx.FormValue("identity")
		//
		//_, _ = ctx.Json(map[string]interface{}{
		//	"added": c.Storage(),
		//})
	})
	// 文档初始信息
	en.GET("/doc/v1", func(ctx *Context) {
		_, _ = ctx.Json(Response{
			Docs:      baseDocs.docs,
			IpAddr:    getIpv4Addr(),
			Port:      port,
			ApiServer: apiServer,
		})
	})

	_ = en.Run("0.0.0.0:" + port)
}

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
