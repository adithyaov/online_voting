import { Component, OnInit, Input } from '@angular/core';
import { BallotService } from '../ballot.service';
import { TokenService } from '../token.service';

import { Candidate, Ballot, Token } from '../types';

@Component({
  selector: 'app-phase-voting',
  templateUrl: './phase-voting.component.html',
  styleUrls: ['./phase-voting.component.css']
})
export class PhaseVotingComponent implements OnInit {

  @Input() ballot: Ballot;
  candidates: Candidate[];
  token: Token;
  voted = false;

  constructor(
    private ballotService: BallotService,
    private tokenService: TokenService
  ) { }

  ngOnInit() {
    this.token = this.tokenService.getToken();
    this.ballotService.getCandidates(this.ballot.code)
      .subscribe(candidates => this.candidates = candidates);
  }
}
