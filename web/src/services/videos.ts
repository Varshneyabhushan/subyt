import axios, { AxiosInstance } from "axios";
import countResource from "./countResource";
import countVideoResource, { CountVideoResource } from "./countVideoResource";
import videosResource from "./videosResource";

export interface VideosServiceConfig {
    apiUrl: string;
}

interface Channel {
    Id: string;
    Title: string;
}

interface Thumbnail {
    Width: number;
    Height: number;
    Url: string;
}

export interface Video {
    Id: string;
    Title: string;
    Description: string;
    PublishedAt: Date;
    CreatedAt: Date;
    Channel: Channel;
    Thumbnails: Thumbnail[];
}

export interface VideoResource {
    read(): Video[];
}

export interface CountResource {
    read(): number;
}

export default class VideosService {
    axios: AxiosInstance
    constructor(config: VideosServiceConfig) {
        this.axios = axios.create({
            baseURL: config.apiUrl,
            validateStatus(status) {
                return status >= 200 && status < 500;
            },
        });
    }

    getVideos(skip: number, limit: number): VideoResource {
        let provider = this.axios.get(`?skip=${skip}&limit=${limit}`)
            .then(res => res.data.Videos || [])

        return videosResource(provider)
    }

    getVideosCount() : CountResource {
        let provider = this.axios.get('count')
            .then(res => res.data.Count || 0)

        return countResource(provider)
    }

    searchVideos(term: string, skip: number, limit: number): CountVideoResource {
        let provider = this.axios.get(`search?term=${term}&skip=${skip}&limit=${limit}`)
            .then(res => {
                console.log(res.data)
                return res.data
            })
            
        return countVideoResource(provider)
    }
}