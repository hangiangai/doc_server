package doc

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//
type Annotation struct {
	Key    string
	Title  string
	Url    string
	Header string
	Method string
	Params [][]string
}

//
type Doc struct {
	Name    string       // 文件名称
	Content []Annotation // 内容
}

// 过滤字符串
func filter(c byte) bool {
	var result bool
	switch c {
	case uint8('\n'):
	case uint8('\r'):
	case uint8(' '):
	case uint8('*'):
	default:
		result = true
	}
	return result
}

var (
	// 文件路径 注释md5 注释
	baseDocs = make(map[string]map[string]Annotation)
)

// 匹配关键字
func match(annotation *Annotation, key string, value []string) {
	switch key {
	case "title":
		annotation.Title = value[0]
	case "url":
		annotation.Url = value[0]
	case "header":
		annotation.Header = value[0]
	case "method":
		annotation.Method = value[0]
	case "param":
		annotation.Params = append(annotation.Params, value)
	}
}

// 提取注释
// con 需要解析的内容
// handler 自定义操作函数
func extractAnnotation(con []byte, handler func(annotation Annotation)) {
	var begin bool
	var conLen = len(con)
	var annotation Annotation
	var annotationCon bytes.Buffer
	for i := 1; i < conLen; i++ {
		if con[i] == '*' && con[i-1] == '/' {
			// 是否开始解析
			begin = true
			// 存储注释内容
			annotationCon = bytes.Buffer{}
			// 存储当前解析的注释
			annotation = Annotation{}
		}
		if filter(con[i]) && begin {
			annotationCon.WriteByte(con[i])
			if con[i] == '@' {
				line, err := annotationCon.ReadBytes('@')
				checkError(err, true)
				toStr := string(line)
				val := strings.Split(toStr[:len(toStr)-1], ":")
				match(&annotation, val[0], val[1:])
			}
		}
		if i+1 < conLen && con[i] == '*' && con[i+1] == '/' {
			val := strings.Split(annotationCon.String(), ":")
			match(&annotation, val[0], val[1:])
			tempAtn := Annotation{
				Title:  annotation.Title,
				Url:    annotation.Url,
				Header: annotation.Header,
				Method: annotation.Method,
				Params: annotation.Params,
				Key:    fmt.Sprintf("%x", md5.Sum([]byte(annotation.Url+annotation.Method))),
			}
			if handler != nil { //执行函数不为nil
				handler(tempAtn)
			}
		}
	}
}

//转化为文档
func toDoc(filepath string) Doc {
	con, err := ioutil.ReadFile(filepath)
	if checkError(err, true) {
		return Doc{}
	}
	doc := make([]Annotation, 0)
	// md5Sum值
	md5Sum := make(map[string]Annotation)
	// 提取文件注释
	extractAnnotation(con, func(annotation Annotation) {
		doc = append(doc, annotation)
		bs, err := json.Marshal(annotation)
		checkError(err, true)
		sum := fmt.Sprintf("%x", md5.Sum(bs))
		md5Sum[sum] = annotation
	})
	// 保存基础文档注释
	baseDocs[filepath] = md5Sum
	// 防护Doc文档
	return Doc{
		Name:    filepath,
		Content: doc,
	}
}
