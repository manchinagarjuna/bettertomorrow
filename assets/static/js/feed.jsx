
// Function that returns a date string.
let convertDateToString = (date) => {
    let dateObj = new Date(date);
    return [dateObj.toLocaleDateString(), dateObj.toLocaleTimeString()];
};

// A React component for each event that shows up on the feed.
class EventPost extends React.Component {
    render() {
        return (
        <div class="card">
            <div class="user-info">
                <img class="user-thumb" src="images/headshot.JPG"/>
                <p class="user-id">Marly H.</p>
            </div>
            <div class="post-data">
                <p class="date">April 6, 2019</p>
                <p class="location">Cambridge, MA</p>
            </div>
            <div class="post-content">
                <h2 class="post-title">Title Goes Here</h2>
                <p class="post-text">Content goes here... Lorem, ipsum dolor sit amet consectetur adipisicing elit. Aspernatur accusantium aperiam impedit, officia reprehenderit labore eveniet suscipit et obcaecati quisquam sequi similique in hic? Ut vero omnis accusantium quas temporibus?</p>
                <img class="post-img" src="images/event-image.jpg"/>
            </div>
            <div class="social-icon">
                <div class="likes">
                    <i class="fab fa-gratipay"></i>
                    <p class="counter">32</p>
                </div>
                <div class="comments">
                    <i class="fas fa-comments"></i>
                    <p class="counter">15</p>
                </div>
            </div>
        </div>
        );
    }
}

// A React component for the feed itself.
class Feed extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            events: []
        };
    }

    updateLikes(index) {
        const events = this.state.events.slice();
        events[index].Numlikes += 1;
        this.setState({events: events});
        
        console.log("Added like");        
    }

    updateSchedule(index) {
        console.log("Adding event to schedule");
        // Todo
    }

    renderEventPost(index) {
        return (
            <EventPost currentEvent={this.state.events[index]}
                       key={this.state.events[index].ID}
                       onLikeClick={() => this.updateLikes(index)}
                       onScheduleClick={() => this.updateSchedule(index)} />
        );
    }
    
    render() {
        // Display a message if no events are available.
        if (this.state.events.length === 0) {
            return (<div>No events to display</div>);
        } else {
            // Return a list of divs corresponding to each post.
            return (
                <div>
                  {
                      // Map over the events and render each event post.
                      this.state.events.map(
                          (val, index) => this.renderEventPost(index))
                  }
                </div>
            );
        }
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
    Description: "Another cleanup drive",
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

let feedComponent =
    ReactDOM.render(<Feed />, document.getElementById("feed_container"));

// Add dummy events and re-render.
let dummyEvents = [dummyEvent1, dummyEvent2, dummyEvent3];
feedComponent.setState({events: dummyEvents});

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

