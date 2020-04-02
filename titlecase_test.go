package titlecase

import "testing"

func TestPortuguese(t *testing.T) {
	var tests = []struct {
		input  string
		expect string
	}{
		{
			input:  "SINAIS E SISTEMAS",
			expect: "Sinais e Sistemas",
		},
		{
			input:  "CÁLCULO II",
			expect: "Cálculo II",
		},
		{
			input:  "Física Teórica 1",
			expect: "Física Teórica 1",
		},
		{
			input:  "ALGORITMOS E ESTRUTURAS DE DADOS I",
			expect: "Algoritmos e Estruturas de Dados I",
		},
		{
			input:  "INTRODUÇÃO À ÁLGEBRA",
			expect: "Introdução à Álgebra",
		},
		{
			input:  "ALGORITMOS E TEORIA DOS GRAFOS",
			expect: "Algoritmos e Teoria dos Grafos",
		},
		{
			input:  "TRABALHO DE CONCLUSÃO DE CURSO 1 EM ORGANIZAÇÃO E ARQUITETURA DE COMPUTADORES",
			expect: "Trabalho de Conclusão de Curso 1 em Organização e Arquitetura de Computadores",
		},
		{
			input:  "SAÚDE SEXUALIDADE NA ADOLESCÊNCIA",
			expect: "Saúde Sexualidade na Adolescência",
		},
		{
			input:  "DETECÇÃO PRECOCE PARA DROGAS DE ABUSO",
			expect: "Detecção Precoce para Drogas de Abuso",
		},
		{
			input:  "ITERAÇÃO HUMANO-COMPUTADOR",
			expect: "Iteração Humano-Computador",
		},
		{
			input:  "CÁLCULO 1A",
			expect: "Cálculo 1A",
		},
		{
			input:  "2345",
			expect: "2345",
		},
		{
			input:  "23a 45b",
			expect: "23A 45B",
		},
	}

	for _, tt := range tests {
		assertEqual(t, tt.expect, Portuguese(tt.input))
	}
}

func assertEqual(t *testing.T, got, expect string) {
	t.Helper()
	if got != expect {
		t.Errorf("got %q but expected %q", got, expect)
	}
}
