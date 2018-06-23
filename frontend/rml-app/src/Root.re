let component = ReasonReact.statelessComponent("Root");

let divider = dataContent =>
  ReasonReact.cloneElement(
    <li className="divider" />,
    ~props={"data-content": dataContent},
    [||],
  );

let make = _children => {
  ...component,
  render: _self =>
    <div className="wrapper">
      <div className="side-bar-wrapper"> <SideBar /> </div>
      <div className="content-wrapper">
        <div className="content">
          <ContentHeading />
          <div className="divider-wrapper"> (divider("BALLOT INFO")) </div>
          <BallotDetails />
        </div>
      </div>
      <div className="ballot-bar-wrapper"> <BallotBar /> </div>
      <div className="cmd-wrapper"> <Cmd /> </div>
    </div>,
};