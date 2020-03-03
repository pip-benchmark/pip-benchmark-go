const _ = require('lodash');

import { BenchmarkBuilder } from '../runner/BenchmarkBuilder';
import { BenchmarkRunner } from '../runner/BenchmarkRunner';
import { ConsoleEventPrinter } from './ConsoleEventPrinter';
import { CommandLineArgs } from './CommandLineArgs';

export class ConsoleBenchmarkBuilder extends BenchmarkBuilder {
 
    public configureWithArgs(args: any): BenchmarkBuilder {
        let _args: CommandLineArgs;

        if (args instanceof CommandLineArgs) {
            _args = <CommandLineArgs>args;
        } else {
            _args = new CommandLineArgs(args);
        }

        // Load modules
        for (let module of _args.modules) {
            this._runner.benchmarks.addSuitesFromModule(module);
        }

        // Load test suites classes
        for (let clazz of _args.classes) {
            this._runner.benchmarks.addSuiteFromClass(clazz);
        }
            
            // Load configuration
        if (_args.configurationFile != null)
            this._runner.parameters.loadFromFile(_args.configurationFile);

        // Set parameters
        if (!_.isEmpty(_args.parameters))
            this._runner.parameters.set(_args.parameters);

        // Select benchmarks
        if (_args.benchmarks.length == 0)
            this._runner.benchmarks.selectAll();
        else 
            this._runner.benchmarks.selectByName(_args.benchmarks);
                        
        // Configure benchmarking
        this._runner.configuration.measurementType = _args.measurementType;
        this._runner.configuration.nominalRate = _args.nominalRate;
        this._runner.configuration.executionType = _args.executionType;
        this._runner.configuration.duration = _args.duration;

        return this;
    }
 
    public create(): BenchmarkRunner {
        let _runner = super.create();
        ConsoleEventPrinter.attach(_runner);
        return _runner;
    }
}