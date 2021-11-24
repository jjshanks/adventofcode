package challenge

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func PartOne(instr string) int {
	keyData := []byte(instr)
	num := 0
	hash := ""
	for !strings.HasPrefix(hash, "00000") {
		num += 1
		numBytes := []byte(strconv.Itoa(num))
		data := append(keyData, numBytes...)
		hash = fmt.Sprintf("%x", md5.Sum(data))
	}
	return num
}
