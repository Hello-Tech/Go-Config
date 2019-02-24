# GoConfig

Go项目的配置文件解析器，目前主要支持Ini文件

## ini格式的文件

.ini 文件是Initialization File的缩写，即初始化文件，是windows的系统配置文件所采用的存储格式，统管windows的各项配置，一般用户就用windows提供的各项图形化管理界面就可实现相同的配置了。但在某些情况，还是要直接编辑ini才方便，一般只有很熟悉windows才能去直接编辑。

INI文件由节、键、值组成：

 - 节：节是一系列Key（键）,Value（值）的集合
 - 参数：一个参数表示一个键值对
 - 注解：凡以`;`,`#`开始的某一行，表示这是一行注释。

 更多Ini文件的细节，可以参考：[https://en.wikipedia.org/wiki/INI_file](https://en.wikipedia.org/wiki/INI_file)


## 用法与用例

```
iniReader := NewIniReader()
if iniReader.LoadIni(path) {
    // read sections and keys
}
```
