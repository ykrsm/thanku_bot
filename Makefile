FILE=prod_ip
IP=`cat $(FILE)`

prod:
	./main -p

test:
	./main -d

dev:
	env GOOS=linux GOARCH=386 go build -o main && docker-compose up --build

deploy:
	git checkout master
	git pull
	git status
	env GOOS=linux GOARCH=386 go build -o main
	scp ./main slack@$(IP):/home/slack/thanku-bot  
	echo Deployed source to IP: $(IP)
