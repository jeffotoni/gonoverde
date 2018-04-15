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
	"strings"
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

//convertendo float para string com virgula
func FloatToStringVirgula(input_num float64, casaDecimal int) string {

	// converte numero para string, manter duas casas decimais
	v := strconv.FormatFloat(input_num, 'f', casaDecimal, 64)
	return strings.Replace(v, ".", ",", -1)
}

// remover caracteres virgula ou ponto
func FloatToStringClean(input_num float64, casaDecimal int) string {

	// converte numero para string, manter duas casas decimais
	v := strconv.FormatFloat(input_num, 'f', casaDecimal, 64)
	return strings.Replace(v, ".", "", -1)
}

// convertendo de Float com casas decimais
func FloatCasasDecimais(input_num float64, casaDecimal int) float64 {

	// converte numero para string, manter duas casas decimais
	strFloat := strconv.FormatFloat(input_num, 'f', casaDecimal, 64)

	return StringToFloat(strFloat)
}

// convertendo para float64 uma string
func StringToSaldoWithDecimal(stringSaldoBody, stringSaldoDecimal string) (Resultado float64) {

	Resultado, errs := strconv.ParseFloat(stringSaldoBody+"."+stringSaldoDecimal, 64)

	if errs != nil {

		log.Println(errs)
	}

	return
}

// convertendo para float64 uma string
func StringToFloat(valorString string) (Resultado float64) {

	Resultado, errs := strconv.ParseFloat(valorString, 64)

	if errs != nil {

		log.Println(errs)
	}

	return
}

// convertendo string int para transformar em float e com casas decimais
// retorna uma string com o formato correto do valor que esta em string
func IdContaSaldoString(linha string) (idConta string, SaldoFloatString string) {

	// linha
	// id conta
	// e saldo
	if linha != "" {

		// vamos transformar o valor em decimal
		// sera um float para que possamos
		// fazer os calculos
		vetorConta := strings.Split(linha, ",")

		// get conta
		idConta = vetorConta[0]

		// get saldo
		stringSaldo := vetorConta[1]

		// colocar ponto nas duas ultimas posicoes
		// gerando casas decimais da string

		// somente o decimal sem as casas decimais
		stringSaldoBody := stringSaldo[:len(stringSaldo)-2]

		// somente o algarismo apos a virgula casas decimais
		stringSaldoDecimal := Substr(stringSaldo, len(stringSaldo)-2, 2)

		// gerando o saldo em float, com as casas decimais para efetuar os calculos
		Saldo := StringToSaldoWithDecimal(stringSaldoBody, stringSaldoDecimal)

		// convertendo float para string,
		// float com duas casas decimais
		SaldoFloatString = FloatToString(Saldo, 2)

		return

	} else {

		return "", ""
	}
}
