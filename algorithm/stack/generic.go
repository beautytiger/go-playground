package stack

import "fmt"

//a queue can contain any type
type Generic struct {
	q *[]interface{}
}

func (q *Generic)init(){
	if q.q == nil {
		q.q = &[]interface{}{}
	}
}

func (q *Generic)Push(data interface{}){
	q.init()
	*q.q = append(*q.q, data)
}

func (q *Generic)Pop() interface{} {
	q.init()
	if q.Len() < 1 {
		return -1
	}
	r := (*q.q)[len(*q.q)-1]
	*q.q = (*q.q)[:len(*q.q)-1]
	return r
}

func (q *Generic)Len() int {
	return len(*q.q)
}

func (q *Generic)Reset(){
	*q.q = (*q.q)[:0]
}

func (q *Generic)Print() {
	fmt.Printf("Length: %d, data: %v\n", q.Len(), *q.q)
}

func NewGeneric() *Generic {
	q := Generic{}
	q.init()
	return &q
}
