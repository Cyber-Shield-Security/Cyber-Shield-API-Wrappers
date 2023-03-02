package main

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	apiKey      = "4IqF7qAxP9JObVr5R16XOxpKePKPMAFpaZ1gft7Ez75IVp1H"
	appVersion  = "APPLICATION_NAME_AND_VERSION"
	csBackend   = "15.204.58.133"
	serverPort  = 17234
)

type licenseID struct {
	AuthStatus bool   `json:"auth_status"`
	ID         string `json:"id"`
	Username   string `json:"username"`
	DiscordID  string `json:"discord_id"`
	Rank       string `json:"rank"`
}

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", csBackend, serverPort))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	licID := &licenseID{
		ID: apiKey,
	}

	message := fmt.Sprintf("%s\n%s\n", appVersion, toJSON(licID))
	if _, err := conn.Write([]byte(message)); err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		data := buf[:n]
		message := string(data)
		switch message {
		case "close":
			fmt.Println("[!] Error, The owner has disconnected you from using Cyber Shield...")
			return
		case "banned":
			fmt.Println("[!] Error, You have been banned from using Cyber Shield...")
			return
		default:
			if err := json.Unmarshal(data, licID); err != nil {
				panic(err)
			}
			fmt.Printf("Successfully authed! %s\n", licID.Rank)
			// start sending or receiving data with server here
		}
	}
}

func toJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}


// github.com/9xN
