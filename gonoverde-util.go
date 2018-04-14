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

func PrintDefaults() {

	var help string

	help = `	
  Use: 
   gonoverde [OPTION]...
   or: gonoverde contas.csv transacoes.csv
`
	fmt.Println(CyanCor(help))
}
