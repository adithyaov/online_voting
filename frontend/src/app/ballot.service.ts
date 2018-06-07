import { Injectable } from '@angular/core';
import { Ballot, Candidate } from './types';
import { BALLOTS, CANDIDATES } from './mock-data';
import { HttpClient } from '@angular/common/http';
import { GETBALLOTSURL, GETBALLOTURL } from './consts';
import { makeHeaders } from './common';

import { Observable, throwError, of } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { TokenService } from './token.service';


@Injectable({
  providedIn: 'root'
})
export class BallotService {

  constructor(
    private http: HttpClient,
    private token: TokenService
  ) { }

  getBallots(): Observable<Ballot[]> {
    // return of(BALLOTS);
    const data = {
      email: this.token.currentToken().user.email,
    };
    return this.http.post<Ballot[]>(GETBALLOTSURL, data, makeHeaders({'Token': this.token.currentToken().jwt_token}))
      .pipe(
        retry(3),
        catchError((error) => {
          console.log('some error dude');
          console.log(error);
          return throwError('error');
        })
      );
  }

  getBallot(code: string): Observable<Ballot> {
    const data = {
      code: code,
    };
    return this.http.post<Ballot>(GETBALLOTURL, data, makeHeaders({'Token': this.token.currentToken().jwt_token}))
      .pipe(
        retry(3),
        catchError((error) => {
          console.log('some error dude');
          console.log(error);
          return throwError('error');
        })
      );
  }

  getCandidates(code: string): Observable<Candidate[]> {
    return of(CANDIDATES);
  }

}
