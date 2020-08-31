package doc

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

type note struct {
	Title   string     //请求标题
	Url     string     //请求地址
	Header  string     //请求请求头
	Method  string     //请求方法
	Params  [][]string //请求参数
	ResHint string     //返回提示
	toReset string
}

/*

 */
func __filter__(c byte) bool {
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

/*
	notes 用于存储解析的内容
	con 读取的内容
	cLen 读取的内容长度
	count 有效解析个数 表示数组下表最大值
	filter 过滤函数
*/
func extractNote(notes *[]note, con []byte, cLen int, count *int, filter func(n *note) bool) {
	var begin bool
	var currentExtractCon bytes.Buffer
	var isAppend = true
	for i := 1; i < cLen; i++ {
		if con[i] == '*' && con[i-1] == '/' {
			begin = true
			currentExtractCon.Reset()
			if isAppend {
				*notes = append(*notes, note{})
			}
		} else if begin && __filter__(con[i]) {
			currentExtractCon.WriteByte(con[i])
			if con[i] == '@' {
				line, err := currentExtractCon.ReadBytes('@')
				checkError(err, true)
				toStr := string(line)
				val := strings.Split(toStr[:len(toStr)-1], ":")
				if len(val[1:]) > 0 {
					match(&(*notes)[*count], val[0], val[1:])
				}
			}
		} else if i+1 < cLen && con[i] == '*' && con[i+1] == '/' {
			begin = false
			val := strings.Split(currentExtractCon.String(), ":")
			match(&(*notes)[*count], val[0], val[1:])
			if filter != nil && filter(&(*notes)[*count]) {
				isAppend = false
			} else {
				*count++
				isAppend = true
			}
		}
	}
}

/*
	读取文件
*/
func readFile(filepath string, buffer *bytes.Buffer) {
	var (
		readBuf = make([]byte, 1024)
		file    *os.File
		err     error
	)
	if file, err = os.Open(filepath); err == nil {
		defer file.Close()
		n, err := file.Read(readBuf)
		for err == nil {
			buffer.Write(readBuf[:n])
			n, err = file.Read(readBuf)
		}
		if err != io.EOF {
			checkError(err, true)
		}
	}
}

/*
	sum值
*/
func md5Sum(c []byte) string {
	return fmt.Sprintf("%x", md5.Sum(c))
}

type Doc struct {
	Sum    string
	ConSum string
	Path   string
	Notes  []note
	Count  int
}

/*
	转化为文档
*/
func (doc *Doc) toDoc(filepath string, buffer *bytes.Buffer) *Doc {
	readFile(filepath, buffer)
	sum := md5Sum(buffer.Bytes())
	if sum == doc.Sum {
		return nil
	}
	var document Doc
	document.Path = filepath
	extractNote(&document.Notes, buffer.Bytes(),
		buffer.Len(), &document.Count, func(n *note) bool {
			if n.Title == "" || n.Url == "" {
				return true
			}
			return false
		})

	return &document
}

/*

 */
func (doc *Doc) paramsFile(filepath string) *Doc {
	var buffer bytes.Buffer
	return doc.toDoc(filepath, &buffer)
}

/*

 */
type Docs struct {
	Doc
	docs  []*Doc
	count int
}

/*

 */
func (ds *Docs) params(filepath []string) {
	var buffer bytes.Buffer
	for _, v := range filepath {
		ds.docs = append(ds.docs, ds.toDoc(v, &buffer))
		ds.count++
		buffer.Reset()
	}
}
