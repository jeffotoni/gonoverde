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

package main

import (
	"fmt"
	. "github.com/jeffotoni/gonoverde"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// gerar arquivo contas
// para simular calculos
// o arquivo tem o fomato
// id conta int
// saldo inicial da conta com centavos int
// Ex: 345,14428
// Gerar Arquivo transacao
// Enquanto nosso programa gera uma
// linha para contas
// ele gera diversas linhas desta conta para transacao

func main() {

	fmt.Println("... gerando arquivos ...")
	RuneTime()

	//go func() {

	// removendo file
	// iniciando novamente
	RemoveFile(FCONTAS)

	// removendo file
	// iniciando novamente
	RemoveFile(FTRANSACOES)

	// if o arquivo nao existe cria e faz append no conteudo adicionado
	f, err := os.OpenFile(FCONTAS, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	// if o arquivo nao existe cria e faz append no conteudo adicionado
	ft, err2 := os.OpenFile(FTRANSACOES, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Println(err2)
	}

	// id conta
	var idconta, j, quantTransacao, countConta, countTrans int

	// set
	// zero
	// J
	j = 1

	countConta = 0
	countTrans = 0

	// set a quantidade de
	// transacoes que pode
	// ter um ID
	quantTransacao = Random(2, 100)

	// id conta
	// inicia em 10
	idconta = 10

	// saldo inicial e transacao
	var saldoi, vtransacao float64

	// string de valores idconta,saldoi
	var idcontaNow, stringsaldoi, stringtrasacao, stringValores string

	// padrao para remocao de ponto dos valores
	var replacer = strings.NewReplacer(".", "")

	// carregando o arquivo
	// com idconta e saldo inicial
	for i := 1; i <= LINHAS; i++ {

		//randValor := Random(10, 10000)
		randSaldo := Random(1, 100000)

		// transacao positiva
		randSaldoTr := Random(1, 1000)

		// transacao negativa
		randSaldoTrN := Random(10, 800)

		// saldo sendo gerado aleatorio
		saldoi = RandomF() * float64(randSaldo)

		// positivo
		if randSaldoTr%2 == 0 {

			// transacao positiva
			vtransacao = RandomF() * float64(randSaldoTr)

		} else {

			// transacao negativo
			vtransacao = (RandomF() * float64(randSaldoTrN)) * -1.0
		}

		// convert float para string
		stringtrasacao = strconv.FormatFloat(vtransacao, 'f', 2, 64)

		// removendo . (ponto) da string
		stringtrasacao = replacer.Replace(stringtrasacao)

		// convert float para string
		stringsaldoi = strconv.FormatFloat(saldoi, 'f', 2, 64)

		// removendo . (ponto) da string
		stringsaldoi = replacer.Replace(stringsaldoi)

		// convertendo inteiro para string
		idcontaNow = strconv.Itoa(idconta)

		// na primeira
		// vez
		if j == 1 {

			//concatenando idconta com saldo inicial e formatando
			//para ser gravado no arquivo conforme o modelo
			stringValores = idcontaNow + "," + stringsaldoi + "\n"

			// gera uma conta e transacao somente uma vez
			// salva contas
			if _, err := f.Write([]byte(stringValores)); err != nil {

				log.Println(err)
			}

			//concatenando idconta com o valor da transacao
			//para ser gravado no arquivo conforme o modelo
			stringValores = idcontaNow + "," + stringtrasacao + "\n"

			// salva transacao
			if _, err := ft.Write([]byte(stringValores)); err != nil {

				log.Println(err)
			}

			countConta++
			countTrans++

		} else {

			// gera uma transacao
			// enquanto j nao
			// for igual quantTransacao
			// ele recebe o mesmo
			// Id, isto servira
			// para manter o mesmo
			// Id da conta
			if j == quantTransacao {

				//concatenando idconta com o valor da transacao
				//para ser gravado no arquivo conforme o modelo
				stringValores = idcontaNow + "," + stringtrasacao + "\n"

				// salva transacao
				if _, err := ft.Write([]byte(stringValores)); err != nil {

					log.Println(err)
				}

				// gera uma transacao
				quantTransacao = Random(2, 100)
				idconta++
				j = 0

				countTrans++

			} else {

				// gera transacao
				//concatenando idconta com o valor da transacao
				//para ser gravado no arquivo conforme o modelo
				stringValores = idcontaNow + "," + stringtrasacao + "\n"

				// salva transacao
				if _, err := ft.Write([]byte(stringValores)); err != nil {

					log.Println(err)
				}

				countTrans++
			}
		}

		j++
	}

	// fechar arquivo
	if err := f.Close(); err != nil {
		log.Println(err)
	}

	// show screen
	fmt.Println("\nArquivo "+FCONTAS+" gerado com sucesso, total de [", countConta, "] linhas")
	fmt.Println("Arquivo "+FTRANSACOES+" gerado com sucesso, total de [", countTrans, "] linhas")
	fmt.Println("Foi gerado " + strconv.Itoa(LINHAS) + " linhas.")

	//}()

	//fim
}

// Randon number
func Random(min, max int) int {

	//time.Sleep(time.Millisecond * 30)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Randon number
func RandomF() float64 {

	//time.Sleep(time.Millisecond * 50)
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
