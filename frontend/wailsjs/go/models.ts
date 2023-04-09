export namespace backend {
	
	export class MakefileEntry {
	    prefix: string;
	    value: string;
	    title: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new MakefileEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.prefix = source["prefix"];
	        this.value = source["value"];
	        this.title = source["title"];
	        this.description = source["description"];
	    }
	}

}

