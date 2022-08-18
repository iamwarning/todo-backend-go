package main

import "os"

func main() {
	a := &App{}
	a.Initialize()
	port := getPort()
	a.Run(port)
}
func getPort() string {
	if _, ok := os.LookupEnv("PORT"); ok {
		return ":" + os.Getenv("PORT")
	}
	return ":8000"
}
