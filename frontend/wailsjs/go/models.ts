export namespace main {
	
	export class FestivalInfo {
	    name: string;
	    days: number;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new FestivalInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.days = source["days"];
	        this.type = source["type"];
	    }
	}

}

