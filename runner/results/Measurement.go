export class Measurement {

    public constructor(currentValue: number, minValue: number,
        averageValue: number, maxValue: number) {
        this.currentValue = currentValue;
        this.minValue = minValue;
        this.averageValue = averageValue;
        this.maxValue = maxValue;
    }

    public currentValue: number;
    public minValue: number;
    public averageValue: number;
    public maxValue;
}