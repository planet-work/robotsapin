import { Component } from '@angular/core';

import { NavController } from 'ionic-angular';
import { MusicService } from '../../providers/music';


@Component({
  selector: 'page-music',
  templateUrl: 'music.html'
})
export class MusicPage {

  constructor(public navCtrl: NavController,
  public music: MusicService) {

    this.music.loadSongs().subscribe(
      x => console.log("Music: songs loaded")
    )
    
  }

}
