CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go  
ssh root@47.93.206.157 "supervisorctl stop go-server"
scp main root@47.93.206.157:/root/projects/clinicSystemGo/ 
ssh root@47.93.206.157 "supervisorctl start go-server"