package main

import (
	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	//"fyne.io/fyne/v2/theme"
	//"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func showGallery() {
	//a := app.New()
	w := myapp.NewWindow("GALLERY")
	w.Resize(fyne.NewSize(900,900))
root_src:="C:\\Users\\Ashita agrawal\\Pictures";
files, err := ioutil.ReadDir(root_src)
if err != nil {
	log.Fatal(err)
}
tabs := container.NewAppTabs()
var piscsarr []string 
for _, file := range files {
	
     if !file.IsDir(){
	ext := strings.Split(file.Name(), ".")[1]
	if ext=="png" || ext=="jpg" {
		piscsarr = append(piscsarr)
		image:=canvas.NewImageFromFile(root_src+"\\"+file.Name())
		tabs.Append(container.NewTabItem(file.Name(), image ))

	}

	}
}


	



	

tabs.SetTabLocation(container.TabLocationLeading)

w.SetContent((container.NewBorder(nil,nil,nil,nil,tabs )))

	w.Show()
}
