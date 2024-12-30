# Curso Full Cycle 3.0 - Módulo Event Drive Architecture

<div>
    <img alt="Criado por Daniel Alecrin" src="https://img.shields.io/badge/criado%20por-Daniel Alecrin-%23f08700">
    <img alt="License" src="https://img.shields.io/badge/license-MIT-%23f08700">
</div>

---

## Descrição

O Curso Full Cycle é uma formação completa para fazer com que pessoas desenvolvedoras sejam capazes de trabalhar em projetos expressivos sendo capazes de desenvolver aplicações de grande porte utilizando de boas práticas de desenvolvimento.

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
