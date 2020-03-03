let async = require('async');

import { ConfigurationManager } from '../config/ConfigurationManager';
import { ResultsManager } from '../results/ResultsManager';
import { ExecutionState } from './ExecutionState';
import { BenchmarkInstance } from '../benchmarks/BenchmarkInstance';
import { BenchmarkSuiteInstance } from '../benchmarks/BenchmarkSuiteInstance';
import { BenchmarkResult } from '../results/BenchmarkResult';

import { ExecutionContext } from './ExecutionContext';
import { ExecutionStrategy } from './ExecutionStrategy';
import { ProportionalExecutionStrategy } from './ProportionalExecutionStrategy';

export class SequencialExecutionStrategy extends ExecutionStrategy {
    private _running: boolean = false;
    private _current: ProportionalExecutionStrategy;
    private _timeout: any;

    public constructor(configuration: ConfigurationManager, results: ResultsManager, 
        execution: any, benchmarks: BenchmarkInstance[]) {
        super(configuration, results, execution, benchmarks);
    }

    public get isStopped(): boolean {
        return !this._running;
    }

    public start(callback?: (err: any) => void): void {
        if (this._configuration.duration <= 0)
            throw new Error("Duration was not set");

        if (this._running) {
            callback(null);
            return;
        }

        this._running = true;

        this.execute(callback);
    }

    public stop(callback?: (err: any) => void): void {
        if (this._timeout != null) {
            clearTimeout(this._timeout);
            this._timeout = null;
        }

        if (this._running) {
            this._running = false;

            if (this._execution)
                this._execution.stop();

            if (this._current != null) {
                this._current.stop(callback);
            } else {
                if (callback) callback(null);
            }
        } else {
            if (callback) callback(null);
        }
    }

    private execute(callback?: (err: any) => void) {
        async.eachSeries(
            this._benchmarks,
            (benchmark, callback) => {
                // Skip if benchmarking was interrupted
                if (!this._running) {
                    callback();
                    return;
                }

                // Start embedded strategy
                this._current = new ProportionalExecutionStrategy(this._configuration, this._results, null, [benchmark]);
                this._current.start();

                // Wait for specified duration and stop embedded strategy
                this._timeout = setTimeout(
                    () => {
                        this._current.stop((err) => {
                            this._current = null;
                            callback(err);
                        });
                    }, 
                    this._configuration.duration * 1000
                );
            },
            (err) => {
                this.stop(callback);
            }
        )
    }

}
