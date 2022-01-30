package main

import (
	"strconv"

	//"fyne.io/fyne/v2"
	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)
type Size struct {
    Width  float32 // The number of units along the X axis.
    Height float32 // The number of units along the Y axis.
}
func showCalci() {
	
//	 c:= Size{1000.0,10005.00}
//
//	a := app.New()
	
	//w.Resize(fyne.Size(c));
	output:=""
	input:= widget.NewLabel(output)
	historystr:=""
	history:=widget.NewLabel(historystr)
	ishistory:=false
	var historyarr[]string
	historybtn :=widget.NewButton("history",func() {
if ishistory{
	historystr=""
}else{
for i:=len(historyarr)-1 ;i>=0 ;i--{
		historystr=historystr+historyarr[i]
		historystr+="\n"
	}}
		ishistory=!ishistory
	history.SetText(historystr)
	


	})
	backbtn :=widget.NewButton("back",func() {
		if len(output)>0{
		output=output[:len(output)-1]
		input.SetText(output)
		}
	})
clearbtn:=widget.NewButton("clear",func() {
	output =""
	input.SetText((output))

	})
	openbtn:=widget.NewButton("(",func() {
		output =output+"("
		input.SetText((output))

	})
	closebtn:=widget.NewButton(")",func() {
		output =output+")"
		input.SetText((output))

	})
	dividebtn:=widget.NewButton("/",func() {
		output =output+"/"
		input.SetText((output))

	})
	sevenbtn:=widget.NewButton("7",func() {
		output =output+"7"
		input.SetText((output))

	})
	eightbtn:=widget.NewButton("8",func() {
		output =output+"8"
		input.SetText((output))

	})
	ninebtn:=widget.NewButton("9",func() {
		output =output+"9"
		input.SetText((output))
	})
	multiplybtn:=widget.NewButton("*",func() {
		output =output+"*"
		input.SetText((output))

	})
	fourbtn:=widget.NewButton("4",func() {
		output =output+"4"
		input.SetText((output))

	})
	fivebtn:=widget.NewButton("5",func() {

		output =output+"5"
		input.SetText((output))
	})
	sixbtn:=widget.NewButton("6",func() {
		output =output+"6"
		input.SetText((output))

	})
	
	minusbtn:=widget.NewButton("-",func() {
		output =output+"-"
		input.SetText((output))

	})
	onebtn:=widget.NewButton("1",func() {
		output =output+"1"
		input.SetText((output))

	})
	twobtn:=widget.NewButton("2",func() {
		output =output+"2"
		input.SetText((output))

	})
	threebtn:=widget.NewButton("3",func() {
		output =output+"3"
		input.SetText((output))

	})
	sumbtn:=widget.NewButton("+",func() {
		output =output+"+"
		input.SetText((output))

	})
	zerobtn:=widget.NewButton("0",func() {
		output =output+"0"
		input.SetText((output))

	})
	dotbtn:=widget.NewButton(".",func() {
		output =output+"."
		input.SetText((output))

	})
	equalbtn:=widget.NewButton("=",func() {
		expression, err := govaluate.NewEvaluableExpression(output);
		if err ==nil{

		
		result, err := expression.Evaluate(nil);
		if err == nil{
			ans:= strconv.FormatFloat(result.(float64),'f',-1,64)
             strToAppend:=output+ "=" +ans;
			 historyarr=append((historyarr),strToAppend)
			 output=ans
		}else{
			output ="error"
		}
	}else{
		output ="error"
	}
	input.SetText(output)

		
	

	})

	
	calcContainer:=(container.NewVBox(
		input ,
		history,//treted as a rowS AND WE MAKE COLNM THEN 
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historybtn,
				backbtn,
			),
			container.NewGridWithColumns(4,
				clearbtn,
				openbtn,
				closebtn,
				
				dividebtn,
			),
			container.NewGridWithColumns(4,
				sevenbtn,
				eightbtn,  
				ninebtn,
				multiplybtn,
			),
			container.NewGridWithColumns(4,
                fourbtn,
                fivebtn,
                sixbtn,
                minusbtn,

			),
			container.NewGridWithColumns(4,
		        onebtn,
	         	twobtn,
                threebtn,
                sumbtn,
			),

		
		     container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
			    zerobtn,
			    dotbtn,),
			    equalbtn, ), 
				),
			 ))
	// 		 w := myapp.NewWindow("CALCULATOR")
	w := myapp.NewWindow("Calculator");
	w.Resize(fyne.NewSize(450,280));
	//r, _ := fyne.LoadResourceFromPath("static\\calclogo.png")
	//w.SetIcon(r)
	w.SetContent(container.NewBorder(nil,nil,nil,nil,calcContainer))
	w.Show()
}