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

  constructor(
    private ballotService: BallotService,
    private token: TokenService
  ) { }

  ngOnInit() {
    this.ballotService.candidatesObservable(this.ballot.code)
      .subscribe(candidates => {
        this.candidates = candidates;
      });
  }

}
