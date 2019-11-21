package configs

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// this function for first load ENV file
func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// InitASCII this function for load in first time, for beauty command line, and link for make it
func InitASCII() {
	initEnv()
	fmt.Println(`
____________ _____   ___ _____ _____ _____      _____ ___________ _   _ _____ _____  _____ 
| ___ \ ___ \  _  | |_  |  ___/  __ \_   _|    / ___|  ___| ___ \ | | |_   _/  __ \|  ___|
| |_/ / |_/ / | | |   | | |__ | /  \/ | |______\ --.| |__ | |_/ / | | | | | | /  \/| |__  
|  __/|    /| | | |   | |  __|| |     | |______|--. \  __||    /| | | | | | | |    |  __| 
| |   | |\ \\ \_/ /\__/ / |___| \__/\ | |     /\__/ / |___| |\ \\ \_/ /_| |_| \__/\| |___ 
\_|   \_| \_|\___/\____/\____/ \____/ \_/     \____/\____/\_| \_|\___/ \___/ \____/\____/ `)
	// http://patorjk.com/software/taag/#p=display&f=Doom&t=PROJECT-SERVICE
}
