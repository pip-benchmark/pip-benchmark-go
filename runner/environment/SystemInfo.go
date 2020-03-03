let os = require('os');

export class SystemInfo {

    public constructor() {
        this.put("Machine Name", os.hostname());
        this.put("User Name", os.userInfo().username);
        this.put("Operating System Name", os.type());
        this.put("Operating System Version", os.release());
        this.put("Operating System Architecture", os.arch());
        this.put("Node.js Name", (<any>process).release.name);
        this.put("Node.js Version", process.version);
    }

    private put(parameter: string, value: string): void {
        this[parameter] = value;
    }
}
