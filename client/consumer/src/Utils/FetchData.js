import {
  AcceptHeader,
  OkStatusCode,
  MultipleChoicesRedirectStatusCode,
  HTMLApplicationContent
} from "../Constants/httpConstants.js";

export const checkStatus = response => {
  if (
    response.status >= OkStatusCode &&
    response.status < MultipleChoicesRedirectStatusCode
  ) {
    return response;
  }
  return response.json().then(json => {
    return Promise.reject({
      status: response.status,
      ok: false,
      statusText: response.statusText,
      body: json
    });
  });
};

export const parseJSON = response => {
  return response.json();
};

export const parseHTML = response => {
  return response.text();
};

const fetchData = async (url, options) => {
  const headers = {
    ...options.headers
  };
  options = {
    ...options,
    headers: headers
  };
  return fetch(url, options)
    .then(checkStatus)
    .then(
      headers[AcceptHeader] && headers[AcceptHeader] === HTMLApplicationContent
        ? parseHTML
        : parseJSON
    )
    .then(data => data)
    .catch(error => error);
};

export default fetchData;
