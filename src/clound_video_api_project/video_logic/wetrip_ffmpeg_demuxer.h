/*
 * wetrip_ffmpeg_demuxer.h
 *
 *  Created on: 2016年6月5日
 *      Author: mok
 */

#ifndef WETRIP_FFMPEG_DEMUXER_H_
#define WETRIP_FFMPEG_DEMUXER_H_


extern int wetrip_video_thum_gps_generate(const char* video_path,const char* thum_video_path, const char* filename,float t1,float t2,float t3,float t4,float t5,float t6,float t7,float t8,float t9,float t10,char* out_json);
extern int wetrip_video_thum_generate(const char* video_full_file_path,const char* thum_full_file_path);


#endif /* WETRIP_FFMPEG_DEMUXER_H_ */
