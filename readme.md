# http text送信

RaspberryPi-pico Wで[LCD描画プログラムを作った](https://github.com/zsuzuki/picosrvtest)ので、
そこにテキストファイルを送信したくてChat GPT-4につくってもらった。

## 使用方法

```shell
go build src/sendtxt/sendtxt.go
./sendtxt http://192.168.1.2/post send.txt
```
