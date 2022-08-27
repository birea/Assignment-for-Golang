package server

import (
	"fmt"
	"log"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
)

//-------------------- REST server based on Gin
type Server struct {
	router *fiber.App
}



// ------------new server that takes advantage of zerolog for logging
// ---------------------app configuration
func NewServer() *Server {
	r := fiber.New()

	r.Get("/flight_points", FlightPoints)

	return &Server{
		router: r,
	}
}

// --------------------------------Start REST
// -------------------------------switch
func (s *Server) Start(port int) {
	err := s.router.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		// -----next port
		if strings.Contains(err.Error(), "address already in use") {
			fmt.Println("")
			log.Printf("PORT ALREADY IN USE::%d", port)
			port++
			log.Printf("TRYING NEXT PORT:%d\n", port)
			s.Start(port)
		} else {
			panic(err)
		}
	}
}
