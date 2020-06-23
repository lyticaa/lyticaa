package types

type SparklineData struct {
	ToolText string  `json:"tooltext"`
	Value    float64 `json:"value"`
}

type Sparkline struct {
	Data []SparklineData `json:"data"`
}

type DataSet struct {
	SeriesName string `json:"seriesname"`
	Data       string `json:"data"`
}

type Category struct {
	Category string `json:"category"`
}

type Line struct {
	Categories []Category `json:"categories"`
	DataSets   []DataSet  `json:"dataset"`
}

type Chart struct {
	Line      Line      `json:"line"`
	Sparkline Sparkline `json:"sparkline"`
}
