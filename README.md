# aws_lambda_events
Save CloudTrail events to DynamoDB

How it works:
EventBridge (CloudTrail Event) - LAMBDA -  DynamoDB 

Events:
- StopInstances
- StartInstances
- DeleteUSer
- CreateUser
> Note: Any API Call is Event that you can procced and save to db

