import { Component, Input, OnInit } from '@angular/core';
import { BarChartPoint } from '../app.component';

@Component({
    selector: 'bar-chart',
    templateUrl: './bar-chart.component.html',
    styleUrls: ['./bar-chart.component.less'],
})
export class BarChartComponent implements OnInit {
    @Input() title: string = '';
    @Input() inputData: BarChartPoint[] = [];
    @Input() xAxisLabel: string = '';
    @Input() yAxisLabel: string = '';
    @Input() colourScheme = {};
    @Input() sortingScheme: string = '';

    public displayData :any[] = [];

    constructor() {}

    ngOnInit(): void {
        //First sort the inputted data
        this.inputData.sort((a, b) => {
            switch (this.sortingScheme) {
                case 'Asc':
                    return a.y - b.y;
                case 'Des':
                    return b.y - a.y;
                default:
                    return 0;
            }
        });

        this.displayData = this.inputData.map((item) =>  {
            return {
                name: item.x,
                value: item.y
            }
        });
    }
}
