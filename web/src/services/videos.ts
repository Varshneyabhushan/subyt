
interface VideosServiceConfig {
    apiUrl : string;
}

interface Video {
    id : string;
}

export default class VideosService {
    apiUrl = ""
    constructor(config : VideosServiceConfig) {
        this.apiUrl = config.apiUrl
    }

    getVideos(skip : number, limit : number) : Promise<Video[]> {
        return Promise.reject("not yet implemented")
    }

    searchVideos(skip : number, limit : number) : Promise<Video[]> {
        return Promise.reject("not yet implemented")
    }
}