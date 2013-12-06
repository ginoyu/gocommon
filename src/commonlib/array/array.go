package array

import "errors"

type Array struct {
	inner []interface{}
}

func New() *Array {
	return &Array{inner: make([]interface{}, 0)}
}

func (p *Array) Push(a interface{}) {
	p.inner = append(p.inner, a)
}

func (p *Array) Pop() (interface{}, error) {
	if len(p.inner) > 0 {
		i := len(p.inner) - 1
		ret := p.inner[i]
		p.inner = p.inner[0:i]
		return ret, nil
	} else {
		err := errors.New("The Array is empty")
		return nil, err
	}
}

func (p *Array) Size() int {
	return len(p.inner)
}

func (p *Array) Array() []interface{} {
	return p.inner
}

func (p *Array) Append(elems ...interface{}) {
	for _, elem := range elems {
		p.inner = append(p.inner, elem)
	}
}
