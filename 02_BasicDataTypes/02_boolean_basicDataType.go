// For documentiation in go file
// use the godoc packege
// to instal in terminal type  :
// > go get -u golang.org/x/tools/cmd/godoc
//after install run it as a web server by this command  :
// > godoc -http=:6060
// and then in browswer go to this url  : http://localhost:6060/pkg/

// boolean types
// :=    used for first  assing value  to  variable that doesnt exist

package main

import "fmt"

func main() {
	a := true
	b := false

	//fmt.Println("",a,b,a==b)
	fmt.Printf("a=%v , b=%v , a==b = %v", a, b, a == b)

}
