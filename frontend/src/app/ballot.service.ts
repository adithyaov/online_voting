import { Injectable } from '@angular/core';
import { Ballot, Candidate } from './types';
import { BALLOTS, CANDIDATES } from './mock-data';
import { HttpClient } from '@angular/common/http';
import { GETBALLOTSURL } from './consts';
import { makeHeaders } from './common';

import { Observable, throwError, of } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BallotService {

  constructor(private http: HttpClient) { }

  getBallots(): Observable<Ballot[]> {
    // return of(BALLOTS);
    const data = {
      email: 'lalala'
    };
    return this.http.post<Ballot[]>(GETBALLOTSURL, data, makeHeaders('lalala'))
      .pipe(
        retry(3),
        catchError(() => throwError('error'))
      );
  }

  getBallot(code: string): Observable<Ballot> {
    return of(BALLOTS.filter(b => b.code === code)[0]);
  }

  getCandidates(code: string): Observable<Candidate[]> {
    return of(CANDIDATES);
  }

}
