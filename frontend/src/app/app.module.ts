import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BallotsComponent } from './ballots/ballots.component';
import { AppRoutingModule } from './/app-routing.module';
import { BallotDetailsComponent } from './ballot-details/ballot-details.component';

@NgModule({
  declarations: [
    AppComponent,
    BallotsComponent,
    BallotDetailsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
