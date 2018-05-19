import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Ballot, Candidate } from './types';
import { BALLOTS, CANDIDATES } from './mock-data';

@Injectable({
  providedIn: 'root'
})
export class BallotService {

  constructor() { }

  getBallots(): Observable<Ballot[]> {
    return of(BALLOTS);
  }

  getBallot(code: string): Observable<Ballot> {
    return of(BALLOTS.filter(b => b.code == code)[0]);
  }

  getCandidates(code: string): Observable<Candidate[]> {
  	return of(CANDIDATES);
  }

}
