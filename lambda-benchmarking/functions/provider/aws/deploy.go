package aws

import (
	"fmt"
	"functions/util"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func (lambda Interface) DeployFunction(i int, language string) string {
	lambda.createFunction(i, language)
	arn := lambda.getFunctionARN(i)

	lambda.createRESTAPI(i)
	apiID := lambda.GetAPIID(i)
	resourceID := lambda.getResourceID(i, apiID)

	lambda.createAPIFunctionIntegration(i, apiID, resourceID, arn)
	lambda.createAPIDeployment(i, apiID)
	lambda.addExecutionPermissions(i)
	return apiID
}

func (lambda Interface) getFunctionARN(i int) string {
	cmd := exec.Command("/usr/local/bin/aws", "lambda", "list-functions", "--query",
		fmt.Sprintf("Functions[?FunctionName==`%s-%v`].FunctionArn", name, i), "--output", "text",
		"--region", region)
	arn := util.RunCommandAndLog(cmd)
	arn, _ = strconv.Unquote(strings.ReplaceAll(strconv.Quote(arn), `\n`, ""))
	log.Printf("ARN of lambda %s-%v is %s", name, i, arn)
	return arn
}

func (lambda Interface) createFunction(i int, language string) {
	log.Printf("Creating producer lambda %s-%v", name, i)
	cmd := exec.Command("/usr/local/bin/aws", "lambda", "create-function", "--function-name",
		fmt.Sprintf("%s-%v", name, i), "--runtime", language, "--role", util.CheckAndReturnEnvVar("AWS_LAMBDA_ROLE"),
		"--handler", "producer-handler", "--zip-file", fmt.Sprintf("fileb://%s.zip", name),
		"--tracing-config", "Mode=PassThrough")
	// Set Mode to Active to sample and trace a subset of incoming requests with AWS X-Ray.PassThrough otherwise.
	util.RunCommandAndLog(cmd)
}

func (lambda Interface) createRESTAPI(i int) {
	log.Printf("Creating corresponding API %s-API-%v (clone of %s)", name, i, cloneAPIID)
	cmd := exec.Command("/usr/local/bin/aws", "apigateway", "create-rest-api", "--name",
		fmt.Sprintf("%s-API-%v", name, i), "--description",
		fmt.Sprintf("The API used to access benchmarking Lambda function %v", i), "--endpoint-configuration",
		"types=REGIONAL", "--region", region, "--clone-from", cloneAPIID)
	util.RunCommandAndLog(cmd)
}

// Note: `items[1].id` for US, `items[0].id` for EU
func (lambda Interface) getResourceID(i int, apiID string) string {
	// items[0].id needed in eu-west-2
	cmd := exec.Command("/usr/local/bin/aws", "apigateway", "get-resources", "--rest-api-id",
		apiID, "--query", "items[1].id", "--output", "text", "--region", region)
	resourceID := util.RunCommandAndLog(cmd)
	resourceID, _ = strconv.Unquote(strings.ReplaceAll(strconv.Quote(resourceID), `\n`, ""))
	log.Printf("RESOURCEID of %s-API-%v is %s", name, i, resourceID)
	return resourceID
}

func (lambda Interface) createAPIFunctionIntegration(i int, apiID string, resourceID string, arn string) {
	log.Printf("Creating integration between lambda %s-%v and API %s-API-%v", name, i, name, i)
	cmd := exec.Command("/usr/local/bin/aws", "apigateway", "put-integration", "--rest-api-id",
		apiID, "--resource-id", resourceID, "--http-method", "ANY", "--type", "AWS_PROXY", "--integration-http-method",
		"ANY", "--uri", fmt.Sprintf("arn:aws:apigateway:%s:lambda:path/2015-03-31/functions/%s/invocations",
			region, arn), "--request-templates",
		`{"application/x-www-form-urlencoded":"{\"body\": $input.json(\"$\")}"}`, "--region", region)
	util.RunCommandAndLog(cmd)
}

func (lambda Interface) createAPIDeployment(i int, apiID string) {
	log.Printf("Creating deployment for API %s-API-%v (stage %s)", name, i, stage)
	cmd := exec.Command("/usr/local/bin/aws", "apigateway", "create-deployment", "--rest-api-id",
		apiID, "--stage-name", stage, "--region", region)
	util.RunCommandAndLog(cmd)
}

func (lambda Interface) addExecutionPermissions(i int) {
	log.Printf("Adding permissions to execute lambda function %s-%v", name, i)
	cmd := exec.Command("/usr/local/bin/aws", "lambda", "add-permission", "--function-name",
		fmt.Sprintf("%s-%v", name, i), "--statement-id", "apigateway-benchmarking", "--action",
		"lambda:InvokeFunction", "--principal", "apigateway.amazonaws.com", "--region", region)
	util.RunCommandAndLog(cmd)
}