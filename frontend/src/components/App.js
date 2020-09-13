import React from "react";
import NavBar from "./NavBar";
import EventList from "./EventList";
import CentralStatus from "./CentralStatus";
import TwitterFeed from "./TwitterFeed";
// import "./App.css";

const App = () => {
  return (
    <div className="content">
      <NavBar />
      <div className="columns">
        <div className="column is-two-thirds">
          <CentralStatus />
          <EventList />
        </div>
        <div className="column">
          <TwitterFeed />
        </div>
      </div>
    </div>
  );
};

export default App;
