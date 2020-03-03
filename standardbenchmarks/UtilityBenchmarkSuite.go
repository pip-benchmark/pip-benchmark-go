import { BenchmarkSuite } from '../BenchmarkSuite';

export class UtilityBenchmarkSuite extends BenchmarkSuite {

    public constructor() {
        super("Utility", "Set of utility benchmark tests");

        this.createBenchmark("Empty", "Does nothing", this.executeEmpty);
        this.createBenchmark("RandomDelay", "Introduces random delay to measuring thread", this.executeRandomDelay);
    }

    private executeEmpty(callback: (err: any) => void): void {
        // This is an empty benchmark
        callback(null);
    }

    private executeRandomDelay(callback: (err: any) => void): void {
        setTimeout(
            () => { callback(null); },
            Math.random() * 1000
        );
    }
}
