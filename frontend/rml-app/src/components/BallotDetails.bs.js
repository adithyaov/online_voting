// Generated by BUCKLESCRIPT VERSION 3.1.5, PLEASE EDIT WITH CARE
'use strict';

var React = require("react");
var ReasonReact = require("reason-react/src/ReasonReact.js");

var component = ReasonReact.statelessComponent("BallotDetails");

function str(str_) {
  return str_;
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
              return React.createElement("div", undefined, React.createElement("div", {
                              className: "title"
                            }, "@ballot_details"), React.createElement("div", {
                              className: "subtitle"
                            }, "GS-2017", React.createElement("span", {
                                  className: "tag is-primary is-pulled-right"
                                }, "NOMINATION")), React.createElement("table", {
                              className: "table is-hoverable is-bordered details-table"
                            }, React.createElement("tbody", undefined, React.createElement("tr", undefined, React.createElement("th", undefined, "Name"), React.createElement("td", undefined, "General secratary elections 2014")), React.createElement("tr", undefined, React.createElement("th", undefined, "Public Key"), React.createElement("td", undefined, "23532423")), React.createElement("tr", {
                                      className: "is-hidden-mobile"
                                    }, React.createElement("th", undefined, "Group Modulus"), React.createElement("td", undefined, React.createElement("div", undefined, "9386498236498241092740937509341235938649823649824109274093750934123593864982364982410927409375093412359386498236498241092740937509341235938649823649824109274093750934123593864982364982410927409375093412359386498236498241092740937509341235938649823649824109274093750934123593864982364982410927409375093412359386498236498241092740937509341235"))))));
            }),
          /* initialState */component[/* initialState */10],
          /* retainedProps */component[/* retainedProps */11],
          /* reducer */component[/* reducer */12],
          /* subscriptions */component[/* subscriptions */13],
          /* jsElementWrapped */component[/* jsElementWrapped */14]
        ];
}

exports.component = component;
exports.str = str;
exports.make = make;
/* component Not a pure module */