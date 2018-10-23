package tipboardsubscriber

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)



//SimpleSubscriber - A struct to maintain a single 'just-value' tile on the MQTT broker tipboard
type SimpleSubscriber struct {
	Dashboard   TipboardDash //dashboard to push the updates to
	TileKey     string       //key to identify the tile to push to
	Title       string       //title to be pushed in the tile data
	Description string       //description text to push to the tile
	Value       int          //the value to be pushed to the tile
}

//curl http://localhost:7272/api/v0.1/13446d45c9544d4da0b981bd946de743/push -X POST -d "tile=just_value" -d "key=clients_total" -d 'data={"title": "Total:", "description": "(Total number of registered clients)", "just-value": "3"}'
func updateTile(s *SimpleSubscriber, value int) (ok bool) {
	fmt.Println("pushing update to tile")
	//push the value to the dashboard
	urlStr := "http://" + s.Dashboard.DashHost + ":" + strconv.Itoa(s.Dashboard.DashPort) + "/api/v0.1/" + s.Dashboard.DashAPIKey + "/push"
	bodyStr := "tile=just_value&key=" + s.TileKey + "&data={\"title\": \"" + s.Title + ":\", \"description\": \"" + s.Description + "\", \"just-value\": \"" + strconv.Itoa(s.Value) + "\"}"
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(bodyStr))
	if err != nil {
		// handle err
		fmt.Println("balls")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Println("feck")
	}
	defer resp.Body.Close()
	return
}

//define a function for the default message handler

var msgRcvd MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
	//updateTile(this, 100)

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    log.SetOutput(file)
    log.Print("Logging to a file in Go!")

}

//Subscribe to the designated broker and topic
//when a message is received, a single integer value is extracted and pushed to the tile as an update.
//func (s *SimpleSubscriber) Subscribe(brokerHost string, brokerPort int, topic string) (ok bool) {
func (s *SimpleSubscriber) Subscribe(c MQTT.Client, topic string) (ok bool) {
	//subscribe to the specified broker and topic

	/*//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.1.6:1883")
	opts.SetClientID("go-simple")
	//opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}*/

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe("test_topic/#", 0, msgRcvd); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Println("subscribed")

	time.Sleep(10 * time.Second)

	/*
			opts := mqtt.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883").SetClientID("sample")
			opts.SetProtocolVersion(4)

			c := mqtt.NewClient(opts)
			if token := c.Connect(); token.Wait() && token.Error() != nil {
			    panic(token.Error())
			}

			var msgRcvd := func(client *mqtt.Client, message mqtt.Message) {
		    	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
			}

			if token := c.Subscribe("example/topic", 0, msgRcvd); token.Wait() && token.Error() != nil {
			    fmt.Println(token.Error())
			}
	*/
	/*//push the value to the dashboard
	urlStr := "http://" + s.Dashboard.DashHost + ":" + strconv.Itoa(s.Dashboard.DashPort) + "/api/v0.1/" + s.Dashboard.DashAPIKey + "/push"
	bodyStr := "tile=just_value&key=" + s.TileKey + "&data={\"title\": \"" + s.Title + ":\", \"description\": \"" + s.Description + "\", \"just-value\": \"" + strconv.Itoa(s.Value) + "\"}"
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(bodyStr))
	if err != nil {
		// handle err
		fmt.Println("balls")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Println("feck")
	}
	defer resp.Body.Close()
	return*/
	//updateTile(s, 100)
	return
}

/*

apiUrl := "https://api.com"
    resource := "/user/"
    data := url.Values{}
    data.Set("name", "foo")
    data.Add("surname", "bar")

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := u.String() // 'https://api.com/user/'

    client := &http.Client{}
    r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
    r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    resp, _ := client.Do(r)
    fmt.Println(resp.Status)

*/
