package main

import (
	"encoding/json"
	"fmt"

	"github.com/manchinagarjuna/bettertomorrow/pkg/event"
	"github.com/manchinagarjuna/bettertomorrow/pkg/organization"
	"github.com/manchinagarjuna/bettertomorrow/pkg/user"
)

func main() {
	var org organization.Organization
	var usr user.User
	var evnt event.Event

	usr.ID = "1234"
	usr.Displayname = "user 1"
	usr.Icon = "<filepath>"
	usr.RadiusInMiles = 20
	usr.Location = "<city name>"

	org.ID = "31215"
	org.Name = "org 1"
	org.Icon = "<filepath>"
	org.Location = "<city name>"
	org.Contacts = []string{usr.ID}
	org.Volunteers = []string{usr.ID}

	evnt.ID = "123"
	evnt.Description = "a generic event decription"
	evnt.Starttime = "123456"
	evnt.Endtime = "674768458975"
	evnt.Location = "<place name>"
	evnt.Org = org.ID
	evnt.Posttime = "5673567223"
	evnt.Numlikes = 10

	org.Events = []string{evnt.ID}
	usr.Orgs = []string{org.ID}

	res, _ := json.Marshal(org)
	fmt.Println(string(res))

	fmt.Println("------------------------------------")

	res, _ = json.Marshal(usr)
	fmt.Println(string(res))

	fmt.Println("------------------------------------")

	res, _ = json.Marshal(evnt)
	fmt.Println(string(res))
}
