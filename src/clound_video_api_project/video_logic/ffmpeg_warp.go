package video_logic

// #cgo CFLAGS: -Wall
// #cgo LDFLAGS: -Wl,-rpath="/home/liuliang/ffmpeg-build/lib"
// #cgo LDFLAGS: -L/home/liuliang/workspace/wetrip_ffmpeg_demuxer/Debug
// #cgo LDFLAGS: -L/home/liuliang/workspace/wetrip_ffmpeg_demuxer
// #cgo LDFLAGS: -lwetrip_ffmpeg_demuxer -lstdc++ -ljpeg  -lpthread -lrt
// #cgo LDFLAGS: -L/home/liuliang/ffmpeg-build/lib
// #cgo  LDFLAGS: -lavformat -lavcodec -lswscale -lavutil -lswresample
// #include <stdio.h>
// #include <stdlib.h>
// #include "wetrip_ffmpeg_demuxer.h"
import "C"

import "fmt"
import "unsafe"
import "clound_video_api_project/models"

func Video_thum(full_file_path string, thum_full_file_path string) int {
	cs_video_path := C.CString(full_file_path)
	defer C.free(unsafe.Pointer(cs_video_path))

	cs_thum_path := C.CString(thum_full_file_path)
	defer C.free(unsafe.Pointer(cs_thum_path))

	result := int(C.wetrip_video_thum_generate(cs_video_path, cs_thum_path))
	return result
}

func Video_thum_gps(filename string, t1 float32, t2 float32, t3 float32, t4 float32, t5 float32, t6 float32, t7 float32, t8 float32, t9 float32, t10 float32) (result int, result_json string) {
	cs_video_path := C.CString(models.Video_Dir)
	defer C.free(unsafe.Pointer(cs_video_path))

	cs_thum_path := C.CString(models.Video_Thumb_Dir)
	defer C.free(unsafe.Pointer(cs_thum_path))

	cs_video_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(cs_video_filename))

	outstr := (*C.char)(C.malloc(512 * 2))
	defer C.free(unsafe.Pointer(outstr))

	fmt.Print("*==========================\n")
	result = int(C.wetrip_video_thum_gps_generate(cs_video_path, cs_thum_path, cs_video_filename, C.float(t1), C.float(t2), C.float(t3), C.float(t4), C.float(t5), C.float(t6), C.float(t7), C.float(t8), C.float(t9), C.float(t10), outstr))
	fmt.Println(result)
	fmt.Print("#==========================\n")
	if result == 0 {
		result_json = C.GoString(outstr)
	}
	fmt.Println(result_json)
	return
}

func main_test() {

	outstr := (*C.char)(C.malloc(200))
	cs := C.CString("1B527A4303010-100MEDIA-06021411_0003.mp4")
	//result := C.wetrip_video_generate(cs, 1.5, 6.4,
	//	outstr)
	//fmt.Print(result)
	fmt.Print(C.GoString(outstr))
	C.free(unsafe.Pointer(cs))
	C.free(unsafe.Pointer(outstr))

}
