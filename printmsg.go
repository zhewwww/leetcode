package leetcode

// FooBar 用于按顺序打印 "foo" 和 "bar"，共 n 次。
type FooBar struct {
	n            int         // 打印的次数
	writeChannel chan int    // 用于线程同步的通道
}

// NewFooBar 创建一个新的 FooBar 实例
func NewFooBar(n int) *FooBar {
	return &FooBar{n: n, writeChannel: make(chan int)}
}

// Foo 方法负责打印 "foo"，并通过通道通知 Bar 方法
func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// 调用传入的函数打印 "foo"
		printFoo()
		// 向通道发送信号，通知 Bar 方法可以运行
		fb.writeChannel <- 0
		// 等待 Bar 方法处理完成
		<-fb.writeChannel
	}
}

// Bar 方法负责打印 "bar"，并通过通道通知 Foo 方法
func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// 等待 Foo 方法发来的信号
		<-fb.writeChannel
		// 调用传入的函数打印 "bar"
		printBar()
		// 向通道发送信号，通知 Foo 方法可以继续运行
		fb.writeChannel <- 0
	}
}