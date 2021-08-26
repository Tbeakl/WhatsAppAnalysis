import { Component, Input, OnInit } from '@angular/core';
import { BarChartPoint, GroupName, summaryDataStructure, userSummaryDataStructure, valueByDate } from '../app.component';

@Component({
	selector: 'data-grid',
	templateUrl: './data-grid.component.html',
	styleUrls: ['./data-grid.component.less'],
})
export class DataGridComponent implements OnInit {
	@Input() summaryData?: summaryDataStructure;
	@Input() dateUsageData?: valueByDate[];
	@Input() dateAverageLengthData?: valueByDate[];
	@Input() numberOfMessagesByUser?: BarChartPoint[];
	@Input() groupNames?: GroupName[];
	@Input() userSummaries?: userSummaryDataStructure[];
	@Input() averageMessageLengthByUsers?: BarChartPoint[];
	@Input() averageMessageLengthAgainstNumberOfMessages?: any[];

	public colorScheme = {
		domain: [
			'#647c8a',
			'#3f51b5',
			'#2196f3',
			'#00b862',
			'#afdf0a',
			'#a7b61a',
			'#f3e562',
			'#ff9800',
			'#ff5722',
			'#ff4514',
		],
	};

	constructor() {}

	ngOnInit(): void {}
}
