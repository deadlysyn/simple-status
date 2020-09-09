"use strict";

// function WarningBanner(props) {
//   if (!props.warn) {
//     return null;
//   }
//   return <div className="warning">Warning!</div>;
// }

// class Page extends React.Component {
//   constructor(props) {
//     super(props);
//     this.state = { showWarning: true };
//     this.handleToggleClick = this.handleToggleClick.bind(this);
//   }

//   handleToggleClick() {
//     this.setState((state) => ({
//       showWarning: !state.showWarning,
//     }));
//   }

//   render() {
//     return (
//       <div>
//         <WarningBanner warn={this.state.showWarning} />
//         <button onClick={this.handleToggleClick}>
//           {this.state.showWarning ? "Hide" : "Show"}
//         </button>
//       </div>
//     );
//   }
// }

// ReactDOM.render(<Page />, document.getElementById("root"));

class Status extends React.Component {
  constructor(props) {
    super(props);
    this.getStatus = this.getStatus.bind(this);
  }

  async getStatus() {
    const request = new Request("/api/events");
    try {
      const response = await fetch(request);
      console.log(response);
    } catch (e) {
      throw new Error("Something went wrong on api server!");
    }
  }

  render() {
    return (
      <div>
        <p>Status</p>
        <button onClick={this.getStatus}>click</button>
      </div>
    );
  }
}

class TwitterFeed extends React.Component {
  render() {
    return <p>Twitter</p>;
  }
}

ReactDOM.render(<Status />, document.getElementById("Status"));
ReactDOM.render(<TwitterFeed />, document.getElementById("TwitterFeed"));
