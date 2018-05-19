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
	role_code: string;
}

class Token {
	jwt_token: string;
	user: User;
}


class Candidate {
	user: User;
	ballot: Ballot;
	document: string;
	nominee1: string;
	nominee2: string;
}


export { Ballot, Token, Candidate, User }

