import { Component, Input, OnInit } from '@angular/core';
import { valueByDate } from '../app.component';
const monthName = new Intl.DateTimeFormat("en-us", { month: "short" });
const weekdayName = new Intl.DateTimeFormat("en-us", { weekday: "short" });

@Component({
	selector: 'heat-map-calendar',
	templateUrl: './heat-map-calendar.component.html',
	styleUrls: ['./heat-map-calendar.component.less']
})
export class HeatMapCalendarComponent implements OnInit {

	@Input() title: string = ""
	@Input() data: valueByDate[] = []

	public outputData: any[] = [];
	public minValue: number = 0;
	public maxValue: number = 100;
	// your color scheme
	colorScheme = {
		domain: [
			'#647c8a', '#3f51b5', '#2196f3', '#00b862', '#afdf0a', '#a7b61a', '#f3e562', '#ff9800', '#ff5722', '#ff4514'
		]
	};
	constructor() {
	}

	ngOnInit(): void {
		this.outputData = this.transformInput(this.data);
		this.minValue = Math.min(...this.data.map(t => t.Value));
		this.maxValue = Math.max(...this.data.map(t => t.Value));
	}

	public transformInput(data: valueByDate[]): any[] {
		// today
		const now = data[data.length - 1].Date;
		const todaysDay = now.getDate();
		const thisDay = new Date(now.getFullYear(), now.getMonth(), todaysDay);

		// Monday
		const thisMonday = new Date(thisDay.getFullYear(), thisDay.getMonth(), todaysDay - thisDay.getDay() + 1);
		const thisMondayDay = thisMonday.getDate();
		const thisMondayYear = thisMonday.getFullYear();
		const thisMondayMonth = thisMonday.getMonth();

		// Get the number of weeks before Monday to show all the data
		const timeSpanOfData = Math.abs(Math.ceil((data[data.length - 1].Date.getTime() - data[0].Date.getTime()) / (7 * 24 * 60 * 60 * 1000)))
		const calendarData = [];
		const getDate = (d: number | undefined) => new Date(thisMondayYear, thisMondayMonth, d);
		for (let week = -timeSpanOfData; week <= 0; week++) {
			const mondayDay = thisMondayDay + week * 7;
			const monday = getDate(mondayDay);

			// one week
			const series = [];
			for (let dayOfWeek = 7; dayOfWeek > 0; dayOfWeek--) {
				const date = getDate(mondayDay - 1 + dayOfWeek);

				// skip future dates
				if (date > now) {
					continue;
				}

				// value
				const value = getValueForDate(date, data);

				series.push({
					date,
					name: weekdayName.format(date),
					value
				});
			}

			calendarData.push({
				name: monday.toString(),
				series
			});
		}

		return calendarData;
	}

	calendarAxisTickFormatting(mondayString: string) {
		const monday = new Date(mondayString);
		const month = monday.getMonth();
		const day = monday.getDate();
		const year = monday.getFullYear();
		const lastSunday = new Date(year, month, day - 1);
		const nextSunday = new Date(year, month, day + 6);
		return (lastSunday.getMonth() !== nextSunday.getMonth()) ? monthName.format(nextSunday) : '';
	}
}


function getValueForDate(date: Date, data: valueByDate[]): number {
	for (let index = 0; index < data.length; index++) {
		const element = data[index];
		if (element.Date.getDate() == date.getDate() && element.Date.getMonth() == date.getMonth() && element.Date.getFullYear() == date.getFullYear()) {
			return element.Value;
		}
	}
	return 0;
}

