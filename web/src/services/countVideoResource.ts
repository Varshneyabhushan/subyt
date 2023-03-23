import { CountResource, Video, VideoResource } from "./videos";

export type CountVideoResource = [CountResource, VideoResource]

export default function countVideoResource(countProvider: Promise<{ Videos?: Video[], Count: number }>)
    : CountVideoResource {
    let status = 'pending'
    let count: number = 0;
    let videos: Video[] = [];
    let responseError: Error;

    const suspender = countProvider.then(
        (res) => {
            status = 'success'
            count = res.Count 
            videos = res.Videos ?? []
        },
        (err) => {
            status = 'error'
            responseError = err
        },
    )

    let countResource = {
        read() {
            switch (status) {
                case 'pending':
                    throw suspender
                case 'error':
                    throw responseError
                default:
                    return count
            }
        }
    }

    let videosResource = {
        read() {
            switch (status) {
                case 'pending':
                    throw suspender
                case 'error':
                    throw responseError
                default:
                    return videos
            }
        }
    }

    return [countResource, videosResource]
}