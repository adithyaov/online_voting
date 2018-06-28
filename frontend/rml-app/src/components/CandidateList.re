let component = ReasonReact.statelessComponent("BallotBar");

let str = str_ => ReasonReact.string(str_);

let candidateCard =
  <div className="card">
    <div className="card-content">
      <div className="media">
        <div className="media-left">
          <figure className="image is-48x48">
            <img
              src="https://picturepan2.github.io/spectre/img/avatar-2.png"
              alt="Placeholder image"
            />
          </figure>
        </div>
        <div className="media-content">
          <p className="title is-4"> (str("John Smith")) </p>
          <p className="subtitle is-6"> (str("@johnsmith")) </p>
        </div>
      </div>
    </div>
    <footer className="card-footer">
      <a href="#" className="card-footer-item"> (str("Vote")) </a>
      <a href="#" className="card-footer-item"> (str("Document")) </a>
    </footer>
  </div>;

let make = _children => {
  ...component,
  render: _self =>
    <div className="columns is-multiline">
      <div className="column is-half"> candidateCard </div>
      <div className="column is-half"> candidateCard </div>
      <div className="column is-half"> candidateCard </div>
      <div className="column is-half"> candidateCard </div>
    </div>,
};