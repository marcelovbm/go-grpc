# go-grpc

Este projeto busca apenas realizar um implementação simples de uma conexão gRPC utilizando a linguagem Go.

O principal objetivo é o de aprender e conseguir visualizar os modos possível de comunição, tais como: 

- (Client) request -> (Server) respose
- (Client) request -> (Server) stream
- (Client) stream -> (Server) response
- (Client) stream -> (Server) stream

Para que o projeto seja executado, é necessário ter a linguagem Go instalada em sua máquina.
- [Go](https://golang.org/)

Para subir o server da aplicação é preciso executar o seguinte comando:

~~~shell
go run cmb/server/server.go
~~~

A ação de executar o client fará com que a função já seja executada. No arquivo `cmd/client/client.go` há os métodos que serão executados.

~~~shell
go run cmb/client/client.go
~~~