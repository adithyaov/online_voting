import { Component } from '@angular/core';
import { TokenService } from './token.service';
import { Token } from './types';
import { LoadingBarService } from '@ngx-loading-bar/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  constructor(
    private token: TokenService,
    private loader: LoadingBarService
  ) { }

  title = 'application';
  infoModel = false;

  activateInfoModel(): void {
    this.infoModel = true;
    console.log(this.infoModel);
    this.loader.start();
  }

  deactivateInfoModel(): void {
    this.infoModel = false;
    console.log(this.infoModel);
    this.loader.complete();
  }


}
