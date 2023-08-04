package main

func main() {
	x := []int{3, 5, 8}
	y := append(x, 11)
	y = append(y, 22, 34, 22)
	println(cap(x))
	println(cap(y))
}
