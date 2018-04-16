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
	"log"
	"os"
	"time"
)

// gerando arquivo de log de erros
// todo erro que ocorrer será
// gravado neste arquivo..
func WriteLog(errostr string) {

	t := time.Now()

	errorString := t.Format("Mon Jan _2 15:04:05 2006")
	errorString = errorString + " " + errostr + "\n"

	// if o arquivo nao existe cria e faz append no conteudo adicionado
	ft, err2 := os.OpenFile(FILELOG, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Println(err2)
	}

	// gravando o log em disco
	if _, err := ft.Write([]byte(errorString)); err != nil {
		log.Println(err)
	}
}

// apagando o arquivo
// de log
func WriteLogClean() {

	RemoveFile(FILELOG)
}
