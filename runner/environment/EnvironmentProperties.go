let path = require('path');

import { Properties } from '../../utilities/Properties';

export class EnvironmentProperties extends Properties {

    private getFilePath(): string {
        return "BenchmarkEnvironment.properties";
    }

    public get cpuBenchmark(): number {
        return super.getAsDouble("CpuBenchmark", 0);
    }

    public set cpuBenchmark(value: number) {
        super.setAsDouble("CpuBenchmark", value);
    }

    public get diskBenchmark(): number {
        return super.getAsDouble("DiskBenchmark", 0);
    }

    public set diskBenchmark(value: number) {
        super.setAsDouble("DiskBenchmark", value);
    }

    public get videoBenchmark(): number {
        return super.getAsDouble("VideoBenchmark", 0);
    }

    public set videoBenchmark(value: number) {
        super.setAsDouble("VideoBenchmark", value);
    }

    public load(): void {
        if (path.existSync(this.getFilePath())) {
            this.loadFromFile(this.getFilePath());
        }
    }

    public save(): void {
        this.saveToFile(this.getFilePath());
    }
}