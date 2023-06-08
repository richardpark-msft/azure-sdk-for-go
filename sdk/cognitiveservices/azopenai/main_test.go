package azopenai

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	endpoint string
	apiKey   string
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: failed to load env file: %s\n", err)
	}
	endpoint = os.Getenv("AOAI_ENDPOINT")
	apiKey = os.Getenv("AOAI_API_KEY")

	os.Exit(m.Run())
}
