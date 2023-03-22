
import { Video } from "./videos"

export default function videosResource(videosProvider: Promise<Video[]>) {
  let status = 'pending'
  let response: Video[] = [];
  let err: Error;

  const suspender = videosProvider.then(
    (res) => {
      status = 'success'
      response = res
    },
    (err) => {
      status = 'error'
      err = err
    },
  )

  return {
    read() {
      switch (status) {
        case 'pending':
          throw suspender
        case 'error':
          throw err
        default:
          return response
      }
    }
  }
}