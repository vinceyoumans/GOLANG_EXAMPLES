package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

func BuildOut00() {
	p := fmt.Println
	//====================================
	p("==========   seems like BuildAll is working ======")
	for i, rec := range TempAllDir {
		p("=====================================")
		p(i)
		p(rec.ID)
		if rec.TTAgentCaptureUnreachable {
			p("====  AgentCaptureUnreachable")
			putR01AgentCaptureUnreachable(rec)
		}
		if rec.TTOneWayAudioRecording {
			p("====  OneWayAudioRecording")
			putR02OneWayAudioRecording(rec)
		}
		if rec.TTPacketsMissed {
			p("====  PacketsMissed")
			putR03PacketsMissed(rec)
		}
		if rec.TTPacketsMissedForRecording {
			p("====  PacketsMissedForRecording")
			putR04PacketsMissedForRecording(rec)
		}
		if rec.TTRecordingError {
			p("====  RecordingError")
			putR05RecordingError(rec)
		}
	}
}

//========================================
//========================================
func putR01AgentCaptureUnreachable(val Fdir) {
	p := fmt.Println
	var ST OutS01AgentCaptureUnreachable

	ST.ID = val.ID
	ST.AlarmName = "AgentCaptureUnreachable"
	ST.AlarminstanceID = val.Alarminstanceid
	ST.ComponentName = val.ComponentName
	ST.TriggerTime = val.EventTrigger
	//tTime := strings.Split(val.EventDateString, "T")
	ST.TSA01Date = val.EventDateString
	ST.TSA01Time = val.EventTimeString
	ST.TSUnix = val.EventTimeStampInt

	byteValue := val.RecString
	// byteValue := val.RecDetails
	var q A00StandardStruct //  this works for all files
	xml.Unmarshal(byteValue, &q)
	//q := val.RecDetails

	p("=====   val.recDetails = ")
	p(q)

	ST.P00 = q.Paramaters.Parameter[0]
	ST.P01 = q.Paramaters.Parameter[1]
	ST.P02_IPAdd = q.Paramaters.Parameter[2]
	ST.P03_port = q.Paramaters.Parameter[3]

	ST.EventTP10 = val.EventTP10
	ST.EventTP20 = val.EventTP20
	ST.EventTP30 = val.EventTP30

	TempOutS01 = append(TempOutS01, ST)
	return
}

//========================================
//========================================
//========================================
func putR02OneWayAudioRecording(val Fdir) {
	p := fmt.Println
	var ST OutS02OneWayAudioRecording

	ST.ID = val.ID
	ST.AlarmName = "OneWayAudioRecording"
	ST.AlarminstanceID = val.Alarminstanceid
	ST.ComponentName = val.ComponentName
	ST.TriggerTime = val.EventTrigger
	ST.TSA01Date = val.EventDateString
	ST.TSA01Time = val.EventTimeString
	ST.TSUnix = val.EventTimeStampInt

	byteValue := val.RecString
	var q A00StandardStruct //  this works for all files
	xml.Unmarshal(byteValue, &q)

	fmt.Print("=====   val.recDetails = ")
	p(q)

	ST.P00 = q.Paramaters.Parameter[0]
	ST.P01_SIP_ID = q.Paramaters.Parameter[1]
	ST.P02_Origin_Phone = q.Paramaters.Parameter[2]
	ST.P04_NoFrikClue = q.Paramaters.Parameter[3]

	t := strings.Split(q.Paramaters.Parameter[6], "-)")
	t1 := strings.Split(t[0], ":")
	t2 := strings.Split(t[1], ":")

	ST.P06_IP_String = q.Paramaters.Parameter[6]

	ST.P06A_IP_PORT = t[0]
	ST.P06A_IP = t1[0]
	ST.P06A_PORT = t1[1]

	ST.P06b_IP_PORT = t[1]
	ST.P06B_IP = t2[0]
	ST.P06B_PORT = t2[1]

	ST.EventTP10 = val.EventTP10
	ST.EventTP20 = val.EventTP20
	ST.EventTP30 = val.EventTP30

	TempOutS02 = append(TempOutS02, ST)
	return
}

//====================================================
//====================================================
//====================================================
func putR03PacketsMissed(val Fdir) {
	p := fmt.Println
	var ST OutS03PacketsMissed

	ST.ID = val.ID
	ST.AlarmName = "PacketsMissed"
	ST.AlarminstanceID = val.Alarminstanceid
	ST.ComponentName = val.ComponentName
	ST.TriggerTime = val.EventTrigger
	ST.TSA01Date = val.EventDateString
	ST.TSA01Time = val.EventTimeString
	ST.TSUnix = val.EventTimeStampInt

	byteValue := val.RecString
	var q A00StandardStruct //  this works for all files
	xml.Unmarshal(byteValue, &q)

	fmt.Print("=====   val.recDetails = ")
	p(q)

	ST.P00 = q.Paramaters.Parameter[0]
	ST.P01 = q.Paramaters.Parameter[1]
	ST.P02 = q.Paramaters.Parameter[2]
	ST.P03 = q.Paramaters.Parameter[3]

	ST.EventTP10 = val.EventTP10
	ST.EventTP20 = val.EventTP20
	ST.EventTP30 = val.EventTP30

	TempOutS03 = append(TempOutS03, ST)
	return
}

//====================================================
//====================================================
//====================================================
func putR04PacketsMissedForRecording(val Fdir) {
	//p := fmt.Println
	var ST OutS04PacketsMissedForRecording

	ST.ID = val.ID
	ST.AlarmName = "PacketsMissedForRecording"
	ST.AlarminstanceID = val.Alarminstanceid
	ST.ComponentName = val.ComponentName
	ST.TriggerTime = val.EventTrigger
	ST.TSA01Date = val.EventDateString
	ST.TSA01Time = val.EventTimeString
	ST.TSUnix = val.EventTimeStampInt

	ST.P00 = val.RecDetails.Paramaters.Parameter[0]
	ST.P01_IP_PORT = val.RecDetails.Paramaters.Parameter[1]
	a := strings.Split(val.RecDetails.Paramaters.Parameter[1], ":")
	ST.P01_IP = a[0]
	ST.P01_PORT = a[1]

	ST.P02_IP_PORT = val.RecDetails.Paramaters.Parameter[2]
	b := strings.Split(val.RecDetails.Paramaters.Parameter[2], ":")
	ST.P02_IP = b[0]
	ST.P02_PORT = b[1]

	ST.P03 = val.RecDetails.Paramaters.Parameter[3]
	ST.P04 = val.RecDetails.Paramaters.Parameter[4]

	ST.EventTP10 = val.EventTP10
	ST.EventTP20 = val.EventTP20
	ST.EventTP30 = val.EventTP30

	TempOutS04 = append(TempOutS04, ST)
	return
}

//===========================================
//===========================================
//===========================================
func putR05RecordingError(val Fdir) {
	p := fmt.Println
	var ST OutS05RecordingError

	ST.ID = val.ID
	ST.AlarmName = "RecordingError"
	ST.AlarminstanceID = val.Alarminstanceid
	ST.ComponentName = val.ComponentName
	ST.TriggerTime = val.EventTrigger
	ST.TSA01Date = val.EventDateString
	ST.TSA01Time = val.EventTimeString
	ST.TSUnix = val.EventTimeStampInt

	byteValue := val.RecString
	var q A00StandardStruct //  this works for all files
	xml.Unmarshal(byteValue, &q)

	fmt.Print("=====   val.recDetails = ")
	p(q)

	ST.P00 = q.Paramaters.Parameter[0]
	ST.P01 = q.Paramaters.Parameter[1]
	ST.P02 = q.Paramaters.Parameter[2]
	ST.P03 = q.Paramaters.Parameter[3]

	// t := strings.Split(q.Paramaters.Parameter[3], "'&apos;'")
	// ST.P04 = t[1]
	// ST.P05 = t[3]

	t := strings.Split(q.Paramaters.Parameter[3], "'")
	ST.P04 = t[1]
	ST.P05 = t[3]

	ST.EventTP10 = val.EventTP10
	ST.EventTP20 = val.EventTP20
	ST.EventTP30 = val.EventTP30

	p("====   ttt  ====")
	p(t)
	TempOutS05 = append(TempOutS05, ST)
	return
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
