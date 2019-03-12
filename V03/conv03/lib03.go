package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/asdine/storm"
)

// Will save Structs to BOLT

func SaveToDBALL() {
	p := fmt.Println
	//  Will save to a BOLT DB

	db, err := storm.Open("./KKdb/KKErrors.db")
	if err != nil {
		p("=======   could not create the DB file  =====")
		os.Mkdir("KKdb", 0777)
	}
	defer db.Close()

	//=================================================
	//save to first FDir
	p("=======   FDir  =====")
	for i, r := range TempAllDir {
		err = db.Save(&r)
		p(i)
		if err != nil {
			p("=======   could not  db.Save(&TempAllDir)  =====")
		}
	}

	var rr0 []Fdir
	err = db.All(&rr0)
	for i, r := range rr0 {
		p(i)
		p(r.ID)
	}

	//=================================================
	//save to first OutS01AgentCaptureUnreachable
	p("=======   OutS01AgentCaptureUnreachable  =====")
	for i, r := range TempOutS01 {
		err = db.Save(&r)
		p(i)
		if err != nil {
			p("=======   could not  db.Save(TempOutS01)  =====")
		}
	}

	var rr1 []OutS01AgentCaptureUnreachable
	err = db.All(&rr1)
	for i, r := range rr1 {
		p(i)
		p(r.ID)
	}

	//=================================================
	//save to first OutS01AgentCaptureUnreachable
	p("=======   OutS02OneWayAudioRecording  =====")
	for i, r := range TempOutS02 {
		err = db.Save(&r)
		p(i)
		if err != nil {
			p("=======   could not  db.Save(TempOutS02)  =====")
		}
	}

	var rr2 []OutS02OneWayAudioRecording
	err = db.All(&rr2)
	for i, r := range rr2 {
		p(i)
		p(r.ID)
	}

	//=================================================
	//save to first OutS03PacketsMissed
	p("=======   OutS03PacketsMissed  =====")
	for i, r := range TempOutS03 {
		err = db.Save(&r)
		p(i)
		if err != nil {
			p("=======   could not  db.Save(TempOutS03)  =====")
		}
	}

	var rr3 []OutS03PacketsMissed
	err = db.All(&rr3)
	for i, r := range rr3 {
		p(i)
		p(r.ID)
	}

	//=================================================
	//save to first OutS04PacketsMissedForRecording
	p("=======   OutS04PacketsMissedForRecording  =====")
	for i, r := range TempOutS04 {
		err = db.Save(&r)
		p(i)
		if err != nil {
			p("=======   could not  db.Save(TempOutS04)  =====")
		}
	}

	var rr4 []OutS04PacketsMissedForRecording
	err = db.All(&rr4)
	for i, r := range rr4 {
		p(i)
		p(r.ID)
	}

	//=================================================
	//save to first OutS04PacketsMissedForRecording
	p("=======   OutS05RecordingError  =====")
	for i, r := range TempOutS05 {
		err = db.Save(&r)
		p(i)
		if err != nil {
			p("=======   could not  db.Save(TempOutS05)  =====")
		}
	}

	var rr5 []OutS05RecordingError
	err = db.All(&rr5)
	for i, r := range rr5 {
		p(i)
		p(r.ID)
	}

	//=================================================
	//=================================================
	//=================================================
	//  MAKE SPREAD SHEET

	xlsx := excelize.NewFile()

	//=================================================
	//=================================================
	//=================================================
	// Create a new sheet.
	index := xlsx.NewSheet("ALL ERRORS")
	xlsx.SetActiveSheet(index)
	// Do Headers
	p(H00[0])
	p(H00)
	for i, h := range H00 {
		p("========  Adding Headers  ========")
		C := WriteCell(i+1, 1)
		p(H00[i])
		xlsx.SetCellValue("ALL ERRORS", C, h)
	}

	p("===========   begining the real loop =======")
	err = db.All(&rr0)
	if err != nil {

	}

	rr0NumField01 := rr0[0]
	rr0NumField02 := reflect.ValueOf(rr0NumField01)
	p("===  number of fields ===")
	rr0NumField03 := rr0NumField02.NumField()
	p(rr0NumField03)

	p("========= ", rr0NumField03) //  number of records
	for i, r := range rr0 {
		vr := reflect.ValueOf(r)
		p("==", strconv.Itoa(i))
		for x := 0; x <= rr0NumField03-1; x++ {
			a := vr.Field(x)
			p("==", x, "==", a)
			C := WriteCell(x+1, i+2)
			xlsx.SetCellValue("ALL ERRORS", C, a)
		}
	}

	//=================================================
	//=================================================
	//=================================================
	// Create a new sheet.
	index = xlsx.NewSheet("AgentCaptureUnreachable")
	xlsx.SetActiveSheet(index)
	// Do Headers
	p(H01[0])
	p(H01)
	for i, h := range H01 {
		p("========  Adding Headers  ========")
		C := WriteCell(i, 1)
		p(H01[i])
		xlsx.SetCellValue("AgentCaptureUnreachable", C, h)
	}

	p("===========   begining the real loop =======")

	rr1NumField01 := rr1[0]
	rr1NumField02 := reflect.ValueOf(rr1NumField01)
	p("===  number of fields ===")
	rr1NumField03 := rr1NumField02.NumField()
	p(rr1NumField03)

	p("========= ", rr1NumField03) //  number of records
	for i, r := range rr1 {
		vr := reflect.ValueOf(r)
		p("==", strconv.Itoa(i))
		p(vr.Field(0))
		for x := 0; x <= rr1NumField03-1; x++ {
			fmt.Print("=====", strconv.Itoa(x), "  ===")
			a := vr.Field(x)
			p(a)
			C := WriteCell(x+1, i+2)
			xlsx.SetCellValue("AgentCaptureUnreachable", C, a)
		}
	}

	//=================================================
	//=================================================
	//=================================================
	// Create a new sheet.  OneWayAudioRecording
	index = xlsx.NewSheet("OneWayAudioRecording")
	xlsx.SetActiveSheet(index)
	// Do Headers
	p(H02[0])
	p(H02)
	for i, h := range H02 {
		p("========  Adding Headers  ========")
		C := WriteCell(i, 1)
		p(H02[i])
		xlsx.SetCellValue("OneWayAudioRecording", C, h)
	}

	p("===========   begining the real loop =======")

	rr2NumField01 := rr2[0]
	rr2NumField02 := reflect.ValueOf(rr2NumField01)
	p("===  number of fields ===")
	rr2NumField03 := rr2NumField02.NumField()
	p(rr2NumField03)

	p("========= ", rr2NumField03) //  number of records
	for i, r := range rr2 {
		vr := reflect.ValueOf(r)
		p("==", strconv.Itoa(i))
		p(vr.Field(0))
		for x := 0; x <= rr2NumField03-1; x++ {
			fmt.Print("=====", strconv.Itoa(x), "  ===")
			a := vr.Field(x)
			p(a)
			C := WriteCell(x+1, i+2)
			xlsx.SetCellValue("OneWayAudioRecording", C, a)
		}
	}
	//=====================================================
	//=====================================================
	//=====================================================
	// Create a new sheet.  PacketsMissed
	index = xlsx.NewSheet("PacketsMissed")
	xlsx.SetActiveSheet(index)
	// Do Headers
	p(H03[0])
	p(H03)
	for i, h := range H03 {
		p("========  Adding Headers  ========")
		C := WriteCell(i, 1)
		p(H03[i])
		xlsx.SetCellValue("PacketsMissed", C, h)
	}

	p("===========   begining the real loop =======")

	rr3NumField01 := rr3[0]
	rr3NumField02 := reflect.ValueOf(rr3NumField01)
	p("===  number of fields ===")
	rr3NumField03 := rr3NumField02.NumField()
	p(rr3NumField03)

	p("========= ", rr3NumField03) //  number of records
	for i, r := range rr3 {
		vr := reflect.ValueOf(r)
		p("==", strconv.Itoa(i))
		p(vr.Field(0))
		for x := 0; x <= rr3NumField03-1; x++ {
			fmt.Print("=====", strconv.Itoa(x), "  ===")
			a := vr.Field(x)
			p(a)
			C := WriteCell(x+1, i+2)
			xlsx.SetCellValue("PacketsMissed", C, a)
		}
	}

	//=====================================================
	//=====================================================
	//=====================================================
	// Create a new sheet.  PacketsMissed
	index = xlsx.NewSheet("PacketsMissedForRecording")
	xlsx.SetActiveSheet(index)
	// Do Headers
	p(H04[0])
	p(H04)
	for i, h := range H04 {
		p("========  Adding Headers  ========")
		C := WriteCell(i, 1)
		p(H04[i])
		xlsx.SetCellValue("PacketsMissedForRecording", C, h)
	}

	p("===========   begining the real loop =======")

	rr4NumField01 := rr4[0]
	rr4NumField02 := reflect.ValueOf(rr4NumField01)
	p("===  number of fields ===")
	rr4NumField03 := rr4NumField02.NumField()
	p(rr4NumField03)

	p("========= ", rr4NumField03) //  number of records
	for i, r := range rr4 {
		vr := reflect.ValueOf(r)
		p("==", strconv.Itoa(i))
		p(vr.Field(0))
		for x := 0; x <= rr4NumField03-1; x++ {
			fmt.Print("=====", strconv.Itoa(x), "  ===")
			a := vr.Field(x)
			p(a)
			C := WriteCell(x+1, i+2)
			xlsx.SetCellValue("PacketsMissedForRecording", C, a)
		}
	}

	//=====================================================
	//=====================================================
	//=====================================================
	// Create a new sheet.  PacketsMissed
	index = xlsx.NewSheet("RecordingError")
	xlsx.SetActiveSheet(index)
	// Do Headers
	p(H05[0])
	p(H05)
	for i, h := range H05 {
		p("========  Adding Headers  ========")
		C := WriteCell(i, 1)
		p(H05[i])
		xlsx.SetCellValue("RecordingError", C, h)
	}

	p("===========   begining the real loop =======")

	rr5NumField01 := rr5[0]
	rr5NumField02 := reflect.ValueOf(rr5NumField01)
	p("===  number of fields ===")
	rr5NumField03 := rr5NumField02.NumField()
	p(rr5NumField03)

	p("========= ", rr5NumField03) //  number of records
	for i, r := range rr5 {
		vr := reflect.ValueOf(r)
		p("==", strconv.Itoa(i))
		p(vr.Field(0))
		for x := 0; x <= rr5NumField03-1; x++ {
			fmt.Print("=====", strconv.Itoa(x), "  ===")
			a := vr.Field(x)
			p(a)
			C := WriteCell(x+1, i+2)
			xlsx.SetCellValue("RecordingError", C, a)
		}
	}
	//=====================================================
	//=====================================================
	//=====================================================
	err = xlsx.SaveAs("./KKErrors.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

//=====================================================
//=====================================================
//=====================================================
func WriteCell(C int, row int) string {
	p := fmt.Println
	Col := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW"}
	cell := Col[C] + strconv.Itoa(row)
	p("========   the Cell is.... ")
	p(cell)
	return cell
}

//====================================================
//====================================================
//====================================================
//  Builds the Headers for Spread Sheet
func buildHeaders() {
	p := fmt.Println
	//=========================
	p("=========   building Spread Sheets  ===========")

	p("======  Build H00  =====")
	t := reflect.TypeOf(TempAllDir[0])
	p(t.NumField())
	//H00 = append(H00, " ")
	for i := 0; i <= t.NumField()-1; i++ {
		field := t.Field(i)
		p(i)
		tag := field.Tag.Get("ss")
		H00 = append(H00, tag)
		p(tag)
	}
	p(H00)

	p("======  Build H01  =====")
	t = reflect.TypeOf(TempOutS01[0])
	p(t.NumField())
	H01 = append(H01, " ")
	for i := 0; i <= t.NumField()-1; i++ {
		field := t.Field(i)
		p(i)
		tag := field.Tag.Get("ss")
		H01 = append(H01, tag)
		p(tag)
	}
	p(H01)

	p("======  Build H02  =====")
	t = reflect.TypeOf(TempOutS02[0])
	p(t.NumField())
	H02 = append(H02, " ")
	for i := 0; i <= t.NumField()-1; i++ {
		field := t.Field(i)
		p(i)
		tag := field.Tag.Get("ss")
		H02 = append(H02, tag)
		p(tag)
	}
	p(H02)

	p("======  Build H03  =====")
	t = reflect.TypeOf(TempOutS03[0])
	p(t.NumField())
	H03 = append(H03, " ")
	for i := 0; i <= t.NumField()-1; i++ {
		field := t.Field(i)
		p(i)
		tag := field.Tag.Get("ss")
		H03 = append(H03, tag)
		p(tag)
	}
	p(H03)

	p("======  Build H04  =====")
	t = reflect.TypeOf(TempOutS04[0])
	p(t.NumField())
	H04 = append(H04, " ")
	for i := 0; i <= t.NumField()-1; i++ {
		field := t.Field(i)
		p(i)
		tag := field.Tag.Get("ss")
		H04 = append(H04, tag)
		p(tag)
	}
	p(H04)

	p("======  Build H05  =====")
	t = reflect.TypeOf(TempOutS05[0])
	p(t.NumField())
	H05 = append(H05, " ")
	for i := 0; i <= t.NumField()-1; i++ {
		field := t.Field(i)
		p(i)
		tag := field.Tag.Get("ss")
		H05 = append(H05, tag)
		p(tag)
	}
	p(H05)

	p("======  doing reflect 3 =====")

}
