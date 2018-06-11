import { Component, OnInit } from '@angular/core';
import { BallotService } from '../ballot.service';
import { TokenService } from '../token.service';
import { Ballot, Token } from '../types';
import { LoadingBarService } from '@ngx-loading-bar/core';

@Component({
  selector: 'app-ballots',
  templateUrl: './ballots.component.html',
  styleUrls: ['./ballots.component.css']
})
export class BallotsComponent implements OnInit {

  constructor(
    private ballotService: BallotService,
    private token: TokenService,
    private loader: LoadingBarService
  ) { }

  ballots: Ballot[];

  ngOnInit() {
    this.getBallots();
  }



  getBallots(): void {
    this.loader.start();
    this.ballotService.ballotsObservable()
      .subscribe(ballots => {
        this.ballots = ballots;
        this.loader.complete();
      });
  }


}
