package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "Tyme", 60, "The Tyme")
	mag2, _ := newPublication("magazine", "NYT", 55, "The NYT")
	news1, _ := newPublication("newspaper", "The News", 30, "The News Publisher")
	news2, _ := newPublication("newspaper", "The Tyme", 30, "The Tyme Publisher")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails(pub iPublication) {
	fmt.Printf("--------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Printf("--------\n")
}
