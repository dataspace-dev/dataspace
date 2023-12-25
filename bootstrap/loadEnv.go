package bootstrap

import "github.com/joho/godotenv"

// LoadEnv loads the environment variables from a .env file.
// It uses the godotenv package to read the .env file and set the environment variables.
// If an error occurs while loading the .env file, it panics.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}