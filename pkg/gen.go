package pkg

import (
	"github.com/ashupednekar/hdata-encoder/internal/spec"
	"math/rand"
)

func randomData(n int, maxStr int) DataInput {
	if n <= 0 {
		n = 1
	}
	count := rand.Intn(n) + 1
	out := make(DataInput, 0, count)
	for range count {
		if rand.Intn(10) == 0 {
			break
		}
		choice := rand.Intn(3)
		switch choice {
		case 0:
			out = append(out, spec.Str(randomString(rand.Intn(maxStr)+1)))
		case 1:
			out = append(out, spec.I32(int32(rand.Intn(10_000))))
		case 2:
			size := rand.Intn(count)
			out = append(out, randomData(size, maxStr))
		}
	}

	return out
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
