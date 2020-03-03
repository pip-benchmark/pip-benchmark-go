import { ConfigurationManager } from './config/ConfigurationManager';
import { ResultsManager } from './results/ResultsManager';
import { BenchmarksManager } from './benchmarks/BenchmarksManager';
import { ParametersManager } from './parameters/ParametersManager';
import { ExecutionManager } from './execution/ExecutionManager';
import { ReportGenerator } from './reports/ReportGenerator';
import { EnvironmentManager } from './environment/EnvironmentManager';

export class BenchmarkRunner {
    private _configuration: ConfigurationManager;
    private _results: ResultsManager;
    private _parameters: ParametersManager;
    private _benchmarks: BenchmarksManager;
    private _execution: ExecutionManager;
    private _report: ReportGenerator;
    private _environment: EnvironmentManager;

    public constructor() {
        this._configuration = new ConfigurationManager();
        this._results = new ResultsManager();
        this._parameters = new ParametersManager(this._configuration);
        this._benchmarks = new BenchmarksManager(this._parameters);
        this._execution = new ExecutionManager(this._configuration, this._results);
        this._environment = new EnvironmentManager();
        this._report = new ReportGenerator(this._configuration, this._results, 
            this._parameters, this._benchmarks, this._environment);
    }

    public get configuration(): ConfigurationManager {
        return this._configuration;
    }

    public get results(): ResultsManager {
        return this._results;
    }

    public get parameters(): ParametersManager {
        return this._parameters;
    }

    public get execution(): ExecutionManager {
        return this._execution;
    }

    public get benchmarks(): BenchmarksManager {
        return this._benchmarks;
    }

    public get report(): ReportGenerator {
        return this._report;
    }

    public get environment(): EnvironmentManager {
        return this._environment;
    }

    public get isRunning(): boolean {
        return this._execution.isRunning;
    }

    public start(): void {
        this._execution.start(this._benchmarks.isSelected);
    }

    public stop(): void {
        this._execution.stop();
    }

    public run(callback: (err: any) => void): void {
        this._execution.run(this._benchmarks.isSelected, callback);
    }
}