class Ballot {
  code: string;
  name: string;
  e: number;
  n: string;
  regex_voter: string;
  regex_candidate: string;
  phase: string;
}

class Token {
	jwt_token: string;
	name: string;
	picture: string;
	email: string;
}


class Candidate {
	name: string;
	email: string;
	document: string;
	picture: string;
	nominee1: string;
	nominee2: string;
}


export { Ballot, Token, Candidate }

