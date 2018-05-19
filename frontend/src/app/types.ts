class Ballot {
  code: string;
  name: string;
  e: number;
  n: string;
  regex_voter: string;
  regex_candidate: string;
  phase: string;
}

class User {
	name: string;
	picture: string;
	email: string;
}

class Token {
	jwt_token: string;
	user: User;
}


class Candidate {
	user: User;
	ballot_code: string;
	document: string;
	nominee1: string;
	nominee2: string;
}


export { Ballot, Token, Candidate, User }

