package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"errors"
	"strconv"
	"regexp"
	//"cep"
)

const (
	path = "./fretes/"
)

type freteQuery struct {
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
	nVlComprimento int //float

	//Altura da encomenda (incluindo embalagem), em centímetros.
	//Se o formato for envelope, informar zero (0).
	nVlAltura int //float

	//Largura da encomenda (incluindo embalagem), em centímetros
	nVlLargura int //float

	//Diâmetro da encomenda (incluindo embalagem), em centímetros.
	nVlDiametro int //float

	//Indica se a encomenda será entregue com o serviço adicional mão própria.
  //Valores possíveis: S ou N (S – Sim, N – Não)
	sCdMaoPropria string

	//Indica se a encomenda será entregue com o serviço 
	//adicional valor declarado. Neste campo deve ser 
	//apresentado o valor declarado desejado, em Reais.
	//Se não optar pelo serviço informar zero.
	nVlValorDeclarado int //float

	//Indica se a encomenda será entregue com o serviço 
	//adicional aviso de recebimento.
	//Valores possíveis: S ou N (S – Sim, N – Não)
	sCdAvisoRecebimento string
}

var fretes = make(map[string]map[string]string)

func handler(w http.ResponseWriter, r *http.Request) {
	defer showError(w)
	query := getQuery(r.URL.Query())
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Header().Set("Cache-Control", "private, max-age=0")
	w.Header().Set("Server", "RittmeFreteManager-1.0")

	fmt.Fprintf(w, "%+v\n", query)

/*	cep1 := FindRange(query.sCepOrigem)
	cep2 := FindRange(query.sCepDestino)

	var rota string
	if(cep1 < cep2) {
		rota = cep1 + "-" + cep2
	} else {
		rota = cep2 + "-" + cep1
	}

	frete := fretes[rota]

	fmt.Fprintf(w, "%+v\n", frete)*/
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

		//_, err := fmt.Scanf("%d", &i)
		//fmt.Printf("%v", fretes["80000000-90000000.json"])

    http.HandleFunc("/calculador/CalcPrecoPrazo.asmx", handler)
    http.ListenAndServe(":8080", nil)
}

//http://localhost:8080/calculador/CalcPrecoPrazo.asmx?nCdServico=41106&sCepOrigem=22221000&sCepDestino=22221000&nVlPeso=5.000&nCdFormato=1&nVlComprimento=30&nVlAltura=30&nVlLargura=30&nVlDiametro=30&sCdMaoPropria=N&nVlValorDeclarado=0&sCdAvisoRecebimento=N&StrRetorno=XML

func getQuery(url url.Values) freteQuery {

		query  := new(freteQuery)
	  params := url
	  var err error = nil

		elem, ok := params["nCdServico"]
		if ok == false { panic(queryError("nCdServico")) }
		query.nCdServico = elem[0]

		elem, ok = params["sCepOrigem"]
		if ok == false { panic(queryError("sCepOrigem")) }
		query.sCepOrigem = elem[0]

		elem, ok = params["sCepDestino"]
		if ok == false { panic(queryError("sCepDestino")) }
		query.sCepDestino = elem[0]

		elem, ok = params["nVlPeso"]
		if ok == false { panic(queryError("nVlPeso")) }
		query.nVlPeso = elem[0]

		elem, ok = params["nCdFormato"]
		if ok == false { panic(queryError("nCdFormato")) }
		query.nCdFormato, err = strconv.Atoi(elem[0])

		elem, ok = params["nVlComprimento"]
		if ok == false { panic(queryError("nVlComprimento")) }
		query.nVlComprimento, err = strconv.Atoi(elem[0])

		elem, ok = params["nVlAltura"]
		if ok == false { panic(queryError("nVlAltura")) }
		query.nVlAltura, err = strconv.Atoi(elem[0])

		elem, ok = params["nVlLargura"]
		if ok == false { panic(queryError("nVlLargura")) }
		query.nVlLargura, err = strconv.Atoi(elem[0])

		elem, ok = params["nVlDiametro"]
		if ok == false { panic(queryError("nVlDiametro")) }
		query.nVlDiametro, err = strconv.Atoi(elem[0])

		elem, ok = params["sCdMaoPropria"]
		if ok == false { panic(queryError("sCdMaoPropria")) }
		query.sCdMaoPropria = elem[0]

		elem, ok = params["nVlValorDeclarado"]
		if ok == false { panic(queryError("nVlValorDeclarado")) }
		query.nVlValorDeclarado, err = strconv.Atoi(elem[0])

		elem, ok = params["sCdAvisoRecebimento"]
		if ok == false { panic(queryError("sCdAvisoRecebimento")) }
		query.sCdAvisoRecebimento = elem[0]

		if(err != nil) { panic(err)}

		return *query
}

func queryError(param string) error {
	return errors.New("Missing parameter : " + param)
}

func showError(w http.ResponseWriter) {
	if r := recover(); r != nil {
		error := fmt.Sprintf("%s", r)
	  http.Error(w, error, http.StatusInternalServerError)
	}
	//http.Redirect(w, r, "/edit/"+title, http.StatusFound)
}

func (q freteQuery) validate() {

	servico = []string {"40010", "40045", "40215", "40290", "41106"}

	matched, _ := regexp.MatchString("[0-9]{5}", q.nCdServico)
	if matched == false {panic("nCdServico invalido")}
	matched, _ = regexp.MatchString("[0-9]{8}", q.sCepOrigem)
	if matched == false {panic("sCepOrigem invalido")}
	matched, _ = regexp.MatchString("[0-9]{8}", q.sCepDestino)
	if matched == false {panic("sCepDestino invalido")}
	matched, _ = regexp.MatchString("[0-9]{1,2}[\\.\\,]?[0-9]{0,5}", q.nVlPeso)
	if matched == false {panic("nVlPeso invalido")}
	matched, _ = regexp.MatchString("[1-3]", q.nCdFormato)
	if matched == false {panic("nCdFormato invalido")}
	matched, _ = regexp.MatchString("[0-9]{2,3}", q.nVlComprimento)
	if matched == false {panic("nVlComprimento invalido")}
	matched, _ = regexp.MatchString("[0-9]{1,3}", q.nVlAltura)
	if matched == false {panic("nVlAltura invalido")}
	matched, _ = regexp.MatchString("[0-9]{2,3}", q.nVlLargura)
	if matched == false {panic("nVlLargura invalido")}
	matched, _ = regexp.MatchString("[0-9]{1,2}", q.nVlDiametro)
	if matched == false {panic("nVlDiametro invalido")}
	matched, _ = regexp.MatchString("[SN]", q.sCdMaoPropria)
	if matched == false {panic("sCdMaoPropria invalido")}
	matched, _ = regexp.MatchString("[0-9]{1,5}", q.nVlValorDeclarado)
	if matched == false {panic("nVlValorDeclarado invalido")}
	matched, _ = regexp.MatchString("[SN]", q.sCdAvisoRecebimento)
	if matched == false {panic("sCdAvisoRecebimento invalido")}

	// 0 < peso < 30

	//valor declarado ??<5000??
	
	switch q.nCdFormato {
	   case "1":
	      // C- 16/105 L- 11/105 A- 2/105
	   		//C+L+A < 200
	   case "2":
	      // C- 16/60 L- 11/60
	   case "3":
	   		//C- 18/105 D- 5/91
	      //C+L+A < 200
	   }
}