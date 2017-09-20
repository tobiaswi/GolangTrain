package main

func main() {
	var sum int
	n, m := 1, 2
	for n <= 4000000 {
		if n%2 == 0 {
			sum += n
		}
		n, m = m, n+m
	}
}
