import { User } from './user';
import { NullString } from './common';

export interface Candidate {
  user: User;
  ballot_code: string;
  details: string;
  nominee1: NullString;
  nominee2: NullString;
}
