// Generated by BUCKLESCRIPT VERSION 3.1.5, PLEASE EDIT WITH CARE
'use strict';

var React = require("react");
var ReasonReact = require("reason-react/src/ReasonReact.js");

var component = ReasonReact.statelessComponent("BallotDetails");

function code(lang, child) {
  return React.cloneElement(React.createElement("div", {
                  className: "code"
                }), {
              "data-lang": lang
            }, child);
}

function make() {
  return /* record */[
          /* debugName */component[/* debugName */0],
          /* reactClassInternal */component[/* reactClassInternal */1],
          /* handedOffState */component[/* handedOffState */2],
          /* willReceiveProps */component[/* willReceiveProps */3],
          /* didMount */component[/* didMount */4],
          /* didUpdate */component[/* didUpdate */5],
          /* willUnmount */component[/* willUnmount */6],
          /* willUpdate */component[/* willUpdate */7],
          /* shouldUpdate */component[/* shouldUpdate */8],
          /* render */(function () {
              return React.createElement("div", {
                          className: "ballot-details"
                        }, code("NAME", React.createElement("code", undefined, React.createElement("b", undefined, "General Secratary Elections 2018"))), React.createElement("div", {
                              className: "hdiv"
                            }), React.createElement("div", {
                              className: "columns"
                            }, React.createElement("div", {
                                  className: "column"
                                }, code("PUBLIC KEY", React.createElement("code", undefined, "67887"))), React.createElement("div", {
                                  className: "column"
                                }, code("PHASE", React.createElement("code", undefined, "Create")))), React.createElement("div", {
                              className: "hdiv"
                            }), code("GROUP MODULUS", React.createElement("code", undefined, "792730197018927098126318263182609128640912864981246019824601298461092846120984612098461209846120986219086476387560925723095809741086410284671826523987461798456182541287542168746152746124762154621")));
            }),
          /* initialState */component[/* initialState */10],
          /* retainedProps */component[/* retainedProps */11],
          /* reducer */component[/* reducer */12],
          /* subscriptions */component[/* subscriptions */13],
          /* jsElementWrapped */component[/* jsElementWrapped */14]
        ];
}

exports.component = component;
exports.code = code;
exports.make = make;
/* component Not a pure module */
