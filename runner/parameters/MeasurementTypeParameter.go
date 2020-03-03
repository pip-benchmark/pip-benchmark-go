import { Parameter } from '../../Parameter';
import { MeasurementType } from '../config/MeasurementType';
import { ConfigurationManager } from '../config/ConfigurationManager';

export class MeasurementTypeParameter extends Parameter {
    private _configuration: ConfigurationManager;

    public constructor(configuration: ConfigurationManager) {
        super(
            "General.Benchmarking.MeasurementType",
            "Performance type: peak or nominal", 
            "Peak"
        );
        this._configuration = configuration;
    }

    public get value(): string {
         return this._configuration.measurementType == MeasurementType.Peak ? "Peak" : "Nominal"; 
   }
    
    public set value(value: string) {
        value = value.toLowerCase();
        this._configuration.measurementType = value.startsWith("p")
            ? MeasurementType.Peak : MeasurementType.Nominal;
    }
}
