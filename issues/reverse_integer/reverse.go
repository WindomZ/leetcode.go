package reverse_integer

//import "runtime"
//
//func init() {
//	runtime.GOMAXPROCS(4)
//}

func reverse_1(x int) (n int) {
	for ; x != 0; x = x / 10 {
		n = n*10 + x%10
	}
	return
}

func reverse_2(x int) (n int) {
	for ; x != 0; x = x / 10 {
		n *= 10
		n += x % 10
	}
	return
}

func reverse_3(x int) (n int) {
	for t := 0; x != 0; x = x / 10 {
		t = x % 10
		n = n*10 + t
	}
	return
}

func reverse_4(x int) (n int) {
	for t := 0; x != 0; x = x / 10 {
		t = x % 10
		n *= 10
		n += t
	}
	return
}
