let assert = require('chai').assert;
let async = require('async');

import { StandardCpuBenchmark } from '../../src/standardbenchmarks/StandardCpuBenchmark';

suite('StandardCpuBenchmark', ()=> {
    let benchmark: StandardCpuBenchmark;

    setup((done) => {
        benchmark = new StandardCpuBenchmark();
        benchmark.setUp(done);
    });

    teardown((done) => {
        benchmark.tearDown(done);
    });

    test('execute', (done) => {
        async.times(
            100, 
            (n, callback) => { benchmark.execute(callback); }, 
            done
        );
    });

});