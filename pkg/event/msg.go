package event

type Event struct {
	ID          string   `json:id`
	Org         string   `json:org`
	Starttime   string   `json:startTime`
	Endtime     string   `json:endTime`
	Location    string   `json:location`
	Description string   `json:description`
	Numlikes    int64    `json:numLikes`
	Reports     []string `json:reports`
	Comments    []string `json:comments`
	Posttime    string   `json:postTime`
}
