package Server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserDetails struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties []map[string]string
}

func (self *ClientConnection) ValidateOnline() {
	url := "https://sessionserver.mojang.com/session/minecraft/hasJoined?username=" + self.Username + "&serverId=" + self.ServerHash

	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	details := new(UserDetails)
	err = json.Unmarshal(data, details)
	if err != nil {
		panic(err)
	}

	if details.ID == "" {
		panic("user could not be verified")
	}

	if len(details.ID) == 32 {
		self.UUID = details.ID[:8] + "-" + details.ID[8:12] + "-" + details.ID[12:16] + "-" + details.ID[16:20] + "-" + details.ID[24:]
	} else if len(details.ID) == 36 {
		self.UUID = details.ID
	} else {
		panic("invalid UUID from session server")
	}

	self.Username = details.Name
	//throw away properties for now

}
