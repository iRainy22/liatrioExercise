package main

//imports needed libraries
import (
    "log"
	"os"
    "time"
    "github.com/gofiber/fiber/v3"
)

//Uses a struct to improve memory efficiency
type Response struct {
  Message   string `json:"message"`
  Timestamp int64  `json:"timestamp"`
}


func main() {

	//Creates a new fiber instance named app
    app := fiber.New()

	//tells the app to display the JSON message at the address /
	app.Get("/", func (c fiber.Ctx) error {
        return c.JSON(Response{
			Message:   "My name is Billy Bob Sr.",
			Timestamp: time.Now().UnixMilli(),
		})
    })
	
	//reads an ENV variable which is saved in a file on our computer
	//if it finds one, it uses that port, otherwise it uses 8080
	//which our docker also uses
	port:= os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	//Runs the app on the designated port
	log.Fatal(app.Listen(":" + port))
}
