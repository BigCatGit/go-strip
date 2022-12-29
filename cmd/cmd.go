package main

import (
	"fmt"
	"gostrip/core"
	"gostrip/gore"
	"log"
	"path/filepath"
)

func showBanner() {
	var logo = `
                       _        _       
                      | |      (_)      
  __ _  ___ ______ ___| |_ _ __ _ _ __  
 / _  |/ _ \______/ __| __| '__| | '_ \
| (_| | (_) |     \__ \ |_| |  | | |_) |
\__, |\___/      |___/\__|_|  |_| .__/
__/ |                          | |
|___/                           |_|
`
	fmt.Println(logo)
}

func printFolderStructures(f *gore.GoFile, pkgs []*gore.Package) {
	for i, p := range pkgs {
		if i != 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("Package %s: %s\n", p.Name, p.Filepath)
		for _, sf := range f.GetSourceFiles(p) {
			sf.Postfix = "\t"
			fmt.Printf("%s\n", sf)
		}
	}
}

func main() {
	showBanner()
	gore.GoOptions = gore.CommandLineParse()
	gore.GoOptions.CheckConfig()

	filename := gore.GoOptions.Filename
	fileStr, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalln("Failed to parse the filepath:", err)
	}
	f, err := gore.Open(fileStr)
	if err != nil {
		log.Fatalln("Error when opening the file:", err)
	}
	defer f.Close()

	err = f.Init()
	if err != nil {
		panic(err)
	}
	// 获取编译器信息
	cmp, err := f.GetCompilerVersion()
	if err != nil {
		log.Println("Error when extracting compiler information:", err)
	}
	// GoRoot信息
	goroot, err := f.GetGoRoot()
	if err != nil {
		log.Println("GoRoot获取失败，err:", err.Error())
	}

	// 默认处理函数
	// 默认处理文件

	_, err = f.GetTypes()
	if err != nil {
		return
	}
	log.Printf("Compiler version: %s (%s)\n", cmp.Name, cmp.Timestamp)
	if gore.GoOptions.IsMassup {
		log.Printf("混淆版本信息")
		newVersion := core.GetRandomString(int(core.TypeStringOffsets.GoVersion.Length))
		f.SetBytes(core.TypeStringOffsets.GoVersion.Offset, core.TypeStringOffsets.GoVersion.Length, []byte(newVersion))

		log.Printf("混淆结构信息,总数:%d\n", len(core.TypeStringOffsets.Datas))
		for _, t := range core.TypeStringOffsets.Datas {
			//newType := core.GetRandomString(int(t.Length))
			newType := core.GenerateSameWords(int(t.Length), 65)
			f.SetBytes(core.TypeStringOffsets.Base+t.Offset, t.Length, []byte(newType))
		}

		log.Printf("混淆文件信息,总数%d\n", len(core.TypeStringOffsets.FileName))
		for _, t := range core.TypeStringOffsets.FileName {
			newFilename := core.GetRandomString(int(t.Length))
			f.SetBytes(t.Offset, t.Length, []byte(newFilename))
		}

		log.Printf("混淆函数信息\n")
		for _, t := range core.TypeStringOffsets.Func {
			newFunc := core.GetRandomString(int(t.Length))
			//newFunc := core.GenerateSameWords(int(t.Length), 65)
			f.SetBytes(t.Offset, t.Length, []byte(newFunc))
		}

		log.Println("混淆BuildID信息")
		l := len(f.BuildID)
		newBuildID := fmt.Sprintf("%sstripd%sby%sgo-strip", core.GetRandomSymbol(), core.GetRandomSymbol(), core.GetRandomSymbol())
		if l > len(newBuildID) {
			newBuildID = core.GetRandomString(l-len(newBuildID)) + newBuildID
		}
		f.Replace([]byte(f.BuildID), []byte(newBuildID), -1)

		log.Println("混淆GoMod信息")
		for _, t := range core.TypeStringOffsets.GoMod {
			n := core.GetRandomString(int(t.Length))
			off, err := f.GetFva(t.Offset)
			if err != nil {
				continue
			}
			f.SetBytes(off, t.Length, []byte(n))
		}

		log.Println("混淆结束")
		newfile, err := filepath.Abs(gore.GoOptions.OutputFilename)
		if err != nil {
			panic(err)
		}
		log.Println("新的文件保存在", newfile)
		f.Save(newfile)
	} else {
		log.Println("GoRoot:", goroot)
		log.Println("BuildID:", f.BuildID)
		modInfo := f.BuildInfo.ModInfo.String()
		fmt.Println(modInfo)

		packages, err := f.GetPackages()
		if err != nil {
			log.Fatalln(err)
		}
		// Err check not needed since the parsing was has already been checked.
		vn, _ := f.GetVendors()
		packages = append(packages, vn...)

		//std, _ := f.GetSTDLib()
		//packages = append(packages, std...)

		unk, _ := f.GetUnknown()
		packages = append(packages, unk...)

		printFolderStructures(f, packages)
	}
}
