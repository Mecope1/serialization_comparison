package main

import (
	"encoding/json"
	"fmt"
	bp "github.com/Mecope1/tin-can-communicator/binprot"
	// tccserver "github.com/Mecope1/tin-can-communicator/server"
	"github.com/vmihailenco/msgpack"
	"net"
)

// Separate fns for encoding and decoding as decoding tends to take a longer time.
// This difference can be quite substantial.

type myStruct struct {
	VersionNum int
	SampleNum float64
	Message string
}


func main() {

	//versionNum := 1
	//sampleNum := 3.14
	//shortMessage := "Hello, World!"

	// contains 457 words or 3147 Bytes
	//longMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean vitae nisl in felis suscipit accumsan quis in magna. Etiam iaculis odio bibendum, suscipit ipsum vitae, suscipit lorem. Mauris interdum mi ultrices efficitur interdum. Nulla quis purus iaculis leo commodo fermentum. In tincidunt consectetur tristique. Aliquam fermentum sollicitudin ultrices. Donec eu magna vel magna viverra consectetur in at sapien. Ut odio tellus, maximus non porttitor eu, consectetur quis nunc. Ut efficitur, urna in molestie scelerisque, purus nulla sagittis nibh, non pharetra ipsum elit vestibulum dui. Ut sed rhoncus mi. Aenean pretium lacus vel massa rhoncus, at finibus massa sodales. Praesent turpis nulla, mollis eget diam id, tempor pulvinar odio. Pellentesque nisi nunc, placerat sed sapien at, porta cursus purus. Sed nec quam est. Suspendisse vitae mauris risus. Vestibulum facilisis tempor orci sed pharetra. Nam id felis id lectus pharetra lacinia. Nullam ac sapien a nisi ornare sollicitudin a sit amet elit. Vivamus feugiat dui sit amet arcu malesuada rhoncus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Phasellus non scelerisque dui. Nulla gravida id neque nec dapibus. Cras in mattis arcu, sit amet facilisis ipsum. Duis quis quam est. Nunc mollis eleifend nisi, in tristique nibh. Sed pharetra nisl id eleifend condimentum. Phasellus scelerisque erat dolor, sed pretium risus vestibulum sit amet. Cras vehicula neque nisl, ornare consectetur nulla pharetra eu. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Quisque consectetur accumsan eleifend. Mauris in neque tellus. Mauris at urna ipsum. Aenean vitae elementum magna. Proin id porta magna. Vivamus maximus efficitur malesuada. Etiam vel blandit tellus. Aliquam mattis lectus non diam aliquam luctus. Fusce condimentum mauris at sem posuere posuere. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque vehicula non nisl ut volutpat. Etiam bibendum, risus et tempus placerat, augue enim sagittis ligula, vitae posuere eros neque ornare risus. Vestibulum dignissim rhoncus sapien, in lobortis nulla pretium dapibus. Mauris laoreet ipsum nec augue posuere vestibulum ut sed nisi. Nam in aliquet odio, at molestie leo. Duis interdum interdum iaculis. Nullam ac feugiat risus, a posuere nulla. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus turpis libero, dignissim eu felis eu, fringilla mollis nibh. Curabitur non condimentum augue, sed sagittis nisl. Phasellus ante tellus, sagittis ac eros egestas, pretium pharetra dolor.	Aliquam erat volutpat. In et consequat quam. Phasellus dapibus mollis lacinia. Aenean dignissim viverra ante nec tristique. Maecenas massa libero, consectetur quis metus sed, luctus pellentesque leo. Curabitur sollicitudin interdum lectus, non semper nulla sodales iaculis. Nullam at mollis nisl. Nunc quis malesuada erat. In iaculis sed erat eu pellentesque. In tincidunt ut erat vel consectetur. Donec efficitur ipsum quis massa pellentesque sagittis. Quisque suscipit massa at facilisis placerat. Ut venenatis ultricies massa et tempor. "



	// typical max message length

	var msgType uint = 0x8B
	chatMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque finibus, libero id iaculis eu."
	user := "username1234"
	var roomID uint = 1

	tlvEnc(msgType, chatMessage, user, roomID)

	var encodedChatMessage = []byte {139, 1, 1, 255, 0, 12, 117, 115, 101, 114, 110, 97, 109, 101, 49, 50, 51, 52, 0, 0, 0, 100, 76, 111, 114, 101, 109, 32, 105, 112, 115, 117, 109, 32, 100, 111, 108, 111, 114, 32, 115, 105, 116, 32, 97, 109, 101, 116, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 100, 105, 112, 105, 115, 99, 105, 110, 103, 32, 101, 108, 105, 116, 46, 32, 80, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 102, 105, 110, 105, 98, 117, 115, 44, 32, 108, 105, 98, 101, 114, 111, 32, 105, 100, 32, 105, 97, 99, 117, 108, 105, 115, 32, 101, 117, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	tlvDec(encodedChatMessage)


	var msgPackChatEnc = []byte {134, 164, 84, 121, 112, 101, 207, 0, 0, 0, 0, 0, 0, 0, 139, 167, 86, 101, 114, 115, 105, 111, 110, 207, 0, 0, 0, 0, 0, 0, 0, 1, 166, 82, 111, 111, 109, 73, 68, 207, 0, 0, 0, 0, 0, 0, 0, 1, 169, 67, 104, 97, 116, 116, 101, 114, 73, 68, 207, 0, 0, 0, 0, 0, 0, 0, 255, 171, 67, 104, 97, 116, 116, 101, 114, 78, 97, 109, 101, 196, 12, 117, 115, 101, 114, 110, 97, 109, 101, 49, 50, 51, 52, 167, 80, 97, 121, 108, 111, 97, 100, 196, 100, 76, 111, 114, 101, 109, 32, 105, 112, 115, 117, 109, 32, 100, 111, 108, 111, 114, 32, 115, 105, 116, 32, 97, 109, 101, 116, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 100, 105, 112, 105, 115, 99, 105, 110, 103, 32, 101, 108, 105, 116, 46, 32, 80, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 102, 105, 110, 105, 98, 117, 115, 44, 32, 108, 105, 98, 101, 114, 111, 32, 105, 100, 32, 105, 97, 99, 117, 108, 105, 115, 32, 101, 117, 46}
	msgpackEnc_Chat(msgType, chatMessage, user, roomID)
	msgpackDec_Chat(msgPackChatEnc)



	var jsonChatDec = []byte {123, 34, 84, 121, 112, 101, 34, 58, 49, 51, 57, 44, 34, 86, 101, 114, 115, 105, 111, 110, 34, 58, 49, 44, 34, 82, 111, 111, 109, 73, 68, 34, 58, 49, 44, 34, 67, 104, 97, 116, 116, 101, 114, 73, 68, 34, 58, 50, 53, 53, 44, 34, 67, 104, 97, 116, 116, 101, 114, 78, 97, 109, 101, 34, 58, 34, 100, 88, 78, 108, 99, 109, 53, 104, 98, 87, 85, 120, 77, 106, 77, 48, 34, 44, 34, 80, 97, 121, 108, 111, 97, 100, 34, 58, 34, 84, 71, 57, 121, 90, 87, 48, 103, 97, 88, 66, 122, 100, 87, 48, 103, 90, 71, 57, 115, 98, 51, 73, 103, 99, 50, 108, 48, 73, 71, 70, 116, 90, 88, 81, 115, 73, 71, 78, 118, 98, 110, 78, 108, 89, 51, 82, 108, 100, 72, 86, 121, 73, 71, 70, 107, 97, 88, 66, 112, 99, 50, 78, 112, 98, 109, 99, 103, 90, 87, 120, 112, 100, 67, 52, 103, 85, 71, 86, 115, 98, 71, 86, 117, 100, 71, 86, 122, 99, 88, 86, 108, 73, 71, 90, 112, 98, 109, 108, 105, 100, 88, 77, 115, 73, 71, 120, 112, 89, 109, 86, 121, 98, 121, 66, 112, 90, 67, 66, 112, 89, 87, 78, 49, 98, 71, 108, 122, 73, 71, 86, 49, 76, 103, 61, 61, 34, 125}
	jsonEnc_Chat(msgType, chatMessage, user, roomID)
	jsonDec_Chat(jsonChatDec)

	fmt.Println("To see the results of the benchmarks, try type  \"go test -bench=.\" in the project directory.")
}

func msgpackEnc(vNum int, sNum float64, msg string)  {

	_, err := msgpack.Marshal(myStruct{VersionNum: vNum, SampleNum: sNum, Message:msg})

	if err != nil {
		panic(err)
	}

	//return bSli
}

func msgpackDec(bSli []byte) {

	var outputStruct myStruct

	err := msgpack.Unmarshal(bSli, &outputStruct)

	if err != nil {
		panic(err)
	}

	//return outputStruct
}

func jsonEnc(vNum int, sNum float64, msg string) {

	_, err := json.Marshal(myStruct{VersionNum:vNum, SampleNum:sNum, Message:msg})

	if err != nil {
		panic(err)
	}

	//return bSli
}

func jsonDec(bSli []byte) {
	var outputStruct myStruct

	err := json.Unmarshal(bSli, &outputStruct)

	if err != nil {
		panic(err)
	}

	//return outputStruct
}

//func gobEnc(vNum int, sNum float64, msg string) {
//
//	var byteBuf bytes.Buffer
//
//	enc := gob.NewEncoder(&byteBuf)
//	dec := gob.NewDecoder(&byteBuf)
//
//	err := enc.Encode(myStruct{VersionNum:vNum, SampleNum:sNum, Message:msg})
//	if err !=nil {
//		panic(err)
//	}
//
//	var outputStruct myStruct
//	err = dec.Decode(&outputStruct)
//	fmt.Println("gob enc output: ", byteBuf)
//	fmt.Println("gob dec output: ", outputStruct)
//}
//
//
//func gobDec(bSli []byte) {
//
//}




type Client struct {
	socket net.Conn
	data chan []byte
	chatterName string
	roomID uint
}


type Record struct {
	Type uint
	Version uint
	RoomID uint
	ChatterID uint
	ChatterName []byte
	Payload []byte
}

//sampleRecord := &Record{
//Type:        msgType,
//Version:     1,
//RoomID:      roomNum,
//ChatterID:   255,
//ChatterName: []byte(usrName),
//Payload:     []byte(msg),
//}




func tlvEnc (msgType uint, msg string, usrName string, roomID uint) {
	_, err := bp.EncodeMsg(msgType, msg, usrName, roomID)

	if err != nil {
		panic(err)
	}

	//bSli := make([]byte, 256)
	//
	//if (err == nil) {
	//	buf.Read(bSli)
	//	fmt.Println(bSli)
	//}
	//ifmt.Println("Error during message construction: ", err.Error())
	//	//} else {
	//	//	_, err := client.socket.Write(buf.Bytes())
	//	//	if err != nil {
	//	//		client.socket.Close()
	//	//		panic(err)
	//	//	}
	//	//}f err != nil {
	//

}

func tlvDec (msgBytes []byte) {
	//msgBytes := make([]byte, 256)
	_, err := bp.DecodeMsg(msgBytes)

	if err != nil {
		panic(err)
	}

	//else {
	//
	//	fmt.Println("Name: " + string(msg.ChatterName))
	//	fmt.Println("Payload: " + string(msg.Payload))
	//	fmt.Println(msg.Version)
	//	fmt.Println("ChatterID: " + string(msg.ChatterID))
	//	fmt.Println(msg.ChatterID)
	//	fmt.Println(msg.Type)
	//	fmt.Println(msg.RoomID)
	//
	//}
}




//sampleRecord := &Record{
//Type:        msgType,
//Version:     1,
//RoomID:      roomNum,
//ChatterID:   255,
//ChatterName: []byte(usrName),
//Payload:     []byte(msg),
//}


//var msgType uint = 0x8B
//chatMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque finibus, libero id iaculis eu."
//user := "username1234"
//var roomID uint = 1


func msgpackEnc_Chat(msgType uint, msg string, usrName string, roomID uint)  {

	var vNum uint = 1
	var chatterID uint = 255

	_, err := msgpack.Marshal(Record{
									Type: msgType,
									Version: vNum,
									RoomID: roomID,
									ChatterID: chatterID,
									ChatterName: []byte(usrName),
									Payload: []byte(msg),
									})
	if err != nil {
		panic(err)
	}

	//else {
	//	fmt.Println("Packed Message: ")
	//	fmt.Println(packedMSG)
	//}

	//return bSli
}

func msgpackDec_Chat(bSli []byte) {

	var outputRecord Record

	err := msgpack.Unmarshal(bSli, &outputRecord)

	if err != nil {
		panic(err)
	}

	//else {
	//	fmt.Println("Name: " + string(outputRecord.ChatterName))
	//	fmt.Println("Payload: " + string(outputRecord.Payload))
	//	fmt.Println(outputRecord.Version)
	//	fmt.Println("ChatterID: " + string(outputRecord.ChatterID))
	//	fmt.Println(outputRecord.ChatterID)
	//	fmt.Println(outputRecord.Type)
	//	fmt.Println(outputRecord.RoomID)
	//}

	//return outputStruct
}
















func jsonEnc_Chat(msgType uint, msg string, usrName string, roomID uint) {


	var vNum uint = 1
	var chatterID uint = 255

	_, err := json.Marshal(Record{
		Type: msgType,
		Version: vNum,
		RoomID: roomID,
		ChatterID: chatterID,
		ChatterName: []byte(usrName),
		Payload: []byte(msg),
	})

	if err != nil {
		panic(err)
	}

	//else {
	//	fmt.Println("JSON MESSAGE: ")
	//	fmt.Println(jsonMSG)
	//}

	//return bSli
}

func jsonDec_Chat(bSli []byte) {
	var outputRecord Record

	err := json.Unmarshal(bSli, &outputRecord)

	if err != nil {
		panic(err)
	}

	//else {
	//	fmt.Println("Name: " + string(outputRecord.ChatterName))
	//	fmt.Println("Payload: " + string(outputRecord.Payload))
	//	fmt.Println(outputRecord.Version)
	//	fmt.Println("ChatterID: " + string(outputRecord.ChatterID))
	//	fmt.Println(outputRecord.ChatterID)
	//	fmt.Println(outputRecord.Type)
	//	fmt.Println(outputRecord.RoomID)
	//}
	//return outputStruct
}
