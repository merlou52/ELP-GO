package main

import (
	utils "ELP-GO/src/client/client_utils"
	"fmt"
	"net"
)

func main() {
	// numéro de port établi au préalable
	PORT := "8080"

	// connexion au serveur
	conn, err := net.Dial("tcp", "localhost:"+PORT)
	defer conn.Close()

	if err != nil {
		return
	}

	// attendre réception liste filtres serveurs
	listeFilters := utils.ReceiveString(conn, '\t')
	fmt.Println(listeFilters)

	// saisie du filtre
	utils.InputFilter(conn)

	// saisie nom fichier image + validation (exist or not)
	image_path, image_path_abs := utils.InputImagePath()

	// envoi nom image on envoie le chemin relatif car on est en local + soucis de creation
	fmt.Println("Envoi du nom de l'image")
	utils.SendString(conn, image_path+"\n")

	// envoi de l'image
	// time.Sleep(1 * time.Second)
	fmt.Println("Envoi de l'image:", image_path_abs)
	utils.UploadFile(conn, image_path)
	//time.Sleep(1 * time.Second)

	// attente réception nom image modifiée
	fmt.Println("Attente réception nom de l'image modifiée")
	//filename_modified := receiveString(conn, '\n')
	//fmt.Println(filename_modified)

	// attente réception image modifiée
	fmt.Println("Attente de l'image modifiée")
	//receiveFile(conn, filename_modified)
	utils.ReceiveFile(conn, image_path[:len(image_path)-4]+"_modified.txt")
}
