# Desafio: Concorrência em Go (Multithreading)

Este projeto tem como objetivo demonstrar a utilização de **concorrência (multithreading)** em Go para resolver um problema de busca de CEP utilizando duas APIs diferentes. A resposta mais rápida é considerada e enviada ao cliente, com timeout de 1 segundo.

---

## Proposto na Pós-graduação Go Expert - Full Cycle

### Enunciado:
> Realizar chamadas concorrentes para as APIs:
> - https://brasilapi.com.br/api/cep/v1/{cep}
> - http://viacep.com.br/ws/{cep}/json/
>
> Retornar apenas a resposta mais rápida ao cliente. Caso nenhuma responda em até 1 segundo, retornar erro de timeout.

---

## Funcionalidades e Conceitos Aplicados

| Conceito / Recurso Go       | Aplicação no projeto                                                                 |
|-----------------------------|----------------------------------------------------------------------------------------|
| **Goroutines**              | Duas goroutines são disparadas para consultar as APIs concorrentes                    |
| **Green Threads (runtime)**| Go gerencia as goroutines como threads leves de forma eficiente                        |
| **Channels (buffered)**     | Canal com buffer é usado para capturar a primeira resposta                           |
| **Select + Timeout**        | Controla o timeout de 1 segundo com `context.WithTimeout`                              |
| **Context**                 | Garante cancelamento das goroutines após timeout                                      |
| **sync/atomic**             | Contadores seguros para registrar qual API respondeu primeiro                         |
| **Logs com timestamps**     | Demonstram execução paralela no terminal com `time.Now()`                           |

---

## Como rodar o projeto

### Requisitos:
- Go 1.21+

### 1. Clonar o projeto
```bash
git clone https://github.com/fjgmelloni/fullcycle/multithreading.git
cd multithreading
```

### 2. Rodar o servidor
```bash
go run main.go
```

### 3. Fazer uma requisição:
```bash
curl http://localhost:8080/cep/01153000
```
Ou acessar pelo navegador:
```
http://localhost:8080/cep/01153000
```

---

## Exemplo de log no terminal:

```
[15:18:00.123] CEP requisitado: 01153000
[15:18:00.123] Iniciando chamada para BrasilAPI
[15:18:00.123] Iniciando chamada para ViaCEP
[15:18:00.456] Resposta recebida da BrasilAPI
```

Isso comprova que as chamadas foram disparadas simultaneamente.

---

## Contadores de métricas

- O sistema registra quantas vezes cada API respondeu primeiro usando `sync/atomic`.
- Pode-se adicionar futuramente um endpoint `/metrics` para retornar essas estatísticas.

---

## Teste de concorrência com goroutines

Para simular múltiplas requisições simultâneas ao servidor, foi adicionado o arquivo `test.go` dentro da pasta `test/`.

Você pode acessar a pasta e executar:
```bash
cd test
go run test.go
```

Esse teste demonstra que o servidor consegue lidar com múltiplas requisições concorrentes.

---

## Conclusão

> Este projeto demonstra de forma clara a utilização de concorrência em Go (multithreading), usando goroutines, channels, contextos com timeout e operações atômicas. O código está preparado para lidar com condições reais de latência, mantendo segurança na escrita e entrega da resposta HTTP.

---

## Autor

- Felício Melloni
- Projeto para a pós-graduação **Go Expert - Full Cycle**

