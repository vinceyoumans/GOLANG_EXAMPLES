package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//  will save structs to various formats

//===========
//Print to screen

func PPFDIR() {
	pp := fmt.Println
	p := fmt.Print

	pp("========  Starting FDIR Print ==========")
	for i, r := range TempAllDir {
		p("==========  fDIR  =======")
		pp(i)
		pp(r)
	}
	pp("========  Ending FDIR Print ==========")
	jsonData, err := json.Marshal(TempAllDir)

	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create("./KKOUT/FDIR.json")
	if err != nil {
		os.Mkdir("KKOUT", 0777)
		pp("=======   You have no KKOUT Dir")
		pp("===  I will create it......")
		pp("=== Please run the Program again")
		log.Fatalf("failed opening directory: %s", err)
		//panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}

func PPT01() {
	pp := fmt.Println
	p := fmt.Print

	pp("========  Starting PPT01 Print ==========")
	for i, r := range TempOutS01 {
		p("==========  PPT01  =======")
		pp(i)
		pp(r)
	}
	pp("========  Ending PPT01 Print ==========")
	jsonData, err := json.Marshal(TempOutS01)

	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create("./KKOUT/S01AgentCaptureUnreachable.json")
	if err != nil {
		os.Mkdir("KKOUT", 0777)
		pp("=======   You have no KKOUT Dir")
		pp("===  I will create it......")
		pp("=== Please run the Program again")
		log.Fatalf("failed opening directory: %s", err)
		//panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}

func PPT02() {
	pp := fmt.Println
	p := fmt.Print

	pp("========  Starting PPT02 Print ==========")
	for i, r := range TempOutS02 {
		p("==========  PPT02  =======")
		pp(i)
		pp(r)
	}
	pp("========  Ending PPT02 Print ==========")
	jsonData, err := json.Marshal(TempOutS02)

	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create("./KKOUT/S02OneWayAudioRecording.json")
	if err != nil {
		os.Mkdir("KKOUT", 0777)
		pp("=======   You have no KKOUT Dir")
		pp("===  I will create it......")
		pp("=== Please run the Program again")
		log.Fatalf("failed opening directory: %s", err)
		//panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}

func PPT03() {
	pp := fmt.Println
	p := fmt.Print

	pp("========  Starting PPT03 Print ==========")
	for i, r := range TempOutS03 {
		p("==========  PPT03  =======")
		pp(i)
		pp(r)
	}
	pp("========  Ending PPT03 Print ==========")
	jsonData, err := json.Marshal(TempOutS03)

	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create("./KKOUT/S03PacketsMissed.json")
	if err != nil {
		os.Mkdir("KKOUT", 0777)
		pp("=======   You have no KKOUT Dir")
		pp("===  I will create it......")
		pp("=== Please run the Program again")
		log.Fatalf("failed opening directory: %s", err)
		//panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}

func PPT04() {
	pp := fmt.Println
	p := fmt.Print

	pp("========  Starting PPT04 Print ==========")
	for i, r := range TempOutS04 {
		p("==========  PPT04  =======")
		pp(i)
		pp(r)
	}
	pp("========  Ending PPT04 Print ==========")
	jsonData, err := json.Marshal(TempOutS04)

	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create("./KKOUT/S04PacketsMissedForRecording.json")
	if err != nil {
		os.Mkdir("KKOUT", 0777)
		pp("=======   You have no KKOUT Dir")
		pp("===  I will create it......")
		pp("=== Please run the Program again")
		log.Fatalf("failed opening directory: %s", err)
		//panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}

func PPT05() {
	pp := fmt.Println
	p := fmt.Print

	pp("========  Starting PPT05 Print ==========")
	for i, r := range TempOutS05 {
		p("==========  PPT05  =======")
		pp(i)
		pp(r)
	}
	pp("========  Ending PPT05 Print ==========")
	jsonData, err := json.Marshal(TempOutS05)

	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create("./KKOUT/S05RecordingError.json")
	if err != nil {
		os.Mkdir("KKOUT", 0777)
		pp("=======   You have no KKOUT Dir")
		pp("===  I will create it......")
		pp("=== Please run the Program again")
		log.Fatalf("failed opening directory: %s", err)
		//panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}
