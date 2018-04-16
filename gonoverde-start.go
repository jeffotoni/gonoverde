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
	"fmt"
	. "github.com/jeffotoni/gcolor"
	"os"
	"strings"
)

func Start() {

	// pode executar
	var canExec int

	var DirPath []string

	var nameFileC, nameFileT string

	// desativando
	// execucao
	canExec = 0

	// START
	if len(os.Args) == 3 {

		argsWithoutProg := os.Args[1:]

		// nome de arquivo contas
		FCo := argsWithoutProg[0]

		// nome de arquivo transacoes
		FTr := argsWithoutProg[1]

		// se passar o path absoluto aceitar
		// tratar quando for path absoluto
		DirPath = strings.Split(FCo, "/")

		// capturar nome files
		if len(DirPath) > 1 {

			nC := DirPath[len(DirPath)-1 : len(DirPath)]
			// capturar nome files
			nameFileC = fmt.Sprintf("%s", nC[0])

		} else {

			// somente nome
			// file contas
			nameFileC = FCo
		}

		// zerando
		DirPath = []string{}

		// quebrando caso tenha
		// path absolulto
		DirPath = strings.Split(FTr, "/")

		// capturar nome files
		if len(DirPath) > 1 {

			// capturar nome files
			nC := DirPath[len(DirPath)-1 : len(DirPath)]
			// capturar nome files
			nameFileT = fmt.Sprintf("%s", nC[0])

		} else {

			// somente nome
			// file transacao
			nameFileT = FTr
		}

		// comparando nome de arquivos
		if strings.ToLower(nameFileC) == FCONTAS && strings.ToLower(nameFileT) == FTRANSACOES {

			// validando se o arquivo existe
			if ExistsFile(FCo) {

				// fmt.Println("Arquivo existe " + FCONTAS)
				canExec++

			} else {

				fmt.Println(CyanCor("Arquivo " + FCONTAS + " não existe!"))
				return
			}

			if ExistsFile(FTr) {

				// fmt.Println("Arquivo existe " + FTRANSACOES)
				canExec++

			} else {

				fmt.Println(CyanCor("Arquivo " + FTRANSACOES + " não existe!"))
				return
			}

			// fmt.Println(YellowCor("Pode Executar"))
			// iniciar o Calculo
			// do saldo de cada
			// conta
			SaldoContaCliente(FCo, FTr)

		} else {

			// o arquivo tem que ser contas.csv e transacoes.csv
			fmt.Println(YellowCor("O nome dos arquivos tem que ser " + FCONTAS + " e " + FTRANSACOES))
		}

	} else {

		// error
		// apresentar
		// opcoes na tela
		PrintDefaults()
	}
}
