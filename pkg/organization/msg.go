package organization

type Organization struct {
	ID         string   `json:id`
	Name       string   `json:name`
	Location   string   `json:location`
	Contacts   []string `json:contacts`
	Volunteers []string `json:volunteers`
	Icon       string   `json:icon`
	Events     []string `json:events`
}
