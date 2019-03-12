package main

import (
	"encoding/xml"
)

const TP10 = (10 * 60)
const TP20 = (20 * 60)
const TP30 = (30 * 60)

//====================================
//====================================
//====================================
type Fdir struct {
	ID                          string            `ss:"ID"`
	FName                       string            `ss:"FNAME"`
	AlarmType                   string            `ss:"AlarmType"`
	Alarminstanceid             string            `ss:"AlarmInstanceID"`
	ComponentName               string            `ss:"ComponentName"`
	EventTrigger                string            `ss:"EventTrigger"`
	EventDateString             string            `ss:"EventDate_01"`
	EventTimeString             string            `ss:"EventTime_01"`
	EventYEAR                   int               `ss:"YEAR"`
	EventMonth                  int               `ss:"MONTH"`
	EventDay                    int               `ss:"DAY"`
	EventHour                   int               `ss:"HOUR"`
	EventMin                    int               `ss:"MIN"`
	EventSecond                 int               `ss:"SECONDS"`
	EventTSInSeconds            int               `ss:"TS_IN_SECONDS"`
	EventTP10                   int               `ss:"EVENT_TP10"` //day is chopped into 10 minute increments.
	EventTP20                   int               `ss:"EVENT_TP20"` //day chopped into 20 minutes increments
	EventTP30                   int               `ss:"EVENT_TP30"` //day choppedinto 30 minute increments
	EventTimeStampInt           int64             `ss:"EVENT_TS_int64"`
	TTAgentCaptureUnreachable   bool              `ss:"TTAgentCaptureUnreachable"`
	TTOneWayAudioRecording      bool              `ss:"TTOneWayAudioRecording"`
	TTPacketsMissed             bool              `ss:"TTPacketsMissed"`
	TTPacketsMissedForRecording bool              `ss:"TTPacketsMissedForRecording"`
	TTRecordingError            bool              `ss:"TTRecordingError"`
	RecString                   []byte            `ss:"RecString"`
	RecDetails                  A00StandardStruct `ss:"RecDetails"` // the actual Record
}

//====================================
type A00StandardStruct struct {
	XMLName         xml.Name      `xml:"alarm" json:"-" ss:"XMLName"`
	AlarmName       string        `xml:"alarmname" json:"alarmname" ss:"AlarmName"`
	AlarmInstanceID string        `xml:"alarminstanceid" json:"alarminstanceid" ss:"AlarmInstanceID"`
	ComponentName   string        `xml:"componentname" json:"componentname" ss:"ComponentName"`
	TriggerTime     string        `xml:"triggertime" json:"triggertime" ss:"TriggerTime"`
	PriorityValue   string        `xml:"priorityvalue" json:"priorityvalue" ss:"PriorityValue"`
	Paramaters      ParametersS00 `xml:"parameters" json:"parameters" ss:"Paramaters"`
}

type ParametersS00 struct {
	XMLName   xml.Name `xml:"parameters" json:"parameters"  ss:"XMLName"`
	Parameter []string `xml:"parameter" json:"parameter"  ss:"Parameter"`
}

//================================================
//================================================
// <parameter>goxsa5187</parameter>
// <parameter>lt320846.fplu.fpl.com</parameter>
// <parameter>10.252.165.74</parameter>
// <parameter>4001</parameter>
//================================================
type OutS01AgentCaptureUnreachable struct {
	ID              string `ss:"ID"`
	AlarmName       string `ss:"AlarmName"`
	AlarminstanceID string `ss:"AlarmInstanceID"` //Index key
	ComponentName   string `ss:"ComponentName"`
	TriggerTime     string `ss:"TriggerTime"` // original TriggerTIme
	TSA01Date       string `ss:"TS01_DATE"`
	TSA01Time       string `ss:"TS01_TIME"`
	TSUnix          int64  `ss:"TIME_int64"`
	EventTP10       int    `ss:"EVENT_TP10"` //day is chopped into 10 minute increments.
	EventTP20       int    `ss:"EVENT_TP20"` //day chopped into 20 minutes increments
	EventTP30       int    `ss:"EVENT_TP30"` //day choppedinto 30 minute increments
	P00             string `ss:"P00_NoClue_P01"`
	P01             string `ss:"P02_DOMAIN_URL"`
	P02_IPAdd       string `ss:"P03_IP"`
	P03_port        string `ss:"P04_PORT?"` //guessing its a port
}

//================================================
//================================================
// <parameter>goxsa5187</parameter>
// <parameter>SIP/48671621@SEPBCF1F2E99053</parameter>
// <parameter>5553050004</parameter>
// <parameter>GO Phones</parameter>
// <parameter>588010008175886</parameter>
// <parameter>10.143.145.41</parameter>
// <parameter>10.143.145.41:26188 -) 10.147.16.201:52438;</parameter>
//================================================
type OutS02OneWayAudioRecording struct {
	ID               string `ss:"ID"`
	AlarmName        string `ss:"AlarmName"`
	AlarminstanceID  string `ss:"AlarmInstanceID"` //Index key
	ComponentName    string `ss:"ComponentName"`
	TriggerTime      string `ss:"TriggerTime"` // original TriggerTIme
	TSA01Date        string `ss:"TS01_DATE"`
	TSA01Time        string `ss:"TS01_TIME"`
	TSUnix           int64  `ss:"TIME_int64"`
	EventTP10        int    `ss:"EVENT_10"`   //day is chopped into 10 minute increments.
	EventTP20        int    `ss:"EVENT_TP20"` //day chopped into 20 minutes increments
	EventTP30        int    `ss:"EVENT_TP30"` //day choppedinto 30 minute increments
	P00              string `ss:"P00"`
	P01_SIP_ID       string `ss:"P01_SIP_ID"`
	P02_Origin_Phone string `ss:"P02-OriginPhone"`
	P04_NoFrikClue   string `ss:"P04_NoCLue"`
	P06_IP_String    string `ss:"P06_IP_STRING"`
	P06A_IP_PORT     string `ss:"P06_A_IP_PORT"`
	P06b_IP_PORT     string `ss:"P06_B_IP_PORT"`
	P06A_IP          string `ss:"P06_A_IP"`
	P06A_PORT        string `ss:"P06_A_Port"`
	P06B_IP          string `ss:"P06_B_IP"`
	P06B_PORT        string `ss:"P06_B_Port"`
}

//================================================
//================================================
// <parameter>goxsa5187</parameter>
// <parameter>12361</parameter>
// <parameter>RTP</parameter>
// <parameter>24</parameter>
//================================================
//================================================
type OutS03PacketsMissed struct {
	ID              string `ss:"ID"`
	AlarmName       string `ss:"AlarmName"`
	AlarminstanceID string `ss:"AlarmInstanceID"` //Index key
	ComponentName   string `ss:"ComponentName"`
	TriggerTime     string `ss:"TriggerTime"` // original TriggerTIme
	TSA01Date       string `ss:"TS01_DATE"`
	TSA01Time       string `ss:"TS01_TIME"`
	TSUnix          int64  `ss:"TIME_int64"`
	EventTP10       int    `ss:"EVENT_10"` //day is chopped into 10 minute increments.
	EventTP20       int    `ss:"EVENT_20"` //day chopped into 20 minutes increments
	EventTP30       int    `ss:"EVENT_30"` //day choppedinto 30 minute increments
	P00             string `ss:"P00"`
	P01             string `ss:"P00"`
	P02             string `ss:"P00"`
	P03             string `ss:"P00"`
}

//================================================
// <parameter>goxsa5187</parameter>
// <parameter>10.131.50.236:17624</parameter>
// <parameter>10.147.16.201:52716</parameter>
// <parameter>7</parameter>
// <parameter>588010008107110</parameter>
//================================================
type OutS04PacketsMissedForRecording struct {
	ID              string `ss:"ID"`
	AlarmName       string `ss:"AlarmName"`
	AlarminstanceID string `ss:"AlarmInstanceID"` //Index key
	ComponentName   string `ss:"ComponentName"`
	TriggerTime     string `ss:"TriggerTime"` // original TriggerTIme
	TSA01Date       string `ss:"TS01_DATE"`
	TSA01Time       string `ss:"TS01_TIME"`
	TSUnix          int64  `ss:"TIME_int64"`
	EventTP10       int    `ss:"EVENT_10"` //day is chopped into 10 minute increments.
	EventTP20       int    `ss:"EVENT_20"` //day chopped into 20 minutes increments
	EventTP30       int    `ss:"EVENT_30"` //day choppedinto 30 minute increments
	P00             string `ss:"P00"`
	P01_IP_PORT     string `ss:"P01_IP_PORT"`
	P01_IP          string `ss:"P01_IP"`
	P01_PORT        string `ss:"P01_PORT"`
	P02_IP_PORT     string `ss:"P02_IP_PORT"`
	P02_IP          string `ss:"P02_IP"`
	P02_PORT        string `ss:"P02_PORT"`
	P03             string `ss:"P03"`
	P04             string `ss:"P04"`
}

//================================================
//<parameters>
// <parameter>588010008139588</parameter>
// <parameter>409</parameter>
// <parameter>goxsa5187</parameter>
// <parameter>Connection to PC &apos;ELP2FPL88221.gcserv.com&apos; is broken. Connection address &apos;10.116.121.66&apos; was used to connect to the PC. Ensure that PC is on the network and reachable. Also make sure that Screen Capture Service is running on that PC.</parameter>
// </parameters>
//================================================
type OutS05RecordingError struct {
	ID              string `ss:"ID"`              //AlarmInstanceID
	AlarmName       string `ss:"AlarmName"`       //recordingError
	AlarminstanceID string `ss:"AlarmInstanceID"` //Index key
	ComponentName   string `ss:"ComponentName"`   //WGUIMediaChannel
	TriggerTime     string `ss:"TriggerTime"`     // original TriggerTIme
	TSA01Date       string `ss:"TS01_DATE"`       //2018-05-24
	TSA01Time       string `ss:"TS01_TIME"`       //20:40:44.302Z
	TSUnix          int64  `ss:"TIME_int64"`
	EventTP10       int    `ss:"EVENT_10"` //day is chopped into 10 minute increments.
	EventTP20       int    `ss:"EVENT_20"` //day chopped into 20 minutes increments
	EventTP30       int    `ss:"EVENT_30"` //day choppedinto 30 minute increments
	P00             string `ss:"P00"`
	P01             string `ss:"P01"`
	P02             string `ss:"P02"`
	P03             string `ss:"P03"`
	P04             string `ss:"P04"` //ELP2FPL88221.gcserv.com
	P05             string `ss:"P05"` //10.116.121.66
}
