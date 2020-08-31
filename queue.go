package doc

type Queue struct {
	cap   int           // 容量
	data  []interface{} // 数据
	size  int           // 个数
	index int           // 下标
}

func NewQueue(n int) Queue {
	return Queue{
		data:  make([]interface{}, n),
		size:  0,
		index: -1,
		cap:   n,
	}
}

func (q *Queue) Add(v interface{}) {
	q.size++
	q.index++
	if q.size > q.cap {
		q.data = append(q.data, v)
	} else {
		q.data[q.index] = v
	}
}

func (q *Queue) Top() interface{} {
	var ret interface{}
	if q.size > 0 {
		ret = q.data[q.index]
		q.size--
		q.index--
	}
	return ret
}

func (q *Queue) All() []interface{} {
	ret := q.data[0:q.size]
	q.Clear()
	return ret
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Clear() {
	q.size = 0
	q.index = -1
}
