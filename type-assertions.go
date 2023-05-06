package main

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random{}
}