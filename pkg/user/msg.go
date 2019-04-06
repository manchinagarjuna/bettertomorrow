package user

type User struct {
	ID            string   `json:id`
	Displayname   string   `json:displayName`
	Icon          string   `json:icon`
	Location      string   `json:location`
	RadiusInMiles uint16   `json:radiusInMiles`
	Orgs          []string `json:orgs`
}
