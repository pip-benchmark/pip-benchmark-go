import { MeasurementType } from '../runner/config/MeasurementType';
import { ExecutionType } from '../runner/config/ExecutionType';
import { Converter } from '../utilities/Converter';

export class CommandLineArgs {

    public constructor(args: string[]) {
        this.processArguments(args);
    }

    private processArguments(args: string[]): void {
        for (let index = 2; index < args.length; index++) {
            let arg = args[index];
            let moreArgs = index < args.length - 1;
            
            if ((arg == "-a" || arg == "-j" || arg == "--module") && moreArgs) {
                let module = args[++index];
                this.modules.push(module);
            } else if ((arg == "-l" || arg == "--class") && moreArgs) {
                let clazz = args[++index];
                this.classes.push(clazz);
            } else if ((arg == "-b" || arg == "--benchmark") && moreArgs) {
                let benchmark = args[++index];
                this.benchmarks.push(benchmark);
            } else if ((arg == "-p" || arg == "--param") && moreArgs) {
                let param = args[++index];
                let pos = param.indexOf('=');
                let key = pos > 0 ? param.substring(0, pos - 1) : param;
                let value = pos > 0 ? param.substring(pos + 1) : null;
                this.parameters.put(key, value);
            } else if ((arg == "-c" || arg == "--config") && moreArgs) {
                this.configurationFile = args[++index];
            } else if ((arg == "-r" || arg == "--report") && moreArgs) {
                this.reportFile = args[++index];
            } else if ((arg == "-d" || arg == "--duration") && moreArgs) {
                this.duration = Converter.stringToLong(args[++index], 60);
            } else if ((arg == "-m" || arg == "--measure") && moreArgs) {
                let measure = args[++index].toLowerCase();
                this.measurementType = measure.startsWith("nom")
                    ? MeasurementType.Nominal : MeasurementType.Peak;
            } else if ((arg == "-x" || arg == "--execute") && moreArgs) {
                let execution = args[++index].toLowerCase();
                this.executionType = execution.startsWith("seq")
                    ? ExecutionType.Sequential : ExecutionType.Proportional;
            } else if ((arg == "-n" || arg == "--nominal") && moreArgs) {
                this.nominalRate = Converter.stringToDouble(args[++index], 1);
            } else if (arg == "-h" || arg == "--help") {
                this.showHelp = true;
            } else if (arg == "-B" || arg == "--show-benchmarks") {
                this.showBenchmarks = true;
            } else if (arg == "-P" || arg == "--show-params") {
                this.showParameters = true;
            } else if (arg == "-R" || arg == "--show-report") {
                this.showReport = true;
            } else if (arg == "-e" || arg == "--environment") {
                this.measureEnvironment = true;
            }
        }
    }

    public modules: string[] = [];
    public classes: string[] = [];
    public benchmarks: string[] = [];
    public parameters: any = {};
    public configurationFile: string;
    public reportFile: string = "BenchmarkReport.txt";
    public duration: number = 60;
    public showHelp: boolean = false;
    public showBenchmarks: boolean = false;
    public showParameters: boolean = false;
    public showReport: boolean = false;
    public measureEnvironment: boolean = false;
    public measurementType: MeasurementType = MeasurementType.Peak;
    public executionType: ExecutionType = ExecutionType.Proportional;
    public nominalRate: number = 1;
}
