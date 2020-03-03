import { BenchmarkSuite } from '../BenchmarkSuite';
import { StandardCpuBenchmark } from './StandardCpuBenchmark';
import { StandardDiskBenchmark } from './StandardDiskBenchmark';
import { StandardVideoBenchmark } from './StandardVideoBenchmark';

export class StandardHardwareBenchmarkSuite extends BenchmarkSuite {
    private _cpuBenchmarkTest: StandardCpuBenchmark;
    private _diskBenchmarkTest: StandardDiskBenchmark;
    private _videoBenchmarkTest: StandardVideoBenchmark;

    public constructor() {
        super("StandardBenchmark", "Standard hardware benchmark");

        this._cpuBenchmarkTest = new StandardCpuBenchmark();
        this.addBenchmark(this._cpuBenchmarkTest);

        this._diskBenchmarkTest = new StandardDiskBenchmark();
        this.addBenchmark(this._diskBenchmarkTest);

        this._videoBenchmarkTest = new StandardVideoBenchmark();
        this.addBenchmark(this._videoBenchmarkTest);
    }

    public get cpuBenchmarkTest(): StandardCpuBenchmark {
        return this._cpuBenchmarkTest;
    }

    public get diskBenchmarkTest(): StandardDiskBenchmark {
        return this._diskBenchmarkTest;
    }

    public get videoBenchmarkTest(): StandardVideoBenchmark {
        return this._videoBenchmarkTest;
    }
}
