# Curso Full Cycle 3.0 - Módulo Event Drive Architecture

<div>
    <img alt="Criado por Daniel Alecrin" src="https://img.shields.io/badge/criado%20por-Daniel Alecrin-%23f08700">
    <img alt="License" src="https://img.shields.io/badge/license-MIT-%23f08700">
</div>

---

## Descrição

O Curso Full Cycle é uma formação completa para fazer com que pessoas desenvolvedoras sejam capazes de trabalhar em projetos expressivos sendo capazes de desenvolver aplicações de grande porte utilizando de boas práticas de desenvolvimento.

---

## Desafio

Desenvolva um microsserviço em sua linguagem de preferência que seja capaz de receber via Kafka os eventos gerados pelo microsserviço "Wallet Core" e persistir no banco de dados os balances atualizados para cada conta.

Crie um endpoint: `/balances/{account_id}` que exibe o balance atualizado.

Requisitos para entrega:

- [x] Tudo deve rodar via Docker / Docker-compose
- [x] Com um único docker-compose up -d todos os microsserviços, incluindo o da wallet core precisam estar disponíveis para que possamos fazer a correção.
- [x] Não esqueça de rodar migrations e popular dados fictícios em ambos bancos de dados (wallet core e o microsserviço de balances) de forma automática quando os serviços subirem.
- [x] Gere o arquivo ".http" para realizarmos as chamadas em seu microsserviço da mesma forma que fizemos no microsserviço "wallet core"
- [x] Disponibilize o microsserviço na porta: 3003.

---

## Instruções

Para fazer o `build` de todos serviços no diretório raiz executo o comando `docker compose up --build` com isso as imagens serão baixadas e executadas.

Após todos containers estiverem em rodando, inicie as aplicações:

### Serviço de Wallet

Execute o comando `docker compose exec wallet-app bash` após acessar o container execute o comando `go run cmd/walletcore/main.go` ele irá rodas as migrations e subir o sevidor na porta `8080`.

### Serviço de Balances

Execute o comando `docker compose exec balances-app bash` após acessar o container execute o comando `go run cmd/balances/main.go` ele irá rodas as migrations e subir o sevidor na porta `3003`.

Os dois serviços já tem seu arquivo de `api/client.http` já com os `IDs` corretos, mas nada impede te criar novos registros e usar os mesmos.

---

## Repositório Pai

https://github.com/DanielAlecrin/full-cycle-3-0-curso

---

### O que são eventos?

- Situações que ocorreram no passado;
- Normalmente causa efeitos colaterais:
  - Ex.: A porta do carro abriu, então deve ligar a luz interna;
- Ele pode trabalhar de forma internalizada no software ou externalizada;
- Domain Events (Eventos de domínio):
  - Mudança no estado interno da aplicação / regra de negócios -> Ex.: agregados;

### Tipos de eventos

1. Event Notification: Forma curta de comunicação, ex.: `{ "order": 1, "status": "approved" }`;
2. Event Carried State Transfer: Formato completo para trafegar as informações, ex.: `stream de dados` `{ "order": 1, "status": "approved", "products": [{ ... }, { ... }], "tax": "1%", "client": "John Doe" }`;
3. Event Sourcing: Armazenamento dos eventos baseado em uma linha do tempo, possibilidade de replay;
