let component = ReasonReact.statelessComponent("BallotDetails");

let str = str_ => ReasonReact.string(str_);

let make = _children => {
  ...component,
  render: _self =>
    <div>
      <div className="title"> (str("@ballot_details")) </div>
      <div className="subtitle">
        (str("GS-2017"))
        <span className="tag is-primary is-pulled-right">
          (str("NOMINATION"))
        </span>
      </div>
      /* <a className="button is-pulled-right is-large is-light is-hidden-mobile">
           <span className="icon is-large">
             <i className="fas fa-download" />
           </span>
         </a> */
      <table className="table is-hoverable is-bordered details-table">
        <tbody>
          <tr>
            <th> (str("Name")) </th>
            <td> (str("General secratary elections 2014")) </td>
          </tr>
          <tr>
            <th> (str("Public Key")) </th>
            <td> (str("23532423")) </td>
          </tr>
          <tr className="is-hidden-mobile">
            <th> (str("Group Modulus")) </th>
            <td>
              <div>
                (
                  str(
                    "9386498236498241092740937509341235938649823649824109274093750934123593864982364982410927409375093412359386498236498241092740937509341235938649823649824109274093750934123593864982364982410927409375093412359386498236498241092740937509341235938649823649824109274093750934123593864982364982410927409375093412359386498236498241092740937509341235",
                  )
                )
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>,
};