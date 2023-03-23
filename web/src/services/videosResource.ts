
import { Video } from "./videos"

export default function videosResource(videosProvider: Promise<Video[]>) {
  let status = 'pending'
  let response: Video[] = [];
  let responseError: Error;

  const suspender = videosProvider.then(
    (res) => {
      status = 'success'
      response = res
    },
    (err) => {
      status = 'error'
      responseError = err
    },
  )

  return {
    read() {
      switch (status) {
        case 'pending':
          throw suspender
        case 'error':
          throw responseError
        default:
          return response
      }
    }
  }
}