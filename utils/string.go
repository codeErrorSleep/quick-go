package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Str2MD5 returns a MD5 hash in string form of the passed-in `s`
func Str2MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// StringConcat returns the concatenation of `base` and `strs` strings seperated by `sep`
func StringConcat(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

func LiteralToMD5[T string | []byte](s T) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
