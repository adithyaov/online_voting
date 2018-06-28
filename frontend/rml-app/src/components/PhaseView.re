let component = ReasonReact.statelessComponent("PhaseView");

let str = str_ => ReasonReact.string(str_);

let make = _children => {
  ...component,
  render: _self =>
    <div>
      <article className="message is-info">
        <div className="message-header"> <p> (str("VIEW")) </p> </div>
        <div className="message-body">
          (str("This ballot is in view phase."))
          <br />
          (
            str(
              "If you think there is a problem with the current phase please inform any of the",
            )
          )
          <b> (str(" admins")) </b>
          (str(" or the"))
          <b> (str(" moderators")) </b>
          (str("."))
          <br />
          (str(" For more information regarding phases please refer to the "))
          <a> (str("documentation")) </a>
          (str("."))
        </div>
      </article>
      <CandidateList />
    </div>,
};