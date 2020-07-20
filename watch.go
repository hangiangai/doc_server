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

func checkError(err error) bool {
	if err != nil {
		log.Fatal(err)
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
	checkError(err)
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
	checkError(err)
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
			checkError(err)
		}
	}
}


func (fw *FWatcher) handleUpdated(queue *Queue) {
	// 接收并处理更新
	for v := range fw.msg {
		if (fw.count + 1 ) % 2 == 0 {
			var doc = Doc{
				Name: v.Event.Name,
				Content: make([]Annotation, 0),
			}
			con, _ := ioutil.ReadFile(v.Event.Name)
			extractAnnotation(con, func(annotation Annotation) {
				doc.Content = append(doc.Content, annotation)
			})
			queue.Add(doc)
		}
		fw.count++
	}
}

func (fw *FWatcher) run(queue *Queue) {
	fw.addFiles(fw.files)
	go fw.startWatchFiles()
	go fw.handleUpdated(queue)
}
