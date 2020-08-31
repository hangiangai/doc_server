package doc

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"time"
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
			fw.msg <- Notify{Event: event}
		case err, ok := <-fw.watcher.Errors:
			if !ok {
				return
			}
			_ = checkError(err, true)
		}
	}
}

func (fw *FWatcher) handleUpdated() {
	for msg := range fw.msg {
		if (fw.count+1)%2 == 0 {
			if con := baseDocs.paramsFile(msg.Event.Name); con != nil {
				for i := 0; i < baseDocs.count; i++ {
					if baseDocs.docs[i].Path == msg.Event.Name {
						baseDocs.docs[i] = con
					}
				}
				store.Set("updated", con, time.Second*61)
				hub.broadcast <- exchangeBody{
					Type: 3,
				}
			}
		}
		fw.count++
	}
}

func (fw *FWatcher) run() {
	fw.addFiles(fw.files)
	// 监听文件
	go fw.startWatchFiles()
	// 处理文件
	go fw.handleUpdated()
}
