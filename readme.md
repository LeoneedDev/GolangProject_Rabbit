Для начала работы нужно создать файл .env в котором необходимо указать:
TOKEN="Ваш токен для бота"
После чего можно запстить с помощью команды "docker build -t go-app ." и "docker run --name=go-tel-app go-app"
ЧТобы остановить процесс необходимо выполнить команд "docker rm go-tel-app"