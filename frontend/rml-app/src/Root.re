let component = ReasonReact.statelessComponent("Root");

let make = _children => {
  ...component,
  render: _self =>
    <div className="wrapper">
      <div className="side-bar-wrapper"> <SideBar /> </div>
      <div className="content-wrapper"> (ReasonReact.string("content")) </div>
      <div className="ballot-bar-wrapper"> <BallotBar /> </div>
      <div className="cmd-wrapper"> (ReasonReact.string("cmd")) </div>
    </div>,
};