<h4 align="center">Sistema de temperatura por CEP by: <a href="https://www.linkedin.com/in/matheuslopes1999/" target="_blank">Matheus Lopes</a>.</h4>
<p align="center">Desafio proposto no módulo de laboratório do curso de pós graduação Pós Go Expert da Fullcycle</p>

<p align="center">
  <img src="https://img.shields.io/badge/Tests-Passing-2ea44f" alt="Tests - Passing">
  <img src="https://img.shields.io/badge/Go-1.21.x-2ea44f" alt="Go - 1.21.x">
</p>

<p align="center">
  <a href="#como-rodar-a-imagem-docker">Como rodar a imagem docker</a> •
  <a href="#consultando-a-api">Consultando a API</a> •
  <a href="#credits">Credits</a> •
  <a href="#related">Related</a> •
  <a href="#license">License</a>
</p>

## Como rodar a imagem docker

Para clonar essa applicação você precisará ter instalado o [git](https://git-scm.com) e o [golang](https://go.dev/) em sua maquina. Insira os seguintes commandos em sua CLI para iniciar e rodar a instancia docker:

```bash
# Clone este repositório
$ git clone https://github.com/Nimbo1999/temperature-challenge.git

# Navegue no repositório
$ cd temperature-challenge

# Adicione a sua Weather API Key na variável de ambiente subistituindo a API_KEY pela sua KEY.
$ echo "$(cat .env.example)API_KEY" > .env

# Inicie o build da imagem docker
$ docker build -t temperature-app .

# Rode a applicação
$ docker run -p 8080:8080 -d temperature-app
```

## Consultando a API

Assim que o projeto estiver rodando em um container docker, você já pode interagir com a applicação por meio de uma requisição http GET. Esse end-point espera receber um CEP como parâmetro e retorna a temperatura de sua cidade em Graus Celsius, Fahrenheit e Kelvin. Veja o example logo abaixo:

```bash
$ curl http://localhost:8080/01001000
```

Ou se preferir, caso tenha instalado a extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client), abra a pasta `/api` e fique avontade para alterar o cep e executar a requisição.
