let util = require('util');

export class Formatter {

    public static padLeft(value: string, length: number, padSymbol: string): string {
        let output = '';
        output += padSymbol;
        output += value;
        output += padSymbol;

        while (output.length < length + 2) {
            output = padSymbol + output;
        }

        return output;
    }

    public static padRight(value: string, length: number, padSymbol: string): string {
        let output = '';
        output += padSymbol;
        output += value;
        output += padSymbol;

        while (output.length < length + 2) {
            output += padSymbol;
        }

        return output;
    }

    public static formatNumber(value: number, decimals: number = 2): string {
        value = value || 0;
        return value.toFixed(decimals || 2);
    }

    public static formatDate(date: Date): string {
        date = date || new Date();
        let value = date.toISOString();
        let pos = value.indexOf('T');
        return value.substring(0, pos);
    }

    public static formatTime(date: Date) {
        date = date || new Date();
        let value = date.toISOString();
        let pos = value.indexOf('T');
        value = value.substring(pos + 1);
        pos = value.indexOf('.');
        return pos > 0 ? value.substring(0, pos) : value;
    }

    public static formatTimeSpan(ticks: number): string {
        let millis = (ticks % 1000).toFixed(0);
        let seconds = ((ticks / 1000) % 60).toFixed(0);
        let minutes = ((ticks / 1000 / 60) % 60).toFixed(0);
        let hours = (ticks / 1000 / 60 / 60).toFixed(0);
        return util.format("%d:%d:%d.%d", hours, minutes, seconds, millis);
    }

}