import React from "react";
// import "./StatusSummary.css";

const StatusSummary = () => {
  return (
    <section className="section">
      <div className="container is-fluid">
        <span className="icon has-text-success">
          <i className="fas fa-check-square"></i>
        </span>
        <span className="icon has-text-warning">
          <i className="fas fa-exclamation-triangle"></i>
        </span>
        Status Summary
      </div>
    </section>
  );
};

export default StatusSummary;
