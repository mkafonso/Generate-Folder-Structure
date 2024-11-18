package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func shouldIgnore(name string, ignoreList []string) bool {
	// ignora arquivos e pastas que começam com ponto (.)
	if strings.HasPrefix(name, ".") {
		return true
	}

	// verifica se o nome está na lista de pastas a serem ignoradas
	for _, ignore := range ignoreList {
		if name == ignore {
			return true
		}
	}
	return false
}

func generateStructure(path string, indent string, ignoreList []string) string {
	var structure strings.Builder
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("erro ao ler o diretório:", err)
		return ""
	}

	// ordena os itens para listagem alfabética
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for i, entry := range entries {
		// verifica se deve ignorar o item
		if shouldIgnore(entry.Name(), ignoreList) {
			continue
		}

		// define o prefixo ├── para itens intermediários e └── para o último item
		prefix := "├── "
		if i == len(entries)-1 {
			prefix = "└── "
		}

		// adiciona o item na estrutura
		structure.WriteString(fmt.Sprintf("%s%s%s\n", indent, prefix, entry.Name()))

		// se for um diretório, chama recursivamente
		if entry.IsDir() {
			subPath := filepath.Join(path, entry.Name())
			subIndent := indent + "│   "
			structure.WriteString(generateStructure(subPath, subIndent, ignoreList))
		}
	}

	return structure.String()
}

func generateReadme(rootPath string, ignoreList []string) {
	readmeContent := "# Estrutura do Projeto\n\n```\n"
	readmeContent += generateStructure(rootPath, "", ignoreList)
	readmeContent += "```\n"

	readmeFile := filepath.Join(rootPath, "README.md")
	err := os.WriteFile(readmeFile, []byte(readmeContent), 0644)
	if err != nil {
		fmt.Println("erro ao criar o README.md:", err)
		return
	}

	fmt.Println("README.md gerado com sucesso!")
}

func main() {
	// flag para pastas a serem ignoradas
	var ignoreDirs string
	flag.StringVar(&ignoreDirs, "ignore", "", "pastas a serem ignoradas, separadas por vírgula (ex: 'node_modules,vendor')")
	flag.Parse()

	// divide a lista de pastas a serem ignoradas
	var ignoreList []string
	if ignoreDirs != "" {
		ignoreList = strings.Split(ignoreDirs, ",")
	}

	// obtém o diretório atual (raiz do projeto)
	rootPath, err := os.Getwd()
	if err != nil {
		fmt.Println("erro ao obter o diretório atual:", err)
		return
	}

	fmt.Println("gerando README.md para o diretório:", rootPath)
	generateReadme(rootPath, ignoreList)
}
