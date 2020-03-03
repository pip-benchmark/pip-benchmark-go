import { Parameter } from '../../Parameter';
import { Converter } from '../../utilities/Converter';
import { ConfigurationManager } from '../config/ConfigurationManager';

export class NominalRateParameter extends Parameter {
    private _configuration: ConfigurationManager;

    public constructor(configuration: ConfigurationManager) {
        super(
            "General.Benchmarking.NominalRate",
            "Rate for nominal benchmarking in TPS", 
            "1"
        );
        this._configuration = configuration;
    }

    public get value(): string {
        return Converter.doubleToString(this._configuration.nominalRate); 
    }
    
    public set value(value: string) {
        this._configuration.nominalRate = Converter.stringToDouble(value, 1);
    }
}
