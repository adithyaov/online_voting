let create = (code, name, axiosConfig) => {
  let data = {"code": code, "name": name};
  Axios.postDatac("/ballot/create", data, axiosConfig);
};

let find = (code, axiosConfig) => {
  let data = {"code": code};
  Axios.postDatac("/ballot/find", data, axiosConfig);
};

let delete = (code, axiosConfig) => {
  let data = {"code": code};
  Axios.postDatac("/ballot/delete", data, axiosConfig);
};

let blindVote = (candidateEmail, axiosConfig) => {
  let data = {"candidate_email": candidateEmail};
  Axios.postDatac("/ballot/blind-vote", data, axiosConfig);
};

let signBytes = (blinded, voterEmail, axiosConfig) => {
  let data = {"blinded": blinded, "voter_email": voterEmail};
  Axios.postDatac("/ballot/sign-bytes", data, axiosConfig);
};

let unblindSign = (signed, unblinder, axiosConfig) => {
  let data = {"signed": signed, "unblinder": unblinder};
  Axios.postDatac("/ballot/unblind-sign", data, axiosConfig);
};

let verifySign = (voteHash, unblinded, axiosConfig) => {
  let data = {"vote_hash": voteHash, "unblinded": unblinded};
  Axios.postDatac("/ballot/verify-sign", data, axiosConfig);
};

let findBallots = (email, axiosConfig) => {
  let data = {"email": email};
  Axios.postDatac("/ballot/find-ballots", data, axiosConfig);
};

let restartOpenBallots = axiosConfig => {
  let data = {"def": 0};
  Axios.postDatac("/ballot/restart", data, axiosConfig);
};

let update = (code, name, regexCandidate, regexVoter, phase, axiosConfig) => {
  let data = {
    "code": code,
    "name": name,
    "regex_candidate": regexCandidate,
    "regex_voter": regexVoter,
    "phase": phase,
  };
  Axios.postDatac("/ballot/update", data, axiosConfig);
};