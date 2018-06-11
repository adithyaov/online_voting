import { HttpHeaders } from '@angular/common/http';

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

function makeHeaders(headerObj: any) {
  return {
    headers: new HttpHeaders({
      'Content-Type':  'application/json',
      ...headerObj
    })
  };
}

export { joinPath, makeHeaders };

