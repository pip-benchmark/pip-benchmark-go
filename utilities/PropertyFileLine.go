export class PropertyFileLine {
    private _line: string;
    private _key: string;
    private _value: string;
    private _comment: string;

    public constructor(key: string, value?: string, comment?: string) {
        if (value == null && comment == null) {
            this.parseLine(key);
        } else {
            this._key = key;
            this._value = value;
            this._comment = comment;
            this.composeNewLine();
        }
    }

    public get key(): string {
        return this._key;
    }

    public get value(): string {
        return this._value;
    }
        
    public set value(value: string) {
        this._value = value;
        this.composeNewLine();
    }

    public get comment(): string {
        return this._comment; 
    }
    
    public set comment(value: string) {
        this._comment = value;
        this.composeNewLine();
    }

    public get line(): string {
        return this._line;
    }

    private composeNewLine(): void {
        this._line = '';
        if (this._key != null && this._key.length > 0) {
            this._line += this.encodeValue(this._key);
            this._line += '=';
            this._line += this.encodeValue(this._value);
        }
        if (this._comment != null && this._comment.length > 0) {
            this._line += " ;";
            this._line += this._comment;
        }
    }

    private parseLine(line: string): void {
        this._line = line;

        // Parse comment
        let commentIndex = this.indexOfComment(line);
        if (commentIndex >= 0) {
            this._comment = line.substring(commentIndex + 1);
            line = line.substring(0, commentIndex);
        }

        // Parse key and value
        let assignmentIndex = line.indexOf('=');
        if (assignmentIndex >= 0) {
            this._value = line.substring(assignmentIndex + 1);
            this._value = this.decodeValue(this._value);
            this._key = line.substring(0, assignmentIndex);
            this._key = this.decodeValue(this._key);
        } else {
            this._key = this.decodeValue(line);
            this._value = "";
        }
    }

    private indexOfComment(value: string): number {
        let partOfString = false;
        let stringDelimiter = ' ';
        for (let index = 0; index < value.length; index++) {
            let chr = value.charAt(index);
            if (partOfString == false && chr == ';') {
                return index;
            } else if (partOfString == true && chr == stringDelimiter) {
                partOfString = false;
            } else if (partOfString == false && (chr == '\'' || chr == '\"')) {
                partOfString = true;
                stringDelimiter = chr;
            }
        }
        return -1;
    }

    private decodeValue(value: string): string {
        value = value.trim();
        if (value.startsWith("'") && value.endsWith("'")) {
            value = value.substring(1, value.length - 1);
            value = value.replace(/\'\'/g, "'");
        }
        if (value.startsWith("\"") && value.endsWith("\"")) {
            value = value.substring(1, value.length - 1);
            value = value.replace(/\"\"/g, "\"");
        }
        return value;
    }

    private encodeValue(value: string): string {
        if (value == null) {
            return value;
        }

        if (value.startsWith(" ") || value.endsWith(" ") || value.indexOf(';') >= 0) {
            value = value.replace(/\"/g, "\"\"");
            value = "\"" + value + "\"";
        }
        return value;
    }

    public toString(): string {
        return this._line;
    }
}
