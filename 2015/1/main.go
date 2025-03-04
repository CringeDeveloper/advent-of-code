package main

func main() {
}

func getFloor(s string) int {
	floor := 0

	for _, c := range s {
		if c == '(' {
			floor++
		} else {
			floor--
		}
	}

	return floor
}

func getBasementIndex(s string) int {
	floor := 0

	for i, c := range s {
		if c == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}

	return -1
}
