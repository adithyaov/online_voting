let component = ReasonReact.statelessComponent("BallotDetails");

let code = (lang, child) =>
  ReasonReact.cloneElement(
    <div className="code" />,
    ~props={"data-lang": lang},
    [|child|],
  );

let make = _children => {
  ...component,
  render: _self =>
    <div className="ballot-details">
      (
        code(
          "NAME",
          <code>
            <b> (ReasonReact.string("General Secratary Elections 2018")) </b>
          </code>,
        )
      )
      <div className="hdiv" />
      <div className="columns">
        <div className="column">
          (code("PUBLIC KEY", <code> (ReasonReact.string("67887")) </code>))
        </div>
        <div className="column">
          (code("PHASE", <code> (ReasonReact.string("Create")) </code>))
        </div>
      </div>
      <div className="hdiv" />
      (
        code(
          "GROUP MODULUS",
          <code>
            (
              ReasonReact.string(
                "792730197018927098126318263182609128640912864981246019824601298461092846120984612098461209846120986219086476387560925723095809741086410284671826523987461798456182541287542168746152746124762154621",
              )
            )
          </code>,
        )
      )
    </div>,
};