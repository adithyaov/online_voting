let component = ReasonReact.statelessComponent("ContentHeading");

let str = str_ => ReasonReact.string(str_);

let make = _children => {
  ...component,
  render: _self =>
    <footer className="footer">
      <div className="content has-text-centered">
        <p>
          <strong> (str("Bulma")) </strong>
          (str("by"))
          <a href="https://jgthms.com"> (str("Jeremy Thomas")) </a>
          (str(". The source code is licensed"))
          <a href="http://opensource.org/licenses/mit-license.php">
            (str("MIT"))
          </a>
          (str(". The website content is licensed"))
          <a href="http://creativecommons.org/licenses/by-nc-sa/4.0/">
            (str("CC BY NC SA 4.0"))
          </a>
          (str("."))
        </p>
      </div>
    </footer>,
};