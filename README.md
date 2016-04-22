# fpanel_controller

```
Usage of ./fpanel_controller pin1 [pin2...]
  -a	control all pin.
  -d int
    	hold on time in seconds. (default 1)

Current support pins
  1  ->  CTS
  2  ->  TXD
  3  ->  RXD
  4  ->  DTR
```

1. flag外的参数 控制输出口．从硬件设备的GND开始，跳过GND和VCC依次就是1 2 3 4
2. -d 持续输出高点平的时间．　模拟正常的ACPI关机使用默认的1s即可. 模拟强制关机使用5s

# TODO
1. 使用74138之类的让一个ft232可以控制8台设备．
2. 配置支持
3. 远程API支持
4. 直接使用libusb，去除依赖libftd1-2


# 编译依赖
ibftdi1-2 (这个库在debian上打包有问题，无法正常编译．需要手动执行

  ```
  cd /usr/lib/*架构目录*/
  sudo ln -sv libftdi1.so.2.2.0 libftdi1.so
  ```

```go get github.com/x-deepin/fpanel_controller```
