package gore

import (
	"flag"
	"log"
)

type Options struct {
	OutputFilename string
	IsMassup       bool
	Filename       string
	ObfOption      string
}

func CommandLineParse() *Options {
	option := new(Options)
	flag.StringVar(&option.OutputFilename, "output", "", "另保存的文件名")
	flag.StringVar(&option.Filename, "f", "", "源文件名")
	flag.BoolVar(&option.IsMassup, "a", false, "是否消除Go的编译信息")
	flag.Parse()
	return option
}
func (o *Options) CheckConfig() {
	if o.Filename == "" {
		log.Fatalln("请选择一个源文件 UseAge:-f binary")
	}
}

var GoOptions *Options
