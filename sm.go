package sm

import (
	"container/list"
	"strconv"
)

type Machine struct {
	ops   []Operation
	stack *list.List
}

type Operation func(*Machine) error

func New(ops []Operation) *Machine {
	return &Machine{ops: ops, stack: list.New()}
}

func (m *Machine) Push(item uint64) {
	m.stack.PushFront(item)
}

func (m *Machine) Pop() uint64 {
	e := m.stack.Front()
	m.stack.Remove(e)
	return e.Value.(uint64)
}

func (m *Machine) Op() error {
	op := m.Pop()
	return m.ops[op](m)
}

func (m *Machine) Size() int {
	return m.stack.Len()
}

func (m *Machine) Run() error {
	for 0 != m.Size() {
		err := m.Op()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Machine) FromStr(s string, opChar byte, opLen int, base int, width int) error {
	for len(s) > 0 {
		if s[0] == opChar {
			s = s[1:]
			err := m.Op()
			if err != nil {
				return err
			}
		} else {
			u, err := strconv.ParseUint(s[:opLen], base, width)
			if err != nil {
				return err
			}
			s = s[opLen:]
			m.stack.PushFront(u)
		}
	}
	return nil
}
