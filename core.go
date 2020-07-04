package doc

import (
	"bytes"
	"io/ioutil"
)

type C struct {
	s    bytes.Buffer
	seek [][]int
}

func readFile(_f string) []map[string][]string {
	con, err := ioutil.ReadFile(_f)
	checkError(err, 200)
	conLen := len(con)
	var readByte bytes.Buffer
	var begin bool //
	var p_index int = -1 //
	var seek int // 偏移量
	p := make([][]int, 21) // 用于标记关键字的下标
	transformData := make([]map[string][]string, 0)
	for i := 1; i < conLen; i++ {
		if con[i] == '*' && con[i-1] == '/' {
			begin = true
			readByte = bytes.Buffer{}
			p_index = -1
			p = make([][]int, 20)
			seek = 0
		}
		if begin && Filter(con[i]) {
			readByte.WriteByte(con[i])
			if con[i] == '@' {
				p_index++
				p[p_index] = append(p[p_index], seek)
			}
			if p_index > -1 && len(p[p_index]) > 0 && con[i] == ':' {
				p[p_index] = append(p[p_index], seek)
			}
			seek++
		}
		if i+1 < conLen && con[i] == '*' && con[i+1] == '/' {
			p[p_index+1] = append(p[p_index+1], len(readByte.String()))
			res := transform(C{s: readByte, seek: p})
			transformData = append(transformData, res)
		}
	}
	return transformData
}

func transform(c C) map[string][]string {
	params := make(map[string][]string)
	for i := 0; len(c.seek[i]) > 1; i++ {
		cStr := c.s.String()
		key := cStr[c.seek[i][0]:c.seek[i][1]]
		val := cStr[c.seek[i][1]+1 : c.seek[i+1][0]]
		params[key] = append(params[key], val)
	}
	return params
}


// 过滤字符
func Filter(c byte) bool {
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