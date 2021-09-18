import { Component, Input, OnInit } from '@angular/core';
import { LineChartPoint } from '../app.component';
import * as shape from 'd3-shape';


@Component({
    selector: 'line-chart',
    templateUrl: './line-chart.component.html',
    styleUrls: ['./line-chart.component.less']
})
export class LineChartComponent implements OnInit {
    @Input() title: string = '';
    @Input() inputData: LineChartPoint[] = [];
    @Input() xAxisLabel: string = '';
    @Input() yAxisLabel: string = '';
    @Input() colourScheme = {};
    @Input() sortingScheme: string = '';
    public curve = shape.curveCardinal;
    public displayData: any[] = [];


    constructor() { }

    ngOnInit(): void {
        this.displayData = [{
            name: "Test",
            series: this.inputData.map(item =>  {
                return {
                    name: item.x,
                    value: item.y
                }
            })
        }];
    }

}
