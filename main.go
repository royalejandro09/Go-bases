package main

import (
	"desafio/ticket"
	"fmt"
)

// Variables const
var madrugada = [2]string{"00:00", "06:00"}
var mañana = [2]string{"07:00", "12:00"}
var tarde = [2]string{"13:00", "19:00"}
var noche = [2]string{"20:00", "23:00"}

func main() {

	fmt.Println("¡START OF EXECUTION!")

	// Leyendo el archivo .csv y obtengo la lista de tickets vendidos.
	data, error := ticket.ReadFile("./tickets.csv")

	//Manejando el error
	if error != nil {
		fmt.Println(error)
	}

	// Requerimiento 1. Obtener la cantidad de tickets por destino.
	count, _ := ticket.GetTicketByDestinationCount("Colombia", data)
	fmt.Printf("La cantidad de tickets vendidos es de %d", count)
	fmt.Println("")

	//Requerimiento 2. Cantidad de personas que viajan entre los rangos de horas()
	// madrugada ("00:00", "06:00"),mañana ("07:00", "12:00"), tarde ("13:00", "19:00"), y noche ("20:00", "23:00").
	cantMadrugada, _ := ticket.GetNumberOfPeopleTravelingBySchedule(data, madrugada)
	fmt.Printf("La cantidad de personas que viajan son %d", cantMadrugada)
	fmt.Println("")

	cantMañana, _ := ticket.GetNumberOfPeopleTravelingBySchedule(data, mañana)
	fmt.Printf("La cantidad de personas que viajan son %d", cantMañana)
	fmt.Println("")

	cantTarde, _ := ticket.GetNumberOfPeopleTravelingBySchedule(data, tarde)
	fmt.Printf("La cantidad de personas que viajan son %d", cantTarde)
	fmt.Println("")

	cantNoche, _ := ticket.GetNumberOfPeopleTravelingBySchedule(data, noche)
	fmt.Printf("La cantidad de personas que viajan son %d", cantNoche)
	fmt.Println("")

	//Requerimiento 3. Calcular el promedio de personas que viajan a un país determinado en un día.
	average := data.GetPeopleAverageByDestination("Colombia")
	fmt.Printf("La cantidad promedio de personas que viajan es de %f", average)
	fmt.Println("")

	fmt.Println("END THE EXECUTION!")
}
