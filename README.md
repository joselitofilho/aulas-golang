# Aulas de programação em Golang

A linguagem de programação Go é um projeto de código aberto para tornar os programadores mais produtivos.

Go é expressivo, conciso, limpo e eficiente. Seus mecanismos de concorrência facilitam a escrita de programas que aproveitam ao máximo as máquinas multicore e em rede, enquanto seu novo tipo de sistema permite a construção de programa flexível e modular. Go compila rapidamente para código de máquina, mas tem a conveniência da coleta de lixo e o poder de reflexão em tempo de execução. É uma linguagem rápida, com tipagem estática e compilada que se parece com uma linguagem interpretada com tipagem dinâmica.

**Índice**

- [Sobre o curso](#sobre-o-curso)
- [Antes de começar...](#antes-de-começar)
  - [Instale o VSCode](#instale-o-vscode)
  - [Instale o Docker](#instale-o-docker)
  - [Instale o plugin para o Docker](#instale-o-plugin-para-o-docker)
- [Estrutura do curso](#estrutura-do-curso)

## Sobre o curso

Neste curso, nós vamos aprender mais sobre a Linguagem Go. Ao longo do curso, serão apresentadas aulas que ensinam desde as características básicas da linguagem, bem como as variáveis e constantes, estruturas de repetições, funções, dentre outras.

## Antes de começar...

Instale a IDE Visual Studio Code - [VSCode](https://code.visualstudio.com/) e o [Docker](https://www.docker.com/). Mas sinta-se a vontade para utilizar as ferrametas que você estiver mais acostumado.

### Instale o VSCode

Para instalar e configurar siga os passos que estão descritos na documentação oficial do VSCode clicando [aqui](https://code.visualstudio.com/docs).

![image](https://user-images.githubusercontent.com/1102589/123715957-ce4b7600-d84f-11eb-82b5-f1b77cebee1a.png)

### Instale o Docker

Para instalar e configurar siga os passos que estão descritos na documentação oficial do Docker clicando [aqui](https://www.docker.com/get-started).

![image](https://user-images.githubusercontent.com/1102589/123716151-44e87380-d850-11eb-8fed-67a4f720caf6.png)

### Instale o plugin para o Docker

Para facilitar o gerenciamento dos containers que vamos criar, você pode instalar um plugin no VSCode para utilizar algumas funções do Docker na própria IDE. 
- Clique [aqui](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker) para instalar o plugin para o Docker.
- Clique [aqui](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) para instalar o plugin para o remote-containers.

### Setup do ambiente de desenvolvimento

Você terá duas opções para fazer a instalação do ambiente de desenvolvimento.
1. Instalando todo o ambiente do GoLang em sua máquina;
2. Utilizando a imagem Docker de GoLang.

Para instalar todo o ambiente GoLang siga os passos que estão descritos na documentação oficial clicando [aqui](https://golang.org/dl/). Mas se você preferir utilizar o container Docker, siga o passo-a-passo a seguir.

_(usando docker)_

#### Build

Windows (PowerShell):
```bash
$ docker build -t go-dev --build-arg GO_VERSION=1.16.5 .
```

Linux ou MacOS:
```bash
$ docker build -t go-dev --build-arg GO_VERSION=1.16.5 .
```

#### Run

Windows (PowerShell):
```bash
$ docker run -v $pwd\:/src -w /src --label com.docker.compose.project=go-dev -itd --name $(Split-Path -Path $pwd -Leaf) go-dev
```

Linux ou MacOS:
```bash
$ docker run -v `pwd`:/src -w /src --label com.docker.compose.project=go-dev -itd --name ${PWD##*/} go-dev
```

## Estrutura do curso

### [Aula 1](aula01): Introdução e fundamentos da linguagem
Link da [video aula](https://drive.google.com/file/d/1b52kKHzRXfG7FqZZOdJmWvxU5-d6yKck/view?usp=sharing)
- Apresentações
- Configuração do ambiente (VSCode + Docker)
- Hello World
- Tipos, Variaveis e Constantes
- Array, Map e Slice
- Estrutura de controle (if / else, switch), estrutura de repetição (for, range)

### [Aula 2](aula02): Ponteiros e estruturas
[Link da video aula](https://drive.google.com/file/d/1vUf2pfGi416pb8tQ0JMXjxpIj02w4Vo0/view?usp=sharing)
- Ponteiros
- Struct
- Interface
- Defer
- Panic
- Recover

### [Aula 3](aula03): Funções
[Link da video aula](https://drive.google.com/file/d/11936Cf_Rr2xjSf3H6K2VNYF_QDEkdqfd/view?usp=sharing)
- Retorno Nomeado
- Variádicas
- Anônimas
- Recursivas
- Closures
- `func init()` / `func main()`

### [Aula 4](aula04): Go (composição) x OO (herança)
[Link da video aula](https://drive.google.com/file/d/1UeicM1GrvwIvov6PEgfdPJGyL-jMq7OB/view?usp=sharing)
- Criação de pacotes
- Instalação de dependências de terceiros (go get, go install, go mod, go vendor)

### [Aula 5](aula05): Concorrência (Goroutines e Channels)
[Link da video aula 5 + 6](https://drive.google.com/file/d/1dbuYzb03xfwz7J6b8dD7H2pKM5JRg4s5/view?usp=sharing)
- Routines
- Waitgroup
- Channels
- Channels com Buffer
- Select

### [Aula 6](aula06): Padrões de Concorrência
[Link da video aula 5 + 6](https://drive.google.com/file/d/1dbuYzb03xfwz7J6b8dD7H2pKM5JRg4s5/view?usp=sharing)
- Worker Pools
- Generator
- Multiplexador

### [Aula 7](aula07): Criar um serviço / API em Go
[Link da video aula 7](https://drive.google.com/file/d/19826yLS-J1ZseZPL-MCTulah6fcR_FPj/view?usp=sharing)
- Acesse o repositório [golang-echo-apigithub](https://github.com/joselitofilho/golang-echo-apigithub) para ver um exemplo de implementação de uma API feita em Go.

## Principais links
- [Effective Go](https://golang.org/doc/effective_go)
- [Playground Go](https://play.golang.org/)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Echo Web Framework](https://echo.labstack.com/)
- [GitHub API](https://docs.github.com/en/rest)
- [Gorm Library](https://gorm.io/)
