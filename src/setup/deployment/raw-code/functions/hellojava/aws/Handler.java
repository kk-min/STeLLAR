package hellojava;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.LambdaLogger;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyRequestEvent;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyResponseEvent;
import java.util.Map;
import java.time.Instant;

public class ResponseEventBody {
	String region;
	String requestId;
	String[] timestampChain;

	public void ResponseEventBody(String region, String requestId, String[] timestampChain) {
		this.region = region;
		this.requestId = requestId;
		this.timestampChain = timestampChain;
	}
}

public class Handler implements RequestHandler<APIGatewayProxyRequestEvent, APIGatewayProxyResponseEvent>{

  @Override
  public Void handleRequest(APIGatewayProxyRequestEvent event, Context context)
  {
    LambdaLogger logger = context.getLogger();
    logger.log("EVENT TYPE: " + event.getClass());
	Gson gson = new Gson();
	int incrementLimit = event.getQueryStringParameters().getOrDefault("incrementLimit", 0);
	this.simulateWork(incrementLimit);
	String requestId = "no-context";
	if (context != null) {
		requestId = context.getAwsRequestId();	
	}


	Instant now = Instant.now();
	String[] timestampChain = new String[]{""+now.getEpochSecond()+now.getNano()};
	ResponseEventBody resBody = new ResponseEventBody(System.getenv("AWS_REGION"), requestId, timestampChain)

	APIGatewayProxyResponseEvent response = new APIGatewayProxyResponseEvent();
	response.setIsBase64Encoded(false);
	response.setStatusCode(200);
	response.setBody(gson.toJson(resBody));

	return response;
  }

  public void simulateWork(int incrementLimit) {
	int i = 0;
	while (i < incrementLimit) {
	  i++;
	}
  }
}
