package main

import (
    "github.com/joho/godotenv"
    "gopkg.in/gookit/color.v1"
)

func main() {
    godotenv.Load()
    color.Red.Println("Red color")
}
