import React from "react";
import EventList from "./EventList";
import CentralStatus from "./CentralStatus";
import TwitterFeed from "./TwitterFeed";
// import "./App.css";

const App = () => {
  return (
    <div className="columns">
      <div className="column is-two-thirds">
        <section className="section">
          <CentralStatus />
          <EventList />
        </section>
      </div>
      <div className="column">
        <section className="section">
          <TwitterFeed />
        </section>
      </div>
    </div>
  );
};

export default App;
