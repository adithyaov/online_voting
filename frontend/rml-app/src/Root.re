let component = ReasonReact.statelessComponent("Root");

let make = _children => {
  ...component,
  render: _self =>
    <div className="wrapper">
      <div className="side-bar-wrapper"> <SideBar /> </div>
      <div className="content-wrapper">
        <div className="content">
          <ContentHeading />
          <BallotDetails />
          <div className="h10" />
          <UserSet />
        </div>
      </div>
      <div className="ballot-bar-wrapper"> <BallotBar /> </div>
      <div className="cmd-wrapper"> <Cmd /> </div>
    </div>,
};