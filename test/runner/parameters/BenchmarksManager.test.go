let assert = require('chai').assert;

import { BenchmarkRunner } from '../../../src/runner/BenchmarkRunner';

suite('BenchmarksManager', ()=> {

    test('loadSuites', () => {
        let runner = new BenchmarkRunner();
        let benchmarks = runner.benchmarks;

        assert.equal(0, benchmarks.suites.length);
        benchmarks.addSuitesFromModule('./obj/src/standardbenchmarks');

        assert.equal(2, benchmarks.suites.length);
    });

    test('addSuiteFromClass', () => {
        let runner = new BenchmarkRunner();
        let benchmarks = runner.benchmarks;

        assert.equal(0, benchmarks.suites.length);
        benchmarks.addSuiteFromClass('./obj/src/standardbenchmarks,UtilityBenchmarkSuite');

        assert.equal(1, benchmarks.suites.length);
    });

    test('selectAll', () => {
        let runner = new BenchmarkRunner();
        let benchmarks = runner.benchmarks;

        benchmarks.addSuitesFromModule('./obj/src/standardbenchmarks');

        assert.equal(0, benchmarks.isSelected.length);

        runner.benchmarks.selectAll();
        assert.equal(5, benchmarks.isSelected.length);

        benchmarks.unselectAll();
        assert.equal(0, benchmarks.isSelected.length);
    });

    test('selectBenchmarkByName', () => {
        let runner = new BenchmarkRunner();
        let benchmarks = runner.benchmarks;

        benchmarks.addSuitesFromModule('./obj/src/standardbenchmarks');

        assert.equal(0, benchmarks.isSelected.length);

        benchmarks.selectByName(['Utility.Empty']);
        assert.equal(1, benchmarks.isSelected.length);

        benchmarks.unselectByName(['Utility.Empty']);
        assert.equal(0, benchmarks.isSelected.length);
    });

});