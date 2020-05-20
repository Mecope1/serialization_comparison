package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v4"
	"encoding/json"
	//"encoding/gob"
	//"github.com/Mecope1/tin-can-communicator/binprot"
	)

// Separate fns for encoding and decoding as decoding tends to take a longer time.

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

	fmt.Println("To see the results of the benchmarks, try type  \"go test -bench=.\" in the project directory.")
}

func msgpackEnc(vNum int, sNum float64, msg string) []byte {

	bSli, err := msgpack.Marshal(&myStruct{VersionNum: vNum, SampleNum: sNum, Message:msg})
	if err != nil {
		panic(err)
	}

	return bSli
}

func msgpackDec(bSli []byte) myStruct {

	var outputStruct myStruct

	err := msgpack.Unmarshal(bSli, &outputStruct)
	if err != nil {
		panic(err)
	}

	return outputStruct
}

func jsonEnc(vNum int, sNum float64, msg string) []byte {

	bSli, err := json.Marshal(myStruct{VersionNum:vNum, SampleNum:sNum, Message:msg})
	if err != nil {
		panic(err)
	}

	return bSli
}

func jsonDec(bSli []byte) myStruct {
	var outputStruct myStruct

	err := json.Unmarshal(bSli, &outputStruct)
	if err != nil {
		panic(err)
	}

	return outputStruct
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


