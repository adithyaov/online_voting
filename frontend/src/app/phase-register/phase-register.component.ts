import { Component, OnInit, Input } from '@angular/core';
import { BallotService } from '../ballot.service';
import { TokenService } from '../token.service';

import { Candidate, Ballot, Token } from '../types';

@Component({
  selector: 'app-phase-register',
  templateUrl: './phase-register.component.html',
  styleUrls: ['./phase-register.component.css']
})
export class PhaseRegisterComponent implements OnInit {

  @Input() ballot: Ballot;
  documentUrl = 'undefined';

  constructor(
    private ballotService: BallotService,
    private token: TokenService
  ) { }

  ngOnInit() {}
}
