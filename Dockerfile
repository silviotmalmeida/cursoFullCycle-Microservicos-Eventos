# definindo a imagem base
FROM golang:1.20

# definindo a pasta de trabalho a ser criada e focada no acesso
WORKDIR /go/src

# comandos necessários
RUN apt-get update && apt-get install -y librdkafka-dev

# comando para manter o container funcionando
CMD ["tail", "-f", "/dev/null"]