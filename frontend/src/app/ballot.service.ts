import { Injectable } from '@angular/core';

import { Ballot } from './types';
import { BALLOTS } from './mock-data';

@Injectable({
  providedIn: 'root'
})
export class BallotService {

  constructor() { }

  getBallots(): Ballot[] {
    return BALLOTS;
  }

}
