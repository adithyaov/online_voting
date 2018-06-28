let component = ReasonReact.statelessComponent("Nav");

let str = str_ => ReasonReact.string(str_);

let make = _children => {
  ...component,
  render: _self =>
    <nav className="navbar is-spaced is-black">
      <div className="container">
        <div className="navbar-brand">
          <a className="navbar-item has-text-light logo" href="#">
            <b> (str("PEN")) </b>
            <span> (str("BALLOT")) </span>
          </a>
          <div className="navbar-burger burger">
            <span />
            <span />
            <span />
          </div>
        </div>
        <div className="navbar-menu">
          <div className="navbar-end">
            <a className="navbar-item" href="/documentation">
              <span> (str("Documentation")) </span>
            </a>
            <a className="navbar-item" href="https://bulma.io/blog/">
              <span> (str("Blog")) </span>
            </a>
            <a className="navbar-item" href="https://bulma.io/expo/">
              <span> (str("Expo")) </span>
            </a>
            <div className="navbar-item has-dropdown is-hoverable">
              <a className="navbar-link" href="https://bulma.io/more">
                (str("More"))
              </a>
              <div id="moreDropdown" className="navbar-dropdown">
                <a
                  className="navbar-item "
                  href="https://bulma.io//bulma-start">
                  <span>
                    <strong> (str("Bulma start")) </strong>
                    <br />
                    (str("A tiny npm package to get started"))
                  </span>
                </a>
                <hr className="navbar-divider" />
                <a
                  className="navbar-item "
                  href="https://bulma.io//made-with-bulma">
                  <span>
                    <strong> (str("Made with Bulma")) </strong>
                    <br />
                    (str("The official community badge"))
                  </span>
                </a>
              </div>
            </div>
            <a
              className="navbar-item"
              href="https://github.com/jgthms/bulma"
              target="_blank">
              <span className="icon">
                <i className="fab fa-lg fa-github-alt" />
              </span>
            </a>
            <a
              className="navbar-item"
              href="https://twitter.com/jgthms"
              target="_blank">
              <span className="icon">
                <i className="fab fa-lg fa-twitter" />
              </span>
            </a>
            <div className="navbar-item">
              <a className="button is-danger is-outlined"> (str("Login")) </a>
            </div>
          </div>
        </div>
      </div>
    </nav>,
};