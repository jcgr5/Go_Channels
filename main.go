package main

import (
	"fmt"
	"github.com/jcgr5/Go_Channels/Maxcd"
	"github.com/jcgr5/Go_Channels/MayorMatriz"
	"github.com/jcgr5/tareaMatriz1/operaciones/generar"
)

func main() {
	fmt.Println(Maxcd.CalMaxcd(20, 5))
	fmt.Println(Maxcd.CalMaxcdCh(20, 5))
	ma := generar.CrearMatriz()
	fmt.Println(ma)
	fmt.Println(MayorMatriz.CalMayor(ma))
	fmt.Println(MayorMatriz.CalMayorDist(ma))
	fmt.Println(MayorMatriz.CalMayorCh(ma))
}
