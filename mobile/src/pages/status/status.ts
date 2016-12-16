import { Component } from '@angular/core';
import { NavController, NavParams } from 'ionic-angular';

import { StatusService } from '../../providers/status';

@Component({
  selector: 'page-status',
  templateUrl: 'status.html'
})
export class StatusPage {

  constructor(public navCtrl: NavController, public navParams: NavParams,
         public status: StatusService) { }

  ionViewDidLoad() {
    console.log('ionViewDidLoad StatusPage');
    this.status.loadStatus().subscribe(
      (x) => console.log("Status : got it")
    );
  }

}
