# Portable Benchmarking Framework in Golang Changelog

## <a name="2.0.0"></a> 2.0.0 (2020-03-06)

Complete refactoring of benchmark runner

### Breaking Changes
* Runner API was completely changed

## <a name="1.0.0"></a> 1.0.0 (2017-01-28)

Initial public release

### Features
* **benchmarks** Code is structured as Benchmark and BenchmarkSuite classes
* **passive** Passive benchmarks that perform and report their own measurements
* **sequencial** Sequencial execution of benchmarks. Each benchmark runs for specified duration
* **proportional** Proportional execution of benchmarks. Each benchmark is called proportionally
* **peak** Peak measurement makes maximum number of calls, one after another without wait
* **nominal** Nominal measurement tries to maintain specified call rate while measuring system load
* **environment** Taking key environment benchmarks (CPU, Video, Disk) to report for objective interpretation of results
* **runner** Console runner to execute benchmarks

### Breaking Changes
No breaking changes since this is the first version

### Bug Fixes
No fixes in this version

