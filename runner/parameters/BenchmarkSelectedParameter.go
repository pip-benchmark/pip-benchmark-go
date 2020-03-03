var util = require('util');

import { Parameter } from '../../Parameter';
import { Converter } from '../../utilities/Converter';
import { BenchmarkInstance } from '../benchmarks/BenchmarkInstance';

export class BenchmarkSelectedParameter extends Parameter {
    private _benchmark: BenchmarkInstance;

    public constructor(benchmark: BenchmarkInstance) {
        super(util.format("%s.%s.Selected", benchmark.suite.name, benchmark.name),
            util.format("Selecting benchmark %s in suite %s", 
            benchmark.name, benchmark.suite.name), "true");
        this._benchmark = benchmark;
    }

    public get value(): string {
        return Converter.booleanToString(this._benchmark.isSelected); 
    }
    
    public set value(value: string) {
        this._benchmark.isSelected = Converter.stringToBoolean(value, false);
    }
}
