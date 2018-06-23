let component = ReasonReact.statelessComponent("ContentHeading");

let make = _children => {
  ...component,
  render: _self =>
    <div className="heading text-center">
      <div className="text-bold">
        (ReasonReact.string("BALLOT DETAILS"))
        <small> (ReasonReact.string("GS2011")) </small>
      </div>
    </div>,
};