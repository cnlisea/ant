package mathx

func Gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	if b == 0 {
		return a
	}

	return Gcd(b, a%b)
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}
