package main

import (
	"fmt"
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Nodo representa um elemento da pilha
type Nodo struct {
	palavra string
	prox    *Nodo
}

// Pilha representa a pilha
type Pilha struct {
	top *Nodo
}

// Empilhar adiciona um elemento ao topo da pilha
func (p *Pilha) Empilhar(palavra string) {
	novoNodo := &Nodo{palavra: palavra, prox: p.top}
	p.top = novoNodo
}

// Desempilhar remove e retorna o elemento do topo da pilha
func (p *Pilha) Desempilhar() (string, error) {
	if p.top == nil {
		return "", fmt.Errorf("a pilha está vazia")
	}
	palavra := p.top.palavra
	p.top = p.top.prox
	return palavra, nil
}

// Topo retorna o elemento do topo da pilha sem removê-lo
func (p *Pilha) Topo() (string, error) {
	if p.top == nil {
		return "", fmt.Errorf("a pilha está vazia")
	}
	return p.top.palavra, nil
}

// EstaVazia verifica se a pilha está vazia
func (p *Pilha) EstaVazia() bool {
	return p.top == nil
}

func main() {
	// Criação da aplicação Fyne
	meuApp := app.New()
	minhaJanela := meuApp.NewWindow("Mini Editor de Texto")

	// Componentes da interface
	rotulo := widget.NewLabel("Mini Editor de Texto")
	entradaTexto := widget.NewMultiLineEntry()
	pilha := &Pilha{}

	// Adicionar o estado inicial na pilha
	pilha.Empilhar("")

	// Função para atualizar a pilha a cada modificação no texto
	entradaTexto.OnChanged = func(conteudo string) {
		pilha.Empilhar(conteudo)
	}

	// Botão de desfazer
	botaoDesfazer := widget.NewButton("Desfazer", func() {
		// Desfazer a última mudança
		if !pilha.EstaVazia() {
			_, _ = pilha.Desempilhar()           // Remove o estado atual
			anterior, err := pilha.Desempilhar() // Obtém o estado anterior
			if err == nil {
				entradaTexto.SetText(anterior)
				pilha.Empilhar(anterior) // Adiciona de volta à pilha após desfazer
			}
		}
	})

	// Configuração do layout
	minhaJanela.SetContent(container.NewVBox(
		rotulo,
		botaoDesfazer,
		entradaTexto,
	))

	// Ajuste do tamanho da janela
	minhaJanela.Resize(fyne.NewSize(400, 300))

	// Exibição da janela
	minhaJanela.ShowAndRun()
}
