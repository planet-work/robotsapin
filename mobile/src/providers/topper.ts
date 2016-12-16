import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';

export class LedSequence {
  id: number;
  data: Array<string>;
}

@Injectable()
export class TopperService {


  public sequences: Array<LedSequence> = [];
  private loaderObs: Observable<Array<LedSequence>> = null;

  constructor(public http: Http) {
    console.log('Hello Topper Provider');
  }

  public loadSequences() {
    this.loaderObs = Observable.create((observer) => {
      this.http.get('/sapi/topper/')
        .map(res => res.json().data)
        .subscribe(
        images => {
          console.log("Images ....");
          console.log(images);
          for (let i of images) {
            console.log(i);
            let img = new LedSequence;
            img.id = i.attributes.id;
            img.data = i.attributes.data;
            this.sequences.push(img);
          }
          observer.next(this.sequences);
        }
        )
    });
    return this.loaderObs;
  }

}
