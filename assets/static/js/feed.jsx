
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

let fetchPosts = (updateFcn) => {

    let remotePage = "api/v1/posts";

    let request = new XMLHttpRequest();
    request.onload = () => {
        console.log(request);
        let text = request.responseText;
        let data = JSON.parse(text);
        let newState = feedComponent.state.events;
        feedComponent.setState({events: newState});
    };
    let requestLocation = window.location.protocol + "//" + 
                          window.location.host + "/" + remotePage;
    request.open("POST", requestLocation);
    request.setRequestHeader("Content-Type", "application/json");
    let x = JSON.stringify({operation: "get", data: ""});
    console.log(x);
    request.send(x);
};



fetchPosts();

