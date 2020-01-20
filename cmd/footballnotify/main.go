package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/thang14/footballnotify/fire"
	"github.com/thang14/footballnotify/types"
)

func main() {

	fire := fire.New()
	events := types.Events{}
	for {
		newEvents, err := getEvents()
		if err != nil {
			log.Printf("err: %s", err)
			continue
		}
		msgs := events.GetNotificationMessages(newEvents)
		fire.SendMsgs(msgs)
		events = newEvents

		// push message
		time.Sleep(5 * time.Second)
	}

}

func getEvents() (types.Events, error) {
	startTime := time.Now().Format("2006-01-02")
	apiKey := os.Getenv("API_KEY")
	endpoint := fmt.Sprintf("http://apiv2.apifootball.com/?action=get_events&APIkey=%s&from=%s&to=%s", apiKey, startTime, startTime)
	log.Printf("get events: %s", endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	events := make(types.Events, 0)
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}
	return events, nil
}
