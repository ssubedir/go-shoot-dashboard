import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { EnqueuedComponent } from './enqueued/enqueued.component';
import {HttpClientModule} from '@angular/common/http';
import { AjaxServiceService } from './services/ajax/ajax-service.service';
import { ScheduledComponent } from './scheduled/scheduled.component';
import { SuccessfulComponent } from './successful/successful.component';
import { HashLocationStrategy, LocationStrategy } from '@angular/common';

@NgModule({
  declarations: [
    AppComponent,
    EnqueuedComponent,
    ScheduledComponent,
    SuccessfulComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [AjaxServiceService,{provide: LocationStrategy, useClass: HashLocationStrategy}],
  bootstrap: [AppComponent]
})
export class AppModule { }
