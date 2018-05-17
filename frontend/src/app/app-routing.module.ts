import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BallotsComponent }     from './ballots/ballots.component';
import { BallotDetailsComponent }     from './ballot-details/ballot-details.component';

const routes: Routes = [
  { path: '', component: BallotsComponent },
  { path: 'ballots', component: BallotsComponent },
  { path: 'ballot/:code', component: BallotDetailsComponent }
];

@NgModule({
  exports: [ RouterModule ],
  imports: [ RouterModule.forRoot(routes) ],
})
export class AppRoutingModule {}