let component = ReasonReact.statelessComponent("BallotBar");

let make = _children => {
  ...component,
  render: _self =>
    <div className="ballot-bar">
      <BallotCard />
      <BallotCard />
      <BallotCard />
    </div>,
};