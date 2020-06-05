import { Component, OnInit } from '@angular/core';
import { AjaxServiceService } from '../services/ajax/ajax-service.service';
import { HttpClient } from '@angular/common/http';
import { Successful } from '../models/successful';
import * as moment from 'moment';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-successful',
  templateUrl: './successful.component.html',
  styleUrls: ['./successful.component.css']
})
export class SuccessfulComponent implements OnInit {

  sdata:Array<Successful>;
  constructor(private asvs: AjaxServiceService, private http: HttpClient,private titleService: Title) { 
    this.sdata = new Array<Successful>();
    this.titleService.setTitle("Successful Tasks");
  }

  ngOnInit(): void {
    this.getSuccessful();
  }

  getDateFiff(date:Date){
    return moment(date).fromNow()
  }
  
  getSuccessful(){
    this.asvs.getSuccessful().subscribe((data) => {
      console.log(data);
      this.sdata = data;
    })
  }
}
