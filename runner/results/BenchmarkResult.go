import { Measurement } from './Measurement';
import { BenchmarkInstance } from '../benchmarks/BenchmarkInstance';

export class BenchmarkResult {	
    public benchmarks: BenchmarkInstance[] = [];
    public startTime: number = Date.now();
    public elapsedTime: number = 0;
    public performanceMeasurement = new Measurement(0, 0, 0, 0);
    public cpuLoadMeasurement = new Measurement(0, 0, 0, 0);
    public memoryUsageMeasurement = new Measurement(0, 0, 0, 0);
    public errors: any[] = [];
}
