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

const events = [
  {
    service: 1,
    status: "OK",
    description: "All systems nominal.",
    updated: "2020-09-09",
  },
  {
    service: 2,
    status: "Degraded",
    description: "Maintenance in progress.",
    updated: "2020-09-09",
  },
  {
    service: 3,
    status: "OK",
    description: "All systems nominal.",
    updated: "2020-09-09",
  },
];

const EventList = () => {
  const renderedList = services.map((service) => {
    console.log(service);
    return (
      <div className="container is-fluid" key={service.id}>
        <span className="icon has-text-success">
          <i className="fas fa-check-square"></i>
        </span>
        <h1>{service.name}</h1>
      </div>
    );
  });

  return <section className="section">{renderedList}</section>;
};

export default EventList;
