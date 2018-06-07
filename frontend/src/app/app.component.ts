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

  token = this.tokenService.setToken();
  title = 'application';
  infoModel = false;

  activateInfoModel(): void {
    this.infoModel = true;
    console.log(this.infoModel);
  }
  deactivateInfoModel(): void {
    this.infoModel = false;
    console.log(this.infoModel);
  }


}
