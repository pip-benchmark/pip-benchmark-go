let util = require('util');

import { BenchmarkRunner } from '../runner/BenchmarkRunner';
import { BenchmarkResult } from '../runner/results/BenchmarkResult';
import { ExecutionState } from '../runner/execution/ExecutionState';

export class ConsoleEventPrinter {

    public static attach(runner: BenchmarkRunner): void {
        runner.execution.addUpdatedListener(ConsoleEventPrinter.onStateUpdated);
        runner.results.addErrorListener(ConsoleEventPrinter.onErrorReported);
        runner.results.addMessageListener(ConsoleEventPrinter.onMessageSent);
        runner.results.addUpdatedListener(ConsoleEventPrinter.onResultUpdated);
    }

    public static onStateUpdated(state: ExecutionState): void {
        if (state == ExecutionState.Running)
            console.log("Measuring....");
        else if (state == ExecutionState.Completed)
            console.log("Completed measuring.");
    }

    public static onResultUpdated(result: BenchmarkResult): void {
        if (result != null) {
            let output = util.format("%s Performance: %d %d>%d>%d CPU Load: %d %d>%d>%d Errors: %d",
                new Date().toISOString(),
                result.performanceMeasurement.currentValue.toFixed(2),
                result.performanceMeasurement.minValue.toFixed(2),
                result.performanceMeasurement.averageValue.toFixed(2),
                result.performanceMeasurement.maxValue.toFixed(2),
                result.cpuLoadMeasurement.currentValue.toFixed(2),
                result.cpuLoadMeasurement.minValue.toFixed(2),
                result.cpuLoadMeasurement.averageValue.toFixed(2),
                result.cpuLoadMeasurement.maxValue.toFixed(2),
                result.errors.length
            );
            console.log(output);
        }
    }

    public static onMessageSent(message: string): void {
        console.log(message);
    }

    public static onErrorReported(message: string): void {
        console.error(message);
    }
}
