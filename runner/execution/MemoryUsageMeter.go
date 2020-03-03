let os = require('os');

import { BenchmarkMeter } from './BenchmarkMeter';

export class MemoryUsageMeter extends BenchmarkMeter {

    public constructor() {
        super();
    }

    protected performMeasurement(): number {
        return (os.totalmem() - os.freemem()) / 1024 / 1024;
    }

}
