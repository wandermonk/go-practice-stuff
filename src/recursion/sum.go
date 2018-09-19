package recursion

//SumOfN is a recursive function to generate sum of N numbers
func SumOfN(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n + SumOfN(n-1)
}
