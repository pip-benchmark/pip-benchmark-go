package benchmarks


// import (
    
// )
// // let _ = require("lodash");
// // let path = require("path");

// // import { BenchmarkSuite } from "../../BenchmarkSuite";
// // import { Benchmark } from "../../Benchmark";
// // import { BenchmarkSuiteInstance } from "./BenchmarkSuiteInstance";
// // import { BenchmarkInstance } from "./BenchmarkInstance";
// // import { ConfigurationManager } from "../config/ConfigurationManager";
// // import { ParametersManager } from "../parameters/ParametersManager";

// type BenchmarksManager struct{
//      parameters ParametersManager;
//      suites []BenchmarkSuiteInstance
// }

//     func NewBenchmarksManager (parameters ParametersManager) *BenchmarksManager {
//         c := BenchmarksManager{}
//         c.parameters = parameters;
//         c.suites = make([]BenchmarkSuiteInstance, 0)
//         return &c
//     }
 
//     func (c*BenchmarksManager ) Suites() []BenchmarkSuiteInstance {
//         return c.suites;
//     }

//     func (c*BenchmarksManager ) IsSelected() []BenchmarkInstance {
//         let benchmarks: BenchmarkInstance[] = [];

//         for (let suite of c.suites) {
//             for (let benchmark of suite.benchmarks) {
//                 if (benchmark.isSelected)
//                     benchmarks.push(benchmark);
//             }
//         }

//         return benchmarks;
//     }

//     func (c*BenchmarksManager ) SelectAll() void {
//        for (let suite of c.suites) {
//             for (let benchmark of suite.benchmarks) {
//                 benchmark.isSelected = true;
//             }
//         }
//     }

//     func (c*BenchmarksManager ) SelectByName(benchmarkNames: string[]) void {
//         for (let suite of c.suites) {
//             for (let benchmark of suite.benchmarks) {
//                 for (let benchmarkName of benchmarkNames) {
//                     if (benchmarkName == benchmark.fullName)
//                         benchmark.isSelected = true;
//                 }
//             }
//         }
//     }

//     func (c*BenchmarksManager ) Select(benchmarks: BenchmarkInstance[]) void {
//         for (let suite of c.suites) {
//             for (let benchmark of suite.benchmarks){
//                 for (let anotherBenchmark of benchmarks) {
//                     if (benchmark == anotherBenchmark)
//                         benchmark.isSelected = true;
//                 }
//             }
//         }
//     }

//     func (c*BenchmarksManager ) UnselectAll() void {
//         for (let suite of c.suites) {
//             for (let benchmark of suite.benchmarks) {
//                 benchmark.isSelected = false;
//             }
//         }
//     }

//     func (c*BenchmarksManager ) UnselectByName(benchmarkNames: string[]) void {
//         for (let suite of c.suites) {
//             for (let benchmark of suite.benchmarks) {
//                 for (let benchmarkName of benchmarkNames) {
//                     if (benchmarkName == benchmark.fullName)
//                         benchmark.isSelected = false;
//                 }
//             }
//         }
//     }

//     func (c*BenchmarksManager ) Unselect(benchmarks: BenchmarkInstance[]) void {
//         for (let suite of c.suites) {
//             for (let benchmark of suite.benchmarks) {
//                 for (let anotherBenchmark of benchmarks) {
//                     if (benchmark == anotherBenchmark)
//                         benchmark.isSelected = false;
//                 }
//             }
//         }
//     }

//     func (c*BenchmarksManager ) AddSuiteFromClass(suiteClassName: string) void {
//         if (suiteClassName == null || suiteClassName.length == 0) 
//             return;

//         let moduleName = suiteClassName;
//         suiteClassName = null;

//         let pos = moduleName.indexOf(",");
//         if (pos >= 0) {
//             let moduleAndClassName = moduleName;
//             moduleName = moduleAndClassName.substring(0, pos);
//             suiteClassName = moduleAndClassName.substring(pos + 1);
//         }

//         if (moduleName.startsWith("."))
//             moduleName = path.resolve(moduleName);

//         let suite = require(moduleName);
//         if (suite == null)
//             throw new Error("Module " + moduleName + " was not found");

//         if (suiteClassName != null && suiteClassName.length > 0)
//             suite = suite[suiteClassName];

//         if (_.isFunction(suite)) {
//             suite = new suite();
//             c.addSuite(suite);
//         }
//     }

//     func (c*BenchmarksManager ) AddSuite(suite: any) void {
//         if (suite instanceof BenchmarkSuite)
//             suite = new BenchmarkSuiteInstance(suite);
//         if (!(suite instanceof BenchmarkSuiteInstance))
//             throw Error("Incorrect suite type");

//         c.suites.push(suite);
//         c.parameters.addSuite(suite);
//     }

//     func (c*BenchmarksManager ) AddSuitesFromModule(moduleName: string) void {
//         if (moduleName.startsWith("."))
//             moduleName = path.resolve(moduleName);

//         let suites = require(moduleName);
//         if (suites == null)
//             throw new Error("Module " + moduleName + " was not found");

//         for (let prop in suites) {
//             let suite = suites[prop];
//             if (_.isFunction(suite) && suite.name.endsWith("Suite")) {
//                 try {
//                     suite = new suite();
//                     if (suite instanceof BenchmarkSuite) {
//                         suite = new BenchmarkSuiteInstance(suite);
//                         c.suites.push(suite);
//                         c.parameters.addSuite(suite);
//                     }
//                 } catch (ex) {
//                     // Ignore
//                 }
//             }
//         }
//     }

//     func (c*BenchmarksManager ) RemoveSuiteByName(suiteName: string) void {
//         let suite = _.find(c.suites, (suite) => {
//             return suite.name == suiteName;
//         });

//         if (suite != null) {
//             c.parameters.removeSuite(suite);
//             c.suites = _.remove(c.suites, (s) => { return s == suite; })
//         }
//     }

//     func (c*BenchmarksManager ) RemoveSuite(suite: any) void {
//         if (suite instanceof BenchmarkSuite)
//             suite = _.find(c.suites, (s) => { return s.suite == suite });

//         if (!(suite instanceof BenchmarkSuiteInstance))
//             throw new Error("Wrong suite type");

//         c.parameters.removeSuite(suite);
//         c.suites = _.remove(c.suites, (s) => s == suite);
//     }

//     func (c*BenchmarksManager ) Clear() void {
//         for (let index = 0; index < c.suites.length; index++) {
//             let suite = c.suites[index];
//             c.parameters.removeSuite(suite);
//         }
        
//         c.suites = [];
//     }

// }
