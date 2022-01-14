package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Weather App Pep")

	//API 
	res , err:= http.Get("https://api.openweathermap.org/data/2.5/weather?q=mumbai&APPID=218268b13c9594c6481913b7fb8ffe4b")
	if err!=nil{
		fmt.Println(err);
	}
	defer res.Body.Close()

	body , err:= ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println(err);
	}

	weather , err:= UnmarshalWeather(body)
	if err!=nil{
		fmt.Println(err);
	}

		img:= canvas.NewImageFromFile("download.jpg")
		img.FillMode = canvas.ImageFillOriginal

		lable1:= canvas.NewText("Weather Details" , color.White)
		lable1.TextStyle = fyne.TextStyle{Bold: true}

		lable2:= canvas.NewText(fmt.Sprintf("country %s" , weather.Sys.Country) , color.Black)
		lable3:= canvas.NewText(fmt.Sprintf("Wind Speed %2f" , weather.Wind.Speed) , color.Black)
		lable4:= canvas.NewText(fmt.Sprintf("Temp %2f" , weather.Main.Temp) , color.Black)
		lable5:= canvas.NewText(fmt.Sprintf("Humidity %2f " , weather.Main.Humidity) , color.Black)
		lable6:= canvas.NewText(fmt.Sprintf("TempMax %2f" , weather.Main.TempMax) , color.Black)
		lable7:= canvas.NewText(fmt.Sprintf("TempMin %2f" , weather.Main.TempMin) , color.Black)



		w.SetContent(
			container.NewVBox(
				lable1,
				img,
				lable2,
				lable3,
				lable4,
				lable5,
				lable6,
				lable7,
			),
		)

	w.ShowAndRun()
}


// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()


func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`     
	Weather    []WeatherElement `json:"weather"`   
	Base       string           `json:"base"`      
	Main       Main             `json:"main"`      
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`      
	Rain       Rain             `json:"rain"`      
	Clouds     Clouds           `json:"clouds"`    
	Dt         int64            `json:"dt"`        
	Sys        Sys              `json:"sys"`       
	Timezone   int64            `json:"timezone"`  
	ID         int64            `json:"id"`        
	Name       string           `json:"name"`      
	Cod        int64            `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
}

type Rain struct {
	The1H float64 `json:"1h"`
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type WeatherElement struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
}
