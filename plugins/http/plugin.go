package main

type httpError struct {
}

func (g httpError) GetErrorMap() map[string]interface{} {
	return map[string]interface{}{
		"400": struct {
			Message string `json:"msg"`
		}{"The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing)."},
		"401": struct {
			Message string `json:"msg"`
		}{"Although the HTTP standard specifies \"unauthorized\", semantically this response means \"unauthenticated\". That is, the client must authenticate itself to get the requested response."},
		"403": struct {
			Message string `json:"msg"`
		}{"The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource. Unlike 401 Unauthorized, the client's identity is known to the server."},
		"404": struct {
			Message string `json:"msg"`
		}{"The server cannot find the requested resource. In the browser, this means the URL is not recognized. In an API, this can also mean that the endpoint is valid but the resource itself does not exist. Servers may also send this response instead of 403 Forbidden to hide the existence of a resource from an unauthorized client. This response code is probably the most well known due to its frequent occurrence on the web."},
	}
}

var Group httpError
