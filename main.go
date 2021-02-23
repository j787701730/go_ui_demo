package main

import (
	"fmt"
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"strconv"
	"time"
)

type TMainForm struct {
	*vcl.TForm
	Button1        *vcl.TButton
	TStatusBar     *vcl.TStatusBar
	TMonthCalendar *vcl.TMonthCalendar
	Text           *vcl.TLabel
	Timer1         *vcl.TTimer
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

func (f *TMainForm) doTimer(sender vcl.IObject) {
	//f.progress++
	//if f.progress > 100 {
	//	f.progress = 0
	//}
	//// 进度值
	//f.taskBar.SetProgressValue(f.progress, 100)
	timer := time.Now()
	sMonth := fmt.Sprintf("%02d", int(timer.Month()))
	sDay := fmt.Sprintf("%02d", timer.Day())
	sHour := fmt.Sprintf("%02d", timer.Hour())
	sMinute := fmt.Sprintf("%02d", timer.Minute())
	sSecond := fmt.Sprintf("%02d", timer.Second())
	f.Text.SetCaption(strconv.Itoa(timer.Year()) + "-" + sMonth + "-" + sDay + " " + sHour + ":" + sMinute + ":" + sSecond)
	//if strings.Contains(sSecond, "5") {
	//	form1.ScreenCenter()
	//	form1.Show()
	//}
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("主窗口")
	f.SetColor(colors.ClWhite)
	// 标题栏不显示
	//f.EnabledSystemMenu(false)
	// 设置边框, 0 会禁止拉伸和移动
	//f.SetBorderStyle(0)

	// 禁止最大化
	//f.EnabledMaximize(false)

	f.SetWidth(600)
	f.SetHeight(400)
	//f.ScreenCenter()
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

	//f.TMonthCalendar = vcl.NewMonthCalendar(f)
	//f.TMonthCalendar.SetParent(f)
	//f.TMonthCalendar.SetTop(100)

	timer := time.Now()
	sMonth := fmt.Sprintf("%02d", int(timer.Month()))
	sDay := fmt.Sprintf("%02d", timer.Day())
	sHour := fmt.Sprintf("%02d", timer.Hour())
	sMinute := fmt.Sprintf("%02d", timer.Minute())
	sSecond := fmt.Sprintf("%02d", timer.Second())
	f.Text = vcl.NewLabel(f)
	f.Text.SetParent(f)
	f.Text.SetWidth(400)
	f.Text.SetHeight(400)
	//f.Text.SetColor(colors.ClBlack)
	f.Text.SetAlignment(types.AsrCenter)
	f.Text.SetLayout(types.AkRight)
	f.Text.SetCaption(strconv.Itoa(timer.Year()) + "-" + sMonth + "-" + sDay + " " + sHour + ":" + sMinute + ":" + sSecond)
	f.Text.Font().SetSize(16)
	f.Text.Font().SetColor(colors.ClBlue)
	f.Text.SetWordWrap(true)
	// 居中
	f.Text.AnchorHorizontalCenterTo(f)
	f.Text.AnchorVerticalCenterTo(f)
	f.Timer1 = vcl.NewTimer(f)
	f.Timer1.SetInterval(1)
	f.Timer1.SetOnTimer(f.doTimer)
}

func (f *TMainForm) OnFormCloseQuery(Sender vcl.IObject, CanClose *bool) {
	*CanClose = vcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
}

func (f *TMainForm) OnButton1Click(object vcl.IObject) {
	form1.Show()
	form1.ScreenCenter()
	f.EnabledMinimize(false)
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
