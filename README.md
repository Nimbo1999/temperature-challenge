<h4 align="center">Sistema de temperatura por CEP by: <a href="https://www.linkedin.com/in/matheuslopes1999/" target="_blank">Matheus Lopes</a>.</h4>
<p align="center">Desafio proposto no módulo de laboratório do curso de pós graduação Pós Go Expert da Fullcycle</p>

<p align="center">
  <img src="https://img.shields.io/badge/Tests-Passing-2ea44f" alt="Tests - Passing">
  <img src="https://img.shields.io/badge/Go-1.21.x-2ea44f" alt="Go - 1.21.x">
</p>

<p align="center">
  <a href="#como-rodar-a-imagem-docker">Como rodar a imagem docker</a> •
  <a href="#consultando-a-api">Consultando a API</a> •
  <a href="#acessar-pelo-cloud-run">Acessar pelo cloud run</a> •
  <a href="#rode-os-testes">Rode os testes</a> •
</p>

## Como rodar a imagem docker

Para clonar essa applicação você precisará ter instalado o [git](https://git-scm.com) e o [golang](https://go.dev/) em sua máquina. Insira os seguintes commandos em sua CLI para iniciar e rodar a instancia docker:

```bash
# Clone este repositório
$ git clone https://github.com/Nimbo1999/temperature-challenge.git

# Navegue no repositório
$ cd temperature-challenge

# Adicione a sua Weather API Key na variável de ambiente subistituindo a API_KEY pela sua KEY.
$ echo "$(cat .env.example)API_KEY" > .env

# Inicie o projeto com o Makefile
$ make start-build

# Ou inicie o docker container que possui as duas aplicações e o zipkin.
$ docker compose up -d --build
# Navegue na pasta /api para executar uma requisição HTTP ou realize um POST para o serviço de cep que ficará disponível na porta :8080 com o cep no payload i.e.: { "cep": "01001-000" }
```

## Consultando a API

Assim que o projeto estiver rodando com todos os container docker, você já pode interagir com a applicação por meio de uma requisição http POST. Esse end-point espera receber um CEP no body e retorna a temperatura de sua cidade em Graus Celsius, Fahrenheit e Kelvin. Veja o example logo abaixo:

```bash
$ curl -X POST \
-H "Content-Type: application/json" \
-d '{"cep":"01001000"}' \
http://localhost:8080
```

Ou se preferir, caso tenha instalado a extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client), abra a pasta `/api` e fique avontade para alterar o cep e executar a requisição.

## Acessando o dashboard do Zipkin

Assim que a primeira requisição for feita, acesse o painel do zipkin para consultar as métricas disponíveis em http://localhost:9411/

## Acessar pelo cloud run

> **Deprecated** Cloud run foi depreciado e não pode ser mais usado já que não faz mais parte desse desafio.

Essa aplicação possui um deploy ativo feito pelo Cloud Run. Interaja com essa aplicação em produção a partir de sua CLI com o seguinte comando:

> **Informação:**
> Note que o valor `01001000` representa o CEP que deseja consultar a temperatura, podendo ser alterado por um outro valor a qualquer momento.

```bash
$ curl https://temperature-challenge-u6rku7jkaq-uc.a.run.app/01001000
```

## Rode os testes

Foi utilizado a biblioteca `stretchr/testify` para auxiliar com os testes da applicação. Execute o seguinte comando em sua CLI para rodar os testes em sua máquina.

```bash
$ go test ./...
```
