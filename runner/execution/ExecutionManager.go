let _ = require('lodash');

import { ConfigurationManager } from '../config/ConfigurationManager';
import { ResultsManager } from '../results/ResultsManager';
import { ExecutionType } from '../config/ExecutionType';
import { BenchmarkInstance } from '../benchmarks/BenchmarkInstance';
import { BenchmarkSuiteInstance } from '../benchmarks/BenchmarkSuiteInstance';
import { ExecutionState } from './ExecutionState';
import { ExecutionStrategy } from './ExecutionStrategy';
import { BenchmarkResult } from '../results/BenchmarkResult';
import { ExecutionCallback } from './ExecutionCallback';

import { ProportionalExecutionStrategy } from './ProportionalExecutionStrategy';
import { SequencialExecutionStrategy } from './SequencialExecutionStrategy';

export class ExecutionManager {
	protected _configuration: ConfigurationManager;
    protected _results: ResultsManager;
    
    private _updatedListeners: ExecutionCallback[] = [];
    private _running: boolean = false;
    private _strategy: ExecutionStrategy = null;

    public constructor(configuration: ConfigurationManager, results: ResultsManager) {
        this._configuration = configuration;
        this._results = results;
    }
    
    public get isRunning(): boolean {
        return this._running;
    }

    public start(benchmarks: BenchmarkInstance[]): void {
        this.run(benchmarks, (err) => {});
    }

    public run(benchmarks: BenchmarkInstance[], callback?: (err: any) => void): void {
        if (benchmarks == null || benchmarks.length == 0) {
            callback(new Error("There are no benchmarks to execute"));
            return;
        }

        if (this._running)
            this.stop();
        this._running = true;

        this._results.clear();
        this.notifyUpdated(ExecutionState.Running);

        // Create requested execution strategy
        if (this._configuration.executionType == ExecutionType.Sequential) {
            this._strategy = new SequencialExecutionStrategy(this._configuration, this._results, this, benchmarks);
        } else {
            this._strategy = new ProportionalExecutionStrategy(this._configuration, this._results, this, benchmarks);
        }

        // Initialize parameters and start 
        this._strategy.start((err) => {
            this.stop();
            if (callback) callback(err);
        });
    }

    public stop(): void {
        if (this._running) {
            this._running = false;

            if (this._strategy != null) {
                this._strategy.stop();
                this._strategy = null;
            }

            this.notifyUpdated(ExecutionState.Completed);
        }
    }

    public addUpdatedListener(listener: ExecutionCallback): void {
        this._updatedListeners.push(listener);
    }

    public removeUpdatedListener(listener: ExecutionCallback): void {
        for (let index = this._updatedListeners.length - 1; index >= 0; index--) {
            if (this._updatedListeners[index] == listener)
                this._updatedListeners = this._updatedListeners.splice(index, 1);
        }
    }

    public notifyUpdated(state: ExecutionState): void {
        for (let index = 0; index < this._updatedListeners.length; index++) {
            try {
                let listener = this._updatedListeners[index];
                listener(state);
            } catch (ex) {
                // Ignore and send a message to the next listener.
            }
        }
    }

}