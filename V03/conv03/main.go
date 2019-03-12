package main

import (
	"fmt"
)

var TempOutS01 []OutS01AgentCaptureUnreachable
var TempOutS02 []OutS02OneWayAudioRecording
var TempOutS03 []OutS03PacketsMissed
var TempOutS04 []OutS04PacketsMissedForRecording
var TempOutS05 []OutS05RecordingError
var TempAllDir []Fdir //  the entire Directory

//  This section for the Spread Sheet Headers
//Col := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW"}
var H00 []string
var H01 []string
var H02 []string
var H03 []string
var H04 []string
var H05 []string

func main() {
	p := fmt.Println

	p("======Step 1.0 ================")
	BuildTempAllDir()

	//  will do the TempOut systems
	BuildOut00()

	//  Buulds headers for Spread Sheet
	buildHeaders()

	p("======Step 5.0 ================")
	PPFDIR()

	PPT01()
	PPT02()
	PPT03()
	PPT04()
	PPT05()

	p("======Step 10.0 ================")
	SaveToDBALL()

	p("======Step 20.0 ================")
	//buildHeaders()
	//p(H00)
}
