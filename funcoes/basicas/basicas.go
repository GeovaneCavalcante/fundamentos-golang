package main

import "fmt"

func f1() {
	fmt.Printf("F1")
}

func f2(p1 string, p2 string) {
	fmt.Println(p1, p2)
}

func f3() string {
	return "f3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf(p1, p2)
}

func f5() (string, string, string) {
	return "a", "b", "c"
}

func main() {
	f1()
	f2("a", "b")
	r3, r4 := f3(), f4("a2", "12")
	fmt.Println(r3, r4)
	r5, _, r7 := f5()
	fmt.Println(r5, r7)

}
