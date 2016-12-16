import { Component } from '@angular/core';

import { NavController, NavParams } from 'ionic-angular';
import { DisplayService } from '../../providers/display';


@Component({
  selector: 'page-display',
  templateUrl: 'display.html'
})
export class DisplayPage {
  selectedItem: any;

  constructor(public navCtrl: NavController, public navParams: NavParams,
      public display: DisplayService) {
  }
}
