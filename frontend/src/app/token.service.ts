import { Injectable } from '@angular/core';
import { Token } from './types';
import { TOKEN } from './mock-data';

import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

import { GETTOKENURL } from './consts';
import { makeHeaders } from './common';

import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class TokenService {

  constructor(
    private http: HttpClient
  ) {}

  token: Token = undefined;

  requestToken(): Token {
    this.http.post<Token>(GETTOKENURL, {}, makeHeaders({'Google-Token': 'ssss'}))
    .pipe(
      retry(3),
      catchError((error) => {
        console.log('some error dude');
        console.log(error);
        return throwError('error');
      })
    ).subscribe(t => this.token = t);
    return this.token;
  }

  destroyToken(): Token {
    this.token = undefined;
    return this.token;
  }

  currentToken(): Token {
    return this.token;
  }

}
