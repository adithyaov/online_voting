let create = (code: string, name: string, axiosConfig) => {
  let data = {"code": code, "name": name};
  Axios.postDatac("/ballot/create", data, axiosConfig);
};

let find = (code: string, axiosConfig) => {
  let data = {"code": code};
  Axios.postDatac("/ballot/find", data, axiosConfig);
};

let delete = (code: string, axiosConfig) => {
  let data = {"code": code};
  Axios.postDatac("/ballot/delete", data, axiosConfig);
};

let blindVote = (candidateEmail: string, axiosConfig) => {
  let data = {"candidate_email": candidateEmail};
  Axios.postDatac("/ballot/blind-vote", data, axiosConfig);
};

let signBytes = (blinded: list(int), voterEmail: string, axiosConfig) => {
  let data = {"blinded": blinded, "voter_email": voterEmail};
  Axios.postDatac("/ballot/sign-bytes", data, axiosConfig);
};

let unblindSign = (signed: list(int), unblinder: list(int), axiosConfig) => {
  let data = {"signed": signed, "unblinder": unblinder};
  Axios.postDatac("/ballot/unblind-sign", data, axiosConfig);
};

let verifySign = (voteHash: list(int), unblinded: list(int), axiosConfig) => {
  let data = {"vote_hash": voteHash, "unblinded": unblinded};
  Axios.postDatac("/ballot/verify-sign", data, axiosConfig);
};

let findBallots = (email: string, axiosConfig) => {
  let data = {"email": email};
  Axios.postDatac("/ballot/find-ballots", data, axiosConfig);
};

let restartOpenBallots = axiosConfig => {
  let data = {"def": 0};
  Axios.postDatac("/ballot/restart", data, axiosConfig);
};

let update =
    (
      code: string,
      name: string,
      regexCandidate: string,
      regexVoter: string,
      phase: char,
      axiosConfig,
    ) => {
  let data = {
    "code": code,
    "name": name,
    "regex_candidate": regexCandidate,
    "regex_voter": regexVoter,
    "phase": phase,
  };
  Axios.postDatac("/ballot/update", data, axiosConfig);
};