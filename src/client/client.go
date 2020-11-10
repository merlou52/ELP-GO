package main

import (
	"ELP-GO/src/elputils"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Get port from argc, or use default value 8080
	var PORT string
	if len(os.Args) == 2 {
		PORT = os.Args[1]
	} else {
		PORT = "8080"
	}

	// Connecting to the server on port PORT
	fmt.Printf("Connecting to server on port %s. You can change the port used by specifying it as an argc value eg 'client 5000'\n", PORT)
	conn, err := net.Dial("tcp", "localhost:"+PORT)

	if err != nil {
		fmt.Printf("Couldn't listen on port %s. Is the server running ?\n", PORT)
		return
	}

	// Connection successful, close the connection when it's over
	fmt.Printf("Connection established with server on port %s\n\n", PORT)
	defer conn.Close()

	// Input filter
	filterList := elputils.ReceiveArray(conn, ";", '\n')
	elputils.InputFilter(conn, filterList)

	// Print current directory
	dir, err := os.Getwd()
	fmt.Println(dir)

	// Input image filename
	imagePath, imagePathAbs := elputils.InputImagePath()

	// Input output file
	fmt.Println("Enter the output file name")
	outputFile := strings.Trim(elputils.InputString(), "\n")

	// Send image filename to the server
	fmt.Println("Sending filename")
	elputils.SendString(conn, imagePath+"\n")

	// Send file
	fmt.Println("Sending image", imagePathAbs)
	elputils.UploadFile(conn, imagePath)

	// Receiving the modified image
	fmt.Println("Waiting for the modified image...")
	elputils.ReceiveFile(conn, outputFile)
	fmt.Println("Transformation complete, output in ", outputFile)
}
