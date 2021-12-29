Gerando os arquivos
```
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```

Testando as request com o evans

```
evans -r repl --host localhost --port 5000
```