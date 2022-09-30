package main

import (
	"hackaton/internal/file"
	"hackaton/internal/service"
)

func main() {
	/* 	var tickets []service.Ticket
	 */ // Funcion para obtener tickets del archivo csv
	/* 	file.Read()
	 */

	/* juanTicket := service.Ticket{Id: 1001, Names: "Juan Sote", Email: "pipi12@gmail.com", Destination: "Miami", Date: "20:05", Price: 1500}
	 */
	rari := file.File{Path: "./tickets.csv"}

	data, err := rari.Read()
	if err != nil {
		panic(err)
	}

	ticks := service.NewBookings(data)

	/* ticks.Create(juanTicket)
	if err != nil {
		panic(err)
	}
	*/
	/* 	ticks.Read(1)
	 */
	/* 	ticks.Update(1001, juanTicket)
	 */ticks.Delete(3)
	/* fmt.Println(ticks) */
	/* ticks.Create(service.Ticket{}) */

}
