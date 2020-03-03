import { ResultsManager } from '../results/ResultsManager';
import { BenchmarkInstance } from '../benchmarks/BenchmarkInstance';
import { BenchmarkResult } from '../results/BenchmarkResult';

import { TransactionMeter } from './TransactionMeter';
import { CpuLoadMeter } from './CpuLoadMeter';
import { MemoryUsageMeter } from './MemoryUsageMeter';

export class ResultAggregator {
    private static readonly MaxErrorCount = 1000;

    private _results: ResultsManager;
    private _benchmarks: BenchmarkInstance[];
    private _transactionCounter: number = 0;
    private _result: BenchmarkResult = null;
    private _transactionMeter: TransactionMeter;
    private _cpuLoadMeter: CpuLoadMeter;
    private _memoryUsageMeter: MemoryUsageMeter;

    public constructor(results: ResultsManager, benchmarks: BenchmarkInstance[]) {
        this._results = results;
        this._benchmarks = benchmarks;

        this._cpuLoadMeter = new CpuLoadMeter();
        this._transactionMeter = new TransactionMeter();
        this._memoryUsageMeter = new MemoryUsageMeter();

        this.start();
    }

    public get result(): BenchmarkResult {
        return this._result;
    }

    public start(): void {
        this._result = new BenchmarkResult();
        this._result.benchmarks = this._benchmarks;
        this._result.startTime = Date.now();

        this._transactionCounter = 0;
        this._transactionMeter.clear();
        this._cpuLoadMeter.clear();
        this._memoryUsageMeter.clear();
    }

    public incrementCounter(increment: number, now?: number): void {
        now = now || Date.now();
        this._transactionCounter += increment;

        // If it's less then a second then wait
        let measureInterval = now - this._transactionMeter.lastMeasuredTime;
        if (measureInterval >= 1000 && this._result != null) {
            // Perform measurements
            this._transactionMeter.setTransactionCounter(this._transactionCounter);
            this._transactionCounter = 0;
            this._transactionMeter.measure();
            this._cpuLoadMeter.measure();
            this._memoryUsageMeter.measure();

            // Store measurement results
            this._result.elapsedTime = now - this._result.startTime;
            this._result.performanceMeasurement = this._transactionMeter.measurement;
            this._result.cpuLoadMeasurement = this._cpuLoadMeter.measurement;
            this._result.memoryUsageMeasurement = this._memoryUsageMeter.measurement;

            this._results.notifyUpdated(this._result);
        }
    }

    public sendMessage(message: string): void {
        this._results.notifyMessage(message);
    }

    public reportError(error: any): void {
        if (this._result.errors.length < ResultAggregator.MaxErrorCount)
            this._result.errors.push(error);

        this._results.notifyError(error);
    }

    public stop() : void {
        this._results.add(this._result);
    }
}