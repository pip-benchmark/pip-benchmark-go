import { BenchmarkSuite } from '../../BenchmarkSuite';
import { DefaultCpuBenchmark } from './DefaultCpuBenchmark';
import { DefaultDiskBenchmark } from './DefaultDiskBenchmark';
import { DefaultVideoBenchmark } from './DefaultVideoBenchmark';

export class StandardBenchmarkSuite extends BenchmarkSuite {
    private _cpuBenchmark: DefaultCpuBenchmark;
    private _diskBenchmark: DefaultDiskBenchmark;
    private _videoBenchmark: DefaultVideoBenchmark;

    public constructor() {
        super("StandardBenchmark", "Measures overall system performance");
        
        this._cpuBenchmark = new DefaultCpuBenchmark();
        this.addBenchmark(this._cpuBenchmark);

        this._diskBenchmark = new DefaultDiskBenchmark();
        this.addBenchmark(this._diskBenchmark);

        this._videoBenchmark = new DefaultVideoBenchmark();
        this.addBenchmark(this._videoBenchmark);

        this.createParameter("FilePath", "Path where test file is located on disk", "");
        this.createParameter("FileSize", "Size of the test file", "102400000");
        this.createParameter("ChunkSize", "Size of a chunk that read or writter from/to test file", "1024000");
        this.createParameter("OperationTypes", "Types of test operations: read, write or all", "all");
    }

    public get cpuBenchmark(): DefaultCpuBenchmark {
        return this._cpuBenchmark;
    }

    public get diskBenchmark(): DefaultDiskBenchmark {
        return this._diskBenchmark;
    }

    public get videoBenchmark(): DefaultVideoBenchmark {
        return this._videoBenchmark;
    }
}
