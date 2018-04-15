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
	"log"
	"os"
	"strconv"
)

// Exists file in disck
func ExistsFile(file string) bool {

	if _, err := os.Stat(file); err != nil {

		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// apresentando na
// tela as opcoes
// quando usuario
// digitar sem os
// parametros
func PrintDefaults() {

	var help string

	help = `	
  Use: 
   gonoverde [OPTION]...
   or: gonoverde contas.csv transacoes.csv
`
	fmt.Println(CyanCor(help))
}

// fazendo um substr da string
func Substr(s string, pos, length int) string {

	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}

	return string(runes[pos:l])
}

// convertendo de Float para String
func FloatToString(input_num float64, casaDecimal int) string {

	// converte numero para string, manter duas casas decimais
	return strconv.FormatFloat(input_num, 'f', casaDecimal, 64)
}

// convertendo para float64 uma string
func StringToSaldoWithDecimal(stringSaldoBody, stringSaldoDecimal string) (Resultado float64) {

	Resultado, errs := strconv.ParseFloat(stringSaldoBody+"."+stringSaldoDecimal, 64)

	if errs != nil {

		log.Println(errs)
	}

	return
}
