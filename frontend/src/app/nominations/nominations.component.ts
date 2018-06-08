import { Component, OnInit } from '@angular/core';
import { UserService } from '../user.service';
import { TokenService } from '../token.service';

import { Candidate, Token } from '../types';

import { LoadingBarService } from '@ngx-loading-bar/core';

@Component({
  selector: 'app-nominations',
  templateUrl: './nominations.component.html',
  styleUrls: ['./nominations.component.css']
})
export class NominationsComponent implements OnInit {

  candidates: Candidate[];

  constructor(
    private userService: UserService,
    private token: TokenService,
    private loader: LoadingBarService
  ) { }

  ngOnInit() {
    this.loader.start();
    this.userService.getNominations()
      .subscribe(candidates => {
        this.candidates = candidates;
        this.loader.complete();
      });
  }

}
