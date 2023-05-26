We use the **`net/http`** package to handle web requests in Golang. It returns a response(as a reference) and an error. The response has a property called **Close**. **It is very important to close the response body yourself, otherwise you might face a lot of errors**

## Usage
```
response, err := http.Get(url)
if err != nil {
	panic(err)
}
// do something with the response
response.Body.Close() // very important
```
