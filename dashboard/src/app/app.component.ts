import { Component } from '@angular/core';
import {environment} from '../environments/environment';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'go-shoot';
  public version: string = environment.appVersion
  public constructor(private titleService: Title ) {
    this.titleService.setTitle("Dashboard");
  }
}
