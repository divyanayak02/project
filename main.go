package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Response struct
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

//Post structure
type ipInfo struct {
	City          string  `json:"city"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	IP            string  `json:"ip"`
	Latitude      float64 `json:"latitude"`
	Location      struct {
		CallingCode             string `json:"calling_code"`
		Capital                 string `json:"capital"`
		CountryFlag             string `json:"country_flag"`
		CountryFlagEmoji        string `json:"country_flag_emoji"`
		CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode"`
		GeonameID               int64  `json:"geoname_id"`
		IsEu                    bool   `json:"is_eu"`
		Languages               []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"languages"`
	} `json:"location"`
	Longitude  float64 `json:"longitude"`
	RegionCode string  `json:"region_code"`
	RegionName string  `json:"region_name"`
	Type       int32   `json:"type"`
	Zip        string  `json:"zip"`
}

//var posts []Post
// func getPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(posts)
// }
// func createPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var post Post
// 	_ = json.NewDecoder(r.Body).Decode(&post)

// 	posts = append(posts, post)
// 	json.NewEncoder(w).Encode(post)
// }
func getIPInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ip := params["ip"]
	var ipStruct ipInfo
	var resp Response

	url := "http://api.ipstack.com/" + ip + "?access_key=4dbe5149a238b4d52d8d7788c7e7c39f"
	//method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
	}
	res, err := client.Do(req)
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)

	}
	//for displaying in postman
	err = json.Unmarshal(body, &ipStruct)
	//fmt.Println(string(body))
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
	} else {
		resp.Success = true
		resp.Data = ipStruct
		json.NewEncoder(w).Encode(resp)

	}

	//}
}

func main() {
	router := mux.NewRouter()
	//posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	//router.HandleFunc("/getdata", getPosts).Methods("GET")
	//router.HandleFunc("/postdata", createPost).Methods("POST")
	router.HandleFunc("/getipinfo/{ip}", getIPInfo).Methods("GET")

	http.ListenAndServe(":8000", router)
}
