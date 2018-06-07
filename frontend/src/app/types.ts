interface Ballot {
  code: string;
  name: string;
  e: number;
  n: string;
  regexpVoter: string;
  regexpCandidate: string;
  phase: string;
}

interface User {
  name: string;
  picture: string;
  email: string;
  roleCode: string;
}

interface Token {
  jwtToken: string;
  user: User;
}


interface Candidate {
  user: User;
  ballot: Ballot;
  document: string;
  nominee1: string;
  nominee2: string;
}


export { Ballot, Token, Candidate, User };

