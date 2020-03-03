# <img src="https://github.com/pip-devs/pip-selenium-ps/raw/master/artifacts/logo.png" alt="Pip.Services Logo" style="max-width:30%"> <br/> Portable Benchmarking Framework in Node.js

This benchmarking framework ported cross multiple languages to obtain comparible performance metrics across different implementations.
In addition to performance benchmarking, it helps in other types of non-functional testing like load, reliability or concurrency.

* Measures performance in **transactions per second** or **TPS** 
* Supports **active** (by calling Execute method) or **passive** (by reporting via Context) measurement methods
* Supports **configuration parameters** to set connection strings or other settings for benchmarks
* Runs benchmarks **sequential** or in **proportional** by allocating % of calls to each benchmark
* Measures **peak** or **nominal** measurement at specified transaction rate
* Measures **system utilization** (RAM and CPU) during benchmarking process
* Measures overall **environment** performance (CPU, Video, Disk) for objective interpretation of results
* Capture and **errors** or **validation** results
* **Console runner** to execute benchmarks

## Usage

To run benchmark do the following
```bash
node ./benchmark.js -a <path to suite> -b <benchmark name> -p <param>=<value>
```

To show available benchmarks
```bash
node ./benchmark.js -a <path to suite> -B
```

To show available parameters
```bash
node ./benchmark.js -a <path to suite> -P
```

To measure environment (CPU, video, disk)
```bash
node ./benchmark.js -e
```

## Installation

TBD...

## Acknowledgements

This module created and maintained by **Sergey Seroukhov**.

