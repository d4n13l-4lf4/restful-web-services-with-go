resource "aws_api_gateway_rest_api" "test" {
  name = "EC2Example"
  description = "Terraform EC2 REST API Example"
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_method" "test" {
  authorization = "NONE"
  http_method = "GET"
  resource_id = aws_api_gateway_rest_api.test.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.test.id
}

resource "aws_api_gateway_method_response" "test" {
  authorization = "NONE"
  http_method = aws_api_gateway_method.test.http_method
  resource_id = aws_api_gateway_rest_api.test.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.test.id
  status_code = "200"
}

resource "aws_api_gateway_integration_response" "MyDemoIntegrationResponse" {
  http_method = aws_api_gateway_method.test.http_method
  resource_id = aws_api_gateway_rest_api.test.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.test.id
  status_code = aws_api_gateway_method_response.test.status_code
}

resource "aws_api_gateway_integration" "test" {
  http_method = aws_api_gateway_method.test.http_method
  resource_id = aws_api_gateway_rest_api.test.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.test.id
  type = "HTTP"
  integration_http_method = "GET"
  uri = "http://${aws_instance.api_server.public_dns}/api/books"
}

resource "aws_api_gateway_deployment" "test" {
  depends_on = [
    aws_api_gateway_integration.test
  ]
  rest_api_id = aws_api_gateway_rest_api.test.id
  stage_name = "test"
}