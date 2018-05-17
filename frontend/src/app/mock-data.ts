
import { Ballot, Candidate } from './types';

const BALLOTS: Ballot[] = [
  { phase: "C", code: "c1", name: 'Mr. Nice 1', e: 22, n: "skjsba", regex_voter: "rv", regex_candidate: "rc" },
  { phase: "V", code: "c2", name: 'Mr. Nice 2', e: 22, n: "skjsba", regex_voter: "rv", regex_candidate: "rc" },
  { phase: "V", code: "c3", name: 'Mr. Nice 3', e: 22, n: "skjsba", regex_voter: "rv", regex_candidate: "rc" },
  { phase: "R", code: "c4", name: 'Mr. Nice 4', e: 22, n: "skjsba", regex_voter: "rv", regex_candidate: "rc" },
  { phase: "R", code: "c5", name: 'Mr. Nice 5', e: 22, n: "skjsba", regex_voter: "rv", regex_candidate: "rc" },
  { phase: "N", code: "c6", name: 'Mr. Nice 6', e: 22, n: "skjsba", regex_voter: "rv", regex_candidate: "rc" },
  { phase: "C", code: "c7", name: 'Mr. Nice 7', e: 22, n: "skjsba", regex_voter: "rv", regex_candidate: "rc" },
  { phase: "C", code: "c7", name: 'Mr. Nice 7', e: 22, n: "skjsba", regex_voter: "rv", regex_candidate: "rc" },
  { phase: "C", code: "c7", name: 'Mr. Nice 7', e: 22, n: "skjsba", regex_voter: "rv", regex_candidate: "rc" }
];

const CANDIDATES: Candidate[] = [
	{name: "Test", nominee1: undefined, nominee2: undefined, email: "lala@gmail.com", document: "oogle.ciom", picture:""},
	{name: "Test", nominee1: "n1", nominee2: "n2", email: "lala@gmail.com", document: "oogle.ciom", picture:""},
	{name: "Test", nominee1: "adithya.creed@gmail.com", nominee2: undefined, email: "lala@gmail.com", document: "oogle.ciom", picture:""}
]

export {
	BALLOTS,
	CANDIDATES
}

