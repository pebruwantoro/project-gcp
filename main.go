package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func initConfig() {
	// Tell Viper to read from a file named ".env"
	viper.SetConfigFile(".env")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		os.Exit(1)
	}

	// Bind environment variables
	viper.AutomaticEnv()
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	version := viper.GetString("VERSION")
	fmt.Fprintf(w, "Hello, World!, Version: %s", version)
}

func main() {
	initConfig()

	port := viper.GetString("PORT")

	http.HandleFunc("/", helloWorldHandler)

	fmt.Printf("Server is running on http://localhost:%s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
