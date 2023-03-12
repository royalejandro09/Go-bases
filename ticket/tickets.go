package ticket

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// 1.Obtener la cantidad de tickets por destino.
func GetTicketByDestinationCount(destination string, data Tickets) (int, error) {
	count := 0

	if data.GetLenght() == 0 {
		log.Fatal("I cannot execute the method GetTicketByDestinationCount(), because the data is null")
	}

	for _, ticket := range data {
		if strings.ToLower(destination) == strings.ToLower(ticket.paisDestino) {
			count++
		}
	}

	fmt.Printf("La cantidad de tickets vendidos para %s es de %d", destination, count)
	fmt.Println("")
	return count, nil
}

// 2. Cantidad de personas que viajan entre los rangos de horas()
// madrugada ("00:00", "06:00"),mañana ("07:00", "12:00"), tarde ("13:00", "19:00"), y noche ("20:00", "23:00").
func GetNumberOfPeopleTravelingBySchedule(data Tickets, start string, end string) (int, error) {
	count := 0

	if data.GetLenght() == 0 {
		log.Fatal("I cannot execute the method GetNumberOfPeopleTravelingBySchedule(), because the data is null")
	}

	startTime, _ := time.Parse("15:04", start)
	endTime, _ := time.Parse("15:04", end)

	for _, ticket := range data {
		if ticket.horaVuelo.After(startTime) && ticket.horaVuelo.Before(endTime) {
			count++
		}
	}
	fmt.Printf("La cantidad de personas que viajan entre las %s & %s horas son: %d", start, end, count)
	fmt.Println("")
	return count, nil
}
