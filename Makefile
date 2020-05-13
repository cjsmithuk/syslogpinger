
build:
	docker build -t syslogpinger:latest .
runtcp:
	docker run -d syslogpinger:latest -proto=tcp -host=192.168.178.220
