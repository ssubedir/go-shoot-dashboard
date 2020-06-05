import { Component, OnInit } from '@angular/core';
import { AjaxServiceService } from '../services/ajax/ajax-service.service';
import { HttpClient } from '@angular/common/http';
import { Enqueued } from '../models/enqueued';
import * as moment from 'moment';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-enqueued',
  templateUrl: './enqueued.component.html',
  styleUrls: ['./enqueued.component.css']
})
export class EnqueuedComponent implements OnInit {

  qdata:Array<Enqueued>;
  constructor(private asvs: AjaxServiceService, private http: HttpClient,private titleService: Title) {
    this.qdata = new Array<Enqueued>();
    this.titleService.setTitle("Enqueued Tasks");
   }


  ngOnInit(): void {
    this.getEnqueue();
    
  }


  getDateFiff(date:Date){
    return moment(date).fromNow()
  }
  

  getEnqueue(){
    this.asvs.getEnqueue().subscribe((data) => {
      console.log(data);
      this.qdata = data;
    })
  }

}
