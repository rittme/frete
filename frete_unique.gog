package main

import "fmt"
import "flag"
import "strings"
import "io/ioutil"
import "regexp"
import "encoding/json"

const (
    default_cepOrigem  = "22221000"
    default_cepDestino = "22221000"
    //default_peso       = "0.300"
    default_metodo     = "41106"
    default_verbose    = false

    usage_cepOrigem   = "Cep de origem"
    usage_cepDestino  = "Cep de destino"
    //usage_peso        = "Peso da encomenda"
    usage_metodo      = "Metodo utilisado da encomenda"
    usage_verbose     = "Mostra os parâmetros e o resultado XML"

    ws = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx/CalcPrecoPrazo?nCdEmpresa=&sDsSenha=&nCdServico=@SERVICO@&sCepOrigem=@ORIGEM@&sCepDestino=@DESTINO@&nVlPeso=@PESO@&nCdFormato=1&nVlComprimento=30&nVlAltura=30&nVlLargura=30&nVlDiametro=30&sCdMaoPropria=N&nVlValorDeclarado=0&sCdAvisoRecebimento=N&StrRetorno=XML"
  )

var f_cepOrigem string
var f_cepDestino string
//var f_peso string
var f_metodo string
var f_verbose bool

var results = make(map[string]string)

func init(){

  flag.StringVar(&f_cepOrigem, "origem", default_cepOrigem, usage_cepOrigem)
  flag.StringVar(&f_cepOrigem, "o", default_cepOrigem, usage_cepOrigem+" (shorthand)")

  flag.StringVar(&f_cepDestino, "destino", default_cepDestino, usage_cepDestino)
  flag.StringVar(&f_cepDestino, "d", default_cepDestino, usage_cepDestino+" (shorthand)")

  /*flag.StringVar(&f_peso, "peso", default_peso, usage_peso)
  flag.StringVar(&f_peso, "p", default_peso, usage_peso+" (shorthand)")*/

  flag.StringVar(&f_metodo, "metodo", default_metodo, usage_metodo)
  flag.StringVar(&f_metodo, "m", default_metodo, usage_metodo+" (shorthand)")

  flag.BoolVar(&f_verbose, "verbose", default_verbose, usage_verbose)
  flag.BoolVar(&f_verbose, "v", default_verbose, usage_verbose+" (shorthand)")
  flag.Parse()
}

func main() {
  pesos := []string{"0.300", "1.000", "2.000", "3.000", "4.000", "5.000", "6.000", "7.000", "8.000", "9.000", "10.000", "11.000", "12.000", "13.000", "14.000", "15.000", "16.000", "17.000", "18.000", "19.000", "20.000", "21.000", "22.000", "23.000", "24.000", "25.000", "26.000", "27.000", "28.000", "29.000", "30.000"}
  c := make(chan int)  

  for _, v := range pesos {
    go request(f_metodo, f_cepOrigem, f_cepDestino, v, c)
  }
  i := 0
  l := len(pesos)
  L: for {
    select {
    case <-c:
      i++
      if(i >= l) {
        break L
      }
    }
  }
  
  output()
}

func request(metodo string, origem string, destino string, peso string, c chan int) {
if(f_verbose) {
    fmt.Println(origem)
    fmt.Println(destino)
    //fmt.Println(peso)
    fmt.Println(metodo)
    fmt.Println()
  }

  resp, err := http.Get(getURL(metodo, origem, destino, peso))
  if err != nil {
    fmt.Printf("Connection error: %s", err)
    return
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Printf("Error reading response: %s", err)
    return
  }
  if(f_verbose) {
    fmt.Printf("%s",body)
    fmt.Println()
  }

  reVal := regexp.MustCompile("<Valor>(.*)</Valor>")
  val := reVal.FindStringSubmatch(string(body))[1]
  
  rePrazo := regexp.MustCompile("<PrazoEntrega>(.*)</PrazoEntrega>")
  prazo := rePrazo.FindStringSubmatch(string(body))[1]

  results["prazo"] = prazo
  results[peso] = val
  c <- 1
}

func getURL(metodo string, origem string, destino string, peso string) string {
  url := ws
  url = strings.Replace(url, "@SERVICO@", metodo, 1)
  url = strings.Replace(url, "@ORIGEM@", origem, 1)
  url = strings.Replace(url, "@DESTINO@", destino, 1)
  url = strings.Replace(url, "@PESO@", peso, 1)
  return url
}

func output() {
  fmt.Printf("%v",results)
  encoded, err := json.Marshal(results)
  if err != nil { panic(err) }
  err = ioutil.WriteFile(f_cepOrigem + "-" + f_cepDestino + ".json", encoded, 0777)
  if err != nil { panic(err) }
}