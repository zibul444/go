package main

func main() {
	f := fib()
	println(f(), f(), f(), f(), f())
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
