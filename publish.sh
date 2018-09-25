CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go  
ssh root@47.92.128.90 "cd /root/projects/clinicSystemGo;git pull"
scp main root@47.92.128.90:/root/projects/
ssh root@47.92.128.90 "supervisorctl stop go-server"
ssh root@47.92.128.90 "mv /root/projects/main /root/projects/clinicSystemGo/main"
ssh root@47.92.128.90 "supervisorctl start go-server"