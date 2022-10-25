package helper

import "github.com/inhies/go-bytesize"

const MaxFileSize = 1 * bytesize.MB

func IsCorrectSize(size uint64) bool {
	return size <= uint64(MaxFileSize)
}
