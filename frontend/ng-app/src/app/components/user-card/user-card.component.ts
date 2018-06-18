import { Component, Input, Output, EventEmitter } from '@angular/core';

import { Candidate } from '../../types';

@Component({
  selector: 'app-user-card',
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.css']
})
export class UserCardComponent {
  @Input() candidate: Candidate;
  @Input() phase: 'C' | 'N' | 'V' | 'S';
  @Input() currentEmail: string;

  @Output() nominateRequest = new EventEmitter<string>();
  @Output() VoteRequest = new EventEmitter<string>();

  nominate(candidateEmail: string) {
    this.nominateRequest.emit(candidateEmail);
  }

  vote(candidateEmail: string) {
    this.VoteRequest.emit(candidateEmail);
  }

  constructor() {}
}
