package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"cyber-shield/util"
)

const VERSION = "1.0.0"

type AuthOptions struct {
	Hostname string
	Port     int
	License  string
}

type License struct {
	AuthStatus bool   `json:"auth_status"`
	ID         string `json:"id"`
	Username   string `json:"username"`
	DiscordID  string `json:"discord_id"`
	Rank       string `json:"rank"`
}

func (auth AuthOptions) create() (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", auth.Hostname, auth.Port))

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (auth AuthOptions) getPayload() (string, error) {
	license := License{
		ID: auth.License,
	}

	json, err := util.Jsonify(license)

	if err != nil {
		return "", err
	}

	return json, nil
}

func (auth AuthOptions) Auth() (*License, error) {
	conn, err := auth.create()
	response := &License{}

	if err != nil {
		return response, err
	}

	defer conn.Close()

	payload, err := auth.getPayload()

	if err != nil {
		return response, err
	}

	if _, err := conn.Write([]byte(fmt.Sprintf("%s\n%s\n", VERSION, payload))); err != nil {
		return response, err
	}

	buffer := make([]byte, 1024)

	for {
		bytes, err := conn.Read(buffer)

		if err != nil {
			return response, err
		}

		data := string(buffer[:bytes])

		switch data {
		case "close":
			return response, errors.New("you have been disconnected")

		case "banned":
			return response, errors.New("you have been banned")

		default:
			if err := json.Unmarshal([]byte(data), response); err != nil {
				return response, err
			}

			return response, nil
		}
	}
}
