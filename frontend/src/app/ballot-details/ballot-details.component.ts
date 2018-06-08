import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';

import { BallotService } from '../ballot.service';
import { Ballot, Candidate, Token } from '../types';
import { TokenService } from '../token.service';

@Component({
  selector: 'app-ballot-details',
  templateUrl: './ballot-details.component.html',
  styleUrls: ['./ballot-details.component.css']
})
export class BallotDetailsComponent implements OnInit {

  ballot: Ballot;

  constructor(
    private route: ActivatedRoute,
    private ballotService: BallotService,
    private location: Location,
    private token: TokenService
  ) { }

  ngOnInit(): void {
    const code = this.route.snapshot.paramMap.get('code');
    this.setBallot(code);
  }

  setBallot(code): void {
    this.ballotService.ballotObservable(code)
      .subscribe(ballot => {
        this.ballot = ballot;
      });
  }


}
