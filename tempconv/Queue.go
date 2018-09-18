package tempconv

//借鉴C++对队列的实现思路， 简单实现
type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	head *Node //头节点
	end  *Node //尾节点
}

//返回一个初化始的队列
func NewQueue() *Queue {
	q := &Queue{nil, nil}
	return q
}

//传入的使指针 那么不用返回什么 结果也会随着指针改变带回去
func (q *Queue) Push(data interface{}) {
	n := &Node{
		data: data,
		next: nil,
	}
	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	return
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.head == nil {
		return nil, false
	}
	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}
	return data, true
}
