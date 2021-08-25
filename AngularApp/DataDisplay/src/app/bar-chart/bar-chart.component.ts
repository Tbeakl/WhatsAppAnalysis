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

    public listOfOption: string[] = [];
    public listOfSelectedValue: string[] = [];

    public displayData :any[] = [];

    constructor() {}

    public ngOnInit(): void {
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
        this.listOfSelectedValue = [];
        this.listOfOption = this.inputData.map(item => item.x);
        this.changeDisplayedData();
    }

    public changeDisplayedData() {
        this.displayData = this.inputData.filter(item => {
            if (this.listOfSelectedValue.length == 0) return true;
            return this.listOfSelectedValue.includes(item.x);
        }).map(item =>  {
            return {
                name: item.x,
                value: item.y
            }
        });
    }
}
