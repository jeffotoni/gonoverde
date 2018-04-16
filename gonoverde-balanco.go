/*
* Go Library (C) 2018 Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
* @project     Noverde
* @package     main
* @author      @jeffotoni
* @size        15/04/2018
 */

package gonoverde

import (
	"bufio"
	// "errors"
	"fmt"
	. "github.com/jeffotoni/gcolor"
	"github.com/jeffotoni/gonoverde/gbolt"
	"log"
	"os"
	"regexp"
	"strings"
	//"time"
)

var errs error

// metodo resposnsavel por gerar os saldos das contas
// dos clientes, ler suas transacoes e fazer o calculo
// de saldo e apresentar na tela
func SaldoContaCliente(ContasFile, TransFile string) {

	// ascii
	// loader
	// RuneTime()

	// limpando o log
	// e iniciando nova
	// execucao
	WriteLogClean()

	// apagando a base
	// de dados para
	// gerar uma nova
	gbolt.DropDatabase()

	// Testing boltdb database
	// Start ping database
	// Creating ping ok
	gbolt.Save("Ping", "ok")

	// Testing whether it was recorded
	// and read on the boltdb, we
	// recorded a Ping and then
	// read it back.
	if gbolt.Get("Ping") != "ok" {

		log.Println("Services Error Data Base!")
		os.Exit(0)
	}

	// boltdb testado e ok
	//fmt.Println(CyanCor("Bolt Testado e funcionando.."))
	//fmt.Println(PurpleCor("Salvando Contas no banco de dados Bolt"))

	// Iremos ler todo arquivo de contas.csv e coloca-lo no banco noSql boltdb
	// Apos gerar o banco de dados contas, iremos abrir e percorrer o arquivo
	// transacoes.csv, neste caso não iremos precisar gravar no banco porque
	// o arquivo encontra-se ordenado por idconta, os saldos dos clientes
	// estão ordenados ...
	// Tudo correndo bem teremos todos os saldos no banco já no formato
	// float da forma que precisamos para calcular
	errs = LerFileSaveDb(ContasFile)

	if errs != nil {

		fmt.Println("Error ao ler aquivo " + ContasFile + " não poderemos continuar!")
		log.Println(errs)
		os.Exit(0)
	}

	// Caso tenhamos que ordenar o arquivo de
	// transacoes, este metodo faz com que
	// o arquivo seja salvo em banco
	// as transacoes serao gravadas
	// para o mesmo id conta, usamos um separador
	// para concatenar os valores
	// errs = LerFileTransactionSaveDb(TransFile)

	// exemplo:
	//fmt.Println("Base de dados carregada transacoes")
	// fmt.Println("Dado: ", gbolt.Get("trans_10", BDTrans))
	// Vetor := strings.Split(gbolt.Get("trans_10", BDTrans), ";")
	// for _, val := range Vetor {
	// 	fmt.Println(val)
	// }

	// if errs != nil {
	// 	fmt.Println("Error ao ler aquivo " + TransFile + " não poderemos continuar!")
	// 	log.Println(errs)
	// 	os.Exit(0)
	// }

	// os.Exit(0)

	// fmt.Println("")
	// fmt.Println("Contas Salvas no Bolt")
	// fmt.Println("Iniciando leitura do arquivo de transações")
	// fmt.Println("")

	// lendo o arquivo de transacoes.csv
	// agora vamos ler as transacoes e fazer os calculos para apresentar na tela
	// para ler as transacoes iremos ler linha a linha para não estourar nossa
	// memória com milhares de registros
	// Vamos quebrar em blocos as transacoes manter em um vetor em memória
	// somente o grupo de conta que estiver lendo no momento do inicio a fim
	// e depois vamos varrer este vetor da conta para efetuar os calculos
	// terminando, limpamos o vetor e inicamos tudo novamente com outra conta
	// ou seja manteremos em memória somente uma conta por vez para que possamos
	// fazer os calculos de saldo corretamente e de forma segura
	// para sabermos o saldo inicial iremos buscar no banco
	errs = CalcularSaldoTransacoes(TransFile)

	if errs != nil {

		fmt.Println("Error ao ler aquivo " + TransFile + ", isto impossibilita de fazer os calculos de saldo!")
		log.Println(errs)
		os.Exit(0)
	}
}

// Esta funcao ira ler o arquivo linha a linha
// não irá jogar tudo na memória, está otimizado
// para salvar linha a linha no banco boltdb
// o banco foi criado para otimizar a leitura
// que iremos precisar fazer quando formos
// ler as transacoes do cliente
func LerFileSaveDb(ContasFile string) error {

	// variaveis declaradas para evitar declaracoes dentro do loop
	var linha, idConta, SaldoInicialFloatString, SaldoInicialString string

	// Abre o arquivo
	arquivo, err := os.Open(ContasFile)

	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return err
	}

	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	scanner := bufio.NewScanner(arquivo)

	//fmt.Println(YellowCor("... Lendo " + ContasFile + " e salvando no banco! ..."))
	// RuneTime()

	// varrendo o arquivo
	for scanner.Scan() {

		// get linha
		linha = scanner.Text()

		// tem que possuir
		// conteudo na linha
		if linha != "" {

			// a linha tem que possuir somente uma virgula
			quantVirg := strings.Count(linha, ",")

			// a linha deve
			// possuir somente
			// uma virgula
			if quantVirg == 1 {

				// retornando dados da linha em string
				idConta, SaldoInicialFloatString, SaldoInicialString = IdContaSaldoString(linha)

				// se nao tiver valores gera error log
				if idConta != "" && SaldoInicialString != "" {

					// salvar somente se for int e nao conter caracteres indesejaveis
					// conta e saldo tem que possuir somente numeros

					re, _ := regexp.Compile(`[^0-9]`)

					// verdadeiro significa que
					// a string possui caracteres
					// falso significa que possui
					// somente numeros
					if !re.MatchString(idConta) && !re.MatchString(SaldoInicialString) {

						// salvar nova banco idConta => Saldo
						gbolt.Save(idConta, SaldoInicialFloatString)

					} else {

						// gerar log de erro
						WriteLog("O Arquivo " + ContasFile + ", foi encontrado o Idconta ou Saldo errados => idConta: " + idConta + " Saldo: " + SaldoInicialString)
					}
				} else {

					//gera log
					WriteLog("O Arquivo " + ContasFile + ", foi encontrado Idconta ou Saldo vazios => idConta: " + idConta + " Saldo: " + SaldoInicialString)
				}
			} else {

				WriteLog("O Arquivo " + ContasFile + " contém varias virgulas, isto não é permitido!")
			}

		} else {

			// gerar log
			WriteLog("O Arquivo " + ContasFile + " contém linha vazia!")
		}
	}

	// encontrando
	//fmt.Println("Contas Salvas com sucesso!")

	return scanner.Err()
}

// fazendo o calculo com os dados da transacao
// a funcao ira percorrer o arquivo de transacoes
// pesquisar se tem saldo inicial
// e apartir dai iniciar os calculos
func CalcularSaldoTransacoes(TransFile string) error {

	// variaveis declaradas
	// para evitar declaracoes
	// dentro do loop
	var linha, idConta, ValorTransacaoStr, idContaTemp, SaldoIString, ValorTransacaoNotFloat string

	var SaldoInicialFloat, SaldoFloatTotal, ValorTransacaoFloat float64

	var VetorTransacao []float64

	var j int

	// set inicio
	idContaTemp = ""
	j = 0
	SaldoFloatTotal = 0

	// Abre o arquivo
	arquivo, err := os.Open(TransFile)

	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return err
	}

	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	scanner := bufio.NewScanner(arquivo)

	// lendo arquivos e mostrando na tela
	// fmt.Println(YellowCor("..... Lendo " + TransFile + " e efetuando os Cálculos de Saldo! ....."))
	// fmt.Println("")

	// RuneTime()
	// varrendo o arquivo
	for scanner.Scan() {

		// get linha
		linha = scanner.Text()

		// testando linha vazia
		if linha != "" {

			// a linha tem que possuir somente uma virgula
			quantVirg := strings.Count(linha, ",")

			// a linha deve
			// possuir somente
			// uma virgula
			if quantVirg == 1 {

				// retornando IdConta e o valor da transacao
				idConta, ValorTransacaoStr, ValorTransacaoNotFloat = IdContaSaldoString(linha)

				if idConta != "" && ValorTransacaoNotFloat != "" {

					// validar o conteudo do arquivo
					re, _ := regexp.Compile(`[^0-9]`)

					// verdadeiro significa que
					// a string possui caracteres
					// falso significa que possui
					// somente numeros
					if !re.MatchString(idConta) && !re.MatchString(ValorTransacaoNotFloat) {

						// pode continuar..
						// convertendo valor da transacao em float
						ValorTransacaoFloat = StringToFloat(ValorTransacaoStr)

						//fmt.Println(idConta)

						// buscar saldo inicial da conta
						SaldoIString = gbolt.Get(idConta)

						// validando se a conta
						// tem saldo
						// caso contrario
						if SaldoIString == "" {

							// mensagem de erro caso nao encontre o id da conta para pegar o saldo inicial
							textError := "O Arquivo [" + TransFile + "] não foi encontrado o saldo da conta [" + idConta + "] não foi encontrado no banco de dados!"
							// err := errors.New(textError)

							// gerar log de erro
							WriteLog(textError)

							// nao existe
							// o Saldo Inicial da conta
							continue
						}

						// entrando
						// pela primeira
						// vez
						if j == 0 {

							// coloca no vetor
							VetorTransacao = append(VetorTransacao, ValorTransacaoFloat)

							// seta e
							// nao ira entrar
							// mais nesta condicao
							j = 1

						} else {

							// se for igual armazena
							// no vetor
							if idContaTemp == idConta {

								// preenchendo o vetor com valor da transacao
								// de um cliente especifico
								// isto só é possivel pq o arquivo está ordenado
								VetorTransacao = append(VetorTransacao, ValorTransacaoFloat)

							} else {

								// buscar saldo inicial da conta
								// SaldoIString = gbolt.Get(idContaTemp)

								// trasforma string em float do saldo
								SaldoInicialFloat = StringToFloat(SaldoIString)

								// saldo total
								SaldoFloatTotal = SaldoFloatTotal + SaldoInicialFloat

								// fazendo o calculo das transacoes
								CalcularBalanco(idContaTemp, VetorTransacao, SaldoFloatTotal)

								// limpar vetor
								// iniciar novamente
								VetorTransacao = []float64{}
								SaldoFloatTotal = 0

								// carregando o vetor com idConta do proximo cliente
								VetorTransacao = append(VetorTransacao, ValorTransacaoFloat)
							}
						}

						// pegar o idConta
						idContaTemp = idConta

					} else {

						// gerar log de erro
						WriteLog("O Arquivo [" + TransFile + "], foi encontrado o Idconta ou Transacao errados => idConta: [" + idConta + "] Transacao: " + ValorTransacaoNotFloat)
					}

				} else {

					WriteLog("O Arquivo [" + TransFile + "] não conseguimos ler o id Conta e o Saldo !")
				}

			} else { // varias virgulas ou uma

				WriteLog("O Arquivo [" + TransFile + "] contém varias virgulas isto não é permitido!")
			}

		} else { // linha vazia

			WriteLog("O Arquivo [" + TransFile + "] contém linha vazia!")
		}

	} // quando ele quebrar o laco precisará fazer o ultimo registro

	// fazendo a ultima posicao do vetor
	if len(VetorTransacao) > 0 {

		// buscar saldo inicial da conta
		SaldoIString = gbolt.Get(idContaTemp)

		if SaldoIString == "" {

			// mensagem de erro caso nao encontre o id da conta para pegar o saldo inicial
			textError := "O Arquivo [" + TransFile + "] não foi encontrado o saldo da conta [" + idContaTemp + "] não foi encontrado no banco de dados!"
			// err := errors.New(textError)

			// gerar log de erro
			WriteLog(textError)

		} else {

			// trasforma string em float do saldo
			SaldoInicialFloat = StringToFloat(SaldoIString)

			// saldo total
			SaldoFloatTotal = SaldoFloatTotal + SaldoInicialFloat

			// fazendo o calculo para escrever na tela
			CalcularBalanco(idContaTemp, VetorTransacao, SaldoFloatTotal)
		}
	}

	return scanner.Err()
}

// Esta funcao ira ler o arquivo linha a linha
// não irá jogar tudo na memória, está otimizado
// para salvar linha a linha no banco boltdb
// o banco foi criado para otimizar a leitura
// que iremos precisar fazer quando formos
// ler as transacoes do cliente
func LerFileTransactionSaveDb(TransFile string) error {

	// variaveis declaradas para evitar declaracoes dentro do loop
	var linha, idConta, SaldoInicialFloatString, SaldoInicialString string

	// Abre o arquivo
	arquivo, err := os.Open(TransFile)

	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return err
	}

	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	scanner := bufio.NewScanner(arquivo)

	// RuneTime()

	// varrendo o arquivo
	for scanner.Scan() {

		// get linha
		linha = scanner.Text()

		if linha != "" {

			// retornando dados da linha em string
			idConta, SaldoInicialFloatString, SaldoInicialString = IdContaSaldoString(linha)

			if idConta != "" && SaldoInicialString != "" {

				// salvar somente se for int e nao conter caracteres indesejaveis
				// conta e saldo tem que possuir somente numeros

				re, _ := regexp.Compile(`[^0-9]`)

				// verdadeiro significa que
				// a string possui caracteres
				// falso significa que possui
				// somente numeros
				if !re.MatchString(idConta) && !re.MatchString(SaldoInicialString) {

					// chave dos ids transaction
					keyT := "trans_" + idConta

					// perguntar se existe primeiro
					stringValores := gbolt.Get(keyT, BDTrans)

					if stringValores == "" {

						// salvar nova banco idConta => Saldo
						gbolt.Save(keyT, SaldoInicialFloatString, BDTrans)

					} else {

						// concatenando os values para o id correspondente..
						stringValores = stringValores + ";" + SaldoInicialFloatString

						// salvando no banco
						gbolt.Save(keyT, stringValores, BDTrans)
					}

				} else {

					// gerar log de erro
					WriteLog("O Arquivo " + TransFile + ", foi encontrado o Idconta ou Saldo errados => idConta: " + idConta + " Saldo: " + SaldoInicialString)
				}
			} else {

				//gera log
				WriteLog("O Arquivo " + TransFile + ", foi encontrado Idconta ou Saldo vazios => idConta: " + idConta + " Saldo: " + SaldoInicialString)
			}
		} else {

			// gerar log
			WriteLog("O Arquivo " + TransFile + " contém linha vazia!")
		}
	}

	// encontrando
	//fmt.Println("Contas Salvas com sucesso!")

	return scanner.Err()
}

// funcao responsavel por varrer o vetor com as transacoes e efetuar os calculos conforme
// a descricao do projeto
func CalcularBalanco(idContaTemp string, VetorTransacao []float64, SaldoFloatTotal float64) {

	// Encontrou outro codigo cliente
	// Percorre o vetor faz os calculos do vetor e do cliente anterior
	// Limpa o vetor e inicia seu preenchimento novamente
	// com um novo codigo cliente
	// Calculo do Saldo
	for _, Tvalor := range VetorTransacao {

		// O saldo de uma conta deve ser calculado a partir de seu saldo inicial, aplicando cada uma das
		// transações relacionadas a esta conta. Depósitos devem aumentar o saldo da conta e débitos
		// devem reduzir esse mesmo saldo, na medida do valor da transação.

		// Uma conta pode​ assumir um valor negativo e não existe limite inferior para o saldo da conta.
		// Contudo, cada transação de débito que termina deixando o saldo da conta negativo implica
		// uma multa de R$ 5,00​ a ser descontada imediatamente. Esta multa se aplica independente da
		// conta se encontrar ou não com saldo negativo antes da transação, mas não se aplica se a
		// transação for um depósito

		SaldoFloatTotal = SaldoFloatTotal + Tvalor

		// converte para duas casas decimais
		//FloatCasasDecimais(ValorSubAdd, 2)

		// saldo da conta negativa
		// multa de R$ 5,00
		// mas aplica somente em Debitos Negativos
		// transacoes de deposito nao  se aplica
		if SaldoFloatTotal < 0 && Tvalor < 0 {

			// converte para duas casas decimais
			SaldoFloatTotal = SaldoFloatTotal - 5
		}
	}

	// apresentando resultado na tela
	fmt.Println(YellowCor(idContaTemp + "," + FloatToStringClean(SaldoFloatTotal, 2)))
}
