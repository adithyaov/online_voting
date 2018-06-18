import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';

import { BallotBtnComponent, UserCardComponent } from './components';

@NgModule({
  declarations: [AppComponent, BallotBtnComponent, UserCardComponent],
  imports: [BrowserModule],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}
