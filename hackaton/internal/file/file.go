package file

import (
	"hackaton/internal/service"
	"os"

	"github.com/gocarina/gocsv"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {

	ticketsFile, err := os.OpenFile(f.Path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer ticketsFile.Close()

	tickets := []*service.Ticket{}
	newtick := []service.Ticket{}

	if err := gocsv.UnmarshalFile(ticketsFile, &tickets); err != nil { // Load clients from file
		panic(err)
	}
	for _, ticket := range tickets {
		newtick = append(newtick, service.Ticket{Id: ticket.Id, Names: ticket.Names, Email: ticket.Email, Destination: ticket.Destination, Date: ticket.Date, Price: ticket.Price})
	}

	return newtick, nil

}

func (f *File) Write(service.Ticket) error {
	return nil
}
