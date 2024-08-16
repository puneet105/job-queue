## Testing the Application

#### Obtain a JWT Token
You need to log in to obtain a JWT token that will be used for subsequent authenticated requests.
1. Login to Get a JWT Token:Use curl to make a POST request to the /login endpoint:bashCopy codecurl -X POST -d "username=admin&password=password" http://localhost:8080/login
2. If the credentials are correct, you'll receive a JWT token as a response.Example response:jsonCopy code"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjEyNTc1MjAwfQ.xvQyFtLdZDBJF6c4z1P1ZJL9A4VpdvTYPjzMfqV4ouY"
3. Save the JWT Token:Save this token for use in the subsequent steps.

#### Publish Jobs to the Queues (Redis and RabbitMQ)
Use the JWT token obtained from the login to publish jobs to RabbitMQ, Kafka, and Redis.
1. Publish a Job:Use curl to make a POST request to the /publish endpoint:bashCopy codecurl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzIzODk4NjM0fQ.V7C8reg3zfYH14rSF5FnT70jox-Lb-N4XMT2B6LxTsw" http://localhost:8080/publish
2. Replace <your_jwt_token> with the token you obtained earlier.
3. Check the Response:If successful, you should see a response:textCopy codePublished messages to all queues


#### Output
You will get an output in below format

```
job-queue-app-1       | 2024/08/16 13:20:53 Server is running on :8080
job-queue-app-1       | Published message To RabbitMQ: Message 1 from Publisher 1
job-queue-app-1       | Published message To RabbitMQ: Message 1 from Publisher 3
job-queue-app-1       | Published message To RabbitMQ: Message 1 from Publisher 2
job-queue-app-1       | Published message to Redis: Message 1 from Publisher 1
job-queue-app-1       | Published message to Redis: Message 1 from Publisher 2
job-queue-app-1       | Published message to Redis: Message 1 from Publisher 3
job-queue-app-1       | Redis: Received a message: Message 1 from Publisher 2
job-queue-app-1       | Redis: Received a message: Message 1 from Publisher 3
job-queue-app-1       | Redis: Received a message: Message 1 from Publisher 1
job-queue-app-1       | RabbitMQ: Received a message: Message 1 from Publisher 1
job-queue-app-1       | RabbitMQ: Received a message: Message 1 from Publisher 3
job-queue-app-1       | RabbitMQ: Received a message: Message 1 from Publisher 2
```