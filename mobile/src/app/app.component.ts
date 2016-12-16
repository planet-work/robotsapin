import { Component, ViewChild } from '@angular/core';
import { Nav, Platform } from 'ionic-angular';
import { StatusBar, Splashscreen } from 'ionic-native';

import { StatusPage } from '../pages/status/status';
import { MusicPage } from '../pages/music/music';
import { DisplayPage } from '../pages/display/display';
import { TopperPage } from '../pages/topper/topper';
import { SensorsPage } from '../pages/sensors/sensors';


import { MusicService } from '../providers/music';
import { DisplayService, DisplayImage } from '../../providers/display';
import { TopperService } from '../../providers/topper';


@Component({
  templateUrl: 'app.html'
})
export class MyApp {
  @ViewChild(Nav) nav: Nav;

  rootPage: any = StatusPage;

  pages: Array<{ title: string, component: any }>;

  constructor(public platform: Platform,
   public music: MusicService) {
    this.initializeApp();

    // used for an example of ngFor and navigation
    this.pages = [
      { title: 'État', component: StatusPage },
      { title: 'Musique', component: MusicPage },
      { title: 'Affichage', component: DisplayPage },
      { title: 'Étoile', component: TopperPage },
      { title: 'Capteurs', component: SensorsPage },
    ];

  }

  initializeApp() {
    this.platform.ready().then(() => {
      // Okay, so the platform is ready and our plugins are available.
      // Here you can do any higher level native things you might need.
      StatusBar.styleDefault();
      Splashscreen.hide();
    });
  }

  openPage(page) {
    // Reset the content nav to have just this page
    // we wouldn't want the back button to show in this scenario
    this.nav.setRoot(page.component);
  }
}
