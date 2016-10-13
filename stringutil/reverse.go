package stringutil

func Reverse(s string) string {
// r := ... is equal to var r <type> = ...  type is set at compile time by the value provided
	// this can only be used inside a function.  Otherwise, must use var x <type> = ...
// rune is a data type consversion that allows you to access each char of a string
	r := []rune(s)
// values can be set to variables simultaneously using commas as in the following lines
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
