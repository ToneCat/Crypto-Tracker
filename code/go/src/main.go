package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/prices", func(w http.ResponseWriter, r *http.Request) { //route
		enableCors(&w) // enables cores function

		// makes call to cryptocompare API and fetches prices
		rs, err := http.Get("https://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS,PAX,BNB,TRX,ETC,ADA,LINK,PAX,XLM,XMR,BSV&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7cahttps://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,BCH,USDT,LTC,EOS&tsyms=USD&api_key=503cf7d09362826984c93be58a5f538a835b6f2ba1548234ed7087674078d7ca")
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")

		io.Copy(w, rs.Body) // sends prices to client end

	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3001", nil) // serves at localhost:3001
}

// enables CORS browser support so React can communicate
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
