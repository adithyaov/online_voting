let authUser = axiosConfig => {
  let data = {"def": 0};
  Axios.postDatac("/ballot/token-service", data, axiosConfig);
};

let delete = (email, axiosConfig) => {
  let data = {"email": email};
  Axios.postDatac("/ballot/delete", data, axiosConfig);
};

let updatePersonal = (email, name, picture, axiosConfig) => {
  let data = {"email": email, "name": name, "picture": picture};
  Axios.postDatac("/ballot/update-personal", data, axiosConfig);
};

let updateRole = (email, roleCode, axiosConfig) => {
  let data = {"email": email, "role_code": roleCode};
  Axios.postDatac("/ballot/update-role", data, axiosConfig);
};