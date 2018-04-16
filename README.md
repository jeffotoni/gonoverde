# gonoverde

Este programa é somente um exercício que tem como objetivo calcular o balanço da conta corrente de um conjunto de clientes.

Os arquivos de entrada estão em formato CSV, com campos delimitados por vírgula, sem aspas, sem cabeçalho.

O programa **gonoverde** gera os dois arquivos contas.csv e transacoes.csv já no padrão para efetuar as operações de débito e depósito.

Uma transação de valor positivo é um depósito​ na conta. Uma transação de valor negativo é um débito​ na conta.

# Exemplo do funcionamento do programa

![image](https://github.com/jeffotoni/gonoverde/blob/master/gifanimation/gonoverde.gif)


### Regras e Lógica do programa

O saldo de uma conta deve ser calculado a partir de seu saldo inicial, aplicando cada uma das transações relacionadas a esta conta. Depósitos devem aumentar o saldo da conta e débitos devem reduzir esse mesmo saldo, na medida do valor da transação. 

Uma conta pode​ assumir um valor negativo e não existe limite inferior para o saldo da conta. Contudo, cada transação de  débito que termina deixando o saldo da conta negativo implica uma **multa de R$ 5,00**​ a ser descontada imediatamente. 

Esta multa se aplica independente da conta se encontrar ou não com saldo  egativo antes da transação, mas não se aplica se a transação for um depósito.

### Estrutura do Programa

O **gonoverde** que é nosso **main.go** ele recebe sempre 2 argumentos, contas.csv e transacoes.csv ou você possui os arquivos ou pode executar o **main-gerar.go** para gerar os arquivos par você.

O arquivo **transacoes.csv** tem que está **ordenado** pelo id da conta do cliente.

O programa está gravando no banco todas as contas.csv onde possui os saldos iniciais de todos os clientes, logo após ele irá ler e fazer o balanço conforme as regras acima.
Toda fez que executar o programa para gerar o balanço ele irá apagar toda base de dados gerada e limpar os logs.

Todo erro encontrado é gerado um arquivo de log o nome dele é **logsys.log**.

Foi usado um noSql para não estourar memória em casos de arquivos muito grandes e com noSql não há necessidade de colocarmos os dados do cliente na memória.

Quando fazemos a leitura dos arquivos não estamos jogando tudo na memória, estamos lendo e percorrendo o arquivo linha a linha, otimizando nosso processo de leitura e cálculos.

Os cálculos quando estamos percorrendo o arquivo de transações é armazenado em memória um bloco contendo as transações de um único cliente, e o mesmo vetor é limpo a cada novo cliente.

```go
- gonoverde
 - gbolt
   - gonoverde-gbolt.go    (biblioteca para abstrar alguns comandos do bolt noSql)
 - src
   - gonoverde
       - dockerfile        (poderá executar o programa em um container)
       - contas.csv        (arquivo contas gerado a partir do main-gerar.go)
       - transacoes.csv    (arquivo transacoes gerado a partir do main-gerar.go)
       - logsys.log        (este arquivo é gerado se existir algum erro)
       - main-gerar.go     (programa responsavel por gerar uma base para simulacao) 
       - main.go           (programa principal responsavel por fazer nosso balanço)
	
   - gonoverde-write-log.go    (biblioteca para gerar log) 
   - gonoverde-runetime.go     (responsavel por fazer um loader ascii)  
   - gonoverde-eviroment.go    (algumas variaveis de ambiente)
   - gonoverde-start.go        (nesta lib irá tratar as entradas cmd)
   - gonoverde-util.go         (lib responsável por conter algumas funções de conversão)
   - gonoverde-balanco.go      (lib responsável por gerar o balanco das contas) 

```

### Executando gonoverde com Docker

Caso desejar poderá baixar a imagem que é de 5M somente e rodar o programa gonoverde
Você também poderá editar arquivo **dockerfile** gerar seu executável e construir sua própria imagem se assim desejar.

Para baixar a imagem basta da pull ou run

```
// baixando a imagem, o comando abaixo 
// caso preferir também baixa a imagem
$ sudo docker pull jeffotoni/gonoverde

// coloquei localmente os arquivos contas.csv e transacoes.csv no /tmp/gonoverde
// no container irá para /tmp
$ sudo docker run -itd --rm --name gonoverde -v /tmp/gonoverde:/tmp jeffotoni/gonoverde:latest

// o comando abaixo ele irá executar o programa e gerar o balanço
$ sudo docker exec <idcontainer> gonoverde /tmp/contas.csv /tmp/transacoes.csv

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

Caso tenha os arquivos, não precisará executar esta etapa, basta copia-los para **src/gonoverde**

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

### Executar e calcular balanço de conta Corrente

```
// entrando no diretorio
// para executar ou compilar
$ cd src/gonoverde

// uma forma de executar com run
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
