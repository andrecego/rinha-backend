# Rinha de Backend

## Rodando o projeto

Primeiro faça o build da imagem com o comando:

```bash
docker build -t api .
```

Em seguida suba o projeto com o comando:

```bash
docker compose up
```

### Teste de stress

Baixe o [gatling](https://gatling.io/open-source/), coloque no path seguindo `GATLING_BIN_DIR` definido
no `stress-test/runTest.sh` e após isso execute o script `stress-test/runTest.sh`.

Os resultados serão gerados em `stress-test/results`.
