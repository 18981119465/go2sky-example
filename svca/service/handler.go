package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	go2sky_http "github.com/tetratelabs/go2sky/plugins/http"

	"github.com/18981119465/go2sky-example/model"
	"github.com/18981119465/go2sky-example/svca/data"
)

// DBClient global variable.
var DBClient data.IData

// GetAccount is handler to get account
func GetAccount(w http.ResponseWriter, r *http.Request) {

	// Read the 'accountID' path parameter from the mux map
	var accountID = mux.Vars(r)["accountID"]

	log.Println("Get account from boltdb")
	time.Sleep(1 * time.Second)
	// Read the account struct BoltDB
	account, err := DBClient.QueryAccount(r.Context(), accountID)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	account.ServedBy = getIP()

	log.Println("Get quote from svcb")
	time.Sleep(1 * time.Second)

	quote, err := getQuote(r.Context(), accountID)
	if err == nil {
		account.Quote = quote
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback then display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address (non loopback). Exiting.")
}

func getQuote(ctx context.Context, accountID string) (model.Quote, error) {
	addr := fmt.Sprintf("http://"+os.Getenv("SVCB_HOST")+":"+os.Getenv("SVCB_PORT")+"/quotes/%s", accountID)

	req, _ := http.NewRequest("GET", addr, nil)
	req = req.WithContext(ctx)

	client, err := go2sky_http.NewClient(Tracer)
	if err != nil {
		log.Fatalf("create client error %v \n", err)
	}

	log.Printf("Request sent to %s", addr)

	resp, err := client.Do(req)

	if err == nil && resp.StatusCode == 200 {
		quote := model.Quote{}
		bytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bytes, &quote)
		return quote, nil
	}
	return model.Quote{}, fmt.Errorf("Some error")

}
