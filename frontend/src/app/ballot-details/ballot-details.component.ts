import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';

import { BallotService } from '../ballot.service';
import { Ballot, Candidate, Token } from '../types'
import { TokenService } from '../token.service';

@Component({
  selector: 'app-ballot-details',
  templateUrl: './ballot-details.component.html',
  styleUrls: ['./ballot-details.component.css']
})
export class BallotDetailsComponent implements OnInit {

	ballot: Ballot;
	token: Token;

  constructor(
    private route: ActivatedRoute,
    private ballotService: BallotService,
    private location: Location,
    private tokenService: TokenService
  ) { }

  ngOnInit(): void {
  	this.token = this.tokenService.getToken();
    const code = this.route.snapshot.paramMap.get('code');
		this.ballotService.getBallot(code)
		  .subscribe(ballot => this.ballot = ballot);
  }


}
