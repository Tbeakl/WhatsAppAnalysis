import { Component, Input, OnInit } from '@angular/core';
import { summaryDataStructure } from '../app.component';

@Component({
  selector: 'basic-summary',
  templateUrl: './basic-summary.component.html',
  styleUrls: ['./basic-summary.component.less']
})
export class BasicSummaryComponent implements OnInit {

  @Input() data?: summaryDataStructure;

  constructor() { }

  ngOnInit(): void {

  }

}
