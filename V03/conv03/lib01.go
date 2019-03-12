package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//====================================
//====================================
//====================================
//=======TempAllDir []Fdir    ========
func BuildTempAllDir() {
	p := fmt.Println
	var z Fdir

	p("=============================================")
	p("=========   ConsumeDB001()     ==============")
	// Step 01  - get list of Files to be imported.
	//os.Mkdir("KKdb", 0777)
	file, err := os.Open("KKIn")
	if err != nil {
		os.Mkdir("KKIn", 0777)
		os.Mkdir("KKdb", 0777)
		p("=======   this must be your first time running")
		p("===  notice I just created a KKIN directory...")
		p("=== Plese put your Error files you collet into this DIR")
		p("=== and run this program again at command line")
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders

	//====================================
	// run though list.  will save all files
	for _, fname := range list {
		p("==============   in loop 1 ======")
		ID := fname
		recStringTemp, _ := ioutil.ReadFile("./kkIn/" + fname)
		ftemp := strings.Split(fname, ".")

		AlarmTypeTemp := ftemp[0]
		AlarmInstanceidTemp := GetAlarmInstanceID(recStringTemp) //ftemp[1]
		ComponentNameTemp := GetComponentName(recStringTemp)

		EventTriggerTimeTemp := GetTriggerTime(recStringTemp)
		EventDateStringTemp := GetEventDateString(recStringTemp)
		EventTimeStringTemp := GetEventTimeString(recStringTemp)

		EventYEARTemp := getEventYEAR(recStringTemp)
		EventMonthTemp := getEventMonth(recStringTemp)
		EventDayTemp := getEventDay(recStringTemp)

		EventHourTemp := getEventHour(recStringTemp)
		EventMinTemp := getEventMin(recStringTemp)
		EventSecondTemp := getEventSecond(recStringTemp)

		EventTSInSecondsTemp := getEventTSInSeconds(recStringTemp)
		EventTP10Temp := getEventTP10(recStringTemp)
		EventTP20Temp := getEventTP20(recStringTemp)
		EventTP30Temp := getEventTP30(recStringTemp)

		EventTimeStampIntTemp := GetUnixTS(recStringTemp)  //int64
		RecDetailsTemp := getRecDetailsTemp(recStringTemp) //returns A00Standard

		TTAgentCaptureUnreachableTemp := isIt_TTAgentCaptureUnreachable(fname)     //bool
		TTOneWayAudioRecordingTemp := isIt_TTOneWayAudioRecording(fname)           //bool
		TTPacketsMissedTemp := isIt_TTPacketsMissed(fname)                         // bool
		TTPacketsMissedForRecordingTemp := isIt_TTPacketsMissedForRecording(fname) //bool
		TTRecordingErrorTemp := isIt_TTRecordingError(fname)                       //bool

		//==============================
		z.ID = ID
		z.FName = fname
		z.AlarmType = AlarmTypeTemp
		z.Alarminstanceid = AlarmInstanceidTemp
		z.ComponentName = ComponentNameTemp
		z.EventTrigger = EventTriggerTimeTemp
		z.EventDateString = EventDateStringTemp
		z.EventTimeString = EventTimeStringTemp
		z.EventYEAR = EventYEARTemp
		z.EventMonth = EventMonthTemp
		z.EventDay = EventDayTemp

		z.EventHour = EventHourTemp
		z.EventMin = EventMinTemp
		z.EventSecond = EventSecondTemp

		z.EventTSInSeconds = EventTSInSecondsTemp
		z.EventTP10 = EventTP10Temp
		z.EventTP20 = EventTP20Temp
		z.EventTP30 = EventTP30Temp

		z.EventTimeStampInt = EventTimeStampIntTemp

		z.TTAgentCaptureUnreachable = TTAgentCaptureUnreachableTemp
		z.TTOneWayAudioRecording = TTOneWayAudioRecordingTemp
		z.TTPacketsMissed = TTPacketsMissedTemp
		z.TTPacketsMissedForRecording = TTPacketsMissedForRecordingTemp
		z.TTRecordingError = TTRecordingErrorTemp

		z.RecString = recStringTemp
		z.RecDetails = RecDetailsTemp

		TempAllDir = append(TempAllDir, z)
	}

}

//====================================
//====================================
//====================================

func GetAlarmInstanceID(xmlFile []byte) string {
	var q A00StandardStruct
	xml.Unmarshal(xmlFile, &q)
	fmt.Print("====  AlarmInstanceID  ==")
	fmt.Println(q.AlarmInstanceID)
	return q.AlarmInstanceID
}

func GetComponentName(xmlFile []byte) string {
	var q A00StandardStruct
	xml.Unmarshal(xmlFile, &q)
	fmt.Print("====  ComponentName  ==")
	fmt.Println(q.ComponentName)
	return q.ComponentName
}

func GetTriggerTime(xmlFile []byte) string {
	var q A00StandardStruct
	xml.Unmarshal(xmlFile, &q)
	fmt.Print("====  triggerTime  ==")
	fmt.Println(q.TriggerTime)
	return q.TriggerTime
}

func GetEventDateString(xmlFile []byte) string {
	var q A00StandardStruct
	xml.Unmarshal(xmlFile, &q)
	fmt.Print("====  triggerTime  ==")
	fmt.Println(q.TriggerTime)
	EventDate := strings.Split(q.TriggerTime, "T")
	return EventDate[0]
}

func GetEventTimeString(xmlFile []byte) string {
	var q A00StandardStruct
	xml.Unmarshal(xmlFile, &q)
	fmt.Print("====  triggerTime  ==")
	fmt.Println(q.TriggerTime)
	EventDate := strings.Split(q.TriggerTime, "T")
	return EventDate[1]
}

func getEventYEAR(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[0], "-")
	fmt.Print("====  year  ==")
	fmt.Println(tt)
	ttt, _ := strconv.Atoi(tt[0])
	return ttt
}

func getEventMonth(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[0], "-")
	fmt.Print("====  month  ==")
	fmt.Println(tt)
	ttt, _ := strconv.Atoi(tt[1])
	return ttt
}
func getEventDay(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[0], "-")
	fmt.Print("====  day  ==")
	fmt.Println(tt)
	ttt, _ := strconv.Atoi(tt[2])
	return ttt
}

func getEventHour(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[1], ":")
	fmt.Print("====  hours  ==")
	fmt.Println(tt)
	ttt, _ := strconv.Atoi(tt[0])
	return ttt
}
func getEventMin(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[1], ":")
	fmt.Print("====  min  ==")
	fmt.Println(tt)
	ttt, _ := strconv.Atoi(tt[1])
	return ttt
}
func getEventSecond(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[1], ":")
	ttt := strings.Split(tt[2], ".")
	fmt.Print("====  sec  ==")
	fmt.Println(ttt[0])
	tttt, _ := strconv.Atoi(ttt[0])
	return tttt
}

func getEventTSInSeconds(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[1], ":")

	tt0_hour, _ := strconv.Atoi(tt[0])
	tt0_hour = tt0_hour * 60 * 60

	tt1_hours, _ := strconv.Atoi(tt[1])
	tt1_hours = tt1_hours * 60

	tt2_min, _ := strconv.Atoi(tt[2])
	ttt := tt0_hour + tt1_hours + tt2_min
	fmt.Print(ttt)
	fmt.Print("=====")
	return ttt
}

func getEventTP10(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[1], ":")

	tt0_hour, _ := strconv.Atoi(tt[0])
	tt0_hour = tt0_hour * 60 * 60

	tt1_hours, _ := strconv.Atoi(tt[1])
	tt1_hours = tt1_hours * 60

	tt2_min, _ := strconv.Atoi(tt[2])
	ttt := tt0_hour + tt1_hours + tt2_min
	fmt.Print(ttt)
	fmt.Print("=====")
	ttt = ttt / TP10
	fmt.Print("====  GetTP10  ==")
	fmt.Println(ttt)
	return ttt
}

func getEventTP20(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[1], ":")

	tt0_hour, _ := strconv.Atoi(tt[0])
	tt0_hour = tt0_hour * 60 * 60

	tt1_hours, _ := strconv.Atoi(tt[1])
	tt1_hours = tt1_hours * 60

	tt2_min, _ := strconv.Atoi(tt[2])
	ttt := tt0_hour + tt1_hours + tt2_min
	fmt.Print(ttt)
	fmt.Print("=====")
	ttt = ttt / TP20
	fmt.Print("====  GetTP10  ==")
	fmt.Println(ttt)
	return ttt
}

func getEventTP30(xmlFile []byte) int {
	t := GetTriggerTime(xmlFile)
	tt := strings.Split(t, "T")
	tt = strings.Split(tt[1], ":")

	tt0_hour, _ := strconv.Atoi(tt[0])
	tt0_hour = tt0_hour * 60 * 60

	tt1_hours, _ := strconv.Atoi(tt[1])
	tt1_hours = tt1_hours * 60

	tt2_min, _ := strconv.Atoi(tt[2])
	ttt := tt0_hour + tt1_hours + tt2_min
	fmt.Print(ttt)
	fmt.Print("=====")
	ttt = ttt / TP30
	fmt.Print("====  GetTP10  ==")
	fmt.Println(ttt)
	return ttt
}

func getRecDetailsTemp(xmlFile []byte) A00StandardStruct {
	var q A00StandardStruct
	xml.Unmarshal(xmlFile, &q)
	fmt.Print("====  getRecDetailsTemp  ==")
	fmt.Println(q)
	return q
}

func GetUnixTS(recStringTemp []byte) int64 {
	//layout := "01/02/2006 3:04:05 PM"
	m := getEventMonth(recStringTemp)
	d := getEventDay(recStringTemp)
	y := getEventYEAR(recStringTemp)
	H := getEventHour(recStringTemp)
	M := getEventHour(recStringTemp)
	S := getEventSecond(recStringTemp)

	// mm := strconv.Itoa(m)
	// dd := strconv.Itoa(d)
	// yy := strconv.Itoa(y)
	// HH := strconv.Itoa(H)
	// MM := strconv.Itoa(M)
	// SS := strconv.Itoa(S)

	//	layout = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	t := time.Date(y, time.Month(m), d, H, M, S, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	secs := t.Unix()
	return secs
}

//*=======================================
//*=======================================
//*=======================================
func isIt_TTAgentCaptureUnreachable(fileID string) bool {
	p := fmt.Println
	p("====================================================")
	p("====================================================")
	if strings.Contains(fileID, "AgentCaptureUnreachable") {
		p("=======  it is an AgentCaptureUnrechable  ====")
		p(fileID)
		p("===================")
		return true
	} else {
		return false
	}
}

//*
//*=======================================
func isIt_TTOneWayAudioRecording(fileID string) bool {
	p := fmt.Println
	p("====================================================")
	p("====================================================")
	if strings.Contains(fileID, "OneWayAudioRecording") {
		p("=======  it is an OneWayAudioRecording  ====")
		p(fileID)
		p("===================")
		return true
	} else {
		return false
	}
}

//*
//*=======================================
func isIt_TTPacketsMissed(fileID string) bool {
	p := fmt.Println
	p("====================================================")
	p("====================================================")
	t := strings.Split(fileID, ".")
	var tt bool
	if len(t[0]) == len("PacketsMissed") {
		tt = true
	} else {
		tt = false
	}
	if (strings.Contains(fileID, "PacketsMissed")) && tt {
		p("=======  it is an PacketsMissedForRecording  ====")
		p(fileID)
		p("===================")
		return true
	} else {
		return false
	}
}

//*
//*=======================================
func isIt_TTPacketsMissedForRecording(fileID string) bool {
	p := fmt.Println
	p("====PacketsMissedForRecording========")
	p("====================================================")
	t := strings.Split(fileID, ".")
	var tt bool
	if len(t[0]) == len("PacketsMissedForRecording") {
		tt = true
	} else {
		tt = false
		return false
	}

	if (strings.Contains(fileID, "PacketsMissedForRecording")) && tt {
		p("=======  it is an PacketsMissedForRecording  ====")
		p(fileID)
		p("===================")
		return true
	} else {
		return false
	}
}

//*
//*=======================================
func isIt_TTRecordingError(fileID string) bool {
	p := fmt.Println
	if strings.Contains(fileID, "RecordingError") {
		p("=======  it is an RecordingError  ====")
		p(fileID)
		p("===================")
		return true
	} else {
		return false
	}
}
