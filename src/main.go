package main

import (
	"io"
	"os"
)

func main() {

	createIndex()

}

func createIndex() {
	const indexMain = `
	<!DOCTYPE HTML>
	<html>
  	<head>
    	<meta charset="utf-8">
		<title>Simple Go Web App</title>
		<link rel="stylesheet" href="css/main.css">
  	</head>
  	<body>
		<div id='root'></div>
		<h1 class='heading'>Some Heading</h1>
  	</body>
	</html>
	`

	const mainCSS=`
		.heading{
			color: red;
		}
	
	`
	//makes folders if not exists
	os.MkdirAll("../main/img", os.ModePerm)
	os.MkdirAll("../main/css", os.ModePerm)
	os.MkdirAll("../main/js", os.ModePerm)


	//make index file
	file, _ := os.Create("../main/fromString.html")
	// checkError(err)
	defer file.Close()
	io.WriteString(file, indexMain)
	// checkError(err)

	//make css file
	cssFile, _ := os.Create("../main/css/main.css")
	// checkError(err)
	defer cssFile.Close()
	io.WriteString(cssFile, mainCSS)

	//make js file


}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
