package service

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id          int    `csv:"ticket_id"`
	Names       string `csv:"ticket_names"`
	Email       string `csv:"ticket_email"`
	Destination string `csv:"ticket_destination"`
	Date        string `csv:"ticket_date"`
	Price       int    `csv:"ticket_price"`
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {

	arrBookings := b.Tickets
	arrBookings = append(arrBookings, Ticket{Id: t.Id, Names: t.Names, Email: t.Email, Destination: t.Destination, Date: t.Date, Price: t.Price})

	ticketStr := fmt.Sprintf("%d,%s,%s,%s,%s,%d", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
	/* ticketByt := []byte(ticketStr) */

	f, err := os.OpenFile("tickets.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintln(f, ticketStr)
	if err != nil {
		f.Close()
		panic(err)
	}
	err = f.Close()

	fmt.Println(arrBookings)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {

	ticks := b.Tickets
	var foundTicket Ticket

	for i := 0; i < len(ticks); i++ {
		if ticks[i].Id == id {
			foundTicket = ticks[i]
			break
		}

	}

	return foundTicket, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {

	allTickets := b.Tickets
	ticks := Ticket{}

	for i, data := range allTickets {
		if data.Id == t.Id {
			ticks = Ticket{t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price}
			allTickets[i] = Ticket{t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price}
		}
	}

	f, err := os.OpenFile("tickets.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("tickets.csv", []byte(""), 0644)

	for _, tickets := range allTickets {

		ticketString := fmt.Sprintf("%d,%s,%s,%s,%s,%d", tickets.Id, tickets.Names, tickets.Email, tickets.Destination, tickets.Date, tickets.Price)

		_, err = fmt.Fprintln(f, ticketString)
		if err != nil {
			f.Close()
			panic(err)
		}

	}
	err = f.Close()
	return ticks, nil
}

func (b *bookings) Delete(id int) (int, error) {
	allTickets := b.Tickets

	allTickets = append(allTickets[:id-1], allTickets[id:]...)

	f, err := os.OpenFile("tickets.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("tickets.csv", []byte(""), 0644)

	for _, tickets := range allTickets {

		ticketString := fmt.Sprintf("%d,%s,%s,%s,%s,%d", tickets.Id, tickets.Names, tickets.Email, tickets.Destination, tickets.Date, tickets.Price)

		_, err = fmt.Fprintln(f, ticketString)
		if err != nil {
			f.Close()
			panic(err)
		}

	}
	err = f.Close()

	return id, nil
}
