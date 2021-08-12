import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { uk_UA } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import uk from '@angular/common/locales/uk';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzGridModule } from 'ng-zorro-antd/grid';
import { NzTableModule } from 'ng-zorro-antd/table';
import { DataGridComponent } from './data-grid/data-grid.component';
import { BasicSummaryComponent } from './basic-summary/basic-summary.component';
import { HeatMapCalendarComponent } from './heat-map-calendar/heat-map-calendar.component';
import { NgxChartsModule } from '@swimlane/ngx-charts';
import { BarChartComponent } from './bar-chart/bar-chart.component';
import { GroupNamesComponent } from './group-names/group-names.component';

registerLocaleData(uk);

@NgModule({
  declarations: [
    AppComponent,
    DataGridComponent,
    BasicSummaryComponent,
    HeatMapCalendarComponent,
    BarChartComponent,
    GroupNamesComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    NzButtonModule,
    NzCardModule,
    NzGridModule,
    NzTableModule,
    NgxChartsModule
  ],
  providers: [{ provide: NZ_I18N, useValue: uk_UA }],
  bootstrap: [AppComponent]
})
export class AppModule { }
