export namespace backend {
	
	export class MakefileEntry {
	    command: string;
	    title: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new MakefileEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.command = source["command"];
	        this.title = source["title"];
	        this.description = source["description"];
	    }
	}

}

