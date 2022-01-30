package main

import (
	"fmt"
	"image/color"

	"io/ioutil"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	"fyne.io/fyne/v2/container"
	"encoding/json"
	"net/http"
)

func  showWeatherApp(w fyne.Window) {
	//app := app.New()

	//w =myapp.NewWindow("Weather app ")
//w.Resize(fyne.NewSize(600,600));
//api part 
res ,err:=http.Get("https://api.openweathermap.org/data/2.5/weather?q=London,uk&APPID=37c25980de1fe71eb381ffd4b11d8b8e")
if err !=nil{
fmt.Println(err)
}
defer res.Body.Close()
body , err:= ioutil.ReadAll(res.Body)
if err !=nil{
	fmt.Println(err)
	}
weather,err:=UnmarshalWelcome(body)
if err !=nil{
	fmt.Println(err)
	}
img :=canvas.NewImageFromFile("images.png")
img.FillMode =canvas.ImageFillOriginal
label1:=canvas.NewText("weather details",color.White)
label1.TextStyle=fyne.TextStyle{Bold: true}
label2:=canvas.NewText(fmt.Sprintf("country %s ",weather.Sys.Country),color.White)
label3:=canvas.NewText(fmt.Sprintf("wind speed %.2f ",weather.Wind.Speed),color.White)
label4:=canvas.NewText(fmt.Sprintf("temp  %2f ",weather.Main.Temp),color.White)
label5:=canvas.NewText(fmt.Sprintf("humidity %d ",weather.Main.Humidity),color.White)
	
weatherContainer:=(container.NewVBox(
 
	label1,
	img,
	label2,
	label3,
	label4,
	label5,


))

	

w.SetContent(container.NewBorder(panelContent,nil,nil,nil,weatherContainer),)
w.Show()
}





func UnmarshalWelcome(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord     `json:"coord"`     
	Weather    []WeatherElement `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod,string"`       
}

type Clouds struct {
	All int64 `json:"all,string"`
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

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id,string"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise,string"`
	Sunset  int64  `json:"sunset,string"` 
}

type WeatherElement struct {
	ID          int64  `json:"id,string"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
}

