package tree

import (
	"errors"
	"fmt"
)

type ArrayTree struct {
	Item *[]int
}

// become useless when add SetLeft
func (a *ArrayTree)Add(data int) error {
	if data == 0 {
		return errors.New("0 is reserved for nil")
	}
	if a.Item == nil {
		v := make([]int, 0)
		a.Item = &v
	}
	*a.Item = append(*a.Item, data)
	return nil
}

// become useless when add SetLeft
func (a *ArrayTree)Pop(){
	if a.Item == nil {
		return
	}
	if len(*a.Item) == 0 {
		return
	}
	*a.Item = (*a.Item)[:len(*a.Item)-1]
}

func (a *ArrayTree)expandArray(maxIndex int) {
	for len(*a.Item) < maxIndex+1 {
		*a.Item = append(*a.Item, 0)
	}
}

func (a *ArrayTree)SetLeft(data int, parent int) error {
	if data == 0 {
		return errors.New("0 is reserved for nil")
	}
	if len(*a.Item) < parent + 1 {
		return errors.New("parent is nil")
	}
	if (*a.Item)[parent] == 0 {
		fmt.Println("parent is nil, please set parent first")
		return errors.New("parent is nil")
	}
	index := 2*parent +1
	a.expandArray(index)
	(*a.Item)[index] = data
	return nil
}

func (a *ArrayTree)SetRight(data int, parent int) error {
	if data == 0 {
		return errors.New("0 is reserved for nil")
	}
	if len(*a.Item) < parent + 1 {
		return errors.New("parent is nil")
	}
	if (*a.Item)[parent] == 0 {
		fmt.Println("parent is nil, please set parent first")
		return errors.New("parent is nil")
	}
	index := 2*parent +2
	a.expandArray(index)
	(*a.Item)[index] = data
	return nil
}

func (a *ArrayTree)Print() {
	if a.Item == nil {
		return
	}
	for _, v := range *a.Item {
		if v == 0 {
			fmt.Print("-")
		} else {
			fmt.Print(v)
		}
		fmt.Print(" ")
	}
}

