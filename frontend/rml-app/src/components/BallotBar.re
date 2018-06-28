let component = ReasonReact.statelessComponent("BallotBar");

let str = str_ => ReasonReact.string(str_);

let make = _children => {
  ...component,
  render: _self =>
    <nav className="panel">
      <p className="panel-heading"> (str("Meine Ballots")) </p>
      <div className="panel-block">
        <p className="control has-icons-left">
          <input
            className="input is-small"
            type_="text"
            placeholder="search"
          />
          <span className="icon is-small is-left">
            <i className="fas fa-search" ariaHidden=true />
          </span>
        </p>
      </div>
      <p className="panel-tabs">
        <a className="is-active"> (str("all")) </a>
        <a> (str("view")) </a>
        <a> (str("nominate")) </a>
        <a> (str("register")) </a>
        <a> (str("vote")) </a>
      </p>
      <a className="panel-block is-active">
        <span className="panel-icon">
          <i className="fas fa-book" ariaHidden=true />
        </span>
        (str("bulma"))
      </a>
      <a className="panel-block">
        <span className="panel-icon">
          <i className="fas fa-book" ariaHidden=true />
        </span>
        (str("marksheet"))
      </a>
    </nav>,
};