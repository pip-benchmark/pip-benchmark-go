import { PassiveBenchmark } from '../../PassiveBenchmark';

export class DefaultVideoBenchmark extends PassiveBenchmark {
    private static readonly MaxLength = 100;

    public constructor() {
        super("Video", "Measures performance of video card");
    }
}