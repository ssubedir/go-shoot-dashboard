import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { EnqueuedComponent } from './enqueued/enqueued.component';
import { ScheduledComponent } from './scheduled/scheduled.component';
import { SuccessfulComponent } from './successful/successful.component';

const routes: Routes = [
  { path: '', component: EnqueuedComponent },
  { path: 'tasks/enqueued', component: EnqueuedComponent },
  { path: 'tasks/scheduled', component: ScheduledComponent },
  { path: 'tasks/successful', component: SuccessfulComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
