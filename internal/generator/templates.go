package generator

import (
	"bytes"
	"text/template"
)

// Estrutura de dados para passar pro template
type ModuleData struct {
	PkgName string // Ex: "user"
}

const entityTemplate = `package {{.PkgName}}

type {{.PkgName | title}} struct {
	ID uint64
	// Adicione os campos da entidade aqui
}
`

const dtoTemplate = `package {{.PkgName}}

type Create{{.PkgName | title}}Input struct {
	// Campos para criação
}

type {{.PkgName | title}}Output struct {
	ID uint64
}
`

const repositoryTemplate = `package {{.PkgName}}

type Repository interface {
	Create(entity *{{.PkgName | title}}) error
	FindByID(id uint64) (*{{.PkgName | title}}, error)
}
`

const usecaseTemplate = `package {{.PkgName}}

type Usecase interface {
	Execute(input Create{{.PkgName | title}}Input) (*{{.PkgName | title}}Output, error)
}

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Execute(input Create{{.PkgName | title}}Input) (*{{.PkgName | title}}Output, error) {
	// Lógica de negócio aqui
	return nil, nil
}
`

// Função auxiliar para renderizar o template
func GenerateFile(tmplString string, data ModuleData) ([]byte, error) {
	// Funções extras no template, como 'title' para deixar a primeira letra maiúscula
	funcMap := template.FuncMap{
		"title": func(s string) string {
			if len(s) == 0 {
				return ""
			}
			// Tratamento simples para capitalizar a primeira letra
			return string(s[0]-32) + s[1:]
		},
	}

	tmpl, err := template.New("module").Funcs(funcMap).Parse(tmplString)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Exportando os templates para usar no comando da CLI
var Templates = map[string]string{
	"entity.go":     entityTemplate,
	"dto.go":        dtoTemplate,
	"repository.go": repositoryTemplate,
	"usecase.go":    usecaseTemplate,
}
