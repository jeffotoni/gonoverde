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
	"github.com/jeffotoni/gonoverde/gbolt"
	"os"
)

func SaldoContaCliente(ContasFile, TransFile string) {

	// Testing boltdb database
	// Start ping database
	// Creating ping ok
	gbolt.Save("Ping", "ok")

	// Testing whether it was recorded
	// and read on the boltdb, we
	// recorded a Ping and then
	// read it back.
	if gbolt.Get("Ping") != "ok" {

		fmt.Println("Services Error Data Base!")
		os.Exit(1)
	}

	fmt.Println(PurpleCor("Debita e deposita na conta do cliente"))
}
