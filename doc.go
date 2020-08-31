package doc

type Response struct {
	Docs      []*Doc
	IpAddr    string
	Port      string
	ApiServer string
}

/*
@title: 数据传输结构体
@param: type:int:传输类型
@param: data:string:传输信息
@param: state:int:传输动作
@param: identity:string:身份信息
*/
type TransmissionData struct {
	Type     int
	Data     interface{}
	Status   int
	Identity string
}

type DataExchangeBody struct {
	Type     int
	Identity string
	Data     interface{}
}

var (
	baseDocs = &Docs{docs: make([]*Doc, 0)}
	config   *Config
	store    = NewStore()
	hub      = newHub()
)

func InitAndRun(path string) {
	store.Run()
	go hub.run()
	config = NewConfig(path)
	baseDocs.params(config.files())
	fWatcher := newFWatcher(config.files())
	fWatcher.run()
	initServer(config.Port, config.ApiServ)
}
