import { Component } from '@angular/core';
import { Ballot, Candidate } from './types';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app';
  test: Candidate = {
    user: {
      name: 'Adithya Kumar',
      email: '111501017@smail.iitpkd.ac.in',
      role_code: 'ccc',
      picture: 'ddd'
    },
    ballot_code: 'CCCC',
    details: 'llll',
    nominee1: { valid: true, string: '111511101@smail.iitpkd.ac.in' },
    nominee2: { valid: true, string: 'kasi@iitpkd.ac.in' }
  };
}
