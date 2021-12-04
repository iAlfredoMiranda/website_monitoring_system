package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//constant of duration for monitoring, change here for more or less time

const monitorings = 4
const delayMonitoring = 15

//function main.
func main() {

	showIntro() //call the function for show the Welcome
	showMenu()  //call the function for show the Menu

	command := readCommand() //var command recive the readCommand

	//Use cases
	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Exibindo Logs")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do Programa")
		os.Exit(0)
	default:
		fmt.Println("Comando Não reconhecido.")

	}

}

//func to Welcome
func showIntro() {
	yourName := "Alfredo"
	version := 1.1

	fmt.Println("Olá, sr.", yourName)
	fmt.Println("Este programa ta na versão", version)

}

//func to read user command
func readCommand() int {

	var readedCommand int

	fmt.Scan(&readedCommand)
	fmt.Println("O comando escolhido foi: ", readedCommand)

	return readedCommand

}

//func to Menu
func showMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

//this func start and do the monitoring
func startMonitoring() {

	fmt.Println("Monitorando...")

	sites := readArchiveSites() //call function with site list

	for i := 0; i < monitorings; i++ { //Loop For do monitoring of sites
		for i, site := range sites {

			fmt.Println("Sites monitorados ", i, ": ", site)
			testSite(site) // call func test
		}
		time.Sleep(delayMonitoring * time.Minute) //Monitoring interval
		fmt.Println("")                           //Jump line
	}

}

// this func show the status of sites
func testSite(site string) {

	resp, err := http.Get(site)

	if err != nil { // If a problem show the error

		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 { // If the site has ok show ok

		fmt.Println("Site: ", site, "Foi carregado com sucesso!", resp.StatusCode)
		recordLog(site, true) // call to func recordLog

	} else { // If the site has a problem show the error status
		fmt.Println("Site: ", site, "Não foi possivel carregar. Status Code: ", resp.StatusCode)
		recordLog(site, false) // call to func recordLog
	}

}

// this func read the archive that has a URL list
func readArchiveSites() []string {

	var sites []string // a slice with urls sites

	arquivo, err := os.Open("sites.txt")

	if err != nil { // If a problem show the error

		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo) // command for read archive

	for { // loop for read urls in line
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

// this func record logs
func recordLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // this command is for read the archive or create or record the archive with logs

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "online: " + strconv.FormatBool(status) + "\n") // this command is for write logs with data, site name and status

	arquivo.Close()
}

// this func show logs
func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
