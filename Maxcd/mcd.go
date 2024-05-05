package Maxcd

import (
	"sync"
	"time"
)

// calcula el maximo comun divisor entre dos numeros enteros de manera secuencial
func CalMaxcd(x, y int) (int, string, float64) {
	tinit := time.Now()
	time.Sleep(1000)
	if x > 0 && y == 0 {
		tend := time.Since(tinit)
		//fmt.Println(x, ", tiempo de ejecusion: ", tend)
		return x, "tiempo de ejecucion secuencial: ", tend.Seconds()
	}
	r := x % y
	x = y
	y = r
	return CalMaxcd(x, y)
}

// calcula el maximo comun divisor entre dos numero enteros
// de manera distribuida usando canales.
func CalMaxcdCh(x, y int) (int, string, float64) {
	tinit := time.Now()
	time.Sleep(1000)
	var wg sync.WaitGroup
	contador := 0
	for contador < 1 {
		if x > 0 && y == 0 {
			tend := time.Since(tinit)
			return x, "tiempo de ejecucion concurrente: ", tend.Seconds()
		} else if y > 0 {
			c := make(chan int, 2)
			c <- x
			c <- y
			wg.Add(1)
			go func() {
				x2 := <-c
				y2 := <-c
				u := x2 % y2
				x2 = y2
				y2 = u
				c <- x2
				c <- y2
				wg.Done()
			}()
			wg.Wait()
			x = <-c
			y = <-c
		}
	}
	tend := time.Since(tinit)
	return x, "tiempo de ejecucion: ", tend.Seconds()
}
