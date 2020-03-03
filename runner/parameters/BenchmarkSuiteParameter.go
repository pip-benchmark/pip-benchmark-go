var util = require('util');

import { Parameter } from '../../Parameter';
import { BenchmarkSuiteInstance } from '../benchmarks/BenchmarkSuiteInstance';

export class BenchmarkSuiteParameter extends Parameter {
    private _originalParameter: Parameter;

    public constructor(suite: BenchmarkSuiteInstance, originalParameter: Parameter) {
        super(util.format("%s.%s", suite.name, originalParameter.name),
          originalParameter.description, originalParameter.defaultValue);
        this._originalParameter = originalParameter;
    }

    public get value(): string {
        return this._originalParameter.value; 
    }
    
    public set value(value: string) {
        this._originalParameter.value = value;
    }
}
