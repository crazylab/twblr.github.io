package control_structures

func isEven(n int) bool{
	return n % 2 == 0
}

func collatzChainLength(n int) int {
	count := 0
	for n != 1 {
		if isEven(n){
			n = n / 2
		} else{
			n = 3 * n + 1
		}
		count++
	}
	return count
}
