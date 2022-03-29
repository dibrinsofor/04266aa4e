package lib

import (
	"math/rand"
)

func GenShortSlug() string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	s := make([]rune, 4)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}
