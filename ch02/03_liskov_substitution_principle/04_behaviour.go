package lsp

type Collection interface {
	Add(item interface{})
	Get(index int) interface{}
}

type CollectionImpl struct {
	items []interface{}
}

func (c *CollectionImpl) Add(item interface{}) {
	c.items = append(c.items, item)
}

func (c *CollectionImpl) Get(index int) interface{} {
	return c.items[index]
}

type ReadOnlyCollection struct {
	CollectionImpl
}

func (ro *ReadOnlyCollection) Add(item interface{}) {
	// intentionally does nothing
}

func testIt(c Collection) {
}

func Caller() {
	ro := &ReadOnlyCollection{}
	testIt(ro)
}

type SimplyAdd interface {
	Add(item interface{})
}

func tryOne(sim SimplyAdd) {

}

func callIt() {
	c := &CollectionImpl{}

	tryOne(c)
}
