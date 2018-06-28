let component = ReasonReact.statelessComponent("Root");

let str = str_ => ReasonReact.string(str_);

let make = _children => {
  ...component,
  render: _self =>
    <div>
      <Nav />
      <section className="section">
        <div className="container">
          <div className="tile is-ancestor">
            <div className="tile is-parent">
              <div className="tile is-child">
                <BallotDetails />
                <hr />
                <PhaseView />
              </div>
            </div>
            <div className="tile is-parent is-3">
              <div className="tile is-child"> <BallotBar /> </div>
            </div>
          </div>
        </div>
      </section>
      <Footer />
    </div>,
};