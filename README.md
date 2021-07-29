#### Basic_Restaurant_Menu
I'm using this to serve menu via a webpage using Go Lang

#### Golang Environment: Installation 
1. Download the latest tarball from the Golan page. 
2. remove any existing Go installation:
  * sudo apt list --installed | grep gcc
  * sudo apt remove gccgo-go 
3. Install the tar file: `sudo tar -C /usr/local -xzf go.1.16.5.linux-amd64.tar.gz`
4. `cd /usr/local/go/bin`
5. `./go version`
6. after verifying system location can run 'go', you are now ready to create the local environments. 

#### Local Environment 
1. directories 
```
mkdir ~/dev_local/go
cd ~/dev_local/go
mkdir src
mkdir tmp
mkdir bin
mkdir pkg
```
2. .bash_profile
```
export GOPATH="/home/prateek/dev_local/go"
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
export GOTMPDIR="/home/prateel/dev_local/go/tmp"
```

#### Creating Service for Webpage to run. 
1. Create a service at `/etc/systemd/system/pch.service`
2. Update the file using nano with below information.
```
[Unit]
Description=pch service
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
User=prateek
ExecStart=/home/prateek/dev_local/go/bin/Basic_Restaurant_Menu

[Install]
WantedBy=multi-user.target

```
3. Start the service using below commands:
```
sudo systemctl daemon-reload
sudo systemctl start pch.service
```
4. Verify the same using `sudo systemctl status pch.service`

#### Bash file to run git commands
1. Create a bash file using below commands with the file name `commit.sh --> customize it to your choice`.
```
#! /bin/zsh
git add .
git commit -m "Message"
git push
```
2. Update permissions for the file using `chmod u+x commit.sh`
3. Execute it using `./commit.sh`

#### Issue you may run into
1. Certs not accessible --> always check for permissions and the path
2. Page not load even with successful executable run --> check to see both Ubuntu firewall ports and External facing firewall ports when using public IP or domain
