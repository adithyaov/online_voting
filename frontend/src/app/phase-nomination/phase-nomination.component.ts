import { Component, OnInit, Input } from '@angular/core';
import { BallotService } from '../ballot.service';
import { TokenService } from '../token.service';

import { Candidate, Ballot, Token } from '../types';

@Component({
  selector: 'app-phase-nomination',
  templateUrl: './phase-nomination.component.html',
  styleUrls: ['./phase-nomination.component.css']
})
export class PhaseNominationComponent implements OnInit {

  @Input() ballot: Ballot;
  candidates: Candidate[];
  token: Token;

  constructor(
    private ballotService: BallotService,
    private tokenService: TokenService
  ) { }

  ngOnInit() {
    this.token = this.tokenService.currentToken();
    this.ballotService.getCandidates(this.ballot.code)
      .subscribe(candidates => this.candidates = candidates);
  }

}
