package repl

import "main/value"

const (
	BreakItem    = "break"
	ContinueItem = "continue"
	ReturnItem   = "return"
)

type CallStackItem struct {
	ReturnValue value.IVOR
	Type        []string
	Action      string
}

func (csi *CallStackItem) IsType(t string) bool {

	for _, i := range csi.Type {
		if i == t {
			return true
		}
	}

	return false
}

func (csi *CallStackItem) IsAction(a string) bool {
	return csi.Action == a
}

type CallStack struct {
	Items []*CallStackItem
}

func (cs *CallStack) Push(item *CallStackItem) {
	cs.Items = append(cs.Items, item)
}

func (cs *CallStack) Pop() *CallStackItem {
	item := cs.Items[len(cs.Items)-1]
	cs.Items = cs.Items[:len(cs.Items)-1]
	return item
}

func (cs *CallStack) Peek() *CallStackItem {
	return cs.Items[len(cs.Items)-1]
}

func (cs *CallStack) In(item *CallStackItem) bool {
	for _, i := range cs.Items {
		if i == item {
			return true
		}
	}
	return false
}

func (cs *CallStack) Clean(item *CallStackItem) {

	if !cs.In(item) {
		return
	}

	for {
		peek := cs.Pop()

		if peek == item {
			break
		}
	}

}

func (cs *CallStack) Len() int {
	return len(cs.Items)
}

func NewCallStack() *CallStack {
	return &CallStack{Items: []*CallStackItem{}}
}
