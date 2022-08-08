package local_aws_proxy

import (
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func WrapperHandler(apiHandlerFunc func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//TODO: add all fields
		proxyRequest := events.APIGatewayProxyRequest{
			Path:                            r.URL.Path,
			HTTPMethod:                      r.Method,
			QueryStringParameters:           getFirstFromMulti(r.URL.Query()),
			Headers:                         getFirstFromMulti(r.Header),
			MultiValueHeaders:               r.Header,
			MultiValueQueryStringParameters: r.URL.Query(),
			Body:                            string(b),
		}
		resp, err := apiHandlerFunc(proxyRequest)
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(resp.Body))
		for key, value := range resp.Headers {
			w.Header().Set(key, value)
		}
	}

}
