package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"scarf/internal/generator" // Ajuste o nome do módulo conforme seu go.mod

	"github.com/spf13/cobra"
)

var moduleName string

var modCmd = &cobra.Command{
	Use:   "mod",
	Short: "Gera um novo módulo de domínio",
	Run: func(cmd *cobra.Command, args []string) {
		if moduleName == "" {
			fmt.Println("Erro: O nome do módulo (--name) é obrigatório.")
			os.Exit(1)
		}

		// Caminho onde o módulo será criado (dentro de internal/)
		targetDir := filepath.Join("internal", moduleName)

		// 1. Cria o diretório
		err := os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Erro ao criar diretório: %v\n", err)
			os.Exit(1)
		}

		data := generator.ModuleData{PkgName: moduleName}

		// 2. Loop para gerar cada arquivo baseado nos templates
		for fileName, tmplStr := range generator.Templates {
			content, err := generator.GenerateFile(tmplStr, data)
			if err != nil {
				fmt.Printf("Erro ao processar template %s: %v\n", fileName, err)
				continue
			}

			filePath := filepath.Join(targetDir, fileName)
			err = os.WriteFile(filePath, content, 0644)
			if err != nil {
				fmt.Printf("Erro ao gravar arquivo %s: %v\n", fileName, err)
				continue
			}
			fmt.Printf("✔ Criado: %s\n", filePath)
		}

		fmt.Printf("\nMódulo '%s' gerado com sucesso em %s! 🚀\n", moduleName, targetDir)
	},
}

func init() {
	// Define a flag --name (ou -n) como obrigatória
	modCmd.Flags().StringVarP(&moduleName, "name", "n", "", "Nome do módulo/pacote a ser criado")
	RootCmd.AddCommand(modCmd)
}
