let _ = require('lodash');
let path = require('path');

import { BenchmarkSuite } from '../../BenchmarkSuite';
import { Benchmark } from '../../Benchmark';
import { BenchmarkSuiteInstance } from './BenchmarkSuiteInstance';
import { BenchmarkInstance } from './BenchmarkInstance';
import { ConfigurationManager } from '../config/ConfigurationManager';
import { ParametersManager } from '../parameters/ParametersManager';

export class BenchmarksManager {
    private _parameters: ParametersManager;
    private _suites: BenchmarkSuiteInstance[] = [];

    public constructor(parameters: ParametersManager) {
        this._parameters = parameters;
    }
 
    public get suites(): BenchmarkSuiteInstance[] {
        return this._suites;
    }

    public get isSelected(): BenchmarkInstance[] {
        let benchmarks: BenchmarkInstance[] = [];

        for (let suite of this._suites) {
            for (let benchmark of suite.benchmarks) {
                if (benchmark.isSelected)
                    benchmarks.push(benchmark);
            }
        }

        return benchmarks;
    }

    public selectAll(): void {
       for (let suite of this._suites) {
            for (let benchmark of suite.benchmarks) {
                benchmark.isSelected = true;
            }
        }
    }

    public selectByName(benchmarkNames: string[]): void {
        for (let suite of this._suites) {
            for (let benchmark of suite.benchmarks) {
                for (let benchmarkName of benchmarkNames) {
                    if (benchmarkName == benchmark.fullName)
                        benchmark.isSelected = true;
                }
            }
        }
    }

    public select(benchmarks: BenchmarkInstance[]): void {
        for (let suite of this._suites) {
            for (let benchmark of suite.benchmarks){
                for (let anotherBenchmark of benchmarks) {
                    if (benchmark == anotherBenchmark)
                        benchmark.isSelected = true;
                }
            }
        }
    }

    public unselectAll(): void {
        for (let suite of this._suites) {
            for (let benchmark of suite.benchmarks) {
                benchmark.isSelected = false;
            }
        }
    }

    public unselectByName(benchmarkNames: string[]): void {
        for (let suite of this._suites) {
            for (let benchmark of suite.benchmarks) {
                for (let benchmarkName of benchmarkNames) {
                    if (benchmarkName == benchmark.fullName)
                        benchmark.isSelected = false;
                }
            }
        }
    }

    public unselect(benchmarks: BenchmarkInstance[]): void {
        for (let suite of this._suites) {
            for (let benchmark of suite.benchmarks) {
                for (let anotherBenchmark of benchmarks) {
                    if (benchmark == anotherBenchmark)
                        benchmark.isSelected = false;
                }
            }
        }
    }

    public addSuiteFromClass(suiteClassName: string): void {
        if (suiteClassName == null || suiteClassName.length == 0) 
            return;

        let moduleName = suiteClassName;
        suiteClassName = null;

        let pos = moduleName.indexOf(',');
        if (pos >= 0) {
            let moduleAndClassName = moduleName;
            moduleName = moduleAndClassName.substring(0, pos);
            suiteClassName = moduleAndClassName.substring(pos + 1);
        }

        if (moduleName.startsWith('.'))
            moduleName = path.resolve(moduleName);

        let suite = require(moduleName);
        if (suite == null)
            throw new Error('Module ' + moduleName + ' was not found');

        if (suiteClassName != null && suiteClassName.length > 0)
            suite = suite[suiteClassName];

        if (_.isFunction(suite)) {
            suite = new suite();
            this.addSuite(suite);
        }
    }

    public addSuite(suite: any): void {
        if (suite instanceof BenchmarkSuite)
            suite = new BenchmarkSuiteInstance(suite);
        if (!(suite instanceof BenchmarkSuiteInstance))
            throw Error('Incorrect suite type');

        this._suites.push(suite);
        this._parameters.addSuite(suite);
    }

    public addSuitesFromModule(moduleName: string): void {
        if (moduleName.startsWith('.'))
            moduleName = path.resolve(moduleName);

        let suites = require(moduleName);
        if (suites == null)
            throw new Error('Module ' + moduleName + ' was not found');

        for (let prop in suites) {
            let suite = suites[prop];
            if (_.isFunction(suite) && suite.name.endsWith('Suite')) {
                try {
                    suite = new suite();
                    if (suite instanceof BenchmarkSuite) {
                        suite = new BenchmarkSuiteInstance(suite);
                        this._suites.push(suite);
                        this._parameters.addSuite(suite);
                    }
                } catch (ex) {
                    // Ignore
                }
            }
        }
    }

    public removeSuiteByName(suiteName: string): void {
        let suite = _.find(this._suites, (suite) => {
            return suite.name == suiteName;
        });

        if (suite != null) {
            this._parameters.removeSuite(suite);
            this._suites = _.remove(this._suites, (s) => { return s == suite; })
        }
    }

    public removeSuite(suite: any): void {
        if (suite instanceof BenchmarkSuite)
            suite = _.find(this._suites, (s) => { return s.suite == suite });

        if (!(suite instanceof BenchmarkSuiteInstance))
            throw new Error('Wrong suite type');

        this._parameters.removeSuite(suite);
        this._suites = _.remove(this._suites, (s) => s == suite);
    }

    public clear(): void {
        for (let index = 0; index < this._suites.length; index++) {
            let suite = this._suites[index];
            this._parameters.removeSuite(suite);
        }
        
        this._suites = [];
    }

}
