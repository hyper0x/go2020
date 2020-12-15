package overlapping

import "io"

// MyReader 代表可读的自定义接口。
type MyReader interface {
	io.ReadCloser
}

// MyWriter 代表可写的自定义接口。
type MyWriter interface {
	io.WriteCloser
}

// MyIO 代表可输入输出的自定义接口。
type MyIO interface {
	MyReader
	MyWriter
	// 下面的这个方法声明的添加将会让编译器报错，
	// 因为它与那些内嵌接口中的 Close 方法不完全重叠。
	// Close()
	// 下面这个内嵌接口是有效的，
	// 因为其中的 Close 方法与其他的完全重叠。
	io.Closer
}
