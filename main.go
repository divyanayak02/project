package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

type Response struct {
	Success bool        `json:"Success"`
	Data    interface{} `json:"Data"`
	Error   string      `json:"Error"`
}
type IPAddress struct {
	Data struct {
		Attributes struct {
			AsOwner             string `json:"as_owner"`
			Asn                 int64  `json:"asn"`
			Continent           string `json:"continent"`
			Country             string `json:"country"`
			LastAnalysisResults interface {
			} `json:"last_analysis_results"`
			LastAnalysisStats struct {
				Harmless   int64 `json:"harmless"`
				Malicious  int64 `json:"malicious"`
				Suspicious int64 `json:"suspicious"`
				Timeout    int64 `json:"timeout"`
				Undetected int64 `json:"undetected"`
			} `json:"last_analysis_stats"`
			LastModificationDate     int64         `json:"last_modification_date"`
			Network                  string        `json:"network"`
			RegionalInternetRegistry string        `json:"regional_internet_registry"`
			Reputation               int64         `json:"reputation"`
			Tags                     []interface{} `json:"tags"`
			TotalVotes               struct {
				Harmless  int64 `json:"harmless"`
				Malicious int64 `json:"malicious"`
			} `json:"total_votes"`
			Whois     string `json:"whois"`
			WhoisDate int64  `json:"whois_date"`
		} `json:"attributes"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Type string `json:"type"`
	} `json:"data"`
}

func validIP4(ipAd string) bool {
	ipAd = strings.Trim(ipAd, " ")
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipAd) {
		return true
	}
	return false
}

func getIPAdress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var respo Response
	ip := params["ip"]
	err := validIP4(ip)
	if err == false {
		respo.Success = false
		respo.Error = "IP Should not match"
		json.NewEncoder(w).Encode(respo)
	} else {
		var ipadd IPAddress
		//var respo Response
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://www.virustotal.com/api/v3/ip_addresses/"+ip, nil)
		if err != nil {
			respo.Success = false
			respo.Error = err.Error()
			json.NewEncoder(w).Encode(respo)
		}
		req.Header.Add("x-apikey", "96bb3b282eaa92a99dc48c0bf0e9f817077845c10b29f3f1534535bca767f03a")
		res, err := client.Do(req)
		if err != nil {
			respo.Success = false
			respo.Error = err.Error()
			json.NewEncoder(w).Encode(respo)
		}
		defer r.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			respo.Success = false
			respo.Error = err.Error()
			json.NewEncoder(w).Encode(respo)
		}
		err = json.Unmarshal(body, &ipadd)
		if err != nil {
			respo.Success = false
			respo.Error = err.Error()
			json.NewEncoder(w).Encode(respo)
		} else {
			json.NewEncoder(w).Encode(ipadd)
		}
	}

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/IPaddress/{ip}", getIPAdress).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
