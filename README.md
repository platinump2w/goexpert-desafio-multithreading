# Desafio 2 - Multithreading
Usar o que aprendemos com Multithreading e APIs.

## Descrição
O objetivo deste desafio é buscar e utilizar o primeiro resultado devolvido entre duas APIs distintas.

## Requisitos
- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta;
- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou;
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Execução
Para consultar um CEP padrão, execute o comando abaixo:
```go
go run main.go
```

Para consultar um CEP específico, execute o comando seguindo o exemplo abaixo:
```go
go run main.go 12345-678
```

