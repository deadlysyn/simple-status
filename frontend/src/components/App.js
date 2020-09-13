import React from "react";
import TwitterFeed from "./TwitterFeed";
// import StatusSummary from "./StatusSummary";
import EventList from "./EventList";
// import "./App.css";

const App = () => {
  return (
    <div className="columns">
      <div className="column is-two-thirds">
        <EventList />
      </div>
      <div className="column">
        <TwitterFeed />
        {/* <StatusSummary /> */}
      </div>
    </div>
  );
};

export default App;
