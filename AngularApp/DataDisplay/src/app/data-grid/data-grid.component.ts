import { Component, Input, OnInit } from '@angular/core';
import { summaryDataStructure, valueByDate } from '../app.component';

@Component({
  selector: 'data-grid',
  templateUrl: './data-grid.component.html',
  styleUrls: ['./data-grid.component.less']
})
export class DataGridComponent implements OnInit {

  @Input() summaryData?: summaryDataStructure;
  @Input() dateUsageData: valueByDate[] = [];
  @Input() dateAverageLengthData: valueByDate[] = [];


  constructor() { }

  ngOnInit(): void {
  }

}
