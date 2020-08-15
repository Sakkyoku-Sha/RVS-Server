package main

import (
	"fmt"
)

func main() {
	var t = []byte{0}
	var request = APIRequest{BadRequest, t}
	fmt.Printf("%d", request.Type)
}
