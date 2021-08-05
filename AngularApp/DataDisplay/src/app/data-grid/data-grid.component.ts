import { Component, Input, OnInit } from '@angular/core';
import { summaryDataStructure } from '../app.component';

@Component({
  selector: 'app-data-grid',
  templateUrl: './data-grid.component.html',
  styleUrls: ['./data-grid.component.less']
})
export class DataGridComponent implements OnInit {

  @Input() summaryData?: summaryDataStructure;

  constructor() { }

  ngOnInit(): void {
  }

}
