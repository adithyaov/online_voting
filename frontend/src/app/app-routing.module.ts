import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BallotsComponent } from './ballots/ballots.component';
import { BallotDetailsComponent } from './ballot-details/ballot-details.component';
import { NominationsComponent } from './nominations/nominations.component';

const routes: Routes = [
  { path: '', component: BallotsComponent },
  { path: 'ballots', component: BallotsComponent },
  { path: 'ballot/:code', component: BallotDetailsComponent },
  { path: 'nominations', component: NominationsComponent }
];

@NgModule({
  exports: [ RouterModule ],
  imports: [ RouterModule.forRoot(routes) ],
})
export class AppRoutingModule {}
