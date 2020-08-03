# aws_lambda_events
Save CloudTrail events to DynamoDB

How it works:
EventBridge (CloudTrail Event) - LAMBDA -  DynamoDB 

Events:
- StopInstances
- StartInstances
- DeleteUSer
- CreateUser
> Note: API Call is Event taht you can procced and save to db

