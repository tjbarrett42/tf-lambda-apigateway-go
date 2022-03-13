package main
import "github.com/aws/aws-lambda-go/lambda"
import "github.com/aws/aws-lambda-go/events"
import "errors"
import "fmt"

func main(){
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	
	if request.HTTPMethod == "GET"{
		var stringResponse string = "Yay a successful Response!!"
		ApiResponse := events.APIGatewayProxyResponse{Body: stringResponse, StatusCode: 200}
		return ApiResponse, nil
	} else {
		err := errors.New("Method Not Allowed!")
		ApiResponse := events.APIGatewayProxyResponse{Body: "Method Not OK", StatusCode: 502}
		return ApiResponse, err
	}
	
	// return events.APIGatewayProxyResponse{
	// 	Body: fmt.Sprintf("Body test"), 
	// 	StatusCode: 200,
	// }, nil
}