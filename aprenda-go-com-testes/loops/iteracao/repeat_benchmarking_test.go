package iteracao

import "testing"

// Não é necessário possuir essa nomenclatura

func BenchmarkRepition(b *testing.B) {

	reptitions := int64(4)
	for i := 0; i < b.N; i++ {
		Repeat("a", reptitions)
	}
}
