# News Importer Micro Services

## Build and run images manually
```shell
# worker
docker build -t news-importer-worker .
docker run -d --rm --net=host news-importer-worker
# scheduler
docker build -t news-importer-scheduler .
docker run -d --rm --net=host news-importer-scheduler
# api
docker build -t news-importer-api .
docker run -d --rm --net=host news-importer-api
```

## Run in docker compose
```shell
docker-compose up -d --build
```

## Run in k8s
Start kubernetes cluster
```shell
minikube start
```

Start Istio
```shell
./setup -s
```

Infrastructure deployment
```shell
./setup -i
```

Apps deployment
```shell
./setup -a
```

Restart PODs
```shell
kubectl get deployment # get deployments
kubectl rollout restart deployment <DEPLOYMENT>
```

Lists the URLs for the services in your local cluster
```shell
minikube service list 
```

MySQL restore
```shell
kubectl port-forward news-importer-mysql-<POD> 3306:3306
```
```shell
mysql -u silver -h 127.0.0.1 -p Noticias2011 < docker/mysqldump.news.sql 
```