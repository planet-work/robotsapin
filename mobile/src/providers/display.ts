import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';

export class DisplayImage {
  name: string;
  filename: string;
  data: string;
}

@Injectable()
export class DisplayService {

  public images: Array<DisplayImage> = [];
  private loaderObs: Observable<Array<DisplayImage>> = null;

  constructor(public http: Http) {
    console.log('Hello Display Provider');
  }

  public loadImages() {
    this.loaderObs = Observable.create((observer) => {
      this.http.get('/sapi/display/')
        .map(res => res.json().data)
        .subscribe(
        images => {
          console.log("Images ....");
          console.log(images);
          for (let i of images) {
            console.log(i);
            let img = new DisplayImage;
            img.name = i.attributes.name;
            img.filename = i.attributes.filename;
            img.data = i.attributes.data;
            this.images.push(img);
          }
          observer.next(this.images);
        }
        )
    });
    return this.loaderObs;
  }

  public clear() {
    return this.http.post('/sapi/display/clear', 'filename=');
  }

  public show(filename: string) {
    return this.http.post('/sapi/display/clear', 'filename=' + filename);
  }
}
