const USER = {
  name: 'Adithya Kumar',
  picture:
    'http://icons.iconarchive.com/icons/paomedia/small-n-flat/256/cat-icon.png',
  email: 'adithya.creed@gmail.com',
  role_code: 'A'
};

const TOKEN = {
  jwt_token: 'ahdvajsvd.asdasa.dasda',
  user: USER
};

const BALLOTS = [
  {
    phase: 'C',
    code: 'c1',
    name: 'Mr. Nice 1',
    e: 22,
    n: 'skjsba',
    regex_voter: 'rv',
    regex_candidate: 'rc'
  },
  {
    phase: 'V',
    code: 'c2',
    name: 'Mr. Nice 2',
    e: 22,
    n: 'skjsba',
    regex_voter: 'rv',
    regex_candidate: 'rc'
  },
  {
    phase: 'V',
    code: 'c3',
    name: 'Mr. Nice 3',
    e: 22,
    n: 'skjsba',
    regex_voter: 'rv',
    regex_candidate: 'rc'
  },
  {
    phase: 'R',
    code: 'c4',
    name: 'Mr. Nice 4',
    e: 22,
    n: 'skjsba',
    regex_voter: 'rv',
    regex_candidate: 'rc'
  },
  {
    phase: 'R',
    code: 'c5',
    name: 'Mr. Nice 5',
    e: 22,
    n: 'skjsba',
    regex_voter: 'rv',
    regex_candidate: 'rc'
  },
  {
    phase: 'N',
    code: 'c6',
    name: 'Mr. Nice 6',
    e: 22,
    n: 'skjsba',
    regex_voter: 'rv',
    regex_candidate: 'rc'
  },
  {
    phase: 'C',
    code: 'c7',
    name: 'Mr. Nice 7',
    e: 22,
    n: 'skjsba',
    regex_voter: 'rv',
    regex_candidate: 'rc'
  },
  {
    phase: 'C',
    code: 'c7',
    name: 'Mr. Nice 7',
    e: 22,
    n: 'skjsba',
    regex_voter: 'rv',
    regex_candidate: 'rc'
  },
  {
    phase: 'C',
    code: 'c7',
    name: 'Mr. Nice 7',
    e: 22,
    n: 'skjsba',
    regex_voter: 'rv',
    regex_candidate: 'rc'
  }
];

const CANDIDATES = [
  {
    user: {
      role_code: 'A',
      name: 'Test',
      email: 'adithya.creed@gmail.com',
      picture:
        'http://icons.iconarchive.com/icons/paomedia/small-n-flat/256/cat-icon.png'
    },
    ballot_code: BALLOTS[5]['code'],
    nominee1: { valid: true, string: 'aaaa' },
    nominee2: { valid: true, string: 'bbbb' },
    details: 'oogle.ciom'
  }
];

export { BALLOTS, CANDIDATES, USER, TOKEN };
