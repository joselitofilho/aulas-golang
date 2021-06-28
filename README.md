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

A IDE que vamos utilizar será o Visual Studio Code - [VSCode](https://code.visualstudio.com/)). Mas sinta-se a vontade para utilizar a que você estiver mais acostumado.

### Instale o VSCode

Para instalar e configurar siga os passos que estão descritos na documentação oficial do VSCode clicando [aqui](https://code.visualstudio.com/docs).

![image](https://user-images.githubusercontent.com/1102589/123715957-ce4b7600-d84f-11eb-82b5-f1b77cebee1a.png)

### Instale o Docker

Para instalar e configurar siga os passos que estão descritos na documentação oficial do Docker clicando [aqui](https://www.docker.com/get-started).

![image](https://user-images.githubusercontent.com/1102589/123716151-44e87380-d850-11eb-8fed-67a4f720caf6.png)

### Instale o plugin para o Docker

Para facilitar o gerenciamento dos containers que vamos criar, você pode instalar um plugin no VSCode para utilizar algumas funções do Docker na própria IDE. Clique [aqui](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker) para instalar.

## Estrutura do curso

### Aula 1: Introdução e fundamentos da linguagem
- Apresentações
- Configuração do ambiente (VSCode + Docker)
- Hello World
- Tipos, Variaveis e Constantes
- Array, Map e Slice
- Estrutura de controle (if / else, switch), estrutura de repetição (for, range)

### Aula 2: Ponteiros e estruturas
- Ponteiros
- Struct
- Interface
- Defer
- Panic
- Recover

### Aula 3: Funções
- Retorno Nomeado
- Variádicas
- Anônimas
- Recursivas
- Closures
- `func init()` / `func main()`

### Aula 4: Go (composição) x OO (herança)
- Criação de pacotes
- Instalação de dependências de terceiros (go get, go install, go mod, go vendor)

### Aula 5: Concorrência (Goroutines e Channels)
- Routines
- Waitgroup
- Channels
- Channels com Buffer
- Select

### Aula 6: Padrões de Concorrência
- Worker Pools
- Generator
- Multiplexador

### Aula 7: Criar um serviço / API em Go
