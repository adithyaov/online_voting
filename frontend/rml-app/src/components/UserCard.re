let component = ReasonReact.statelessComponent("UserCard");

let make = _children => {
  ...component,
  render: _self =>
    <div className="user-card">
      <div className="display-picture">
        <figure className="avatar avatar-xl">
          <img
            src="https://picturepan2.github.io/spectre/img/avatar-2.png"
            alt="Adithya Kumar"
          />
        </figure>
      </div>
      <div className="card">
        <div className="card-header">
          <div className="card-title h6">
            (ReasonReact.string("Adithya Kumar"))
          </div>
          <div className="card-subtitle text-gray">
            (ReasonReact.string("111501017@smail.iitpkd.ac.in"))
          </div>
        </div>
        <div className="card-footer">
          <button className="btn btn-primary">
            (ReasonReact.string("Vote"))
          </button>
          <a className="btn"> (ReasonReact.string("Document")) </a>
        </div>
      </div>
    </div>,
};