import { BenchmarkRunner } from './BenchmarkRunner';
import { BenchmarkSuite } from '../BenchmarkSuite';
import { MeasurementType } from './config/MeasurementType';
import { ExecutionType } from './config/ExecutionType';
import { timingSafeEqual } from 'crypto';

export class BenchmarkBuilder {
    protected _runner: BenchmarkRunner = new BenchmarkRunner();

    public forceContinue(isForceContinue: boolean = false): BenchmarkBuilder {
        this._runner.configuration.forceContinue = isForceContinue;
        return this;
    }

    public measureAs(measurementType: MeasurementType): BenchmarkBuilder {
        this._runner.configuration.measurementType = measurementType;
        return this;
    }

    public withNominalRate(nominalRate: number): BenchmarkBuilder {
        this._runner.configuration.nominalRate = nominalRate;
        return this;
    } 

    public executeAs(executionType: ExecutionType): BenchmarkBuilder {
        this._runner.configuration.executionType = executionType;
        return this;
    }

    public forDuration(duration: number): BenchmarkBuilder {
        this._runner.configuration.duration = duration;
        return this;
    }

    public addSuite(suite: BenchmarkSuite): BenchmarkBuilder {
        this._runner.benchmarks.addSuite(suite);
        return this;
    }

    public withParameter(name: string, value: any): BenchmarkBuilder {
        let parameters: any = {};
        parameters[name] = value;
        this._runner.parameters.set(parameters);
        return this;
    }

    public withBenchmark(name: string): BenchmarkBuilder {
        this._runner.benchmarks.selectByName([name]);
        return this;
    }

    public withAllBenchmarks(): BenchmarkBuilder {
        this._runner.benchmarks.selectAll();
        return this;
    }

    public create(): BenchmarkRunner {
        let result = this._runner;
        this._runner = new BenchmarkRunner();
        return result;
    }
}