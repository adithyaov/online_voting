import { HttpHeaders, HttpErrorResponse } from '@angular/common/http';

function joinPath(paths: string[]): string {
  let finalPath: string;
  const trimmed = paths.map((path) => {
    let sliceFrom = 0;
    let sliceTo = path.length;
    if (path.charAt(0) === '/') {
        sliceFrom += 1;
    }
    if (path.charAt(path.length - 1) === '/') {
        sliceTo -= 1;
    }
    return path.slice(sliceFrom, sliceTo);
  });
  trimmed.forEach(path => {
    finalPath += (path + '/');
  });
  return finalPath;
}

function makeHeaders(token: string) {
  return {
    headers: new HttpHeaders({
      'Content-Type':  'application/json',
      'Token': token
    })
  };
}

function handleError(error: HttpErrorResponse) {
  if (error.error instanceof ErrorEvent) {
    // A client-side or network error occurred. Handle it accordingly.
    console.error('An error occurred:', error.error.message);
  } else {
    // The backend returned an unsuccessful response code.
    // The response body may contain clues as to what went wrong,
    console.error(
      `Backend returned code ${error.status}, ` +
      `body was: ${error.error}`);
  }
  // return an observable with a user-facing error message
  // return throwError(
  //   'Something bad happened; please try again later.');
}

export { joinPath, makeHeaders, handleError };

