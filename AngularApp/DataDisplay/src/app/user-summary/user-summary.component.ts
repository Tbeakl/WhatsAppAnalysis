import { Component, Input, OnInit } from '@angular/core';
import { summaryDataStructure, userSummaryDataStructure } from '../app.component';

@Component({
    selector: 'user-summary',
    templateUrl: './user-summary.component.html',
    styleUrls: ['./user-summary.component.less']
})
export class UserSummaryComponent implements OnInit {

    @Input() importedData?: userSummaryDataStructure[];
    public userOptions = () => this.importedData ? this.importedData.filter(item => {
        if (this.username == '' || this.username == null) return true;
        return item.User.trim().toUpperCase().includes(this.username.trim().toUpperCase());
    }).map(item => item.User) : [];
    public userToDisplay?: userSummaryDataStructure;

    public username = "";

    constructor() { }

    public updateUser(){
        this.userToDisplay = this.importedData?.map(item => item.User.trim().toUpperCase()).includes(this.username.trim().toUpperCase()) ? this.importedData?.filter(item => item.User.trim().toUpperCase() == this.username.trim().toUpperCase())[0] : undefined;
    }

    ngOnInit(): void {

    }

}
