package doc

import (
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
)

type Notify struct {
	Event fsnotify.Event
	Name  string
}

func checkError(err error, show bool) bool {
	if err != nil {
		if show {
			log.Println(err)
		}
		return true
	}
	return false
}

type FWatcher struct {
	count   int
	files   []string
	watcher *fsnotify.Watcher
	msg     chan Notify
}

func newFWatcher(files []string) *FWatcher {
	watcher, err := fsnotify.NewWatcher()
	_ = checkError(err, true)
	return &FWatcher{
		watcher: watcher,
		msg:     make(chan Notify),
		files:   files,
	}
}

func (fw *FWatcher) addFiles(files []string) {
	for _, filepath := range files {
		fw.addFile(filepath)
	}
}

func (fw *FWatcher) addFile(filepath string) {
	fw.files = append(fw.files, filepath)
	err := fw.watcher.Add(filepath)
	checkError(err, true)
}

func (fw *FWatcher) startWatchFiles() {
	for {
		select {
		case event, ok := <-fw.watcher.Events:
			if !ok {
				return
			}
			// 通知更新
			fw.msg <- Notify{Event: event}
		case err, ok := <-fw.watcher.Errors:
			if !ok {
				return
			}
			_ = checkError(err, true)
		}
	}
}


func (fw *FWatcher)handleConfigUpdate(){

}

// 针对配置文件的增加和删除
// add
// remove
// 调用toDoc


func (fw *FWatcher) handleUpdated(p *chan bool) {
	// 接收并处理更新
	for v := range fw.msg {
		if (fw.count+1)%2 == 0 {
			var doc = Doc{
				Name:    v.Event.Name,
				Content: make([]Annotation, 0),
			}
			con, _ := ioutil.ReadFile(v.Event.Name)
			// 提取注释
			extractAnnotation(con, func(annotation Annotation) {
				doc.Content = append(doc.Content, annotation)
			})
			//添加更新消息
			updatedData.Add(doc)
			// 通知更新
			*p <- true
		}
		fw.count++
	}
}

func (fw *FWatcher) run(p *chan bool) {
	fw.addFiles(fw.files)
	go fw.startWatchFiles()
	go fw.handleUpdated(p)
}
