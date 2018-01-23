package restoreipaddresses

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_restoreIPAddresses(t *testing.T) {
	assert.Equal(t, []string{}, restoreIPAddresses("000"))
	assert.Equal(t, []string{"0.0.0.0"}, restoreIPAddresses("0000"))
	assert.Equal(t, []string{"1.1.1.1"}, restoreIPAddresses("1111"))
	assert.Equal(t, []string{"0.10.0.10", "0.100.1.0"}, restoreIPAddresses("010010"))
	assert.Equal(t, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}, restoreIPAddresses("101023"))
	assert.Equal(t, []string{"255.255.11.135", "255.255.111.35"}, restoreIPAddresses("25525511135"))
	assert.Equal(t, []string{"11.111.111.111", "111.11.111.111", "111.111.11.111", "111.111.111.11"}, restoreIPAddresses("11111111111"))
	assert.Equal(t, []string{"2.5.62.65", "2.56.2.65", "2.56.26.5", "25.6.2.65", "25.6.26.5", "25.62.6.5"}, restoreIPAddresses("256265"))
}

func Benchmark_restoreIPAddresses(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			restoreIPAddresses("111")
			restoreIPAddresses("1111")
			restoreIPAddresses("101023")
			restoreIPAddresses("25525511135")
		}
	})
}
