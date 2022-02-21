package filesystem

func (f file) IsHidden() bool {
	if f.Name()[0:1] == "." {
		return true
	} else {
		return false
	}
}
