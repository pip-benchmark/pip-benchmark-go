let _ = require('lodash');
let async = require('async');
let util = require('util');

import { BenchmarkRunner } from '../runner/BenchmarkRunner';
import { CommandLineArgs } from './CommandLineArgs';
import { ConsoleEventPrinter } from './ConsoleEventPrinter';

export class ConsoleRunner {
    private _args: CommandLineArgs;
    private _runner: BenchmarkRunner;

    public constructor() { }

    private start(args: string[]): void {
        this._args = new CommandLineArgs(args);
        this._runner = new BenchmarkRunner();
        
        ConsoleEventPrinter.attach(this._runner);

        this.executeBatchMode();
    }

    public stop(): void {
        this._runner.stop();
    }

    private executeBatchMode(): void {
        try {
            if (this._args.showHelp) {
                this.printHelp();
                return;
            }

            // Load modules
            for (let module of this._args.modules) {
                this._runner.benchmarks.addSuitesFromModule(module);
            }

            // Load test suites classes
            for (let clazz of this._args.classes) {
                this._runner.benchmarks.addSuiteFromClass(clazz);
            }
            
            // Load configuration
            if (this._args.configurationFile != null)
                this._runner.parameters.loadFromFile(this._args.configurationFile);

            // Set parameters
            if (!_.isEmpty(this._args.parameters))
                this._runner.parameters.set(this._args.parameters);

            // Select benchmarks
            if (this._args.benchmarks.length == 0)
                this._runner.benchmarks.selectAll();
            else 
                this._runner.benchmarks.selectByName(this._args.benchmarks);
            
            if (this._args.showParameters) {
                this.printParameters();
                return;
            }
            
            if (this._args.showBenchmarks) {
                this.printBenchmarks();
                return;
            }
            
            async.series([
                (callback) => {
                    // Benchmark the environment
                    if (this._args.measureEnvironment) {
                        console.log("Measuring Environment (wait up to 2 mins)...");
                        this._runner.environment.measure(true, true, false, (err) => {
                            let output = util.format(
                                "CPU: %d, Video: %d, Disk: %d",
                                (this._runner.environment.cpuMeasurement || 0).toFixed(2),
                                (this._runner.environment.videoMeasurement || 0).toFixed(2),
                                (this._runner.environment.diskMeasurement || 0).toFixed(2)
                            );
                            console.log(output);

                            callback(err);
                        });
                    } else callback();
                },
                (callback) => {
                    // Configure benchmarking
                    this._runner.configuration.measurementType = this._args.measurementType;
                    this._runner.configuration.nominalRate = this._args.nominalRate;
                    this._runner.configuration.executionType = this._args.executionType;
                    this._runner.configuration.duration = this._args.duration;

                    // Perform benchmarking
                    this._runner.run((err) => {
                        if (this._runner.results.all.length > 0)
                            console.log(this._runner.results.all[0].performanceMeasurement.averageValue.toFixed(2));

                        // Generate report and save to file
                        if (this._args.reportFile != null)
                            this._runner.report.saveToFile(this._args.reportFile);

                        // Show report in console
                        if (this._args.showReport)
                            console.log(this._runner.report.generate());

                        callback(err);
                    });
                }
            ], (err) => {
                // Do nothing
            });
        }
        catch (ex) {
            console.error(ex);
        }
    }

    public printHelp(): void {
        console.log("Pip.Benchmark Console Runner. (c) Conceptual Vision Consulting LLC 2017");
        console.log();
        console.log("Command Line Arguments:");
        console.log("-a <module>    - Module with benchmarks to be loaded. You may include multiple modules");
        console.log("-p <param>=<value> - Set parameter value. You may include multiple parameters");
        console.log("-b <benchmark>   - Name of benchmark to be executed. You may include multiple benchmarks");
        console.log("-c <config file> - File with parameters to be loaded");
        console.log("-r <report file> - File to save benchmarking report");
        console.log("-d <seconds>     - Benchmarking duration in seconds");
        console.log("-h               - Display this help screen");
        console.log("-B               - Show all available benchmarks");
        console.log("-P               - Show all available parameters");
        console.log("-R               - Show report");
        console.log("-e               - Measure environment");
        console.log("-x [prop|seq]    - Execution type: Proportional or Sequencial");
        console.log("-m [peak|nominal] - Measurement type: Peak or Nominal");
        console.log("-n <rate>        - Nominal rate in transactions per second");
    }

    private printBenchmarks(): void {
        console.log("Pip.Benchmark Console Runner. (c) Conceptual Vision Consulting LLC 2017");
        console.log();
        console.log("Benchmarks:");

        let suites = this._runner.benchmarks.suites;
        _.each(suites, (suite) => {
            _.each(suite.benchmarks, (benchmark) => {
                console.log(benchmark.fullName + ' - ' + benchmark.description);
            });
        });
    }

    private printParameters(): void {
        console.log("Pip.Benchmark Console Runner. (c) Conceptual Vision Consulting LLC 2017");
        console.log();
        console.log("Parameters:");

        let parameters = this._runner.parameters.userDefined;
        for (let index = 0; index < parameters.length; index++) {
            let parameter = parameters[index];
            let defaultValue = parameter.defaultValue;
            defaultValue = defaultValue == null || defaultValue.length == 0 
                    ? "" : " (Default: " + defaultValue + ")";
            console.log('' + parameter.name + ' - ' + parameter.description + defaultValue);
        }
    }
        
    public static run(args: string[]): void {
        let runner = new ConsoleRunner();

        // Log uncaught exceptions
        process.on('uncaughtException', (ex) => {
            console.error(ex);
            console.error("Process is terminated");
            process.exit(1);
        });

        // Gracefully shutdown
        process.on('exit', function () {
            runner.stop();
            //console.log("Goodbye!");
        });

        runner.start(args || process.argv);
    }

}