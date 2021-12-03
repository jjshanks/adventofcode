package challenge

func PartTwo(instr string) string {
	oldPassword := PartOne(instr)
	newPassword := nextPassword([]byte(oldPassword))
	return newPassword
}
