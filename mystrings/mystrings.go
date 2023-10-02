package mystrings


// function signature has to be capitalized else it will not be exported
func Reverse(s string) string {
	result := ""
	for _,v := range s {
		result = string(v) + result
	}
	return result
}