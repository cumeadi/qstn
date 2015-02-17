package num

import "math/rand"
import "time"

func init() {
	rand.Seed(time.Now().Unix())
}

func Rand(n int) int {
	return rand.Intn(n)
}

func RandBetween(a, b int) int {
	return rand.Intn(b-a) + a
}
