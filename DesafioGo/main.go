package main

import (
	"fmt"
	"log"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	//Requerimiento 1
	totalTickets, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("El total de tiquetes con destino Brazil:", totalTickets)

	//Requerimiento 2
	totalByPeriod, err := tickets.GetCountByPeriod("noche")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El total de viajes realizados en la noche son:", totalByPeriod)

	//Requerimiento 3
	countRecords, err := tickets.LenData()
	if err != nil {
		log.Fatal(err)
	} else {
		average, err := tickets.AverageDestination("Brazil", countRecords)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("El número promedio de personas que viajan a Brazil es:", average)

		}
	}

}
