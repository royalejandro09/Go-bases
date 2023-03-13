package ticket

import (
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
		if strings.ToLower(destination) == strings.ToLower(ticket.destinationCountry) {
			count++
		}
	}

	//fmt.Printf("La cantidad de tickets vendidos para %s es de %d", destination, count)
	//fmt.Println("")
	return count, nil
}

// 2. Cantidad de personas que viajan entre los rangos de horas()
// madrugada ("00:00", "06:00"),mañana ("07:00", "12:00"), tarde ("13:00", "19:00"), y noche ("20:00", "23:00").
func GetNumberOfPeopleTravelingBySchedule(data Tickets, rangeTime [2]string) (int, error) {
	count := 0

	if data.GetLenght() == 0 {
		log.Fatal("I cannot execute the method GetNumberOfPeopleTravelingBySchedule(), because the data is null")
	}

	startTime, _ := time.Parse("15:04", rangeTime[0])
	endTime, _ := time.Parse("15:04", rangeTime[1])

	for _, ticket := range data {
		if ticket.flightTime.After(startTime) && ticket.flightTime.Before(endTime) {
			count++
		}
	}
	//fmt.Printf("La cantidad de personas que viajan entre las %s & %s horas son: %d", rangeTime[0], rangeTime[1], count)
	//fmt.Println("")
	return count, nil
}
