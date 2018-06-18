export interface Ballot {
  name: string;
  code: string;
  regex_voter: string;
  regex_candidate: string;
  e: number;
  n: string;
  phase: 'C' | 'N' | 'V' | 'S' | 'R';
}
