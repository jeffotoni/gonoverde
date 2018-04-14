# gonoverde

Este programa é somente um exercício que tem como objetivo calcular o balanço da conta corrente de um conjunto de clientes.

Os arquivos de entrada estão em formato CSV, com campos delimitados por vírgula, sem aspas, sem cabeçalho.

O programa gerar-contas-transacoes.go gera os dois arquivos contas.csv e transacoes.csv já no padrão para efetuar as operações de débito e depósito.

Uma transação de valor positivo é um depósito​ na conta. Uma transação de valor negativo é um débito​ na conta.

### Regras e Lógica do programa

O saldo de uma conta deve ser calculado a partir de seu saldo inicial, aplicando cada uma das transações relacionadas a esta conta. Depósitos devem aumentar o saldo da conta e débitos devem reduzir esse mesmo saldo, na medida do valor da transação. 

Uma conta pode​ assumir um valor negativo e não existe limite inferior para o saldo da conta. Contudo, cada transação de  débito que termina deixando o saldo da conta negativo implica uma multa de R$ 5,00​ a ser descontada imediatamente. 

Esta multa se aplica independente da conta se encontrar ou não com saldo  egativo antes da transação, mas não se aplica se a transação for um depósito.


### Install Dependencies

```
go get -v github.com/jeffotoni/gcolor

```

### Start App with Run or Compile

```
git clone github.com/jeffotoni/gonoverde

go run gonoverde.go

go build gonoverde.go

./gonoverde

```