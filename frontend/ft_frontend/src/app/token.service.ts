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

  observable(googleToken: string): Observable<Token> {
    return this.http.post<Token>(GETTOKENURL, {}, makeHeaders({'Google-Token': googleToken}))
    .pipe(
      retry(3),
      catchError((error) => {
        console.log(error);
        return throwError('error');
      })
    );
  }

  subscribe(googleToken: string): void {
    this.loader.start();
    this.observable(googleToken).subscribe(t => {
      this.token = t;
      this.loader.complete();
    });
  }

  destroy(): void {
    this.token = undefined;
  }

}
