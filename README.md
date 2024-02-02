# go-notes
A blogging taking app written in go

Work In Progress

```bash
sudo docker-compose up
```

```bash
# register user POST /api/user/register
curl -X POST -H "Content-Type: application/json" -d '{"name":"Biswa Prakash","email": "biswaprakash444@gmail.com","password":"pass","username":"biswa_1"}' localhost:8080/api/user/register

{"message":"User Created Successfully","registration_date":"2024-02-02T13:18:22.702018Z"}
```

```bash
# login user POST /api/user/login
curl -X POST -H "Content-Type: application/json" -d '{"email": "biswaprakash444@gmail.com","password":"pass"}' localhost:8080/api/user/login

{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJpc3dhcHJha2FzaDQ0NEBnbWFpbC5jb20iLCJleHAiOjE3MDY4OTI0OTUsImlhdCI6MTcwNjg5MTU5NX0.k4P3nvrfaBw29kpfcIJfuc3yNVSe7RzP1GW80Tmu6iI","message":"User Loggedin Successfully"}
```