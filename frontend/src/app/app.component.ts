import { Component } from '@angular/core';
import { TokenService } from './token.service';
import { Token } from './types';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(private tokenService: TokenService) { }

  token: Token = this.tokenService.setToken();

  title = 'application';
}
