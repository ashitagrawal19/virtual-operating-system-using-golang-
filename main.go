package main

import (
	
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
)
var myapp fyne.App =app.New()
 var mywindow fyne.Window=myapp.NewWindow("os")

 var btn1 fyne.Widget
 var btn2 fyne.Widget
  var btn3 fyne.Widget
  var btn4 fyne.Widget
 var img fyne.CanvasObject

 var Deskbt fyne.Widget

 var panelContent *fyne.Container
func main() {
	myapp.Settings().SetTheme(theme.DarkTheme())
	img =canvas.NewImageFromFile("C:\\Users\\Ashita agrawal\\Desktop\\web dev\\os\\download.jpg")

	btn1=widget.NewButtonWithIcon("weather app",theme.InfoIcon(),func(){
		 showWeatherApp(mywindow)

	})
	btn2=widget.NewButtonWithIcon("Calculator",theme.ContentAddIcon(),func(){
	   showCalci()

  })
  btn3=widget.NewButtonWithIcon("Gallery",theme.StorageIcon(),func(){
   showGallery()

})
btn4=widget.NewButtonWithIcon("Texteditor",theme.DocumentIcon(),func(){
   showText()

})
Deskbt =widget.NewButtonWithIcon("home",theme.HomeIcon(),func(){
	mywindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
})

panelContent =container.NewVBox((container.NewGridWithColumns(5,Deskbt,btn1,btn2,btn3,btn4)))
mywindow.Resize(fyne.NewSize(700,700))
mywindow.CenterOnScreen();
mywindow.SetContent(
	container.NewBorder(panelContent,nil,nil,nil,img),
)
mywindow.ShowAndRun()
	

}
