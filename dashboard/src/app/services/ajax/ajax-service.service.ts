import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AjaxServiceService {

  constructor(private http: HttpClient) { }

  getEnqueue():Observable<any>{
    return this.http.get(environment.enqueuedUri)
  }
  getSuccessful():Observable<any>{
    return this.http.get(environment.successfulUri)
  }
  getScheduled():Observable<any>{
    return this.http.get(environment.scheduledUri)
  }
}
