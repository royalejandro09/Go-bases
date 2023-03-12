package main

import (
	"desafio/ticket"
	"fmt"
)

func main() {

	// Leyendo el archivo .csv
	data, error := ticket.ReadFile("./tickets.csv")

	//Manejando el error
	if error != nil {
		fmt.Println(error)
	}
	
	fmt.Println(data.GetLenght())

	// Requerimiento 1. Obtener la cantidad de tickets por destino.
	ticket.GetTicketByDestinationCount("Colombia", data)

	//Requerimiento 2. Cantidad de personas que viajan entre los rangos de horas()
	// madrugada ("00:00", "06:00"),mañana ("07:00", "12:00"), tarde ("13:00", "19:00"), y noche ("20:00", "23:00").s
	ticket.GetTickestByEarlyMorningTrip(data, "00:00", "06:00")

	//Requerimiento 3.
}
