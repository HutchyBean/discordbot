package main

import (
	"fmt"
	"github.com/HutchyBean/discordbot/Commands"
	"github.com/HutchyBean/discordbot/DCH"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)
func init() {
	godotenv.Load()
}

func main() {

	 ds, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	 if err != nil {
	 	log.Fatalf("Could not start the discord bot: %v\n", err)
	 }

	 err = ds.Open()
	 if err != nil {
	 	log.Fatalf("Could not open the discord bot: %v\n", err)
	 }
	cHandler := DCH.Init("^", ds)
	Commands.Load(cHandler)
	if err != nil {
		log.Printf("WARN: %v\n", err)
	}

	fmt.Println("Am loaded")
	select {}
}
