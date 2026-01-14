package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
	x := 1
	xPtr := &x
	var p *int

	fmt.Printf("%T %+v\n", xPtr, *xPtr)
	fmt.Printf("%T %+v\n", p, p)

	ps := struct{ x, y int }{1, 3}
	pPtr := &ps

	fmt.Println(pPtr.x) // (*pPtr).x
	fmt.Println((*pPtr).y)
	ps.x = 2
	fmt.Println((*pPtr).x)
	pPtr.x = 2
	fmt.Println((*pPtr).x)

	//fmt.Println(pPtr.x)

	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(i)
	fmt.Println(i[0:3])
	//	i := AppendUniq

	//ch := make(chan int)
	//
	//for i := 0; i < 3; i++ {
	//	i := i
	//	go func() {
	//		ch <- i
	//		//fmt.Println(i)
	//	}()
	//}

	//for d := range ch {
	//	fmt.Println(d)
	//}
	fmt.Println(DOCA)
	fmt.Printf("Hash: %s\n", hash(DOCA))
	fmt.Println(DOCB)
	fmt.Printf("Hash: %s\n", hash(DOCA))

}

// NB These docs are strictly-speaking the same.
const DOCA = "{ \"foo\": 1.23e1, \"bar\": { \"baz\": true, \"abc\": 12 } }"
const DOCB = "{ \"bar\": { \"abc\": 12, \"baz\": true }, \"foo\": 12.3 }"

func hash(doc string) string {
	// Dumb af, but it's a cheap way to specific the most generic thing
	// you can :-/
	var v interface{}
	json.Unmarshal([]byte(doc), &v) // NB: You should handle errors :-/
	cdoc, _ := json.Marshal(v)
	sum := sha256.Sum256(cdoc)
	return hex.EncodeToString(sum[0:])
}
