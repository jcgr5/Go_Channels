package ProductoCartesianoSlices

import (
	"fmt"
	"strconv"
	"sync"
)

func PcInt(a []int, b []int) []string {
	var A, B, C []string
	wg := sync.WaitGroup{}
	fmt.Println(a, "\n", b)
	wg.Add(1)
	//convertir los slices de enteros en strings//////////
	go func() {
		for _, i := range a {
			A = append(A, strconv.Itoa(i))
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for _, j := range b {
			B = append(B, strconv.Itoa(j))
		}
		wg.Done()
	}()
	wg.Wait()
	//////////////////////////////////////////
	for _, i := range A {
		for _, j := range B {
			z := "(" + i + "," + j + ")" //concatena cada valor del slice A con cada valor de B
			C = append(C, z)
		}
	}
	return C
}

func PcChar(a []byte, b []byte) []byte {
	var A, B, C []byte
	var z byte
	A, B = a, b
	for i := range A {
		for j := range B {
			z = byte(i + j)
			C = append(C, z)
		}
	}
	return C
}

func PcStr(a []string, b []string) []string {
	var A, B, C []string
	var z string
	A, B = a, b
	for i := range A {
		for j := range B {
			z = string(i + j)
			C = append(C, z)
		}
	}
	return C
}
