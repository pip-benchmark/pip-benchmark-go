import { Parameter } from '../../Parameter';
import { ExecutionType } from '../config/ExecutionType';
import { ConfigurationManager } from '../config/ConfigurationManager';

export class ExecutionTypeParameter extends Parameter {
    private _configuration: ConfigurationManager;

    public constructor(configuration: ConfigurationManager) {
        super(
            "General.Benchmarking.ExecutionType",
            "Execution type: proportional or sequencial",
            "Proportional"
        );
        this._configuration = configuration;
    }

    public get value(): string {
        return this._configuration.executionType == ExecutionType.Proportional 
            ? "Proportional" : "Sequencial"; 
    }

    public set value(value: string) {
        value = value.toLowerCase();
        this._configuration.executionType = value.startsWith("p")
            ? ExecutionType.Proportional : ExecutionType.Sequential;
    }
}
