# osx activated

- build

```bash
make
```

- run

```bash
./bin/activate
```

# autostart

使用`cc.guzal.activated.plist`文件,修改key: ProgramArguments里面的字段，改成app所在位置，精确到app名称。
然后复制plist文件到`~/Library/LaunchAgents/`目录.
```bash
cp cc.guzal.activated.plist ~/Library/LaunchAgents/cc.guzal.activated.plist
```