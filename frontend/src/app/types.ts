interface Ballot {
  code: string;
  name: string;
  e: number;
  n: string;
  regexp_voter: string;
  regexp_candidate: string;
  phase: string;
}

interface User {
  name: string;
  picture: string;
  email: string;
  role_code: string;
}

interface Token {
  jwt_token: string;
  user: User;
  status: string;
}


interface Candidate {
  user: User;
  ballot: Ballot;
  document: string;
  nominee1: string;
  nominee2: string;
}


export { Ballot, Token, Candidate, User };

