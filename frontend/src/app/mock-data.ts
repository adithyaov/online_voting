
import { Ballot, Candidate, Token, User } from './types';

const USER: User = {
  name: 'Adithya Kumar',
  picture: 'http://icons.iconarchive.com/icons/paomedia/small-n-flat/256/cat-icon.png',
  email: 'adithya.creed@gmail.com',
  roleCode: 'A'
};

const TOKEN: Token = {
  jwtToken: 'ahdvajsvd.asdasa.dasda',
  user: USER
};


const BALLOTS: Ballot[] = [
  { phase: 'C', code: 'c1', name: 'Mr. Nice 1', e: 22, n: 'skjsba', regexpVoter: 'rv', regexpCandidate: 'rc' },
  { phase: 'V', code: 'c2', name: 'Mr. Nice 2', e: 22, n: 'skjsba', regexpVoter: 'rv', regexpCandidate: 'rc' },
  { phase: 'V', code: 'c3', name: 'Mr. Nice 3', e: 22, n: 'skjsba', regexpVoter: 'rv', regexpCandidate: 'rc' },
  { phase: 'R', code: 'c4', name: 'Mr. Nice 4', e: 22, n: 'skjsba', regexpVoter: 'rv', regexpCandidate: 'rc' },
  { phase: 'R', code: 'c5', name: 'Mr. Nice 5', e: 22, n: 'skjsba', regexpVoter: 'rv', regexpCandidate: 'rc' },
  { phase: 'N', code: 'c6', name: 'Mr. Nice 6', e: 22, n: 'skjsba', regexpVoter: 'rv', regexpCandidate: 'rc' },
  { phase: 'C', code: 'c7', name: 'Mr. Nice 7', e: 22, n: 'skjsba', regexpVoter: 'rv', regexpCandidate: 'rc' },
  { phase: 'C', code: 'c7', name: 'Mr. Nice 7', e: 22, n: 'skjsba', regexpVoter: 'rv', regexpCandidate: 'rc' },
  { phase: 'C', code: 'c7', name: 'Mr. Nice 7', e: 22, n: 'skjsba', regexpVoter: 'rv', regexpCandidate: 'rc' }
];

const CANDIDATES: Candidate[] = [
  {user: {roleCode: 'A', name: 'Test', email: 'lala@gmail.com', picture: undefined},
  ballot: BALLOTS[5], nominee1: undefined, nominee2: undefined, document: 'oogle.ciom'},
];

export {
  BALLOTS,
  CANDIDATES,
  USER,
  TOKEN
};

