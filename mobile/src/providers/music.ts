import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';

export class Song {
  name: string;
  filename: string;
}

@Injectable()
export class MusicService {

  public songs: Array<Song> = [];
  private loaderObs: Observable<Array<Song>> = null;

  constructor(public http: Http) {
    console.log('Hello Music Provider');
  }

  public loadSongs() {
    this.loaderObs = Observable.create((observer) => {
      this.http.get('/sapi/music/')
        .map(res => res.json().data)
        .subscribe(
        images => {
          console.log("Images ....");
          console.log(images);
          for (let i of images) {
            console.log(i);
            let img = new Song;
            img.name = i.attributes.name;
            img.filename = i.attributes.filename;
            this.songs.push(img);
          }
          observer.next(this.songs);
        }
        )
    });
    return this.loaderObs;
  }

  public play(songname: string) {
    this.http.post('/sapi/music/' + songname, '').subscribe(
      (x) => console.log("Play " + songname)
    );
  }

  public volUp() {
    this.http.put('/sapi/music/volume+', '').subscribe(
      (x) => console.log("Music volume +")
    )
  }

  public volDown() {
    this.http.put('/sapi/music/volume-', '').subscribe(
      (x) => console.log("Music volume -")
    )
  }

  public stop() {
    this.http.put('/sapi/music/stop', '').subscribe(
      (x) => console.log("Music stop")
    )
  }

  public pause() {
    this.http.put('/sapi/music/pause', '').subscribe(
      (x) => console.log("Music pause")
    )
  }


}
