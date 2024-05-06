package MayorMatriz

import (
	"sync"
	"time"
)

// Recibe una matriz por parametro y calcula el mayor elemento de manera secuencial
func CalMayor(m [][]int) (int, string, float64) {
	tinit := time.Now()
	time.Sleep(1000)
	mayor := 0
	for cont := 0; cont < len(m); cont++ {
		fila := m[cont]
		for i := 0; i < len(fila); i++ {
			if mayor < fila[i] {
				mayor = fila[i]
			}
		}
	}
	tend := time.Since(tinit)
	return mayor, "Tiempo de ejecucion secuencial: ", tend.Seconds()
}

// Recibe una matriz por parametro y calcula el mayor elemento de manera distribuida
func CalMayorDist(m [][]int) (int, string, float64) {
	tinit := time.Now()
	time.Sleep(1000)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	mayorV := make([]int, 1, len(m))
	mayor := 0
	for cont := 0; cont < len(m); cont++ {
		fila := m[cont]
		wg.Add(1)
		go func() {
			mayorTemp := 0
			for i := 0; i < len(fila); i++ {
				if mayorTemp < fila[i] {
					mayorTemp = fila[i]
				}
			}
			mu.Lock()
			mayorV = append(mayorV, mayorTemp)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	for i := 0; i < len(mayorV); i++ {
		if mayor < mayorV[i] {
			mayor = mayorV[i]
		}
	}
	tend := time.Since(tinit)
	return mayor, "Tiempo de ejecucion concurrente: ", tend.Seconds()
}

// Recibe una matriz por parametro y calcula el mayor elemento
// de manera distribuida usando canales
func CalMayorCh(m [][]int) (int, string, float64) {
	tinit := time.Now()
	time.Sleep(1000)
	wg := sync.WaitGroup{}
	c := make(chan int, len(m))
	mayorV := make([]int, 1, len(m))
	mayor := 0
	for cont := 0; cont < len(m); cont++ {
		fila := m[cont]
		wg.Add(1)
		go func() {
			mayorTemp := 0
			for i := 0; i < len(fila); i++ {
				if mayorTemp < fila[i] {
					mayorTemp = fila[i]
				}
			}
			c <- mayorTemp
			mayorV = append(mayorV, <-c)
			wg.Done()
		}()
	}
	wg.Wait()
	for i := 0; i < len(mayorV); i++ {
		if mayor < mayorV[i] {
			mayor = mayorV[i]
		}
	}
	tend := time.Since(tinit)
	return mayor, "Tiempo de ejecucion concurrente con canales: ", tend.Seconds()
}
