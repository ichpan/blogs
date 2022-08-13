package unittest

func CalcAdd(n int) (sum int) {
	for i := 1; i <= n-1; i++ {
		sum += i
	}
	return sum
}
