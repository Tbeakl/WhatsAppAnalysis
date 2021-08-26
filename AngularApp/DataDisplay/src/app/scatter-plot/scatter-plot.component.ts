import { Component, Input, OnInit } from '@angular/core';

@Component({
    selector: 'scatter-plot',
    templateUrl: './scatter-plot.component.html',
    styleUrls: ['./scatter-plot.component.less']
})
export class ScatterPlotComponent implements OnInit {

    @Input() displayedData?: any[];
	@Input() title: string = "";
	@Input() colourScheme = {};
    
    @Input() xAxisLabel: string = "";
    @Input() yAxisLabel: string = "";

    constructor() { }

    ngOnInit(): void {
    }

}
