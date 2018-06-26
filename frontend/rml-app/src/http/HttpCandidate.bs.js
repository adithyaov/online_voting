// Generated by BUCKLESCRIPT VERSION 3.1.5, PLEASE EDIT WITH CARE
'use strict';

var Axios = require("axios");

function create(ballotCode, userEmail, details, axiosConfig) {
  var data = {
    ballot_code: ballotCode,
    user_email: userEmail,
    details: details
  };
  return Axios.post("/ballot/create", data, axiosConfig);
}

function addNominee(ballotCode, userEmail, nomineeEmail, axiosConfig) {
  var data = {
    ballot_code: ballotCode,
    user_email: userEmail,
    nominee_email: nomineeEmail
  };
  return Axios.post("/ballot/add-nominee", data, axiosConfig);
}

function updateDetails(ballotCode, userEmail, details, axiosConfig) {
  var data = {
    ballot_code: ballotCode,
    user_email: userEmail,
    details: details
  };
  return Axios.post("/ballot/update-details", data, axiosConfig);
}

function $$delete(ballotCode, userEmail, axiosConfig) {
  var data = {
    ballot_code: ballotCode,
    user_email: userEmail
  };
  return Axios.post("/ballot/delete", data, axiosConfig);
}

function ballotCandidates(code, axiosConfig) {
  var data = {
    code: code
  };
  return Axios.post("/ballot/per-ballot", data, axiosConfig);
}

exports.create = create;
exports.addNominee = addNominee;
exports.updateDetails = updateDetails;
exports.$$delete = $$delete;
exports.ballotCandidates = ballotCandidates;
/* axios Not a pure module */
