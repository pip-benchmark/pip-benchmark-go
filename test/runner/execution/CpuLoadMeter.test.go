let assert = require('chai').assert;

import { CpuLoadMeter } from '../../../src/runner/execution/CpuLoadMeter';

suite('CpuLoadMeter', ()=> {

    test('measure', (done) => {
        let meter = new CpuLoadMeter();
        let measure = meter.measure();
        assert.isTrue(measure == 0);

        setTimeout(() => {
            measure = meter.measure();
            assert.isTrue(measure > 0);

            done();
        }, 100);
    });

});