import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BallotsComponent } from './ballots/ballots.component';
import { AppRoutingModule } from './/app-routing.module';
import { BallotDetailsComponent } from './ballot-details/ballot-details.component';
import { NominationsComponent } from './nominations/nominations.component';
import { PhaseRegisterComponent } from './phase-register/phase-register.component';
import { PhaseVotingComponent } from './phase-voting/phase-voting.component';
import { PhaseNominationComponent } from './phase-nomination/phase-nomination.component';

@NgModule({
  declarations: [
    AppComponent,
    BallotsComponent,
    BallotDetailsComponent,
    NominationsComponent,
    PhaseRegisterComponent,
    PhaseVotingComponent,
    PhaseNominationComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }