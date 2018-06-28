let create =
    (ballotCode: string, userEmail: string, details: string, axiosConfig) => {
  let data = {
    "ballot_code": ballotCode,
    "user_email": userEmail,
    "details": details,
  };
  Axios.postDatac("/ballot/create", data, axiosConfig);
};

let addNominee =
    (ballotCode: string, userEmail: string, nomineeEmail: string, axiosConfig) => {
  let data = {
    "ballot_code": ballotCode,
    "user_email": userEmail,
    "nominee_email": nomineeEmail,
  };
  Axios.postDatac("/ballot/add-nominee", data, axiosConfig);
};

let updateDetails =
    (ballotCode: string, userEmail: string, details: string, axiosConfig) => {
  let data = {
    "ballot_code": ballotCode,
    "user_email": userEmail,
    "details": details,
  };
  Axios.postDatac("/ballot/update-details", data, axiosConfig);
};

let delete = (ballotCode: string, userEmail: string, axiosConfig) => {
  let data = {"ballot_code": ballotCode, "user_email": userEmail};
  Axios.postDatac("/ballot/delete", data, axiosConfig);
};

let ballotCandidates = (code: string, axiosConfig) => {
  let data = {"code": code};
  Axios.postDatac("/ballot/per-ballot", data, axiosConfig);
};