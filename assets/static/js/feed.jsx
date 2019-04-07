
// EVENT COMPONENT
class EventPost extends React.Component {
    render() {
        return (
            <div>
              <p>{this.props.currentEvent.Description}</p>
            </div>
        );
    }
}

// Generate some dummy data for testing the feeds view.
// We should eventually get this from the back-end database.
let dummyEvent1 = {
    ID: 123,
    Description: "Cleanup drive",
    // TODO: change start-time to a timezone-aware timestamp!
    // currently it's a int past the POSIX epoch, ignoring leap seconds.
    Starttime: 1554645600,
    Endtime: 1554649200,
    Location: "Cambridge, MA",
    Org: 283994,
    Posttime: 1554631200,
    Numlikes: 25,
};

let dummyEvent2 = {
    ID: 1234,
    Description: "Cleanup drive 3",
    Starttime: 1554645600,
    Endtime: 1554649200,
    Location: "Cambridge, MA",
    Org: 283994,
    Posttime: 1554631200,
    Numlikes: 29,
};

let dummyEvent3 = {
    ID: 12345,
    Description: "Cleanup drive 3",
    Starttime: 1554645600,
    Endtime: 1554649200,
    Location: "Cambridge, MA",
    Org: 283994,
    Posttime: 1554631200,
    Numlikes: 29,
};

ReactDOM.render(<EventPost currentEvent={dummyEvent1} />,
                document.getElementById('feed_container'));

let dummyEvents = [dummyEvent1, dummyEvent2, dummyEvent3];

// usr.ID = "1234"
// usr.Displayname = "user 1"
// usr.Icon = "<filepath>"
// usr.RadiusInMiles = 20
// usr.Location = "<city name>"
// 
// org.ID = "31215"
// org.Name = "org 1"
// org.Icon = "<filepath>"
// org.Location = "<city name>"
// org.Contacts = []string{usr.ID}
// org.Volunteers = []string{usr.ID}
// 
// evnt.ID = "123"
// evnt.Description = "a generic event decription"
// evnt.Starttime = "123456"
// evnt.Endtime = "674768458975"
// evnt.Location = "<place name>"
// evnt.Org = org.ID
// evnt.Posttime = "5673567223"
// evnt.Numlikes = 10
// 
// org.Events = []string{evnt.ID}
// usr.Orgs = []string{org.ID}


