let assert = require('chai').assert;

import { MemoryUsageMeter } from '../../../src/runner/execution/MemoryUsageMeter';

suite('MemoryUsageMeter', ()=> {

    test('measure', (done) => {
        let meter = new MemoryUsageMeter();
        let measure = meter.measure();
        assert.isTrue(measure > 0);

        setTimeout(() => {
            measure = meter.measure();
            assert.isTrue(measure > 0);

            done();
        }, 100);
    });

});