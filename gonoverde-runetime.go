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
	"os"
	"os/signal"
	"time"
)

// construindo
// rune loader
// com caracteres
func RuneTime() {

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fmt.Println("\nfim!")
		fmt.Print("\033[?25h")
		os.Exit(0)
	}()

	go func() {

		fmt.Print("\033[?25l")

		timer := time.Tick(time.Duration(50) * time.Millisecond)

		//s := []rune(`|/=\|=`)
		s := []rune(`|/~\`)
		i := 0

		for {
			<-timer
			fmt.Print("\r")
			fmt.Print("\033[0;33m" + string(s[i]) + "\033[0m")
			i++
			if i == len(s) {
				i = 0
			}
		}
	}()

}
