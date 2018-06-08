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

interface NullString {
  valid: boolean;
  string: string;
}

interface Candidate {
  user: User;
  ballot_code: string;
  details: string;
  nominee1: NullString;
  nominee2: NullString;
}


export { Ballot, Token, Candidate, User };

