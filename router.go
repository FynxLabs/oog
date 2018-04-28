package main

// import (
// 	"log"
// 	"net/http"
// )

// // fun router()
// func router() {
// 	router := mux.NewRouter()
// 	// chat routes
// 	router.HandleFunc("/v1/attachment", chat_attachment).Methods("POST")
// 	router.HandleFunc("/v1/emote", chat_emote).Methods("POST")
// 	router.HandleFunc("/v1/plain", chat_plain).Methods("POST")
// 	// brain routes
// 	router.HandleFunc("/v1/save", brain_save).Methods("POST")
// 	router.HandleFunc("/v1/delete", brain_delete).Methods("POST")
// 	router.HandleFunc("/v1/query", brain_query).Methods("POST")
// 	// plugin routes
// 	router.HandleFunc("/v1/reload", plugin_reload).Methods("POST")
// 	log.Fatal(http.ListenAndServe(":8000", router))
// }
