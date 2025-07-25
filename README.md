# go_project

## preRequist:
1. Docker
2. kubectl
3. minikube

## Steps
1. build and push image to docker hub using 
    docker build -t arunmittal53/go_project:latest .
    docker push arunmittal53/go_project:latest
2. execute all these deployments (dependencies)
    kubectl apply -f postgres-deployment.yaml
    kubectl apply -f redis-deployment.yaml
    kubectl apply -f app-deployment.yaml
3. Port-forward directly to the pod
    kubectl port-forward svc/go-app 8080:8080
4. test using:
    http://localhost:8080/...

## Inside redis container
1. kubectl get pods | grep redis
2. kubectl exec -it <redis-pod-name> -- sh
OR directly using UI
3. redis-cli
4. Normal redis commands

## Inside postgres container
1. kubectl get pods | grep postgres
2. kubectl exec -it <postgres-pod-name> -- sh
OR directly using UI
3. psql -U <DB_USER> -d <DB_NAME> i.e. postgres & gorm_db