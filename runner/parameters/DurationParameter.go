import { Parameter } from '../../Parameter';
import { Converter } from '../../utilities/Converter';
import { ConfigurationManager } from '../config/ConfigurationManager';

export class DurationParameter extends Parameter {
    private _configuration: ConfigurationManager;

    public constructor(configuration: ConfigurationManager) {
        super(
            "General.Benchmarking.Duration", 
            "Duration of benchmark execution in seconds", 
            "60"
        );
        this._configuration = configuration;
    }

    public get value(): string {
        return Converter.integerToString(this._configuration.duration); 
    }
    
    public set value(value: string) {
        this._configuration.duration = Converter.stringToInteger(value, 60);
    }
}
