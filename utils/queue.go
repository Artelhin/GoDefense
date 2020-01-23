package utils

type Queue struct {
	data []interface{}
}

func (q *Queue) Push(i interface{}) {
	q.data = append(q.data, i)
}

func (q *Queue) Pop() interface{} {
	if len(q.data) == 0 {
		return nil
	}
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *Queue) Len() int {
	return len(q.data)
}

func NewQueue() *Queue {
	return &Queue{
		data:make([]interface{},0),
	}
}