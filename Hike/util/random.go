package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(lowerCharSet)

	for i := 0; i < n; i++ {
		c := lowerCharSet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomUserName generates a random user first name
func RandomUserName() string {
	return RandomString(8)
}

// RandomAge generates a random user age from 14 to 85
func RandomAge() int32 {
	return int32(RandomInt(14, 85))
}

// RandomEmail generates a random user email
func RandomEmail() string {
	mail := []string{"gmail.com", "mail.ru", "icloud.com", "yandex.ru"}
	n := len(mail)

	email := RandomString(int(RandomInt(6, 10)))
	email += "@"
	email += mail[rand.Intn(n)]

	return email
}

// RandomPassword generates a random user password
func RandomPassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(upperCharSet + lowerCharSet + numberSet)
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

// RandomRole generates a random user role
func RandomRole() string {
	role := []string{"admin", "hiker", "photographer", "bloger"}
	n := len(role)
	return role[rand.Intn(n)]
}
