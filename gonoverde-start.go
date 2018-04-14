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

		if strings.ToLower(FCo) == FCONTAS && strings.ToLower(FTr) == FTRANSACOES {

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

			fmt.Println(YellowCor("Pode Executar"))
			DebClient()

		} else {

			// o arquivo tem que ser contas.csv e transacoes.csv
			fmt.Println(YellowCor("O nome dos arquivos tem que ser " + FCONTAS + " e " + FTRANSACOES))
		}

		//fmt.Println(argsWithoutProg[0])
		//fmt.Println(argsWithoutProg[1])

	} else {

		// error
		// apresentar
		// opcoes na tela
		PrintDefaults()
	}
}

// Exists file in disck
func ExistsFile(file string) bool {

	if _, err := os.Stat(file); err != nil {

		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func PrintDefaults() {

	var help string

	help = `	
  Use: 
   gonoverde [OPTION]...
   or: gonoverde contas.csv transacoes.csv
`
	fmt.Println(CyanCor(help))
}
