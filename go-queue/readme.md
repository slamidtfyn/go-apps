# rabbitmq server

docker run -d --hostname my-rabbit --name some-rabbit -p 5672:5672 rabbitmq:3

# server

go run send.go msg_1 msg_2 msg_3 msg_n

# client

go run receive.go client_name
