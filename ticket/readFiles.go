package ticket

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/* My structures */

type Tickets []Ticket

type Ticket struct {
	id                 string
	name               string
	email              string
	destinationCountry string
	flightTime         time.Time
	price              float64
}

/* Methods */
func (t Tickets) GetLenght() int {
	return len(t)
}

// Requerimiento 3. Calcular el promedio de personas que viajan a un país determinado en un día.
func (t Tickets) GetPeopleAverageByDestination(destination string) float64 {
	var count float64

	if t == nil {
		log.Fatal("I cannot execute the method GetPeopleAverageByDestination(), because the data is null")
	}

	for _, ticket := range t {
		if strings.ToLower(destination) == strings.ToLower(ticket.destinationCountry) {
			count++
		}
	}

	average := count / float64(len(t))

	fmt.Printf("La cantidad promedio de personas que viajan a %s es de %f", destination, average)
	fmt.Println("")
	return average
}

/* Functions */

// Leer el archivo csv y retornar la data
func ReadFile(path string) (Tickets, error) {
	allTickets := Tickets{}
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("The file %s not found", path)
	}

	dataOk := strings.TrimSpace(string(data))
	ticketsByRow := strings.Split(dataOk, "\n")

	for _, ticket := range ticketsByRow {
		row := strings.Split(ticket, ",")

		flightTime, errParse := time.Parse("15:04", row[4])
		if errParse != nil {
			log.Fatal(errParse)
		}

		price, errParse := strconv.ParseFloat(row[5], 64)
		if errParse != nil {
			log.Fatal(errParse)
		}

		newTicket := Ticket{
			id:                 row[0],
			name:               row[1],
			email:              row[2],
			destinationCountry: row[3],
			flightTime:         flightTime,
			price:              price,
		}

		allTickets = append(allTickets, newTicket)

	}

	return allTickets, nil
}
