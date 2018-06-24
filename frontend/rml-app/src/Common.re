let divider = dataContent =>
  ReasonReact.cloneElement(
    <li className="divider" />,
    ~props={"data-content": dataContent},
    [||],
  );