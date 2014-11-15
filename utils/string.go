package utils

import (
	"math/rand"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyz01234567ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	i, l := 0, len(chars)
	b := make([]rune, n)
	for i < n {
		b[i] = chars[rand.Intn(l)]
		i++
	}
	return string(b)
}
