# gonoverde

Este programa é somente um exercício que tem como objetivo calcular o balanço da conta corrente de um conjunto de clientes.

Os arquivos de entrada estão em formato CSV, com campos delimitados por vírgula, sem aspas, sem cabeçalho.

O programa gerar-contas-transacoes.go gera os dois arquivos contas.csv e transacoes.csv já no padrão para efetuar as operações de débito e depósito.

Uma transação de valor positivo é um depósito​ na conta. Uma transação de valor negativo é um débito​ na conta.

### Regras e Lógica do programa

O saldo de uma conta deve ser calculado a partir de seu saldo inicial, aplicando cada uma das transações relacionadas a esta conta. Depósitos devem aumentar o saldo da conta e débitos devem reduzir esse mesmo saldo, na medida do valor da transação. 

Uma conta pode​ assumir um valor negativo e não existe limite inferior para o saldo da conta. Contudo, cada transação de  débito que termina deixando o saldo da conta negativo implica uma multa de R$ 5,00​ a ser descontada imediatamente. 

Esta multa se aplica independente da conta se encontrar ou não com saldo  egativo antes da transação, mas não se aplica se a transação for um depósito.

### Estrutura do Programa

```go
- gonoverde
	- gbolt
		- gonoverde-gbolt.go
	- src
		- gonoverde
			- main-gerar.go
			- main.go
			- contas.csv
			- transacoes.csv
	
	- gonoverde-write-log.go
	- gonoverde-runetime.go
	- gonoverde-eviroment.go
	- gonoverde-start.go
	- gonoverde-util.go
	- gonoverde-saldocontacliente.go

```

### Instalar Dependencias para o projeto

```
$ go get -v github.com/jeffotoni/gcolor

$ go get -u github.com/boltdb/bolt

```

### Baixando o projeto no GitHub

```
$ git clone github.com/jeffotoni/gonoverde

```

### Gerando Arquivos contas.csv e transacoes.csv

Caso tenha os arquivos, não precisará executar esta etapa, basta copia-los para src/gonoverde

Caso necessite gera-los, o programa gera os dois arquivos, onde contas terá os saldos iniciais das contas e não irá repetir no arquivo de contas, o arquivo de transações irá possuir varias transações de cada conta, os arquivos estão ordenados pelo id da conta.

O arquivo de transacoes.csv está ordenado pelo IdConta.

```
// entrando no src
$ cd src/gonoverde

// executar e gerar 
// arquivos .csv
$ go run main-gerar.go

// caso queira compilar
$ go build main-gerar.go

// gerando os arquivos
// contas.csv e transacoes.csv 
$ ./main-gerar

```

### Calcular Balanço de conta Corrente

```
// uma forma de executar
$ go run main.go contas.csv transacoes.csv

// pode compilar
$ go build main.go

$ ./main contas.csv transacoes.csv

// ou pode deixar no ambiente
$ sudo cp main /usr/bin/gonoverde

// lembrando que os arquivos tem que 
// ser passados os paths ou os seus
// respectivos nomes mas neste caso
// os arquivos devem estar onde irá
// executar

// os arquivos tem que esta no path
// que está executando
$ gonoverde contas.csv transacoes.csv

// passando path dos files não precisa
// que os arquivos estejam no mesmo
// path
$ gonoverde /home/user/files/contas.csv /home/user/files/transacoes.csv

```