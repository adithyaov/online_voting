import { Component, Input, Output, EventEmitter } from '@angular/core';

import { Ballot } from '../../types';
import { phaseIconProps } from '../../common';

@Component({
  selector: 'app-ballot-btn',
  templateUrl: './ballot-btn.component.html',
  styleUrls: ['./ballot-btn.component.css']
})
export class BallotBtnComponent {
  @Input() ballot: Ballot;
  @Output() openBallot = new EventEmitter<string>();

  phaseIconProps = phaseIconProps;

  open(ballotName: string) {
    this.openBallot.emit(ballotName);
  }

  constructor() {}
}
