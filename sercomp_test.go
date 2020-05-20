package main

import "testing"

// Comparison for the encoding and decoding of several formats including Go's gob, a messagepack implementation, JSON,
// and the encoding scheme that's used in another project of mine.


// Tests on shorter data
func BenchmarkMsgpackEnc(b *testing.B) {
	versionNum := 1
	sampleNum := 3.14
	shortMessage := "Hello, World!"

	for n:=0; n<b.N; n++ {
		msgpackEnc(versionNum, sampleNum, shortMessage)
	}
}

func BenchmarkMsgpackDec(b *testing.B) {
	msgpackEncoded := []byte {131, 170, 86, 101, 114, 115, 105, 111, 110, 78, 117, 109, 211, 0, 0, 0, 0, 0, 0, 0, 1, 169, 83, 97, 109, 112, 108, 101, 78, 117, 109, 203, 64, 9, 30, 184, 81, 235, 133, 31, 167, 77, 101, 115, 115, 97, 103, 101, 173, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}

	for n:=0; n<b.N; n++ {
		msgpackDec(msgpackEncoded)
	}
}

func BenchmarkJsonEnc(b *testing.B) {
	versionNum := 1
	sampleNum := 3.14
	shortMessage := "Hello, World!"

	for n:=0; n<b.N; n++ {
		jsonEnc(versionNum, sampleNum, shortMessage)
	}
}

func BenchmarkJsonDec(b *testing.B) {
	jsonEncoded := []byte {123, 34, 86, 101, 114, 115, 105, 111, 110, 78, 117, 109, 34, 58, 49, 44, 34, 83, 97, 109, 112, 108, 101, 78, 117, 109, 34, 58, 51, 46, 49, 52, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33, 34, 125}

	for n:=0; n<b.N; n++ {
		jsonDec(jsonEncoded)
	}
}


// Tests involving larger data. The string message now is 457 words or 3147 Bytes
// This is ~240 times larger than the hello world message used in the previous tests.
func BenchmarkMsgpackEncLong(b *testing.B) {
	versionNum := 1
	sampleNum := 3.14
	longMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean vitae nisl in felis suscipit accumsan quis in magna. Etiam iaculis odio bibendum, suscipit ipsum vitae, suscipit lorem. Mauris interdum mi ultrices efficitur interdum. Nulla quis purus iaculis leo commodo fermentum. In tincidunt consectetur tristique. Aliquam fermentum sollicitudin ultrices. Donec eu magna vel magna viverra consectetur in at sapien. Ut odio tellus, maximus non porttitor eu, consectetur quis nunc. Ut efficitur, urna in molestie scelerisque, purus nulla sagittis nibh, non pharetra ipsum elit vestibulum dui. Ut sed rhoncus mi. Aenean pretium lacus vel massa rhoncus, at finibus massa sodales. Praesent turpis nulla, mollis eget diam id, tempor pulvinar odio. Pellentesque nisi nunc, placerat sed sapien at, porta cursus purus. Sed nec quam est. Suspendisse vitae mauris risus. Vestibulum facilisis tempor orci sed pharetra. Nam id felis id lectus pharetra lacinia. Nullam ac sapien a nisi ornare sollicitudin a sit amet elit. Vivamus feugiat dui sit amet arcu malesuada rhoncus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Phasellus non scelerisque dui. Nulla gravida id neque nec dapibus. Cras in mattis arcu, sit amet facilisis ipsum. Duis quis quam est. Nunc mollis eleifend nisi, in tristique nibh. Sed pharetra nisl id eleifend condimentum. Phasellus scelerisque erat dolor, sed pretium risus vestibulum sit amet. Cras vehicula neque nisl, ornare consectetur nulla pharetra eu. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Quisque consectetur accumsan eleifend. Mauris in neque tellus. Mauris at urna ipsum. Aenean vitae elementum magna. Proin id porta magna. Vivamus maximus efficitur malesuada. Etiam vel blandit tellus. Aliquam mattis lectus non diam aliquam luctus. Fusce condimentum mauris at sem posuere posuere. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque vehicula non nisl ut volutpat. Etiam bibendum, risus et tempus placerat, augue enim sagittis ligula, vitae posuere eros neque ornare risus. Vestibulum dignissim rhoncus sapien, in lobortis nulla pretium dapibus. Mauris laoreet ipsum nec augue posuere vestibulum ut sed nisi. Nam in aliquet odio, at molestie leo. Duis interdum interdum iaculis. Nullam ac feugiat risus, a posuere nulla. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus turpis libero, dignissim eu felis eu, fringilla mollis nibh. Curabitur non condimentum augue, sed sagittis nisl. Phasellus ante tellus, sagittis ac eros egestas, pretium pharetra dolor.	Aliquam erat volutpat. In et consequat quam. Phasellus dapibus mollis lacinia. Aenean dignissim viverra ante nec tristique. Maecenas massa libero, consectetur quis metus sed, luctus pellentesque leo. Curabitur sollicitudin interdum lectus, non semper nulla sodales iaculis. Nullam at mollis nisl. Nunc quis malesuada erat. In iaculis sed erat eu pellentesque. In tincidunt ut erat vel consectetur. Donec efficitur ipsum quis massa pellentesque sagittis. Quisque suscipit massa at facilisis placerat. Ut venenatis ultricies massa et tempor. "

	for n:=0; n<b.N; n++ {
		msgpackEnc(versionNum, sampleNum, longMessage)
	}
}

func BenchmarkMsgpackDecLong(b *testing.B) {
	longmsgpackEncoded := []byte {131, 170, 86, 101, 114, 115, 105, 111, 110, 78, 117, 109, 211, 0, 0, 0, 0, 0, 0, 0, 1, 169, 83, 97, 109, 112, 108, 101, 78, 117, 109, 203, 64, 9, 30, 184, 81, 235, 133, 31, 167, 77, 101, 115, 115, 97, 103, 101, 218, 12, 80, 76, 111, 114, 101, 109, 32, 105, 112, 115, 117, 109, 32, 100, 111, 108, 111, 114, 32, 115, 105, 116, 32, 97, 109, 101, 116, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 100, 105, 112, 105, 115, 99, 105, 110, 103, 32, 101, 108, 105, 116, 46, 32, 65, 101, 110, 101, 97, 110, 32, 118, 105, 116, 97, 101, 32, 110, 105, 115, 108, 32, 105, 110, 32, 102, 101, 108, 105, 115, 32, 115, 117, 115, 99, 105, 112, 105, 116, 32, 97, 99, 99, 117, 109, 115, 97, 110, 32, 113, 117, 105, 115, 32, 105, 110, 32, 109, 97, 103, 110, 97, 46, 32, 69, 116, 105, 97, 109, 32, 105, 97, 99, 117, 108, 105, 115, 32, 111, 100, 105, 111, 32, 98, 105, 98, 101, 110, 100, 117, 109, 44, 32, 115, 117, 115, 99, 105, 112, 105, 116, 32, 105, 112, 115, 117, 109, 32, 118, 105, 116, 97, 101, 44, 32, 115, 117, 115, 99, 105, 112, 105, 116, 32, 108, 111, 114, 101, 109, 46, 32, 77, 97, 117, 114, 105, 115, 32, 105, 110, 116, 101, 114, 100, 117, 109, 32, 109, 105, 32, 117, 108, 116, 114, 105, 99, 101, 115, 32, 101, 102, 102, 105, 99, 105, 116, 117, 114, 32, 105, 110, 116, 101, 114, 100, 117, 109, 46, 32, 78, 117, 108, 108, 97, 32, 113, 117, 105, 115, 32, 112, 117, 114, 117, 115, 32, 105, 97, 99, 117, 108, 105, 115, 32, 108, 101, 111, 32, 99, 111, 109, 109, 111, 100, 111, 32, 102, 101, 114, 109, 101, 110, 116, 117, 109, 46, 32, 73, 110, 32, 116, 105, 110, 99, 105, 100, 117, 110, 116, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 46, 32, 65, 108, 105, 113, 117, 97, 109, 32, 102, 101, 114, 109, 101, 110, 116, 117, 109, 32, 115, 111, 108, 108, 105, 99, 105, 116, 117, 100, 105, 110, 32, 117, 108, 116, 114, 105, 99, 101, 115, 46, 32, 68, 111, 110, 101, 99, 32, 101, 117, 32, 109, 97, 103, 110, 97, 32, 118, 101, 108, 32, 109, 97, 103, 110, 97, 32, 118, 105, 118, 101, 114, 114, 97, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 105, 110, 32, 97, 116, 32, 115, 97, 112, 105, 101, 110, 46, 32, 85, 116, 32, 111, 100, 105, 111, 32, 116, 101, 108, 108, 117, 115, 44, 32, 109, 97, 120, 105, 109, 117, 115, 32, 110, 111, 110, 32, 112, 111, 114, 116, 116, 105, 116, 111, 114, 32, 101, 117, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 113, 117, 105, 115, 32, 110, 117, 110, 99, 46, 32, 85, 116, 32, 101, 102, 102, 105, 99, 105, 116, 117, 114, 44, 32, 117, 114, 110, 97, 32, 105, 110, 32, 109, 111, 108, 101, 115, 116, 105, 101, 32, 115, 99, 101, 108, 101, 114, 105, 115, 113, 117, 101, 44, 32, 112, 117, 114, 117, 115, 32, 110, 117, 108, 108, 97, 32, 115, 97, 103, 105, 116, 116, 105, 115, 32, 110, 105, 98, 104, 44, 32, 110, 111, 110, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 105, 112, 115, 117, 109, 32, 101, 108, 105, 116, 32, 118, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 100, 117, 105, 46, 32, 85, 116, 32, 115, 101, 100, 32, 114, 104, 111, 110, 99, 117, 115, 32, 109, 105, 46, 32, 65, 101, 110, 101, 97, 110, 32, 112, 114, 101, 116, 105, 117, 109, 32, 108, 97, 99, 117, 115, 32, 118, 101, 108, 32, 109, 97, 115, 115, 97, 32, 114, 104, 111, 110, 99, 117, 115, 44, 32, 97, 116, 32, 102, 105, 110, 105, 98, 117, 115, 32, 109, 97, 115, 115, 97, 32, 115, 111, 100, 97, 108, 101, 115, 46, 32, 80, 114, 97, 101, 115, 101, 110, 116, 32, 116, 117, 114, 112, 105, 115, 32, 110, 117, 108, 108, 97, 44, 32, 109, 111, 108, 108, 105, 115, 32, 101, 103, 101, 116, 32, 100, 105, 97, 109, 32, 105, 100, 44, 32, 116, 101, 109, 112, 111, 114, 32, 112, 117, 108, 118, 105, 110, 97, 114, 32, 111, 100, 105, 111, 46, 32, 80, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 110, 105, 115, 105, 32, 110, 117, 110, 99, 44, 32, 112, 108, 97, 99, 101, 114, 97, 116, 32, 115, 101, 100, 32, 115, 97, 112, 105, 101, 110, 32, 97, 116, 44, 32, 112, 111, 114, 116, 97, 32, 99, 117, 114, 115, 117, 115, 32, 112, 117, 114, 117, 115, 46, 32, 83, 101, 100, 32, 110, 101, 99, 32, 113, 117, 97, 109, 32, 101, 115, 116, 46, 32, 83, 117, 115, 112, 101, 110, 100, 105, 115, 115, 101, 32, 118, 105, 116, 97, 101, 32, 109, 97, 117, 114, 105, 115, 32, 114, 105, 115, 117, 115, 46, 32, 86, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 102, 97, 99, 105, 108, 105, 115, 105, 115, 32, 116, 101, 109, 112, 111, 114, 32, 111, 114, 99, 105, 32, 115, 101, 100, 32, 112, 104, 97, 114, 101, 116, 114, 97, 46, 32, 78, 97, 109, 32, 105, 100, 32, 102, 101, 108, 105, 115, 32, 105, 100, 32, 108, 101, 99, 116, 117, 115, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 108, 97, 99, 105, 110, 105, 97, 46, 32, 78, 117, 108, 108, 97, 109, 32, 97, 99, 32, 115, 97, 112, 105, 101, 110, 32, 97, 32, 110, 105, 115, 105, 32, 111, 114, 110, 97, 114, 101, 32, 115, 111, 108, 108, 105, 99, 105, 116, 117, 100, 105, 110, 32, 97, 32, 115, 105, 116, 32, 97, 109, 101, 116, 32, 101, 108, 105, 116, 46, 32, 86, 105, 118, 97, 109, 117, 115, 32, 102, 101, 117, 103, 105, 97, 116, 32, 100, 117, 105, 32, 115, 105, 116, 32, 97, 109, 101, 116, 32, 97, 114, 99, 117, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 32, 114, 104, 111, 110, 99, 117, 115, 46, 32, 80, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 104, 97, 98, 105, 116, 97, 110, 116, 32, 109, 111, 114, 98, 105, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 32, 115, 101, 110, 101, 99, 116, 117, 115, 32, 101, 116, 32, 110, 101, 116, 117, 115, 32, 101, 116, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 32, 102, 97, 109, 101, 115, 32, 97, 99, 32, 116, 117, 114, 112, 105, 115, 32, 101, 103, 101, 115, 116, 97, 115, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 110, 111, 110, 32, 115, 99, 101, 108, 101, 114, 105, 115, 113, 117, 101, 32, 100, 117, 105, 46, 32, 78, 117, 108, 108, 97, 32, 103, 114, 97, 118, 105, 100, 97, 32, 105, 100, 32, 110, 101, 113, 117, 101, 32, 110, 101, 99, 32, 100, 97, 112, 105, 98, 117, 115, 46, 32, 67, 114, 97, 115, 32, 105, 110, 32, 109, 97, 116, 116, 105, 115, 32, 97, 114, 99, 117, 44, 32, 115, 105, 116, 32, 97, 109, 101, 116, 32, 102, 97, 99, 105, 108, 105, 115, 105, 115, 32, 105, 112, 115, 117, 109, 46, 32, 68, 117, 105, 115, 32, 113, 117, 105, 115, 32, 113, 117, 97, 109, 32, 101, 115, 116, 46, 32, 78, 117, 110, 99, 32, 109, 111, 108, 108, 105, 115, 32, 101, 108, 101, 105, 102, 101, 110, 100, 32, 110, 105, 115, 105, 44, 32, 105, 110, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 32, 110, 105, 98, 104, 46, 32, 83, 101, 100, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 110, 105, 115, 108, 32, 105, 100, 32, 101, 108, 101, 105, 102, 101, 110, 100, 32, 99, 111, 110, 100, 105, 109, 101, 110, 116, 117, 109, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 115, 99, 101, 108, 101, 114, 105, 115, 113, 117, 101, 32, 101, 114, 97, 116, 32, 100, 111, 108, 111, 114, 44, 32, 115, 101, 100, 32, 112, 114, 101, 116, 105, 117, 109, 32, 114, 105, 115, 117, 115, 32, 118, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 115, 105, 116, 32, 97, 109, 101, 116, 46, 32, 67, 114, 97, 115, 32, 118, 101, 104, 105, 99, 117, 108, 97, 32, 110, 101, 113, 117, 101, 32, 110, 105, 115, 108, 44, 32, 111, 114, 110, 97, 114, 101, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 110, 117, 108, 108, 97, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 101, 117, 46, 32, 80, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 104, 97, 98, 105, 116, 97, 110, 116, 32, 109, 111, 114, 98, 105, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 32, 115, 101, 110, 101, 99, 116, 117, 115, 32, 101, 116, 32, 110, 101, 116, 117, 115, 32, 101, 116, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 32, 102, 97, 109, 101, 115, 32, 97, 99, 32, 116, 117, 114, 112, 105, 115, 32, 101, 103, 101, 115, 116, 97, 115, 46, 32, 81, 117, 105, 115, 113, 117, 101, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 99, 99, 117, 109, 115, 97, 110, 32, 101, 108, 101, 105, 102, 101, 110, 100, 46, 32, 77, 97, 117, 114, 105, 115, 32, 105, 110, 32, 110, 101, 113, 117, 101, 32, 116, 101, 108, 108, 117, 115, 46, 32, 77, 97, 117, 114, 105, 115, 32, 97, 116, 32, 117, 114, 110, 97, 32, 105, 112, 115, 117, 109, 46, 32, 65, 101, 110, 101, 97, 110, 32, 118, 105, 116, 97, 101, 32, 101, 108, 101, 109, 101, 110, 116, 117, 109, 32, 109, 97, 103, 110, 97, 46, 32, 80, 114, 111, 105, 110, 32, 105, 100, 32, 112, 111, 114, 116, 97, 32, 109, 97, 103, 110, 97, 46, 32, 86, 105, 118, 97, 109, 117, 115, 32, 109, 97, 120, 105, 109, 117, 115, 32, 101, 102, 102, 105, 99, 105, 116, 117, 114, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 46, 32, 69, 116, 105, 97, 109, 32, 118, 101, 108, 32, 98, 108, 97, 110, 100, 105, 116, 32, 116, 101, 108, 108, 117, 115, 46, 32, 65, 108, 105, 113, 117, 97, 109, 32, 109, 97, 116, 116, 105, 115, 32, 108, 101, 99, 116, 117, 115, 32, 110, 111, 110, 32, 100, 105, 97, 109, 32, 97, 108, 105, 113, 117, 97, 109, 32, 108, 117, 99, 116, 117, 115, 46, 32, 70, 117, 115, 99, 101, 32, 99, 111, 110, 100, 105, 109, 101, 110, 116, 117, 109, 32, 109, 97, 117, 114, 105, 115, 32, 97, 116, 32, 115, 101, 109, 32, 112, 111, 115, 117, 101, 114, 101, 32, 112, 111, 115, 117, 101, 114, 101, 46, 32, 76, 111, 114, 101, 109, 32, 105, 112, 115, 117, 109, 32, 100, 111, 108, 111, 114, 32, 115, 105, 116, 32, 97, 109, 101, 116, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 100, 105, 112, 105, 115, 99, 105, 110, 103, 32, 101, 108, 105, 116, 46, 32, 81, 117, 105, 115, 113, 117, 101, 32, 118, 101, 104, 105, 99, 117, 108, 97, 32, 110, 111, 110, 32, 110, 105, 115, 108, 32, 117, 116, 32, 118, 111, 108, 117, 116, 112, 97, 116, 46, 32, 69, 116, 105, 97, 109, 32, 98, 105, 98, 101, 110, 100, 117, 109, 44, 32, 114, 105, 115, 117, 115, 32, 101, 116, 32, 116, 101, 109, 112, 117, 115, 32, 112, 108, 97, 99, 101, 114, 97, 116, 44, 32, 97, 117, 103, 117, 101, 32, 101, 110, 105, 109, 32, 115, 97, 103, 105, 116, 116, 105, 115, 32, 108, 105, 103, 117, 108, 97, 44, 32, 118, 105, 116, 97, 101, 32, 112, 111, 115, 117, 101, 114, 101, 32, 101, 114, 111, 115, 32, 110, 101, 113, 117, 101, 32, 111, 114, 110, 97, 114, 101, 32, 114, 105, 115, 117, 115, 46, 32, 86, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 100, 105, 103, 110, 105, 115, 115, 105, 109, 32, 114, 104, 111, 110, 99, 117, 115, 32, 115, 97, 112, 105, 101, 110, 44, 32, 105, 110, 32, 108, 111, 98, 111, 114, 116, 105, 115, 32, 110, 117, 108, 108, 97, 32, 112, 114, 101, 116, 105, 117, 109, 32, 100, 97, 112, 105, 98, 117, 115, 46, 32, 77, 97, 117, 114, 105, 115, 32, 108, 97, 111, 114, 101, 101, 116, 32, 105, 112, 115, 117, 109, 32, 110, 101, 99, 32, 97, 117, 103, 117, 101, 32, 112, 111, 115, 117, 101, 114, 101, 32, 118, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 117, 116, 32, 115, 101, 100, 32, 110, 105, 115, 105, 46, 32, 78, 97, 109, 32, 105, 110, 32, 97, 108, 105, 113, 117, 101, 116, 32, 111, 100, 105, 111, 44, 32, 97, 116, 32, 109, 111, 108, 101, 115, 116, 105, 101, 32, 108, 101, 111, 46, 32, 68, 117, 105, 115, 32, 105, 110, 116, 101, 114, 100, 117, 109, 32, 105, 110, 116, 101, 114, 100, 117, 109, 32, 105, 97, 99, 117, 108, 105, 115, 46, 32, 78, 117, 108, 108, 97, 109, 32, 97, 99, 32, 102, 101, 117, 103, 105, 97, 116, 32, 114, 105, 115, 117, 115, 44, 32, 97, 32, 112, 111, 115, 117, 101, 114, 101, 32, 110, 117, 108, 108, 97, 46, 32, 76, 111, 114, 101, 109, 32, 105, 112, 115, 117, 109, 32, 100, 111, 108, 111, 114, 32, 115, 105, 116, 32, 97, 109, 101, 116, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 100, 105, 112, 105, 115, 99, 105, 110, 103, 32, 101, 108, 105, 116, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 116, 117, 114, 112, 105, 115, 32, 108, 105, 98, 101, 114, 111, 44, 32, 100, 105, 103, 110, 105, 115, 115, 105, 109, 32, 101, 117, 32, 102, 101, 108, 105, 115, 32, 101, 117, 44, 32, 102, 114, 105, 110, 103, 105, 108, 108, 97, 32, 109, 111, 108, 108, 105, 115, 32, 110, 105, 98, 104, 46, 32, 67, 117, 114, 97, 98, 105, 116, 117, 114, 32, 110, 111, 110, 32, 99, 111, 110, 100, 105, 109, 101, 110, 116, 117, 109, 32, 97, 117, 103, 117, 101, 44, 32, 115, 101, 100, 32, 115, 97, 103, 105, 116, 116, 105, 115, 32, 110, 105, 115, 108, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 97, 110, 116, 101, 32, 116, 101, 108, 108, 117, 115, 44, 32, 115, 97, 103, 105, 116, 116, 105, 115, 32, 97, 99, 32, 101, 114, 111, 115, 32, 101, 103, 101, 115, 116, 97, 115, 44, 32, 112, 114, 101, 116, 105, 117, 109, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 100, 111, 108, 111, 114, 46, 9, 65, 108, 105, 113, 117, 97, 109, 32, 101, 114, 97, 116, 32, 118, 111, 108, 117, 116, 112, 97, 116, 46, 32, 73, 110, 32, 101, 116, 32, 99, 111, 110, 115, 101, 113, 117, 97, 116, 32, 113, 117, 97, 109, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 100, 97, 112, 105, 98, 117, 115, 32, 109, 111, 108, 108, 105, 115, 32, 108, 97, 99, 105, 110, 105, 97, 46, 32, 65, 101, 110, 101, 97, 110, 32, 100, 105, 103, 110, 105, 115, 115, 105, 109, 32, 118, 105, 118, 101, 114, 114, 97, 32, 97, 110, 116, 101, 32, 110, 101, 99, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 46, 32, 77, 97, 101, 99, 101, 110, 97, 115, 32, 109, 97, 115, 115, 97, 32, 108, 105, 98, 101, 114, 111, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 113, 117, 105, 115, 32, 109, 101, 116, 117, 115, 32, 115, 101, 100, 44, 32, 108, 117, 99, 116, 117, 115, 32, 112, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 108, 101, 111, 46, 32, 67, 117, 114, 97, 98, 105, 116, 117, 114, 32, 115, 111, 108, 108, 105, 99, 105, 116, 117, 100, 105, 110, 32, 105, 110, 116, 101, 114, 100, 117, 109, 32, 108, 101, 99, 116, 117, 115, 44, 32, 110, 111, 110, 32, 115, 101, 109, 112, 101, 114, 32, 110, 117, 108, 108, 97, 32, 115, 111, 100, 97, 108, 101, 115, 32, 105, 97, 99, 117, 108, 105, 115, 46, 32, 78, 117, 108, 108, 97, 109, 32, 97, 116, 32, 109, 111, 108, 108, 105, 115, 32, 110, 105, 115, 108, 46, 32, 78, 117, 110, 99, 32, 113, 117, 105, 115, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 32, 101, 114, 97, 116, 46, 32, 73, 110, 32, 105, 97, 99, 117, 108, 105, 115, 32, 115, 101, 100, 32, 101, 114, 97, 116, 32, 101, 117, 32, 112, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 46, 32, 73, 110, 32, 116, 105, 110, 99, 105, 100, 117, 110, 116, 32, 117, 116, 32, 101, 114, 97, 116, 32, 118, 101, 108, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 46, 32, 68, 111, 110, 101, 99, 32, 101, 102, 102, 105, 99, 105, 116, 117, 114, 32, 105, 112, 115, 117, 109, 32, 113, 117, 105, 115, 32, 109, 97, 115, 115, 97, 32, 112, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 115, 97, 103, 105, 116, 116, 105, 115, 46, 32, 81, 117, 105, 115, 113, 117, 101, 32, 115, 117, 115, 99, 105, 112, 105, 116, 32, 109, 97, 115, 115, 97, 32, 97, 116, 32, 102, 97, 99, 105, 108, 105, 115, 105, 115, 32, 112, 108, 97, 99, 101, 114, 97, 116, 46, 32, 85, 116, 32, 118, 101, 110, 101, 110, 97, 116, 105, 115, 32, 117, 108, 116, 114, 105, 99, 105, 101, 115, 32, 109, 97, 115, 115, 97, 32, 101, 116, 32, 116, 101, 109, 112, 111, 114, 46, 32}

	for n:=0; n<b.N; n++ {
		msgpackDec(longmsgpackEncoded)
	}
}

func BenchmarkJsonEncLong(b *testing.B) {
	versionNum := 1
	sampleNum := 3.14
	longMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean vitae nisl in felis suscipit accumsan quis in magna. Etiam iaculis odio bibendum, suscipit ipsum vitae, suscipit lorem. Mauris interdum mi ultrices efficitur interdum. Nulla quis purus iaculis leo commodo fermentum. In tincidunt consectetur tristique. Aliquam fermentum sollicitudin ultrices. Donec eu magna vel magna viverra consectetur in at sapien. Ut odio tellus, maximus non porttitor eu, consectetur quis nunc. Ut efficitur, urna in molestie scelerisque, purus nulla sagittis nibh, non pharetra ipsum elit vestibulum dui. Ut sed rhoncus mi. Aenean pretium lacus vel massa rhoncus, at finibus massa sodales. Praesent turpis nulla, mollis eget diam id, tempor pulvinar odio. Pellentesque nisi nunc, placerat sed sapien at, porta cursus purus. Sed nec quam est. Suspendisse vitae mauris risus. Vestibulum facilisis tempor orci sed pharetra. Nam id felis id lectus pharetra lacinia. Nullam ac sapien a nisi ornare sollicitudin a sit amet elit. Vivamus feugiat dui sit amet arcu malesuada rhoncus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Phasellus non scelerisque dui. Nulla gravida id neque nec dapibus. Cras in mattis arcu, sit amet facilisis ipsum. Duis quis quam est. Nunc mollis eleifend nisi, in tristique nibh. Sed pharetra nisl id eleifend condimentum. Phasellus scelerisque erat dolor, sed pretium risus vestibulum sit amet. Cras vehicula neque nisl, ornare consectetur nulla pharetra eu. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Quisque consectetur accumsan eleifend. Mauris in neque tellus. Mauris at urna ipsum. Aenean vitae elementum magna. Proin id porta magna. Vivamus maximus efficitur malesuada. Etiam vel blandit tellus. Aliquam mattis lectus non diam aliquam luctus. Fusce condimentum mauris at sem posuere posuere. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque vehicula non nisl ut volutpat. Etiam bibendum, risus et tempus placerat, augue enim sagittis ligula, vitae posuere eros neque ornare risus. Vestibulum dignissim rhoncus sapien, in lobortis nulla pretium dapibus. Mauris laoreet ipsum nec augue posuere vestibulum ut sed nisi. Nam in aliquet odio, at molestie leo. Duis interdum interdum iaculis. Nullam ac feugiat risus, a posuere nulla. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus turpis libero, dignissim eu felis eu, fringilla mollis nibh. Curabitur non condimentum augue, sed sagittis nisl. Phasellus ante tellus, sagittis ac eros egestas, pretium pharetra dolor.	Aliquam erat volutpat. In et consequat quam. Phasellus dapibus mollis lacinia. Aenean dignissim viverra ante nec tristique. Maecenas massa libero, consectetur quis metus sed, luctus pellentesque leo. Curabitur sollicitudin interdum lectus, non semper nulla sodales iaculis. Nullam at mollis nisl. Nunc quis malesuada erat. In iaculis sed erat eu pellentesque. In tincidunt ut erat vel consectetur. Donec efficitur ipsum quis massa pellentesque sagittis. Quisque suscipit massa at facilisis placerat. Ut venenatis ultricies massa et tempor. "

	for n:=0; n<b.N; n++ {
		jsonEnc(versionNum, sampleNum, longMessage)
	}
}

func BenchmarkJsonDecLong(b *testing.B) {
	longjsonEncoded := []byte {123, 34, 86, 101, 114, 115, 105, 111, 110, 78, 117, 109, 34, 58, 49, 44, 34, 83, 97, 109, 112, 108, 101, 78, 117, 109, 34, 58, 51, 46, 49, 52, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 76, 111, 114, 101, 109, 32, 105, 112, 115, 117, 109, 32, 100, 111, 108, 111, 114, 32, 115, 105, 116, 32, 97, 109, 101, 116, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 100, 105, 112, 105, 115, 99, 105, 110, 103, 32, 101, 108, 105, 116, 46, 32, 65, 101, 110, 101, 97, 110, 32, 118, 105, 116, 97, 101, 32, 110, 105, 115, 108, 32, 105, 110, 32, 102, 101, 108, 105, 115, 32, 115, 117, 115, 99, 105, 112, 105, 116, 32, 97, 99, 99, 117, 109, 115, 97, 110, 32, 113, 117, 105, 115, 32, 105, 110, 32, 109, 97, 103, 110, 97, 46, 32, 69, 116, 105, 97, 109, 32, 105, 97, 99, 117, 108, 105, 115, 32, 111, 100, 105, 111, 32, 98, 105, 98, 101, 110, 100, 117, 109, 44, 32, 115, 117, 115, 99, 105, 112, 105, 116, 32, 105, 112, 115, 117, 109, 32, 118, 105, 116, 97, 101, 44, 32, 115, 117, 115, 99, 105, 112, 105, 116, 32, 108, 111, 114, 101, 109, 46, 32, 77, 97, 117, 114, 105, 115, 32, 105, 110, 116, 101, 114, 100, 117, 109, 32, 109, 105, 32, 117, 108, 116, 114, 105, 99, 101, 115, 32, 101, 102, 102, 105, 99, 105, 116, 117, 114, 32, 105, 110, 116, 101, 114, 100, 117, 109, 46, 32, 78, 117, 108, 108, 97, 32, 113, 117, 105, 115, 32, 112, 117, 114, 117, 115, 32, 105, 97, 99, 117, 108, 105, 115, 32, 108, 101, 111, 32, 99, 111, 109, 109, 111, 100, 111, 32, 102, 101, 114, 109, 101, 110, 116, 117, 109, 46, 32, 73, 110, 32, 116, 105, 110, 99, 105, 100, 117, 110, 116, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 46, 32, 65, 108, 105, 113, 117, 97, 109, 32, 102, 101, 114, 109, 101, 110, 116, 117, 109, 32, 115, 111, 108, 108, 105, 99, 105, 116, 117, 100, 105, 110, 32, 117, 108, 116, 114, 105, 99, 101, 115, 46, 32, 68, 111, 110, 101, 99, 32, 101, 117, 32, 109, 97, 103, 110, 97, 32, 118, 101, 108, 32, 109, 97, 103, 110, 97, 32, 118, 105, 118, 101, 114, 114, 97, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 105, 110, 32, 97, 116, 32, 115, 97, 112, 105, 101, 110, 46, 32, 85, 116, 32, 111, 100, 105, 111, 32, 116, 101, 108, 108, 117, 115, 44, 32, 109, 97, 120, 105, 109, 117, 115, 32, 110, 111, 110, 32, 112, 111, 114, 116, 116, 105, 116, 111, 114, 32, 101, 117, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 113, 117, 105, 115, 32, 110, 117, 110, 99, 46, 32, 85, 116, 32, 101, 102, 102, 105, 99, 105, 116, 117, 114, 44, 32, 117, 114, 110, 97, 32, 105, 110, 32, 109, 111, 108, 101, 115, 116, 105, 101, 32, 115, 99, 101, 108, 101, 114, 105, 115, 113, 117, 101, 44, 32, 112, 117, 114, 117, 115, 32, 110, 117, 108, 108, 97, 32, 115, 97, 103, 105, 116, 116, 105, 115, 32, 110, 105, 98, 104, 44, 32, 110, 111, 110, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 105, 112, 115, 117, 109, 32, 101, 108, 105, 116, 32, 118, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 100, 117, 105, 46, 32, 85, 116, 32, 115, 101, 100, 32, 114, 104, 111, 110, 99, 117, 115, 32, 109, 105, 46, 32, 65, 101, 110, 101, 97, 110, 32, 112, 114, 101, 116, 105, 117, 109, 32, 108, 97, 99, 117, 115, 32, 118, 101, 108, 32, 109, 97, 115, 115, 97, 32, 114, 104, 111, 110, 99, 117, 115, 44, 32, 97, 116, 32, 102, 105, 110, 105, 98, 117, 115, 32, 109, 97, 115, 115, 97, 32, 115, 111, 100, 97, 108, 101, 115, 46, 32, 80, 114, 97, 101, 115, 101, 110, 116, 32, 116, 117, 114, 112, 105, 115, 32, 110, 117, 108, 108, 97, 44, 32, 109, 111, 108, 108, 105, 115, 32, 101, 103, 101, 116, 32, 100, 105, 97, 109, 32, 105, 100, 44, 32, 116, 101, 109, 112, 111, 114, 32, 112, 117, 108, 118, 105, 110, 97, 114, 32, 111, 100, 105, 111, 46, 32, 80, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 110, 105, 115, 105, 32, 110, 117, 110, 99, 44, 32, 112, 108, 97, 99, 101, 114, 97, 116, 32, 115, 101, 100, 32, 115, 97, 112, 105, 101, 110, 32, 97, 116, 44, 32, 112, 111, 114, 116, 97, 32, 99, 117, 114, 115, 117, 115, 32, 112, 117, 114, 117, 115, 46, 32, 83, 101, 100, 32, 110, 101, 99, 32, 113, 117, 97, 109, 32, 101, 115, 116, 46, 32, 83, 117, 115, 112, 101, 110, 100, 105, 115, 115, 101, 32, 118, 105, 116, 97, 101, 32, 109, 97, 117, 114, 105, 115, 32, 114, 105, 115, 117, 115, 46, 32, 86, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 102, 97, 99, 105, 108, 105, 115, 105, 115, 32, 116, 101, 109, 112, 111, 114, 32, 111, 114, 99, 105, 32, 115, 101, 100, 32, 112, 104, 97, 114, 101, 116, 114, 97, 46, 32, 78, 97, 109, 32, 105, 100, 32, 102, 101, 108, 105, 115, 32, 105, 100, 32, 108, 101, 99, 116, 117, 115, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 108, 97, 99, 105, 110, 105, 97, 46, 32, 78, 117, 108, 108, 97, 109, 32, 97, 99, 32, 115, 97, 112, 105, 101, 110, 32, 97, 32, 110, 105, 115, 105, 32, 111, 114, 110, 97, 114, 101, 32, 115, 111, 108, 108, 105, 99, 105, 116, 117, 100, 105, 110, 32, 97, 32, 115, 105, 116, 32, 97, 109, 101, 116, 32, 101, 108, 105, 116, 46, 32, 86, 105, 118, 97, 109, 117, 115, 32, 102, 101, 117, 103, 105, 97, 116, 32, 100, 117, 105, 32, 115, 105, 116, 32, 97, 109, 101, 116, 32, 97, 114, 99, 117, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 32, 114, 104, 111, 110, 99, 117, 115, 46, 32, 80, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 104, 97, 98, 105, 116, 97, 110, 116, 32, 109, 111, 114, 98, 105, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 32, 115, 101, 110, 101, 99, 116, 117, 115, 32, 101, 116, 32, 110, 101, 116, 117, 115, 32, 101, 116, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 32, 102, 97, 109, 101, 115, 32, 97, 99, 32, 116, 117, 114, 112, 105, 115, 32, 101, 103, 101, 115, 116, 97, 115, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 110, 111, 110, 32, 115, 99, 101, 108, 101, 114, 105, 115, 113, 117, 101, 32, 100, 117, 105, 46, 32, 78, 117, 108, 108, 97, 32, 103, 114, 97, 118, 105, 100, 97, 32, 105, 100, 32, 110, 101, 113, 117, 101, 32, 110, 101, 99, 32, 100, 97, 112, 105, 98, 117, 115, 46, 32, 67, 114, 97, 115, 32, 105, 110, 32, 109, 97, 116, 116, 105, 115, 32, 97, 114, 99, 117, 44, 32, 115, 105, 116, 32, 97, 109, 101, 116, 32, 102, 97, 99, 105, 108, 105, 115, 105, 115, 32, 105, 112, 115, 117, 109, 46, 32, 68, 117, 105, 115, 32, 113, 117, 105, 115, 32, 113, 117, 97, 109, 32, 101, 115, 116, 46, 32, 78, 117, 110, 99, 32, 109, 111, 108, 108, 105, 115, 32, 101, 108, 101, 105, 102, 101, 110, 100, 32, 110, 105, 115, 105, 44, 32, 105, 110, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 32, 110, 105, 98, 104, 46, 32, 83, 101, 100, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 110, 105, 115, 108, 32, 105, 100, 32, 101, 108, 101, 105, 102, 101, 110, 100, 32, 99, 111, 110, 100, 105, 109, 101, 110, 116, 117, 109, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 115, 99, 101, 108, 101, 114, 105, 115, 113, 117, 101, 32, 101, 114, 97, 116, 32, 100, 111, 108, 111, 114, 44, 32, 115, 101, 100, 32, 112, 114, 101, 116, 105, 117, 109, 32, 114, 105, 115, 117, 115, 32, 118, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 115, 105, 116, 32, 97, 109, 101, 116, 46, 32, 67, 114, 97, 115, 32, 118, 101, 104, 105, 99, 117, 108, 97, 32, 110, 101, 113, 117, 101, 32, 110, 105, 115, 108, 44, 32, 111, 114, 110, 97, 114, 101, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 110, 117, 108, 108, 97, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 101, 117, 46, 32, 80, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 104, 97, 98, 105, 116, 97, 110, 116, 32, 109, 111, 114, 98, 105, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 32, 115, 101, 110, 101, 99, 116, 117, 115, 32, 101, 116, 32, 110, 101, 116, 117, 115, 32, 101, 116, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 32, 102, 97, 109, 101, 115, 32, 97, 99, 32, 116, 117, 114, 112, 105, 115, 32, 101, 103, 101, 115, 116, 97, 115, 46, 32, 81, 117, 105, 115, 113, 117, 101, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 99, 99, 117, 109, 115, 97, 110, 32, 101, 108, 101, 105, 102, 101, 110, 100, 46, 32, 77, 97, 117, 114, 105, 115, 32, 105, 110, 32, 110, 101, 113, 117, 101, 32, 116, 101, 108, 108, 117, 115, 46, 32, 77, 97, 117, 114, 105, 115, 32, 97, 116, 32, 117, 114, 110, 97, 32, 105, 112, 115, 117, 109, 46, 32, 65, 101, 110, 101, 97, 110, 32, 118, 105, 116, 97, 101, 32, 101, 108, 101, 109, 101, 110, 116, 117, 109, 32, 109, 97, 103, 110, 97, 46, 32, 80, 114, 111, 105, 110, 32, 105, 100, 32, 112, 111, 114, 116, 97, 32, 109, 97, 103, 110, 97, 46, 32, 86, 105, 118, 97, 109, 117, 115, 32, 109, 97, 120, 105, 109, 117, 115, 32, 101, 102, 102, 105, 99, 105, 116, 117, 114, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 46, 32, 69, 116, 105, 97, 109, 32, 118, 101, 108, 32, 98, 108, 97, 110, 100, 105, 116, 32, 116, 101, 108, 108, 117, 115, 46, 32, 65, 108, 105, 113, 117, 97, 109, 32, 109, 97, 116, 116, 105, 115, 32, 108, 101, 99, 116, 117, 115, 32, 110, 111, 110, 32, 100, 105, 97, 109, 32, 97, 108, 105, 113, 117, 97, 109, 32, 108, 117, 99, 116, 117, 115, 46, 32, 70, 117, 115, 99, 101, 32, 99, 111, 110, 100, 105, 109, 101, 110, 116, 117, 109, 32, 109, 97, 117, 114, 105, 115, 32, 97, 116, 32, 115, 101, 109, 32, 112, 111, 115, 117, 101, 114, 101, 32, 112, 111, 115, 117, 101, 114, 101, 46, 32, 76, 111, 114, 101, 109, 32, 105, 112, 115, 117, 109, 32, 100, 111, 108, 111, 114, 32, 115, 105, 116, 32, 97, 109, 101, 116, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 100, 105, 112, 105, 115, 99, 105, 110, 103, 32, 101, 108, 105, 116, 46, 32, 81, 117, 105, 115, 113, 117, 101, 32, 118, 101, 104, 105, 99, 117, 108, 97, 32, 110, 111, 110, 32, 110, 105, 115, 108, 32, 117, 116, 32, 118, 111, 108, 117, 116, 112, 97, 116, 46, 32, 69, 116, 105, 97, 109, 32, 98, 105, 98, 101, 110, 100, 117, 109, 44, 32, 114, 105, 115, 117, 115, 32, 101, 116, 32, 116, 101, 109, 112, 117, 115, 32, 112, 108, 97, 99, 101, 114, 97, 116, 44, 32, 97, 117, 103, 117, 101, 32, 101, 110, 105, 109, 32, 115, 97, 103, 105, 116, 116, 105, 115, 32, 108, 105, 103, 117, 108, 97, 44, 32, 118, 105, 116, 97, 101, 32, 112, 111, 115, 117, 101, 114, 101, 32, 101, 114, 111, 115, 32, 110, 101, 113, 117, 101, 32, 111, 114, 110, 97, 114, 101, 32, 114, 105, 115, 117, 115, 46, 32, 86, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 100, 105, 103, 110, 105, 115, 115, 105, 109, 32, 114, 104, 111, 110, 99, 117, 115, 32, 115, 97, 112, 105, 101, 110, 44, 32, 105, 110, 32, 108, 111, 98, 111, 114, 116, 105, 115, 32, 110, 117, 108, 108, 97, 32, 112, 114, 101, 116, 105, 117, 109, 32, 100, 97, 112, 105, 98, 117, 115, 46, 32, 77, 97, 117, 114, 105, 115, 32, 108, 97, 111, 114, 101, 101, 116, 32, 105, 112, 115, 117, 109, 32, 110, 101, 99, 32, 97, 117, 103, 117, 101, 32, 112, 111, 115, 117, 101, 114, 101, 32, 118, 101, 115, 116, 105, 98, 117, 108, 117, 109, 32, 117, 116, 32, 115, 101, 100, 32, 110, 105, 115, 105, 46, 32, 78, 97, 109, 32, 105, 110, 32, 97, 108, 105, 113, 117, 101, 116, 32, 111, 100, 105, 111, 44, 32, 97, 116, 32, 109, 111, 108, 101, 115, 116, 105, 101, 32, 108, 101, 111, 46, 32, 68, 117, 105, 115, 32, 105, 110, 116, 101, 114, 100, 117, 109, 32, 105, 110, 116, 101, 114, 100, 117, 109, 32, 105, 97, 99, 117, 108, 105, 115, 46, 32, 78, 117, 108, 108, 97, 109, 32, 97, 99, 32, 102, 101, 117, 103, 105, 97, 116, 32, 114, 105, 115, 117, 115, 44, 32, 97, 32, 112, 111, 115, 117, 101, 114, 101, 32, 110, 117, 108, 108, 97, 46, 32, 76, 111, 114, 101, 109, 32, 105, 112, 115, 117, 109, 32, 100, 111, 108, 111, 114, 32, 115, 105, 116, 32, 97, 109, 101, 116, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 97, 100, 105, 112, 105, 115, 99, 105, 110, 103, 32, 101, 108, 105, 116, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 116, 117, 114, 112, 105, 115, 32, 108, 105, 98, 101, 114, 111, 44, 32, 100, 105, 103, 110, 105, 115, 115, 105, 109, 32, 101, 117, 32, 102, 101, 108, 105, 115, 32, 101, 117, 44, 32, 102, 114, 105, 110, 103, 105, 108, 108, 97, 32, 109, 111, 108, 108, 105, 115, 32, 110, 105, 98, 104, 46, 32, 67, 117, 114, 97, 98, 105, 116, 117, 114, 32, 110, 111, 110, 32, 99, 111, 110, 100, 105, 109, 101, 110, 116, 117, 109, 32, 97, 117, 103, 117, 101, 44, 32, 115, 101, 100, 32, 115, 97, 103, 105, 116, 116, 105, 115, 32, 110, 105, 115, 108, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 97, 110, 116, 101, 32, 116, 101, 108, 108, 117, 115, 44, 32, 115, 97, 103, 105, 116, 116, 105, 115, 32, 97, 99, 32, 101, 114, 111, 115, 32, 101, 103, 101, 115, 116, 97, 115, 44, 32, 112, 114, 101, 116, 105, 117, 109, 32, 112, 104, 97, 114, 101, 116, 114, 97, 32, 100, 111, 108, 111, 114, 46, 92, 116, 65, 108, 105, 113, 117, 97, 109, 32, 101, 114, 97, 116, 32, 118, 111, 108, 117, 116, 112, 97, 116, 46, 32, 73, 110, 32, 101, 116, 32, 99, 111, 110, 115, 101, 113, 117, 97, 116, 32, 113, 117, 97, 109, 46, 32, 80, 104, 97, 115, 101, 108, 108, 117, 115, 32, 100, 97, 112, 105, 98, 117, 115, 32, 109, 111, 108, 108, 105, 115, 32, 108, 97, 99, 105, 110, 105, 97, 46, 32, 65, 101, 110, 101, 97, 110, 32, 100, 105, 103, 110, 105, 115, 115, 105, 109, 32, 118, 105, 118, 101, 114, 114, 97, 32, 97, 110, 116, 101, 32, 110, 101, 99, 32, 116, 114, 105, 115, 116, 105, 113, 117, 101, 46, 32, 77, 97, 101, 99, 101, 110, 97, 115, 32, 109, 97, 115, 115, 97, 32, 108, 105, 98, 101, 114, 111, 44, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 32, 113, 117, 105, 115, 32, 109, 101, 116, 117, 115, 32, 115, 101, 100, 44, 32, 108, 117, 99, 116, 117, 115, 32, 112, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 108, 101, 111, 46, 32, 67, 117, 114, 97, 98, 105, 116, 117, 114, 32, 115, 111, 108, 108, 105, 99, 105, 116, 117, 100, 105, 110, 32, 105, 110, 116, 101, 114, 100, 117, 109, 32, 108, 101, 99, 116, 117, 115, 44, 32, 110, 111, 110, 32, 115, 101, 109, 112, 101, 114, 32, 110, 117, 108, 108, 97, 32, 115, 111, 100, 97, 108, 101, 115, 32, 105, 97, 99, 117, 108, 105, 115, 46, 32, 78, 117, 108, 108, 97, 109, 32, 97, 116, 32, 109, 111, 108, 108, 105, 115, 32, 110, 105, 115, 108, 46, 32, 78, 117, 110, 99, 32, 113, 117, 105, 115, 32, 109, 97, 108, 101, 115, 117, 97, 100, 97, 32, 101, 114, 97, 116, 46, 32, 73, 110, 32, 105, 97, 99, 117, 108, 105, 115, 32, 115, 101, 100, 32, 101, 114, 97, 116, 32, 101, 117, 32, 112, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 46, 32, 73, 110, 32, 116, 105, 110, 99, 105, 100, 117, 110, 116, 32, 117, 116, 32, 101, 114, 97, 116, 32, 118, 101, 108, 32, 99, 111, 110, 115, 101, 99, 116, 101, 116, 117, 114, 46, 32, 68, 111, 110, 101, 99, 32, 101, 102, 102, 105, 99, 105, 116, 117, 114, 32, 105, 112, 115, 117, 109, 32, 113, 117, 105, 115, 32, 109, 97, 115, 115, 97, 32, 112, 101, 108, 108, 101, 110, 116, 101, 115, 113, 117, 101, 32, 115, 97, 103, 105, 116, 116, 105, 115, 46, 32, 81, 117, 105, 115, 113, 117, 101, 32, 115, 117, 115, 99, 105, 112, 105, 116, 32, 109, 97, 115, 115, 97, 32, 97, 116, 32, 102, 97, 99, 105, 108, 105, 115, 105, 115, 32, 112, 108, 97, 99, 101, 114, 97, 116, 46, 32, 85, 116, 32, 118, 101, 110, 101, 110, 97, 116, 105, 115, 32, 117, 108, 116, 114, 105, 99, 105, 101, 115, 32, 109, 97, 115, 115, 97, 32, 101, 116, 32, 116, 101, 109, 112, 111, 114, 46, 32, 34, 125}

	for n:=0; n<b.N; n++ {
		jsonDec(longjsonEncoded)
	}
}
