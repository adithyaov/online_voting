import { Injectable } from '@angular/core';
import { Token } from './types';

@Injectable({
  providedIn: 'root'
})
export class TokenService {

	token: Token = undefined;

	setToken(): Token {
		this.token = {
			jwt_token: "ahdvajsvd.asdasa.dasda",
			name: "Adithya Kumar",
			picture: "http://icons.iconarchive.com/icons/paomedia/small-n-flat/256/cat-icon.png"
		}
		return this.token;
	}

	getToken(): Token {
		return this.token;
	}

  constructor() { }
}
