# fpanel_control

```
usage: fpanel_control {A|B|C|D} [delay seconds]
```

目前只支持简单的命令行参数．必须两个参数
1. 控制的输出口．从硬件设备的GND开始，跳过GND和VCC依次就是ABCD.
2. 持续输出高点平的时间．　模拟正常的ACPI关机使用1即可(1s). 模拟强制关机使用5(5s)

# TODO
1. 使用74138之类的让一个ft232可以控制8台设备．
2. 配置支持
3. 远程API支持
4. 直接使用libusb，去除依赖libftd1-2



# 编译依赖
- libftdi1-2 (这个库在debian上打包有问题，无法正常编译．需要手动执行

  ```
  cd ln -sv /usr/lib/架构目录/
  sudo ln -sv libftdi1.so.2.2.0 libftdi1.so
  ```

