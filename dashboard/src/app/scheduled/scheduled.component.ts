import { Component, OnInit } from '@angular/core';
import { Scheduled } from '../models/scheduled';
import { AjaxServiceService } from '../services/ajax/ajax-service.service';
import { HttpClient } from '@angular/common/http';
import * as moment from 'moment';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-scheduled',
  templateUrl: './scheduled.component.html',
  styleUrls: ['./scheduled.component.css']
})
export class ScheduledComponent implements OnInit {

  sdata:Array<Scheduled>;
  constructor(private asvs: AjaxServiceService, private http: HttpClient,private titleService: Title) { 
    this.sdata = new Array<Scheduled>();
    this.titleService.setTitle("Scheduled Tasks");
  }

  ngOnInit(): void {
    this.getScheduled();
  }

  getDateFiff(date:Date){
    return moment(date).fromNow()
  }

  getScheduled(){
    this.asvs.getScheduled().subscribe((data) => {
      console.log(data);
      this.sdata = data;
    })
  }
}
