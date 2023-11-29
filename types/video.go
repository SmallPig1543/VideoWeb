package types

type VideoCreateRequest struct {
	Types    string `json:"types" form:"types"`
	Title    string `json:"title" form:"title"`
	FilePath string `json:"file_path" form:"file_path"`
}

type VideoWatchRequest struct {
	Vid uint `json:"vid" form:"vid"`
}

type VideoRankRequest struct {
}
