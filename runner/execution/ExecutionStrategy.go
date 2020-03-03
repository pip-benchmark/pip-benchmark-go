let _ = require('lodash');

import { ConfigurationManager } from '../config/ConfigurationManager';
import { ResultsManager } from '../results/ResultsManager';
import { ExecutionState } from './ExecutionState';
import { BenchmarkInstance } from '../benchmarks/BenchmarkInstance';
import { BenchmarkSuiteInstance } from '../benchmarks/BenchmarkSuiteInstance';
import { BenchmarkResult } from '../results/BenchmarkResult';

export abstract class ExecutionStrategy {
    protected _configuration: ConfigurationManager;
    protected _results: ResultsManager;
    protected _execution: any;

    protected _benchmarks: BenchmarkInstance[];
    protected _activeBenchmarks: BenchmarkInstance[];
    protected _suites: BenchmarkSuiteInstance[];

    protected constructor(configuration: ConfigurationManager, 
        results: ResultsManager, execution: any, benchmarks: BenchmarkInstance[]) {

        this._configuration = configuration;
        this._results = results;
        this._execution = execution;

        this._benchmarks = benchmarks;
        this._activeBenchmarks = this.getActiveBenchmarks(benchmarks);
        this._suites = this.getBenchmarkSuites(benchmarks);
    }

    private getActiveBenchmarks(benchmarks: BenchmarkInstance[]): BenchmarkInstance[] {
        return _.filter(benchmarks, (benchmark) => {
            return !benchmark.passive;
        });
    }

    private getBenchmarkSuites(benchmarks: BenchmarkInstance[]): BenchmarkSuiteInstance[] {
        let suites: BenchmarkSuiteInstance[] = [];
        for (let index = 0; index < benchmarks.length; index++) {
            let benchmark = benchmarks[index];
            let suite = benchmark.suite;
            if (!_.some(suites, (s) => s == suite))
                suites.push(suite);
        }
        return suites;
    }

    public abstract isStopped: boolean;
    public abstract start(callback?: (err: any) => void): void;
    public abstract stop(callback?: (err: any) => void): void;

}
