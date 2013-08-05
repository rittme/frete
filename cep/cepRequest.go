package cep

import(
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
  "regexp"
)

const (
  //ws = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx/CalcPrecoPrazo?nCdEmpresa=&sDsSenha=&nCdServico=@SERVICO@&sCepOrigem=@ORIGEM@&sCepDestino=@DESTINO@&nVlPeso=@PESO@&nCdFormato=1&nVlComprimento=30&nVlAltura=30&nVlLargura=30&nVlDiametro=30&sCdMaoPropria=N&nVlValorDeclarado=0&sCdAvisoRecebimento=N&StrRetorno=XML"
  ws = "http://200.228.16.53/calculador/CalcPrecoPrazo.asmx/CalcPrecoPrazo?nCdEmpresa=&sDsSenha=&nCdServico=@SERVICO@&sCepOrigem=@ORIGEM@&sCepDestino=@DESTINO@&nVlPeso=@PESO@&nCdFormato=1&nVlComprimento=30&nVlAltura=30&nVlLargura=30&nVlDiametro=30&sCdMaoPropria=N&nVlValorDeclarado=0&sCdAvisoRecebimento=N&StrRetorno=XML"
)

func Request(metodo string, origem string, destino string, peso string, c chan int, results map[string]string) {
  fmt.Println("getting")

  var resp *http.Response
  var err error = nil
  for {
    resp, err = http.Get(getURL(metodo, origem, destino, peso))
    if err != nil {
      fmt.Printf("Connection error: %s", err)
      return
    } else {
      break
    }
  }
  fmt.Println(err)
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)


  if err != nil {
    fmt.Printf("Error reading response: %s", err)
    return
  }

  reVal := regexp.MustCompile("<Valor>(.*)</Valor>")
  val := reVal.FindStringSubmatch(string(body))[1]

  rePrazo := regexp.MustCompile("<PrazoEntrega>(.*)</PrazoEntrega>")
  prazo := rePrazo.FindStringSubmatch(string(body))[1]

  fmt.Println("%v : %v", val, prazo)
  results["prazo"] = prazo
  _, ok := results[peso]
  if(!ok) {
    results[peso] = val
  }
  fmt.Println("done")
  //c <- 1
}

func getURL(metodo string, origem string, destino string, peso string) string {
  url := ws
  url = strings.Replace(url, "@SERVICO@", metodo, 1)
  url = strings.Replace(url, "@ORIGEM@", origem, 1)
  url = strings.Replace(url, "@DESTINO@", destino, 1)
  url = strings.Replace(url, "@PESO@", peso, 1)
  return url
}

