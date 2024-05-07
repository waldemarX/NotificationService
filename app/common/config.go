package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)


func Config(key string) string {
    env_error := godotenv.Load(".env")
    if env_error != nil {
        fmt.Print("Error loading .env file")
    }
    return os.Getenv(key)
}
