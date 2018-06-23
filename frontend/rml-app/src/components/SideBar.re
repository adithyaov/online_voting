let component = ReasonReact.statelessComponent("SideBar");

let divider = dataContent =>
  ReasonReact.cloneElement(
    <li className="divider" />,
    ~props={"data-content": dataContent},
    [||],
  );

let make = _children => {
  ...component,
  render: _self =>
    <div className="side-bar">
      <div className="columns">
        <div className="column col-8">
          <div className="logo rounded">
            <span className="text-bold"> (ReasonReact.string("PEN")) </span>
            (ReasonReact.string("BALLOT"))
          </div>
        </div>
        <div className="column col-4">
          <button className="btn btn-block">
            (ReasonReact.string("Bye?"))
          </button>
        </div>
      </div>
      <ul className="nav">
        (divider("ADITHYA KUMAR"))
        <li className="nav-item">
          <a href="#"> (ReasonReact.string("Profile")) </a>
        </li>
        <li className="nav-item">
          <a href="#"> (ReasonReact.string("Nominations")) </a>
        </li>
        (divider("SOCKET"))
        <li className="nav-item">
          <a href="#"> (ReasonReact.string("Online Help")) </a>
        </li>
        <li className="nav-item">
          (
            ReasonReact.cloneElement(
              <a className="badge" />,
              ~props={"data-badge": 6},
              [|ReasonReact.string("Announcements")|],
            )
          )
        </li>
        (divider("DEV"))
        <li className="nav-item">
          <a href="#"> (ReasonReact.string("Documentation")) </a>
        </li>
        <li className="nav-item">
          <a href="#"> (ReasonReact.string("Github")) </a>
        </li>
        <li className="nav-item">
          <a href="#"> (ReasonReact.string("Contributers")) </a>
        </li>
        <li className="nav-item">
          <a href="#"> (ReasonReact.string("Guidelines")) </a>
        </li>
        (divider("LEGAL"))
        <li className="nav-item">
          <a href="#"> (ReasonReact.string("Terms of usage")) </a>
        </li>
      </ul>
    </div>,
};