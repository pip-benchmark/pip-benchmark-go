var util = require('util');

import { Parameter } from '../../Parameter';
import { Converter } from '../../utilities/Converter';
import { BenchmarkInstance } from '../benchmarks/BenchmarkInstance';

export class BenchmarkProportionParameter extends Parameter {
    private _benchmark: BenchmarkInstance;

    public constructor(benchmark: BenchmarkInstance) {
        super(util.format("%s.%s.Proportion", benchmark.suite.name, benchmark.name),
            util.format("Sets execution proportion for benchmark %s in suite %s", 
            benchmark.name, benchmark.suite.name), "100");
        this._benchmark = benchmark;
    }

    public get value(): string {
        return Converter.integerToString(this._benchmark.proportion); 
    }
    
    public set value(value: string) {
        this._benchmark.proportion = Converter.stringToInteger(value, 100);
    }
}
