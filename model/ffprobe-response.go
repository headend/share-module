package model

type Streams struct {
	Index	int	`json:"index"`
	Codec_name	string	`json:"codec_name"`
	Codec_long_name	string	`json:"codec_long_name"`
	Profile	string	`json:"profile"`
	Codec_type	string	`json:"codec_type"`
	Codec_time_base	string	`json:"codec_time_base"`
	Codec_tag_string	string	`json:"codec_tag_string"`
	Codec_tag	string	`json:"codec_tag"`
	/*
	use for video stream
	 */
	Width	int	`json:"width"`
	Height	int	`json:"height"`
	Coded_width	int	`json:"coded_width"`
	Coded_height	int	`json:"coded_height"`
	Has_b_frames	int	`json:"has_b_frames"`
	Sample_aspect_ratio	string	`json:"sample_aspect_ratio"`
	Display_aspect_ratio	string	`json:"display_aspect_ratio"`
	Pix_fmt	string	`json:"pix_fmt"`
	Level	int	`json:"level"`
	Chroma_location	string	`json:"chroma_location"`
	Field_order	string	`json:"field_order"`
	Refs	int	`json:"refs"`
	Is_avc	string	`json:"is_avc"`
	Nal_length_size	string	`json:"nal_length_size"`
	Id	string	`json:"id"`
	R_frame_rate	string	`json:"r_frame_rate"`
	Avg_frame_rate	string	`json:"avg_frame_rate"`
	Time_base	string	`json:"time_base"`
	Start_pts	int	`json:"start_pts"`
	Start_time	string	`json:"start_time"`
	Bits_per_raw_sample	string	`json:"bits_per_raw_sample"`
	Tags	Tags	`json:"tags"`
	/*
	Only use for audio stream
	 */
	Sample_fmt	string	`json:"sample_fmt"`
	Sample_rate	string	`json:"sample_rate"`
	Channels	int	`json:"channels"`
	Channel_layout	string	`json:"channel_layout"`
	Bits_per_sample	int	`json:"bits_per_sample"`
	Bit_rate	string	`json:"bit_rate"`

}

type Tags struct {
	Language	string	`json:"language"`
}

type Format struct {
	Filename	string	`json:"filename"`
	Nb_streams	int	`json:"nb_streams"`
	Nb_programs	int	`json:"nb_programs"`
	Format_name	string	`json:"format_name"`
	Format_long_name	string	`json:"format_long_name"`
	Start_time	string	`json:"start_time"`
	Probe_score	int	`json:"probe_score"`
}

type FfprobeResponse struct {
	Streams	[]Streams	`json:"streams"`
	Format	Format	`json:"format"`
}