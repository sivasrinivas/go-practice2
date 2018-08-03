package bits

func countSetBits(n int) (count int) {
	var one = 1
	for n > 0 {
		set := (n & one) > 0
		if set {
			count++
		}
		n = n >> 1
	}
	return
}
