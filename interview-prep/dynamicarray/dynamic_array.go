package dynamicarray

type (
	DynamicArray struct {
		arry [4]int
	}
)

func New() *DynamicArray {
	return &DynamicArray{}
}

func (d *DynamicArray) Set(idx int, val int) {
	d.arry[idx] = val
}

func (d *DynamicArray) Get(idx int) int {
	return d.arry[idx]
}
