import { Component, OnInit } from '@angular/core';
import summaryDataImport from '../../data/Summary.json'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less']
})
export class AppComponent implements OnInit {
  public title = 'WhatsApp Data Analysis';
  public summaryData: summaryDataStructure = summaryDataImport;
  
  ngOnInit() {

  }
}


export interface summaryDataStructure {
  MostPopularMessage: string;
  MostPopularMessageCount: number;
  TotalNumberOfMessages: number;
}