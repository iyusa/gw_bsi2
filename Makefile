app := bjb
ssh_host := 180.250.246.5
ssh_port := 22102
ssh_pass := secret
doc := xref/Spek\ Virtual\ Account\ Web\ Service\ API\ v.1.1.pdf
mysql_host := microsoft.com
mysql_user := gates
tele_file := bin/$(app).bin

help:
	cat Makefile

test:
	@echo ${doc}	

build:
	@go clean
	@env GOOS=linux go build -o bin/$(app).bin main.go
	@cp $(app).yaml bin/
	@echo "bin/$(app).bin has been created ..."
	
build-local:
	@go clean
	@go build -o bin/$(app).bin main.go
	@cp $(app).yaml bin/
	@echo "local bin/$(app).bin has been created ..."
	
run: 
	@cd bin; ./$(app).bin

dev: build-local run

 
upgrade:	
	scp -P$(ssh_port) bin/$(app).bin root@$(ssh_host):/opt

ssh:
	ssh -p$(ssh_port) root@$(ssh_host)	

git:
	git config credential.helper store	

doc:
	open -a preview $(doc)

mysql:
	ssh $(mysql_user)@$(mysql_host) -L 3306:127.0.0.1:3306 -N    
    
telegram:
	curl -F document=@"$(tele_file)" https://api.telegram.org/bot712486832:AAGwSjxN72iyBQsSyCDY7MHOJDzPHH7JcaQ/sendDocument?chat_id=59575204    	


