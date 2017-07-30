package reverse_integer

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

var test_x int = 1234567890

func Test_reverse(t *testing.T) {
	assert.Equal(t, 987654321, reverse_1(test_x))
	assert.Equal(t, 987654321, reverse_2(test_x))
	assert.Equal(t, 987654321, reverse_3(test_x))
	assert.Equal(t, 987654321, reverse_4(test_x))
}

// Benchmark_A
func Benchmark_A_reverse_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_1(test_x)
	}
}
func Benchmark_A_reverse_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_2(test_x)
	}
}
func Benchmark_A_reverse_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_3(test_x)
	}
}
func Benchmark_A_reverse_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_4(test_x)
	}
}

// Benchmark_B
func Benchmark_B_reverse_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_3(test_x)
	}
}
func Benchmark_B_reverse_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_4(test_x)
	}
}
func Benchmark_B_reverse_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_1(test_x)
	}
}
func Benchmark_B_reverse_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_2(test_x)
	}
}

// Benchmark_C
func Benchmark_C_reverse_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_1(test_x)
	}
}
func Benchmark_C_reverse_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_3(test_x)
	}
}
func Benchmark_C_reverse_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_2(test_x)
	}
}
func Benchmark_C_reverse_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_4(test_x)
	}
}
