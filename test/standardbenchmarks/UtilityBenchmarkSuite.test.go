let assert = require('chai').assert;
let async = require('async');

import { UtilityBenchmarkSuite } from '../../src/standardbenchmarks/UtilityBenchmarkSuite';

suite('UtilityBenchmarkSuite', ()=> {
    let suite: UtilityBenchmarkSuite;

    setup((done) => {
        suite = new UtilityBenchmarkSuite();
        suite.setUp(done);
    });

    teardown((done) => {
        suite.tearDown(done);
    });

    test('emptyBenchmark', (done) => {
        assert.equal(2, suite.benchmarks.length);
        let benchmark = suite.benchmarks[0];
        async.times(
            100, 
            (n, callback) => { benchmark.execute(callback); }, 
            done
        );
    });

    test('randomBenchmark', (done) => {
        assert.equal(2, suite.benchmarks.length);
        let benchmark = suite.benchmarks[1];
        async.times(
            1, 
            (n, callback) => { benchmark.execute(callback); }, 
            done
        );
    });

});