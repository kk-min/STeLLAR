package hellojava;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.LambdaLogger;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyRequestEvent;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyResponseEvent;
import java.util.Map;

public class Handler implements RequestHandler<APIGatewayProxyRequestEvent, APIGatewayProxyResponseEvent>{

  @Override
  public Void handleRequest(APIGatewayProxyRequestEvent event, Context context)
  {
    LambdaLogger logger = context.getLogger();
    logger.log("EVENT TYPE: " + event.getClass());
	int incrementLimit = event.getQueryStringParameters().getOrDefault("incrementLimit", 0);
	this.simulateWork(incrementLimit);
	String requestId = "no-context";
	if (context != null) {
		requestId = context.getAwsRequestId();	
	}
	APIGatewayProxyResponseEvent response = new APIGatewayProxyResponseEvent();
	response.setIsBase64Encoded(false);
	response.setStatusCode(200);

  }

  public void simulateWork(int incrementLimit) {
	int i = 0;
	while (i < incrementLimit) {
	  i++;
	}
  }
}
