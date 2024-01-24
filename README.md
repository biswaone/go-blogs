# go-notes
A blogging taking app written in go

Work In Progress

```bash
sudo docker run -p 5432:5432 -e POSTGRES_DB=goblogs -e POSTGRES_USER=goblogs -e POSTGRES_PASSWORD=goblogs --name pg postgres

sudo docker build -t goblogs .

sudo docker run -p 8080:8080  --network host goblogs
```

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Biswa Prakash","email": "biswaprakash444@gmail.com","password":"pass","username":"biswa_1"}' localhost:8080/api/user/register
```