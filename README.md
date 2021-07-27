#### Basic_Restaurant_Menu
I'm using this to serve menu via a webpage using Go Lang

#### Golang Environment: Installation 
1. Download the latest tarball from the Golan page. 
2. remove any existing Go installation
  * sudo apt list --installed | grep gcc
  * sudo apt remove gccgo-go 
3. Install the tar file: 'sudo tar -C /usr/local -xzf go.1.16.5.linux-amd64.tar.gz'
4. 'cd /usr/local/go/bin'
5. './go version/
6. after verifying system location can run 'go', you are now ready to create the local environments. 

#### Local Environment 
1. directories 
'''
mkdir ~/dev_local/go
cd ~/dev_local/go
mkdir src
mkdir tmp
mkdir bin
mkdir pkg
'''
2. .bash_profile
'''
export GOPATH="/home/prateek/dev_local/go"
export PATH=$PATH:/usr/local/go/bin:$GOPATH
export GOTMPDIR="/home/prateel/dev_local/go/tmp"
'''

