import { Injectable } from '@angular/core';
import { Token } from './types';
import { TOKEN } from './mock-data';

@Injectable({
  providedIn: 'root'
})
export class TokenService {

	token: Token = undefined;

	setToken(): Token {
		this.token = TOKEN
		return this.token;
	}

	getToken(): Token {
		return this.token;
	}

  constructor() { }
}
