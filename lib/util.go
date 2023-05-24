package lib

import "math/rand"

func Shuffle[T any](slice []T, source rand.Source) {
	random := rand.New(source)
	for i := len(slice) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
