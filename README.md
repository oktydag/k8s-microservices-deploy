# k8s-microservices-deploy

- This shows you how to create, deploy and scale your microservices via k8s.

###  Download or Clone this repository

```
https://github.com/oktydag/k8s-microservices-deploy.git
```
### Architecture for single pod

![diagram-single-pod](https://raw.githubusercontent.com/oktydag/k8s-microservices-deploy/master/contents/diagram-single-pod.PNG)


### Run Applications

<pre>$ docker run -d -p 5005:5005 --name cs-app oktydag/customer-services-app:1.0
</pre>

<pre>$ docker run -d -p 6379 :6379 --name redis redis
</pre>

<pre>$ docker run -d -p 5006:5006 --name ns-app oktydag/notification-services-app:1.0
</pre>

Alternatively, you can use **docker compose** in location of docker-compose.yml file;
<pre>$ docker-compose up --build
</pre>

### Deploy into k8s
<pre>$ cd k8s-files
</pre>

**for customer-services-app;**

<pre>$ kubectl create -f customer-services-app-deployment.yml
$ kubectl create -f customer-services-app-services.yml
</pre>

**for redis;**

<pre>$ kubectl create -f redis-deployment.yml
$ kubectl create -f redis-services.yml
</pre>

**for notification-services-app;**

<pre>$ kubectl create -f notification-services-app-deployment.yml
$ kubectl create -f notification-services-app-services.yml
</pre>


### Monitor your pods and services

If you set replicas parameter equals to 2 for customer&notification services; you will see that; 

<pre>$ kubectl get pods, svc
</pre>

![k8s-get-pods-svc](https://raw.githubusercontent.com/oktydag/k8s-microservices-deploy/master/contents/k8s-get-pods-svc.PNG)

### Scale your application

If you want to scale your applcation, for example customer-services-app, from 2 to 4, you can run this;

<pre>$  kubectl scale deployment customer-services-app-deployment --replicas=4
</pre>

And you can see scaled application;
<pre>$  kubectl get deployment
</pre>

![k8s-get-pods-svc](https://raw.githubusercontent.com/oktydag/k8s-microservices-deploy/master/contents/k8s-get-scaled-deployment-list.PNG)

and your architecture will be like this;

![diagram-slaced-pods-replicas](https://raw.githubusercontent.com/oktydag/k8s-microservices-deploy/master/contents/diagram-slaced-pods-replicas.PNG)

