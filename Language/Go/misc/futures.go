package misc

// 所谓Futures就是指：有时候在你使用某一个值之前需要先对其进行计算。
// 这种情况下，你就可以在另一个处理器上进行该值的计算，到使用时，
// 该值就已经计算完毕了。Futures模式通过闭包和通道可以很容易实现，
// 类似于生成器，不同地方在于Futures需要返回一个值。
// ref: http://www.golangpatterns.info/concurrency/futures
//

type Matrix struct{}

func Product(a Matrix, b Matrix) Matrix {
	//
}

func Inverse(a Matrix) Matrix {
	//
}

func InverseProduct(a Matrix, b Matrix) {
	a_inv_future := InverseFuture(a)
	b_inv_future := InverseFuture(b)
	a_inv := <-a_inv_future
	b_inv := <-b_inv_future
	return Product(a_inv, b_inv)
}

func InverseFuture(a Matrix) chan Matrix {
	future := make(chan Matrix)
	go func() {
		future <- Inverse(a)
	}()
	return future
}
