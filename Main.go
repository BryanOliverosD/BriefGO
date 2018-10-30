package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

)
// struct final
type Brief struct{
	Categoria string
	Subcategoria string
	Tipollamado string
	Descripcion string
	Marca string
	Sku string
	Precionormal string
	Preciooferta string
	Ou string
	PorcentajeDscto int
	Url string
}
// Users struct which contains
// an array of users
type State struct {
	State []resultList `json:"resultList"`
}
type Dato struct {
	Data State `json:"state"`
}
type resultList struct {
	//variables que se rescatan del json
	Marca   		string `json:"brand"`
	Descripcion  	string `json:"title"`
	Url    			string   `json:"url"`
	SubCategoria 	string `json:"backendCategory"`
	Sku				string `json:"skuId"`
	Precio 			[]Precios `json:"prices"`
	//Subcategoria string `json:"url"`
}
type Precios struct	{
		Label			string   `json:"label"`
		Valor_precio	string	`json:"originalPrice"`
		OU 				bool `json:"opportunidadUnica"`
}

func request (){
	var Lista_objetos []Brief
	var Unidad Brief

	// request
	url := "https://www.falabella.com/rest/model/falabella/rest/browse/BrowseActor/get-product-record-list"

	payload := strings.NewReader("{\"navState\":\"/search/?Ntt=6041709\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("Postman-Token", "27c86025-3816-45cc-85df-df1182082623")

	res, _ := http.DefaultClient.Do(req)

	t := new(Dato)
	if err := json.NewDecoder(res.Body).Decode(t); err != nil {
		log.Fatalln(err)
	}
	//fmt.Printf("Format: %s\n", t.Sku.State[0].Marca)
	// Se asignan variables que no necesitan ser procesadas
	Unidad.Subcategoria = t.Data.State[0].SubCategoria
	Unidad.Marca = t.Data.State[0].Marca
	Unidad.Descripcion = t.Data.State[0].Descripcion
	Unidad.Url = "www.falabella.com"+ string(t.Data.State[0].Url)
	Unidad.Sku = t.Data.State[0].Sku

	//se extraen los primeros 3 caracteres de la subcategoría
	AuxCategoria := []rune(t.Data.State[0].SubCategoria)
	Unidad.Categoria = string(AuxCategoria[0:3])

	//Lista de los objetos
	Lista_objetos = append(Lista_objetos,Unidad)
	fmt.Println(Lista_objetos)
	fmt.Println(Lista_objetos[0].Marca,Lista_objetos[0].Precionormal)


}
func main() {
	request()

}
	/*for i := 0; i < 3; i++ {
		fmt.Println("User marca: " + datos.Sku.resultList[i].Marca)
		fmt.Println("User precio: " + strconv.Itoa(datos.resultList[i].Precio))
		fmt.Println("User descrip: " + datos.resultList[i].Descripcion)

	}
	/*
	xlsx, err := excelize.OpenFile("./Brief_Evento.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Create a new sheet.
	index := xlsx.NewSheet("Sheet2")
	// Set value of a cell.
	xlsx.SetCellValue("Hoja1", "C5", "Hello world.")
	xlsx.SetCellValue("Hoja1", "B5", 100)
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err = xlsx.SaveAs("./Brief_Evento.xlsx")
	if err != nil {
		fmt.Println(err)
	}*/