Com certeza! Aqui está o conteúdo do `README.md` adaptado totalmente para o português do Brasil (PT-BR), mantendo toda a formatação profissional para o seu repositório:

```markdown
# Scarf 🧣

O **Scarf** é uma ferramenta de linha de comando (CLI) desenvolvida em Go para automatizar a criação de boilerplate de código em projetos que seguem os princípios da **Clean Architecture** (Arquitetura Limpa) e **Domain-Driven Design (DDD)**, adaptada especialmente para o padrão idiomático do ecossistema Go.

A ferramenta automatiza a criação de módulos isolados (Vertical Slices) dentro do diretório `internal/`, gerando toda a estrutura de camadas necessária para um novo domínio de negócio em segundos.

---

## 🏗️ Arquitetura dos Módulos Gerados

Ao criar un novo módulo, o **Scarf** gera uma estrutura padronizada baseada em separação de conceitos e inversão de dependências:

```text
internal/
└── <modulo>/
    ├── dto.go        # Data Transfer Objects (estruturas de entrada e saída da aplicação)
    ├── entity.go     # Entidades do Domínio (regras de negócio puras e core da aplicação)
    ├── repository.go # Interface de abstração da camada de persistência de dados
    └── usecase.go    # Casos de Uso (lógica de aplicação e fluxos de negócio)

```

### Fluxo de Dependência

A arquitetura segue rigorosamente a regra de que o core da aplicação (Entidades e Casos de Uso) não deve conhecer detalhes de infraestrutura (como bases de dados ou rotas HTTP):

$$\text{Infra (HTTP/DB)} \longrightarrow \text{Usecase (Aplicação)} \longrightarrow \text{Entity (Domínio)}$$

---

## 🚀 Instalação

Você pode instalar o **Scarf** globalmente na sua máquina diretamente a partir do repositório através do comando `go install`:

```bash
go install [github.com/p-v-dev/scarf@latest](https://github.com/p-v-dev/scarf@latest)

```

> **Nota para usuários Windows:** Garanta que o diretório `$GOPATH/bin` (geralmente `C:\Users\Seu-Usuario\go\bin`) esteja adicionado às Variáveis de Ambiente (`PATH`) do seu sistema ou configurado no seu `$PROFILE`.

---

## 🛠️ Como Utilizar

Navegue até a raiz do seu projeto de API Go e execute o comando `mod` passando o nome do domínio que deseja criar.

```bash
scarf mod --name user

```

### Atalho

Você pode utilizar a flag curta `-n` em vez de `--name`:

```bash
scarf mod -n product

```

### Exemplo de Output no Terminal:

```text
✔ Criado: internal/user/entity.go
✔ Criado: internal/user/dto.go
✔ Criado: internal/user/repository.go
✔ Criado: internal/user/usecase.go

Módulo 'user' gerado com sucesso em internal/user! 🚀

```

---

## 💻 Desenvolvimento Local

Se você quiser clonar o repositório para fazer alterações ou estender os templates de código:

1. Clone o repositório:
```bash
git clone [https://github.com/p-v-dev/scarf.git](https://github.com/p-v-dev/scarf.git)
cd scarf

```


2. Instale as dependências (Cobra CLI):
```bash
go mod tidy

```


3. Compile e instale localmente a partir da pasta atual:
```bash
go install .

```



## 🛠️ Tecnologias Utilizadas

* [Go (Golang)](https://go.dev/)
* [Cobra CLI](https://github.com/spf13/cobra) - Framework para criação de aplicações de linha de comando em Go.
* `text/template` - Pacote nativo do Go para renderização dinâmica de arquivos de código.

---

Desenvolvido com ☕ por Pedro Vitor Rodrigues Brito.

```

```
