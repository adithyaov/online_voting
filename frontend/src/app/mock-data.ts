
import { Ballot, Candidate, Token, User } from './types';

const USER: User = {
  name: 'Adithya Kumar',
  picture: 'http://icons.iconarchive.com/icons/paomedia/small-n-flat/256/cat-icon.png',
  email: 'adithya.creed@gmail.com',
  role_code: 'A'
};

const TOKEN: Token = {
  jwt_token: 'ahdvajsvd.asdasa.dasda',
  user: USER
};


const BALLOTS: Ballot[] = [
  { phase: 'C', code: 'c1', name: 'Mr. Nice 1', e: 22, n: 'skjsba', regex_voter: 'rv', regex_candidate: 'rc' },
  { phase: 'V', code: 'c2', name: 'Mr. Nice 2', e: 22, n: 'skjsba', regex_voter: 'rv', regex_candidate: 'rc' },
  { phase: 'V', code: 'c3', name: 'Mr. Nice 3', e: 22, n: 'skjsba', regex_voter: 'rv', regex_candidate: 'rc' },
  { phase: 'R', code: 'c4', name: 'Mr. Nice 4', e: 22, n: 'skjsba', regex_voter: 'rv', regex_candidate: 'rc' },
  { phase: 'R', code: 'c5', name: 'Mr. Nice 5', e: 22, n: 'skjsba', regex_voter: 'rv', regex_candidate: 'rc' },
  { phase: 'N', code: 'c6', name: 'Mr. Nice 6', e: 22, n: 'skjsba', regex_voter: 'rv', regex_candidate: 'rc' },
  { phase: 'C', code: 'c7', name: 'Mr. Nice 7', e: 22, n: 'skjsba', regex_voter: 'rv', regex_candidate: 'rc' },
  { phase: 'C', code: 'c7', name: 'Mr. Nice 7', e: 22, n: 'skjsba', regex_voter: 'rv', regex_candidate: 'rc' },
  { phase: 'C', code: 'c7', name: 'Mr. Nice 7', e: 22, n: 'skjsba', regex_voter: 'rv', regex_candidate: 'rc' }
];

const CANDIDATES: Candidate[] = [
  {user: {role_code: 'A', name: 'Test', email: 'lala@gmail.com', picture: undefined},
  ballot: BALLOTS[5], nominee1: undefined, nominee2: undefined, document: 'oogle.ciom'},
];

export {
  BALLOTS,
  CANDIDATES,
  USER,
  TOKEN
};

