package ticket

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

/* My structures */

type Tickets []Ticket

type Ticket struct {
	id          string
	nombre      string
	email       string
	paisDestino string
	horaVuelo   time.Time
	precio      float64
}

/* Methods */
func (t Tickets) GetLenght() int {
	return len(t)
}

/* Functions */

// Leer el archivo csv y retornar la data
func ReadFile(path string) (Tickets, error) {
	allTickets := Tickets{}
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, errors.New("File not found")
	}

	dataOk := strings.TrimSpace(string(data))
	ticketsByRow := strings.Split(dataOk, "\n")

	for _, ticket := range ticketsByRow {
		row := strings.Split(ticket, ",")

		horaV, _ := time.Parse("15:04", row[4])
		price, _ := strconv.ParseFloat(row[5], 64)

		newTicket := Ticket{id: row[0], nombre: row[1], email: row[2], paisDestino: row[3], horaVuelo: horaV, precio: price}
		allTickets = append(allTickets, newTicket)
		//fmt.Println(row)
		//fmt.Println(ticket)
	}

	return allTickets, nil
}
