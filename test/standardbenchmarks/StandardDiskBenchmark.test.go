let assert = require('chai').assert;
let async = require('async');

import { StandardDiskBenchmark } from '../../src/standardbenchmarks/StandardDiskBenchmark';

suite('StandardDiskBenchmark', ()=> {
    let benchmark: StandardDiskBenchmark;

    setup((done) => {
        benchmark = new StandardDiskBenchmark();
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