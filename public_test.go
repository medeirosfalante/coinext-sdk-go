package coinnext_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/coinnext-sdk-go"
)

func TestGetTicker(t *testing.T) {
	godotenv.Load()
	client := coinnext.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Public().OrderBook(nil)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}

}
