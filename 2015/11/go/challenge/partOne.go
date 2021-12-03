package challenge

func PartOne(instr string) string {
	password := []byte(instr)
	newPassword := nextPassword(password)
	return newPassword
}
