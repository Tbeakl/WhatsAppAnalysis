import { Component, Input, OnInit } from '@angular/core';
import { GroupName } from '../app.component';

@Component({
    selector: 'group-names',
    templateUrl: './group-names.component.html',
    styleUrls: ['./group-names.component.less'],
})
export class GroupNamesComponent implements OnInit {
    
    @Input() data?: GroupName[];
    
    constructor() {}

    ngOnInit(): void {}

    public convertDateForOutput(date: Date){
        return `${date.toLocaleDateString()} - ${date.toLocaleTimeString().substring(0,date.toLocaleTimeString().length - 3)}`;
    }
}
