let component = ReasonReact.statelessComponent("BallotCard");

let make = _children => {
  ...component,
  render: _self =>
    <div className="ballot-card">
      <div className="tile tile-centered">
        <div className="tile-icon">
          (
            ReasonReact.cloneElement(
              <div
                className="phase bg-secondary rounded tooltip tooltip-right float-right h3"
              />,
              ~props={"data-tooltip": "Lorem Ispum"},
              [|ReasonReact.string("C")|],
            )
          )
        </div>
        <div className="tile-content">
          <div className="tile-title">
            (ReasonReact.string("General Secretary Elections"))
          </div>
          <div className="tile-subtitle text-gray">
            (ReasonReact.string("GS-2015"))
          </div>
        </div>
        <div className="tile-action">
          <button className="btn"> (ReasonReact.string("Open")) </button>
        </div>
      </div>
    </div>,
};