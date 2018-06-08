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
  voted = false;

  constructor(
    private ballotService: BallotService,
    private token: TokenService
  ) { }

  ngOnInit() {
    this.ballotService.candidatesObservable(this.ballot.code)
      .subscribe(candidates => {
        this.candidates = candidates;
        console.log(this.candidates);
      });
  }
}
