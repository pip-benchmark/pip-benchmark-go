import { MeasurementType } from './MeasurementType';
import { ExecutionType } from './ExecutionType';
import { ConfigurationCallback } from './ConfigurationCallback';

export class ConfigurationManager {
    private _measurementType: MeasurementType = MeasurementType.Peak;
    private _nominalRate: number = 1;
    private _executionType: ExecutionType = ExecutionType.Proportional;
    private _duration: number = 60;
    private _forceContinue: boolean = false;
    private _changeListeners: ConfigurationCallback[] = [];
    
    public constructor() {}

     public get measurementType(): MeasurementType {
        return this._measurementType;
    }
    
    public set measurementType(value: MeasurementType) {
        this._measurementType = value; 
        this.notifyChanged();
    }

    public get nominalRate() {
        return this._nominalRate; 
    }
    
    public set nominalRate(value: number) {
        this._nominalRate = value;
        this.notifyChanged();
    }

    public get executionType(): ExecutionType {
        return this._executionType; 
    }
    
    public set executionType(value: ExecutionType) {
        this._executionType = value;
        this.notifyChanged();
    }

    public get duration(): number {
        return this._duration; 
    }
    
    public set duration(value: number) {
        this._duration = value;
        this.notifyChanged();
    }
    
    public get forceContinue(): boolean {
        return this._forceContinue;
    }
    
    public set forceContinue(value: boolean) {
        this._forceContinue = value;
        this.notifyChanged();
    }

    public addChangeListener(listener: ConfigurationCallback): void {
        this._changeListeners.push(listener);
    }

    public removeChangeListener(listener: ConfigurationCallback): void {
        for (let index = this._changeListeners.length - 1; index >= 0; index--) {
            if (this._changeListeners[index] == listener)
                this._changeListeners = this._changeListeners.splice(index, 1);
        }
    }

    public notifyChanged(): void {
        for (let index = 0; index < this._changeListeners.length; index++) {
            try {
                let listener = this._changeListeners[index];
                listener();
            } catch (ex) {
                // Ignore and send a message to the next listener.
            }
        }
    }
}