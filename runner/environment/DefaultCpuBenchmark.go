import { Benchmark } from '../../Benchmark';

export class DefaultCpuBenchmark extends Benchmark {
    private static readonly NumberOfAttempts = 20000;

    public constructor() {
        super("CPU", "Measures CPU performance");
    }

   public execute(callback: (err: any) => void): void {
        // Count increment, comparison and goto for 1 arithmetic operation
        for (let value = 0; value < DefaultCpuBenchmark.NumberOfAttempts; value++) {
            // #1
            let result1 = value + value;
            let result2 = result1 - value;
            let result3 = result1 * result2;
            let result4 = result2 / result3;
            Math.log(result4);

            // #2
            result1 = value + value;
            result2 = result1 - value;
            result3 = result1 * result2;
            result4 = result2 / result3;
            Math.log(result4);

            // #3
            result1 = value + value;
            result2 = result1 - value;
            result3 = result1 * result2;
            result4 = result2 / result3;
            Math.log(result4);

            // #4
            result1 = value + value;
            result2 = result1 - value;
            result3 = result1 * result2;
            result4 = result2 / result3;
            Math.log(result4);

            // #5
            result1 = value + value;
            result2 = result1 - value;
            result3 = result1 * result2;
            result4 = result2 / result3;
            Math.log(result4);

            // #6
            result1 = value + value;
            result2 = result1 - value;
            result3 = result1 * result2;
            result4 = result2 / result3;
            Math.log(result4);

            // #7
            result1 = value + value;
            result2 = result1 - value;
            result3 = result1 * result2;
            result4 = result2 / result3;
            Math.log(result4);

            // #8
            result1 = value + value;
            result2 = result1 - value;
            result3 = result1 * result2;
            result4 = result2 / result3;
            Math.log(result4);

            // #9
            result1 = value + value;
            result2 = result1 - value;
            result3 = result1 * result2;
            result4 = result2 / result3;
            Math.log(result4);

            // #10
            result1 = value + value;
            result2 = result1 - value;
            result3 = result1 * result2;
            result4 = result2 / result3;
            Math.log(result4);
        }

        callback(null);
    }

}
