package ticket

import (
	"fmt"
	"strings"
	"time"
)

// 1.Obtener la cantidad de tickets por destino.
func GetTicketByDestinationCount(destination string, data Tickets) int {
	count := 0

	for _, ticket := range data {
		if ticket.paisDestino == destination {
			count++
		}
		//fmt.Println(row)
		//fmt.Println(ticket)
	}
	fmt.Printf("La cantidad de tickets vendidos para %s es de %d", destination, count)
	fmt.Println("")
	return count
}

// 2. Cantidad de personas que viajan entre los rangos de horas()
// madrugada (0 → 6),mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).
func GetTickestByEarlyMorningTrip(data Tickets, start string, end string) int {
	count := 0
	startTime, _ := time.Parse("15:04", start)
	endTime, _ := time.Parse("15:04", end)

	for _, ticket := range data {
		if ticket.horaVuelo.After(startTime) && ticket.horaVuelo.Before(endTime) {
			count++
		}
	}
	fmt.Printf("La cantidad de personas que viajan entre las %s & %s horas son: %d", start, end, count)
	fmt.Println("")
	return count
}

func GetTickestByEarlyMorningTrip1(data string) int {
	startTime, _ := time.Parse("15:04", "0:00")
	endTime, _ := time.Parse("15:04", "6:00")
	count := 0

	rowTickets := strings.Split(data, "\n")

	for _, ticket := range rowTickets {
		row := strings.Split(ticket, ",")

		timeTicket, _ := time.Parse("15:04", row[5])

		if len(row) == 1 {
			continue
		} else if timeTicket.After(startTime) && timeTicket.Before(endTime) {
			count++
		}
		//fmt.Println(row)
		//fmt.Println(ticket)
	}
	return count
}
