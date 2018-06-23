let component = ReasonReact.statelessComponent("BallotDetails");

let make = _children => {
  ...component,
  render: _self =>
    <div className="cmd bg-secondary">
      (
        ReasonReact.cloneElement(
          <div className="has-icon-left tooltip tooltip-top" />,
          ~props={"data-tooltip": "type `help` to learn more"},
          [|
            <i className="form-icon icon icon-arrow-right" />,
            <input
              type_="text"
              placeholder="execute commands, eg: ballot.open.[GS2017]"
              className="form-input input-lg"
            />,
          |],
        )
      )
    </div>,
};