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

import _regeneratorRuntime from "babel-runtime/regenerator";

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _asyncToGenerator(fn) { return function () { var gen = fn.apply(this, arguments); return new Promise(function (resolve, reject) { function step(key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { return Promise.resolve(value).then(function (value) { step("next", value); }, function (err) { step("throw", err); }); } } return step("next"); }); }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var Status = function (_React$Component) {
  _inherits(Status, _React$Component);

  function Status(props) {
    _classCallCheck(this, Status);

    var _this = _possibleConstructorReturn(this, (Status.__proto__ || Object.getPrototypeOf(Status)).call(this, props));

    _this.getStatus = _this.getStatus.bind(_this);
    return _this;
  }

  _createClass(Status, [{
    key: "getStatus",
    value: function () {
      var _ref = _asyncToGenerator( /*#__PURE__*/_regeneratorRuntime.mark(function _callee() {
        var request, response;
        return _regeneratorRuntime.wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                request = new Request("/api/events");
                _context.prev = 1;
                _context.next = 4;
                return fetch(request);

              case 4:
                response = _context.sent;

                console.log(response);
                _context.next = 11;
                break;

              case 8:
                _context.prev = 8;
                _context.t0 = _context["catch"](1);
                throw new Error("Something went wrong on api server!");

              case 11:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, this, [[1, 8]]);
      }));

      function getStatus() {
        return _ref.apply(this, arguments);
      }

      return getStatus;
    }()
  }, {
    key: "render",
    value: function render() {
      return React.createElement(
        "div",
        null,
        React.createElement(
          "p",
          null,
          "Status"
        ),
        React.createElement(
          "button",
          { onClick: this.getStatus },
          "click"
        )
      );
    }
  }]);

  return Status;
}(React.Component);

var TwitterFeed = function (_React$Component2) {
  _inherits(TwitterFeed, _React$Component2);

  function TwitterFeed() {
    _classCallCheck(this, TwitterFeed);

    return _possibleConstructorReturn(this, (TwitterFeed.__proto__ || Object.getPrototypeOf(TwitterFeed)).apply(this, arguments));
  }

  _createClass(TwitterFeed, [{
    key: "render",
    value: function render() {
      return React.createElement(
        "p",
        null,
        "Twitter"
      );
    }
  }]);

  return TwitterFeed;
}(React.Component);

ReactDOM.render(React.createElement(Status, null), document.getElementById("Status"));
ReactDOM.render(React.createElement(TwitterFeed, null), document.getElementById("TwitterFeed"));