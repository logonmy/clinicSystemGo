CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go  
ssh root@47.93.206.157 "cd /root/projects/clinicSystemGo;git pull"
scp main root@47.93.206.157:/root/projects/clinicSystemGo/main1 
ssh root@47.93.206.157 "supervisorctl stop go-server"
ssh root@47.93.206.157 "mv /root/projects/clinicSystemGo/main1 /root/projects/clinicSystemGo/main"
ssh root@47.93.206.157 "supervisorctl start go-server"