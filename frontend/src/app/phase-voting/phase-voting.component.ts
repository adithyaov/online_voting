import { Component, OnInit, Input } from '@angular/core';
import { BallotService } from '../ballot.service';
import { TokenService } from '../token.service';

import { Candidate, Ballot, Token } from '../types';
import { LoadingBarService } from '@ngx-loading-bar/core';

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
    private token: TokenService,
    private loader: LoadingBarService
  ) { }

  ngOnInit() {
    this.loader.start();
    this.ballotService.candidatesObservable(this.ballot.code)
      .subscribe(candidates => {
        this.candidates = candidates;
        this.loader.complete();
      });
  }
}
