package challenge

var (
	a byte = []byte("a")[0]
	z byte = []byte("z")[0]
	i byte = []byte("i")[0]
	o byte = []byte("o")[0]
	l byte = []byte("o")[0]
)

func nextPassword(password []byte) string {
	newPassword := make([]byte, 8)
	copy(newPassword, password)
	endIdx := len(newPassword) - 1
	// if current password is valid increment it
	if isValid(newPassword) {
		incIdx(newPassword, endIdx)
	}
	// increment until valid
	for !isValid(newPassword) {
		incIdx(newPassword, endIdx)
	}
	return string(newPassword)
}

func incIdx(password []byte, idx int) {
	password[idx] += 1
	if password[idx] > z {
		password[idx] = a
		incIdx(password, idx-1)
	}
}

func isValid(password []byte) bool {
	hasStraight := false
	dupIdx := make([]int, 0)
	last := byte(0)
	for idx, b := range password {
		// can't contain o, i, or l
		if b == o || b == i || b == l {
			return false
		}
		// must contain 3 consective letters in increasing order
		if idx+2 < len(password) && b+1 == password[idx+1] && b+2 == password[idx+2] {
			hasStraight = true
		}
		// must contain two non overlapping pairs
		if b == last && !contains(dupIdx, idx-1) {
			dupIdx = append(dupIdx, idx-1, idx)
		}
		last = b
	}
	return hasStraight && len(dupIdx) >= 4
}

func contains(stack []int, needle int) bool {
	for _, i := range stack {
		if needle == i {
			return true
		}
	}
	return false
}
