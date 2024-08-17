## To publish a Helm chart on GitHub Pages and enable GitHub Pages on the gh-pages branch, follow these detailed steps:

1. Create the gh-pages Branch of your repository

2. Navigate to the directory containing your Helm chart (Chart.yaml) and run:

```
helm package job-queue/
```

This will create a .tgz file (e.g., your-app-0.1.0.tgz).

3. Generate index.yaml: Move the .tgz file to the root of your repository or create a directory structure you prefer. Generate an index.yaml file:

```
helm repo index .
git add .
git commit -m "Add Helm chart"
git push origin gh-pages
```

4. Enable GitHub Pages
Go to your repository settings:

On GitHub, navigate to your repository page.
Click on the Settings tab.
Find the GitHub Pages section:

Scroll down to the GitHub Pages section in the settings.
Configure the GitHub Pages settings:

In the Source section, select the gh-pages branch from the dropdown menu.
Click Save.


Copy the GitHub Pages URL:

After saving, GitHub Pages will provide a URL where your pages are hosted. It will typically be https://your-username.github.io/your-helm-charts.


5. Add the Repository to Helm. Once GitHub Pages is set up, you need to add your repository to Helm on your local machine:

```
helm repo add job-queue https://puneet105.github.io/job-queue/
```

6. Update Helm Repositories. Update your local Helm repositories to pull in the latest charts:

```
helm repo update
```

7. Install the Helm Chart. You can now install your Helm chart from the repository:

```
helm install -n job-test job-queue job-queue --create-namespace=true
```

## Testing the Application

#### Port forwarding 

```
kubectl -n job-test port-forward svc/job-queue 8080:8080

```

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
Command to check an output.

```
kubectl -n job-test logs <APP_POD_NAME> -f
```

You will get an output in below format. 

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
