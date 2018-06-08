import { Injectable } from '@angular/core';
import { Token } from './types';
import { TOKEN } from './mock-data';

import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

import { GETTOKENURL } from './consts';
import { makeHeaders } from './common';

import { HttpClient } from '@angular/common/http';

import { LoadingBarService } from '@ngx-loading-bar/core';

@Injectable({
  providedIn: 'root'
})
export class TokenService {

  constructor(
    private http: HttpClient,
    private loader: LoadingBarService
  ) {}

  token: Token = undefined;

  requestToken(): Token {
    this.loader.start();
    this.http.post<Token>(GETTOKENURL, {}, makeHeaders({'Google-Token': 'ssss'}))
    .pipe(
      retry(3),
      catchError((error) => {
        this.loader.complete();
        console.log(error);
        return throwError('error');
      })
    ).subscribe(t => {
      this.token = t;
      this.loader.complete();
      return this.token;
    });
  }

  destroyToken(): Token {
    this.token = undefined;
    return this.token;
  }

  currentToken(): Token {
    return this.token;
  }

}
