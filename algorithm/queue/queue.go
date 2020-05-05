package queue

import "fmt"

type Queue struct {
	q *[]int
}

func (q *Queue)init(){
	if q.q == nil {
		q.q = &[]int{}
	}
}

func (q *Queue)Push(data int){
	q.init()
	*q.q = append(*q.q, data)
}

func (q *Queue)Pop() int {
	q.init()
	if q.Len() < 1 {
		return -1
	}
	r := (*q.q)[0]
	*q.q = (*q.q)[1:]
	return r
}

func (q *Queue)Len() int {
	return len(*q.q)
}

func (q *Queue)Reset(){
	*q.q = (*q.q)[:0]
}

func (q *Queue)Print() {
	fmt.Printf("Length: %d, data: %v\n", q.Len(), *q.q)
}

func NewQueue() *Queue {
	q := Queue{}
	q.init()
	return &q
}
