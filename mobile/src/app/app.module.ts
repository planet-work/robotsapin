import { NgModule, ErrorHandler } from '@angular/core';
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';
import { MyApp } from './app.component';

import { StatusPage } from '../pages/status/status';
import { MusicPage } from '../pages/music/music';
import { DisplayPage } from '../pages/display/display';
import { TopperPage } from '../pages/topper/topper';
import { SensorsPage } from '../pages/sensors/sensors';

import { StatusService } from '../providers/status';
import { MusicService } from '../providers/music';
import { DisplayService } from '../providers/display';
import { TopperService } from '../providers/topper';



@NgModule({
  declarations: [
    MyApp,
    MusicPage,
    DisplayPage,
    StatusPage,
    TopperPage,
    SensorsPage
  ],
  imports: [
    IonicModule.forRoot(MyApp)
  ],
  bootstrap: [IonicApp],
  entryComponents: [
    MyApp,
    MusicPage,
    DisplayPage,
    StatusPage,
    TopperPage,
    SensorsPage
  ],
  providers: [{ provide: ErrorHandler, useClass: IonicErrorHandler },
                MusicService, DisplayService, TopperService, StatusService]
})
export class AppModule { }
