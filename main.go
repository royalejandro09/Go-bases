package main

import (
	"desafio/ticket"
	"fmt"
)

func main() {

	fmt.Println("¡START OF EXECUTION!")

	// Leyendo el archivo .csv y obtengo la lista de tickets vendidos.
	data, error := ticket.ReadFile("./tickets.csv")

	//Manejando el error
	if error != nil {
		fmt.Println(error)
	}

	// Requerimiento 1. Obtener la cantidad de tickets por destino.
	ticket.GetTicketByDestinationCount("Colombia", data)

	//Requerimiento 2. Cantidad de personas que viajan entre los rangos de horas()
	// madrugada ("00:00", "06:00"),mañana ("07:00", "12:00"), tarde ("13:00", "19:00"), y noche ("20:00", "23:00").
	ticket.GetNumberOfPeopleTravelingBySchedule(data, "00:00", "06:00")

	//Requerimiento 3. Calcular el promedio de personas que viajan a un país determinado en un día.
	data.GetPeopleAverageByDestination("Colombia")

	fmt.Println("END THE EXECUTION!")
}
