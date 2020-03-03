import { BenchmarkMeter } from './BenchmarkMeter';

export class TransactionMeter extends BenchmarkMeter {
    private _transactionCounter: number;

    public constructor() {
        super();
    }

    public incrementTransactionCounter(): void {
        this._transactionCounter++;
    }

    public setTransactionCounter(value: number): void {
        this._transactionCounter = value;
    }

    protected performMeasurement(): number {
        let currentTime = Date.now();
        let durationInMsecs = currentTime - this._lastMeasuredTime;
        let result = this._transactionCounter * 1000 / durationInMsecs;
        this._lastMeasuredTime = currentTime;
        this._transactionCounter = 0;
        return result;
    }
}
