package main

import (
	"fmt"
	"github.com/ying32/govcl/vcl"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl/types"
)

type TMainForm struct {
	*vcl.TForm
	Button1        *vcl.TButton
	TStatusBar     *vcl.TStatusBar
	TMonthCalendar *vcl.TMonthCalendar
}

type TForm1 struct {
	*vcl.TForm
	Button1 *vcl.TButton
	TImage  *vcl.TImage
}

var (
	mainForm *TMainForm
	form1    *TForm1
)

func main() {
	vcl.RunApp(&mainForm, &form1)
}

// --------------MainForm -----------------
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("主窗口")
	f.EnabledMaximize(false)
	f.SetWidth(600)
	f.SetHeight(400)
	f.ScreenCenter()
	f.SetLeft(vcl.Screen.WorkAreaWidth() - 600 - 10)
	f.SetTop(vcl.Screen.WorkAreaHeight() - 400 - 30)
	fmt.Println(vcl.Screen.WorkAreaHeight())
	f.Button1 = vcl.NewButton(f)
	f.Button1.SetWidth(200)
	f.Button1.SetParent(f)
	f.Button1.SetCaption("后宫佳丽三千, 点开有惊喜哦")
	f.Button1.SetLeft(50)
	f.Button1.SetTop(50)
	f.Button1.SetOnClick(f.OnButton1Click)

	f.TStatusBar = vcl.NewStatusBar(f)
	f.TStatusBar.SetSimpleText("我是状态栏: 老牛逼了")
	f.TStatusBar.SetParent(f)

	f.TMonthCalendar = vcl.NewMonthCalendar(f)
	f.TMonthCalendar.SetParent(f)
	f.TMonthCalendar.SetTop(100)

}

func (f *TMainForm) OnFormCloseQuery(Sender vcl.IObject, CanClose *bool) {
	*CanClose = vcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
}

func (f *TMainForm) OnButton1Click(object vcl.IObject) {
	form1.ShowModal()
	form1.ScreenCenter()
	//fmt.Println()
}

// ---------- Form1 ----------------

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("后宫佳丽三千")
	fmt.Println("onCreate")
	f.Button1 = vcl.NewButton(f)
	fmt.Println("f.Button1:", f.Button1.Instance())
	//f.Button1.SetParent(f)
	//f.Button1.SetName("Button1")
	//f.Button1.SetCaption("我是按钮")
	//f.Button1.SetOnClick(f.OnButton1Click)

	f.TImage = vcl.NewImage(f)
	f.TImage.SetParent(f)
	f.TImage.SetWidth(200)
	f.TImage.SetHeight(400)
	f.TImage.Picture().LoadFromFile("./wallhaven-x1vd3d.jpg")
	f.TImage.SetProportional(true)
}

func (f *TForm1) OnButton1Click(object vcl.IObject) {
	vcl.ShowMessage("Click")
}
