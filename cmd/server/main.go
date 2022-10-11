package main

import (
	"desafio-goweb-francopesenda/cmd/server/handler"
	"desafio-goweb-francopesenda/internal/domain"
	"desafio-goweb-francopesenda/internal/tickets"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	list, err := LoadTicketsFromFile("./tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}
	//TICKETS
	repository := tickets.NewRepository(list)
	service := tickets.NewService(repository)
	tickets := handler.NewTicket(service)
	//ROUTER
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	// Rutas a desarollar:

	// GET - “/ticket/getByCountry/:dest”
	// GET - “/ticket/getAverage/:dest”
	routerTickets := r.Group("/ticket")
	{
		routerTickets.GET("/getByCountry/:dest", tickets.GetTicketsByCountry())
		routerTickets.GET("/getAverage/:dest", tickets.AverageDestination())
	}
	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
