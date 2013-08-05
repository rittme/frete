package main

import "fmt"
import "flag"
import "io/ioutil"
import "encoding/json"
import "github.com/rittme/frete/cep"

const (
    default_path    = "fretes/"
    default_metodo  = "41106"
    default_verbose = false

    usage_path    = "Pasta de estocagem dos arquivos"
    usage_metodo  = "Metodo utilisado da encomenda"
    usage_verbose = "Mostra os parâmetros e o resultado XML"
  )

var pesos = []string{"0.300", "1.000", "2.000", "3.000", "4.000", "5.000", "6.000", "7.000", "8.000", "9.000", "10.000", "11.000", "12.000", "13.000", "14.000", "15.000", "16.000", "17.000", "18.000", "19.000", "20.000", "21.000", "22.000", "23.000", "24.000", "25.000", "26.000", "27.000", "28.000", "29.000", "30.000"}

var f_path string
var f_metodo string
var f_verbose bool

func init(){

  flag.StringVar(&f_path, "path", default_path, usage_path)
  flag.StringVar(&f_path, "p", default_path, usage_path+" (shorthand)")

  flag.StringVar(&f_metodo, "metodo", default_metodo, usage_metodo)
  flag.StringVar(&f_metodo, "m", default_metodo, usage_metodo+" (shorthand)")

  flag.BoolVar(&f_verbose, "verbose", default_verbose, usage_verbose)
  flag.BoolVar(&f_verbose, "v", default_verbose, usage_verbose+" (shorthand)")
  flag.Parse()
}

func main() {
  ceps := cep.GetRanges()
  base := cep.GetBases()
  for index_orig, cep_orig := range base {
    for index_dest, cep_dest := range base[index_orig:] {
      channel := make(chan int)
      results := make(map[string]string)
      for _, peso := range pesos {
        if(f_verbose) {
          fmt.Println(f_metodo)
          fmt.Println(cep_orig + " : " + cep_dest)
        }
        cep.Request(f_metodo, cep_orig, cep_dest, peso, channel, results)
      }
      /*i := 0
      l := len(pesos)
      L: for {
        select {
        case <-channel:
          i++
          if(i >= l) {
            break L
          }
        }
      }*/
      if(f_verbose) {
        fmt.Printf("%v",results)
      }
      output(ceps[index_orig], ceps[index_orig+index_dest], results)
    }
  }
}

func output(origem string, destino string, results map[string]string) {
  encoded, err := json.Marshal(results)
  if err != nil { panic(err) }
  err = ioutil.WriteFile(f_path + origem + "-" + destino + ".json", encoded, 0777)
  if err != nil { panic(err) }
}