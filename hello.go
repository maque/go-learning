/*
Package main is something

The s is

	re:
		aa

named return parameters


resources:
https://golang.org/ref/spec
https://golang.org/doc/effective_go.html
https://golang.org/doc/#articles
https://github.com/golang/protobuf
https://github.com/golang/go/wiki
https://golang.org/pkg/
https://gobyexample.com/
https://github.com/golang/go/wiki
https://golang.org/doc/faq
https://github.com/michaeljs1990/val
https://talks.golang.org/2015/json.slide
https://golang.org/cmd/go/
https://golang.org/doc/cmd
*/
package main

import "fmt"

type ApiConnecter interface {
	Connect()
}

const (
	UseMixedCaps string = "A"
)

func (s string) FromString() int {
	return 3
}

func main() {
	i:=1
	k, i := 1,2

	//for key, value := range {}
	//for key := range
	// for _, value := range
	// for i:=0; i<10; i++ {}

	switch {
		case i==0, i==2:
			if k==1 {

				break
			}
			fmt.Println("0")
		case i==2, k==1:
			fmt.Println("1")
	}
	fmt.Println("d".FromString())

	fmt.Printf("hello %v %v %v", "world", i ,k)
}
