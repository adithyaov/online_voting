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

  info_model: boolean = false;
  activate_info_model(): void {
  	this.info_model = true;
  	console.log(this.info_model)
  }
  deactivate_info_model(): void {
  	this.info_model = false;
  	console.log(this.info_model)
  }

  token: Token = this.tokenService.setToken();

  title = 'application';
}
