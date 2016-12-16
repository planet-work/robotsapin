import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';

export class Status {
    display: any;
    topper: any;
    music: any;
    sensors: any;
}

@Injectable()
export class StatusService {

  public status: Status;
  private loaderObs: Observable<Array<Status>> = null;

  constructor(public http: Http) {
    console.log('Hello Display Provider');
  }

  public loadStatus() {
    this.loaderObs = Observable.create((observer) => {
      this.http.get('/sapi/status')
        .map(res => res.json().data)
        .subscribe(
        status => {
          console.log("Status ....");
          console.log(status);
          let s = new Status;
          s.display = status.attributes.display;
          s.topper = status.attributes.topper;
          s.music = status.attributes.music;
          this.status = s;
          observer.next(this.status);
        }
         )
    });
    return this.loaderObs;
  }
}
