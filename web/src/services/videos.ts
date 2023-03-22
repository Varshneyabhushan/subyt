import videosResource from "./videosResource";

export interface VideosServiceConfig {
    apiUrl : string;
}

interface Channel {
    Id : string;
    Title : string;
}

interface Thumbnail {
    Width : number;
    Height : number;
    Url : string;
}

export interface Video {
    Id : string;
    Title : string;
    Description : string;
    PublishedAt : Date;
    CreatedAt : Date;
    Channel : Channel;
    Thumbnails : Thumbnail[];
}

export interface VideoResource {
    read() : Video[];
  }

export default class VideosService {
    apiUrl = ""
    constructor(config : VideosServiceConfig) {
        this.apiUrl = config.apiUrl
    }

    getVideos(skip : number, limit : number) : VideoResource {
        return videosResource(Promise.reject("not yet implemented"))
    }

    searchVideos(term : string, skip : number, limit : number) : VideoResource {
        return videosResource(Promise.reject("not yet implemented"))
    }
}