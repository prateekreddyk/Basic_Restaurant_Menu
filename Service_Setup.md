Please find the information on the service that was setup to run HTTPS

[Unit]
Description=PCH service
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
