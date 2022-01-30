package main

import (
	
	"strconv"

	//"strings"
	//	"log"
"io"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/cmd/fyne_settings/data"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)
var count int  =1
func  showText() {
	//a := app.New()
	w := myapp.NewWindow("editor")
	w.Resize(fyne.NewSize(500,300))

content:=container.NewVBox(
	container.NewVBox(
		widget.NewLabel("pep text editor"),
	),
)
content. Add(widget.NewButton("add new file",func() {

	content.Add(widget.NewLabel("New File"+ strconv.Itoa(count) ))
	count++;

}))
input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	

	savebtn :=  widget.NewButton("Save", func() {
	
	savefiledialog:= dialog.NewFileSave(
		func(uc fyne.URIWriteCloser, e error){
			textdata:=[]byte (input.Text)
			uc.Write(textdata)
		},w)
		savefiledialog.SetFileName(("New File"+ strconv.Itoa(count-1)+"txt" ))
		savefiledialog.Show();
	})

	openbutton:=widget.NewButton("open file",func() {
		
		openfiledialog:=dialog.NewFileOpen( 
			func(r fyne.URIReadCloser, e error){
				textdata ,_:=io.ReadAll(r)
			output:=fyne.NewStaticResource("New File",textdata)
			viewtext:=widget.NewMultiLineEntry()
			viewtext.SetText(string(output.StaticContent))
			w:=fyne.CurrentApp().NewWindow(
				string(output.StaticName))

w.SetContent(container.NewScroll(viewtext))
w.Resize(fyne.NewSize(100,300))	
w.Show()		
},w)
openfiledialog.SetFilter(
	storage.NewExtensionFileFilter([] string{".txt"}))
openfiledialog.Show()

		
	})



w.SetContent(
	container.NewVBox(
		content,
	input,
		container.NewHBox(
		savebtn,
		openbutton,
	),
),
)
 

	w.Show()
}