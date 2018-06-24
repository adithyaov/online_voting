let component = ReasonReact.statelessComponent("UserSet");

let make = _children => {
  ...component,
  render: _self =>
    <div>
      <div className="divider-wrapper"> (Common.divider("CANDIDATES")) </div>
      <div className="h10" />
      <div className="user-set"> <UserCard /> <UserCard /> <UserCard /> </div>
    </div>,
};