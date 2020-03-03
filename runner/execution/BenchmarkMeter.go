import { Measurement } from '../results/Measurement';

export abstract class BenchmarkMeter {
    protected _lastMeasuredTime: number;
    private _currentValue:  number;
    private _minValue: number;
    private _maxValue: number;
    private _averageValue: number;
    private _sumOfValues: number;
    private _numberOfMeasurements: number;

    public constructor() {
        this.clear();
    }

    public get measurement(): Measurement {
        return new Measurement(this.currentValue, this.minValue, 
            this.averageValue, this.maxValue);
    }

    public get lastMeasuredTime(): number {
        return this._lastMeasuredTime; 
    }
    
    public set lastMeasuredTime(value: number) {
        this._lastMeasuredTime = value;
    }

    public get currentValue(): number {
        return this._currentValue;
    }
    
    public set currentValue(value: number) {
        this._currentValue = value;
    }

    public get minValue(): number {
        return this._minValue < Number.MAX_VALUE ? this._minValue : 0;
    }
    
    public set minValue(value: number) {
        this._minValue = value;
    }

    public get maxValue(): number {
        return this._maxValue > Number.MIN_VALUE ? this._maxValue : 0; 
    }
    
    public set maxValue(value: number) {
        this._minValue = value;
    }

    public get averageValue(): number {
        return this._averageValue; 
    }
    
    public set averageValue(value: number) {
        this._averageValue = value;
    }

    public clear(): void {
        this._lastMeasuredTime = Date.now();
        this._currentValue = this.performMeasurement();
        this._minValue = Number.MAX_VALUE;
        this._maxValue = Number.MIN_VALUE;
        this._averageValue = 0;
        this._sumOfValues = 0;
        this._numberOfMeasurements = 0;
    }

    protected calculateAggregates(): void {
        this._sumOfValues += this._currentValue;
        this._numberOfMeasurements++;
        this._averageValue = this._sumOfValues / this._numberOfMeasurements;
        this._maxValue = Math.max(this._maxValue, this._currentValue);
        this._minValue = Math.min(this._minValue, this._currentValue);
    }

    public measure(): number {
        this._currentValue = this.performMeasurement();
        this._lastMeasuredTime = Date.now();
        this.calculateAggregates();
        return this._currentValue;
    }

    protected abstract performMeasurement(): number;
}
