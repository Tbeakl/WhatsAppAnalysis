import { Component, OnInit } from '@angular/core';
import summaryDataImport from '../../data/Summary.json';
import dateUsageDataImport from '../../data/DateSummary.json';
import userSummaryImport from '../../data/UserSummary.json'
import groupNameImport from '../../data/GroupNames.json'

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.less'],
})
export class AppComponent implements OnInit {
    public title = 'WhatsApp Data Analysis';
    public summaryData: summaryDataStructure = summaryDataImport;
    public dateUsageData: valueByDate[] = this.getNumberOfMessagesSent(dateUsageDataImport);
    public dateAverageLengthData: valueByDate[] = this.getAverageMessageLength(dateUsageDataImport);
    public numberOfMessagesByUser: BarChartPoint[] = this.getNumberOfMessagesSentByUser(userSummaryImport);
    public averageMessageLengthByUsers: BarChartPoint[] = this.getAverageMessageLengthByUser(userSummaryImport);
    public averageMessageLengthAgainstNumberOfMessages: any[] = this.getAverageMessageLengthAgaistNumberOfMessages(userSummaryImport);
    public groupNamesImported: GroupName[] = groupNameImport;
    public groupNames: GroupName[] = [];
    public userSummaries: userSummaryDataStructure[] = userSummaryImport;

    public ngOnInit() {
        this.groupNames = this.groupNamesImported.map(item => {
            return {
                Name: item.Name,
                Date: new Date(item.Date),
                LengthOfTime: item.LengthOfTime
            }
        });
    }

    public getNumberOfMessagesSent(dateUsageDataImport: dateData[]): valueByDate[] {
        const output: valueByDate[] = [];
        for (let index = 0; index < dateUsageDataImport.length; index++) {
            const element = dateUsageDataImport[index];
            const outputValue: valueByDate = {
                Date: new Date(element.Date),
                Value: element.NumberOfMessages,
            };
            output.push(outputValue);
        }
        return output;
    }

    public getAverageMessageLength(dateUsageDataImport: dateData[]): valueByDate[] {
        const output: valueByDate[] = [];
        for (let index = 0; index < dateUsageDataImport.length; index++) {
            const element = dateUsageDataImport[index];
            const outputValue: valueByDate = {
                Date: new Date(element.Date),
                Value: element.AverageMessageLengthWords,
            };
            output.push(outputValue);
        }
        return output;
    }

    public getNumberOfMessagesSentByUser(userSummaryImport: userSummaryDataStructure[]): BarChartPoint[] {
        let output: BarChartPoint[] = [];
        userSummaryImport.forEach(user => {
            output.push({x: user.User, y:user.NumberOfMessagesSent});
        });
        return output;
    }

    public getAverageMessageLengthByUser(userSummaryImport: userSummaryDataStructure[]): BarChartPoint[] {
        let output: BarChartPoint[] = [];
        userSummaryImport.forEach(user => {
            output.push({x: user.User, y:user.AverageMessageLength});
        });
        return output;
    }

    public getAverageMessageLengthAgaistNumberOfMessages(userSummaryImport: userSummaryDataStructure[]): any[] {
        let output: any[] = [];
        userSummaryImport.forEach(user => {
            output.push({
                "name": user.User,
                "series": [
                    {
                        "name": "",
                        "x": user.NumberOfMessagesSent,
                        "y": user.AverageMessageLength,
                        "r": 1
                    }
                ]
            });
        });
        return output;
    }
}

export interface summaryDataStructure {
    MostPopularMessage: string;
    MostPopularMessageCount: number;
    NumberOfMediaMessages: number;
    NumberOfDeletedMessages: number;
    NumberOfMessagesSent: number;
    AverageMessageLength: number;
    NumberOfDaysWithActivity: number;
    NumberOfDaysAnalysed: number;
    MostActiveUser: string;
    NumberOfMessagesByMostActiveUser: number;
}

export interface userSummaryDataStructure {
    User: string;
    NumberOfMessagesSent: number;
    NumberOfMessagesDeleted: number;
    NumberOfMediaMessages: number;
    AverageMessageLength: number;
    MostCommonMessage: string;
    MostCommonMessageCount: number;
}

export interface valueByDate {
    Date: Date;
    Value: number;
}

export interface dateData {
    Date: Date;
    NumberOfMessages: number;
    AverageMessageLengthWords: number;
    UserWhoSentLongestMessage: string;
    LengthOfLongestMessage: number;
}

export interface BarChartPoint {
    x: string;
    y: number;
}

export interface GroupName {
    Name: string;
    Date: Date;
    LengthOfTime: number;
}