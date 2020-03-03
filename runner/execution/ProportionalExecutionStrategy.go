let _ = require('lodash');
let async = require('async');

import { MeasurementType } from '../config/MeasurementType';
import { ConfigurationManager } from '../config/ConfigurationManager';
import { ResultsManager } from '../results/ResultsManager';
import { ExecutionState } from './ExecutionState';
import { BenchmarkInstance } from '../benchmarks/BenchmarkInstance';
import { BenchmarkSuiteInstance } from '../benchmarks/BenchmarkSuiteInstance';
import { BenchmarkResult } from '../results/BenchmarkResult';

import { ExecutionContext } from './ExecutionContext';
import { ExecutionStrategy } from './ExecutionStrategy';
import { ResultAggregator } from './ResultAggregator';

export class ProportionalExecutionStrategy extends ExecutionStrategy {
    private _running: boolean = false;
    private _aggregator: ResultAggregator;
    private _ticksPerTransaction: number = 0;
    private _lastExecutedTime: number;
    private _stopTime: number;
    private _benchmarkCount: number;
    private _onlyBenchmark: BenchmarkInstance;
    private _timeout: any;
    
    public constructor(configuration: ConfigurationManager, results: ResultsManager, 
        execution: any, benchmarks: BenchmarkInstance[]) {
        
        super(configuration, results, execution, benchmarks);

        this._aggregator = new ResultAggregator(results, benchmarks);
    }

    public start(callback?: (err: any) => void): void {
        if (this._running) {
            callback(null);
            return;
        }
        
        this._running = true;
        this._aggregator.start();

        this.calculateProportionalRanges();

        if (this._configuration.measurementType == MeasurementType.Nominal)
            this._ticksPerTransaction = 1000.0 / this._configuration.nominalRate;

        // Initialize and start
        async.each(
            this._suites, 
            (suite, callback) => {
                let context = new ExecutionContext(suite, this._aggregator, this);
                suite.setUp(context, callback);
            },
            (err) => {
                // Abort if initialization failed
                if (err) {
                    this._aggregator.reportError(err);
                    callback(err);
                    return;
                }
                
                // Execute benchmarks
                this.execute(callback);
            }
        );
    }

    public get isStopped(): boolean {
        return !this._running;
    }

    public stop(callback?: (err: any) => void): void {
        // Interrupt any wait
        if (this._timeout != null) {
            clearTimeout(this._timeout);
            this._timeout = null;
        }
        
        // Stop and cleanup execution
        if (this._running) {
            this._running = false;
            this._aggregator.stop();

            if (this._execution)
                this._execution.stop();

            // Deinitialize tests
            async.each(
                this._suites, 
                (suite, callback) => {
                    suite.tearDown(callback);
                },
                (err) => {
                    if (callback) callback(err);
                }
            );
        } else {
            if (callback) callback(null);
        }
    }

    private calculateProportionalRanges(): void {
        let totalProportion = 0;
        for (let benchmark of this._activeBenchmarks) {
            totalProportion += benchmark.proportion;
        }

        let startRange = 0;
        for (let benchmark of this._activeBenchmarks) {
            let normalizedProportion = benchmark.proportion / totalProportion;
            benchmark.startRange = startRange;
            benchmark.endRange = startRange + normalizedProportion;
            startRange += normalizedProportion;
        }
    }

    private chooseBenchmarkProportionally(): BenchmarkInstance {
        let proportion = Math.random();

        return _.find(this._activeBenchmarks, (benchmark) => {
            return benchmark.withinRange(proportion);
        });
    }

    private executeDelay(delay: number, callback: (err: any) => void): void {
        this._timeout = setTimeout(() => {
            this._timeout = null;
            this._lastExecutedTime = Date.now();

            callback(null)
        }, delay);
    }

    private executeBenchmark(benchmark: BenchmarkInstance, callback: (err: any) => void): void {
        try {
            if (benchmark == null || benchmark.isPassive) {
                // Delay if benchmarks are passive
                this.executeDelay(500, callback);
            } else {
                // Execute active benchmark
                benchmark.execute((err) => {
                    // Process force continue
                    if (err != null && this._configuration.forceContinue) {
                        this._aggregator.reportError(err);
                        err = null;
                    }

                    // Increment counter
                    let now = Date.now();
                    if (err == null)
                        this._aggregator.incrementCounter(1, now);

                    // Introduce delay to keep nominal rate
                    if (err == null && this._configuration.measurementType == MeasurementType.Nominal) {
                        let delay = this._ticksPerTransaction - (now - this._lastExecutedTime);
                        this._lastExecutedTime = now;

                        if (delay > 0) 
                            this.executeDelay(delay, callback);
                        else callback(err);
                    } else {
                        this._lastExecutedTime = now;
                        callback(err);
                    }
                });
            }
        } catch (ex) {
            // Process force continue
            if (this._configuration.forceContinue) {
                this._aggregator.reportError(ex);
                callback(null);
            } else {
                callback(ex);
            }
        }
    }

    private execute(callback?: (err: any) => void): void {
        this._lastExecutedTime = Date.now();
        let duration = this._configuration.duration > 0 ? this._configuration.duration : 365 * 24 * 36000;
        this._stopTime = Date.now() + duration * 1000;

        this._benchmarkCount = this._benchmarks.length;
        this._onlyBenchmark = this._benchmarkCount == 1 ? this._benchmarks[0] : null;

        // Execute benchmarks
        async.whilst(
            () => {
                return this._running && this._lastExecutedTime < this._stopTime;
            },
            (callback) => {
                let benchmark = this._onlyBenchmark != null 
                    ? this._onlyBenchmark : this.chooseBenchmarkProportionally();
                let called = 0;
                this.executeBenchmark(
                    benchmark, 
                    (err) => { process.nextTick(callback, err); }
                );
            }, 
            (err) => {
                this.stop((err2) => {
                    if (callback) callback(err);
                });
            }
        );
    }
}
