## Go-strip
- 支持 Windows、Linux、MacOS（amd64、386）的二进制文件。
- 源代码进行了工程化相关的处理，支持一键自动化测试，自动化编译[github action]。

**功能**
- 支持混淆
    - 函数名称
    - 函数路径
    - Go Struct
    - Type
    - Go Compiler Version
    - Go BuildID
    - Go Root
    - Go ModInfo
  
## 代码结构
- cmd 生成命令行主程序
- core 一些自己的函数方法
- gore gore的包，目前同步的版本是 `https://github.com/goretk/gore/commit/141f38f24c76dffd798c5cbad48d05129ce9b9b5`
- test 自动化测试