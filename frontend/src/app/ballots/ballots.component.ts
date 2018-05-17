import { Component, OnInit } from '@angular/core';
import { BallotService } from '../ballot.service';
import { TokenService } from '../token.service';
import { Ballot, Token } from '../types';

@Component({
  selector: 'app-ballots',
  templateUrl: './ballots.component.html',
  styleUrls: ['./ballots.component.css']
})
export class BallotsComponent implements OnInit {

  constructor(
    private ballotService: BallotService,
    private tokenService: TokenService
  ) { }


  token: Token

  ngOnInit() {
    this.token = this.tokenService.getToken();
    // GoOn depending on the token
    console.log(this.token == null)
  	this.getBallots();
  }

  ballots: Ballot[];

  getBallots(): void {
  	this.ballotService.getBallots()
      .subscribe(ballots => this.ballots = ballots);      
  }


}
