package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

const (
	path = "./fretes/"
)

type freteRequest struct {
	//Código Serviço
	//40010 SEDEX Varejo
	//40045 SEDEX a Cobrar Varejo
	//40215 SEDEX 10 Varejo
	//40290 SEDEX Hoje Varejo
	//41106 PAC Varejo
	nCdServico string 

	//CEP de Origem sem hífen.Exemplo: 05311900
	sCepOrigem string

	//CEP de Destino sem hífen 
	sCepDestino string

	//Peso da encomenda, incluindo sua embalagem.
  //O peso deve ser informado em quilogramas. 
  //Se o formato for Envelope, o valor máximo permitido será 1 kg
	nVlPeso string

	//Formato da encomenda (incluindo embalagem).
	//Valores possíveis: 1, 2 ou 3
	//1 – Formato caixa/pacote
	//2 – Formato rolo/prisma
	//3 - Envelope
	nCdFormato int

	//Comprimento da encomenda (incluindo embalagem), em centímetros.
	nVlComprimento int

	//Altura da encomenda (incluindo embalagem), em centímetros.
	//Se o formato for envelope, informar zero (0).
	nVlAltura int

	//Largura da encomenda (incluindo embalagem), em centímetros
	nVlLargura int 

	//Diâmetro da encomenda (incluindo embalagem), em centímetros.
	nVlDiametro int

	//Indica se a encomenda será entregue com o serviço adicional mão própria.
  //Valores possíveis: S ou N (S – Sim, N – Não)
	sCdMaoPropria string

	//Indica se a encomenda será entregue com o serviço 
	//adicional valor declarado. Neste campo deve ser 
	//apresentado o valor declarado desejado, em Reais.
	//Se não optar pelo serviço informar zero.
	nVlValorDeclarado int

	//Indica se a encomenda será entregue com o serviço 
	//adicional aviso de recebimento.
	//Valores possíveis: S ou N (S – Sim, N – Não)
	sCdAvisoRecebimento string
}

var fretes = make(map[string]map[string]string)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", r.URL.Query())
}

func main() {
		files, err := ioutil.ReadDir(path)
		if err != nil { panic(err) }

		for _, file := range files {
			fmt.Println("Reading ", file.Name())
			data, err := ioutil.ReadFile(path+file.Name())
			if err != nil { panic(err) }

			fmt.Printf(".")
			var results map[string]string
			fmt.Printf(".")
			err = json.Unmarshal(data, &results)
			fmt.Printf(".")
			if err != nil { panic(err) }
			fretes[file.Name()] = results
			fmt.Println("done")
		}

		//fmt.Printf("%v", fretes["80000000-90000000.json"])

    http.HandleFunc("/calculador/CalcPrecoPrazo.asmx", handler)
    http.ListenAndServe(":8080", nil)
}