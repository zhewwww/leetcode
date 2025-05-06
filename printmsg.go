package leetcode
//依次打印n次foo，bar
type FooBar struct {
	n int
    writeChannel chan int
}

func NewFooBar(n int) *FooBar {
    return &FooBar{n: n, writeChannel: make(chan int)}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
        printFoo()
        fb.writeChannel <- 0
        <- fb.writeChannel
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
        <- fb.writeChannel
        printBar() 
        fb.writeChannel <- 0
	}
}