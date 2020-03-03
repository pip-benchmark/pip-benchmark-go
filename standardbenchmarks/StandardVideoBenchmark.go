import { PassiveBenchmark } from '../PassiveBenchmark';

export class StandardVideoBenchmark extends PassiveBenchmark {
    public constructor() {
        super("Video", "Measures speed of drawing graphical primitives");
    }
}