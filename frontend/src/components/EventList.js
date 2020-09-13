import React from "react";
// import "./StatusSummary.css";

const services = [
  {
    id: 1,
    name: "Nexus IQ Data Services",
    description: "Nexus IQ Data Services",
  },
  {
    id: 2,
    name: "Repositories",
    description: "oss.sonatype.org, repository.sonatype.org",
  },
  {
    id: 3,
    name: "Sonatype Entity Services",
    description: "Corporate Site, Downloads, Blog, etc.",
  },
];

const events = {
  1: {
    status: "OK",
    description: "All systems nominal.",
    updated: "2020-09-09",
  },
  2: {
    status: "Degraded",
    description: "Maintenance in progress.",
    updated: "2020-09-09",
  },
  3: {
    status: "Down",
    description: "All systems nominal.",
    updated: "2020-09-09",
  },
};

const EventList = () => {
  // api call to get "services" list mocked above
  // api call to get latest event for each service ?

  const renderedList = services.map((service) => {
    let statusColor;
    let statusIcon;
    switch (events[service.id].status) {
      case "Down":
        statusColor = "has-text-danger";
        statusIcon = "fa-times-circle";
        break;
      case "Degraded":
        statusColor = "has-text-warning";
        statusIcon = "fa-exclamation-circle";
        break;
      default:
        statusColor = "has-text-success";
        statusIcon = "fa-check-circle";
    }

    return (
      <div className="box" key={service.id}>
        <span className={`icon ${statusColor}`}>
          <i className={`fas ${statusIcon}`}></i>
        </span>
        <span className="subtitle">{service.name}</span>
        <p>
          Description: {service.description}
          <br />
          Last Updated: {events[service.id].updated}
        </p>
      </div>
    );
  });

  return <div className="container is-fluid">{renderedList}</div>;
};

export default EventList;
