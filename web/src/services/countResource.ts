

export default function countResource(countProvider: Promise<number>) {
  let status = 'pending'
  let response: number = 0;
  let responseError: Error;

  const suspender = countProvider.then(
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