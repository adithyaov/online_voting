import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Candidate } from './types';
import { CANDIDATES } from './mock-data';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  getNominations(): Observable<Candidate[]> {
    return of(CANDIDATES);
  }

	constructor() { }

}
