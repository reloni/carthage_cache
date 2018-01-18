# Carthage Cache [![Build Status](https://travis-ci.org/buildben/carthage_cache.svg?branch=master)](https://travis-ci.org/buildben/carthage_cache)

Community-driven centralised *Carthage* cache for dynamic & static binaries of your frameworks. 
Each one has specifically built for your Xcode & Swift versions to speed up your day.

## Install
There're two ways: using precompiled binary or building from sources.
### Download binary
Paste that at a Terminal prompt:
```bash
wget https://static.buildben.io/carthage/client -O /usr/local/bin/bb_carthage_cache | chmod +x /usr/local/bin/bb_carthage_cache 
```
This script will download compiled binary and place it in your path.

### Build from source
**Requirements:** Go 1.8, OSX 10.10+
1. Download source code to **$GOPATH**/src/buildben/carthage_cache/client
2. In the source folder:
```bash
GOOS=darwin GOARCH=amd64 go build -o /usr/local/bin/bb_carthage_cache cmd/carthage_cache.go | chmod +x /usr/local/bin/bb_carthage_cache 
```
3. Done. Now you can use it as any other tool. 

## Usage

### Dynamic linking
Run in your Cartfile.resolved location:
```bash
bb_carthage_cache
```
or if you want osx version:
```bash
bb_carthage_cache -platform mac
```
This command will download and extract binaries to Carthage/Build path:
![image](https://habrastorage.org/webt/hf/q-/jc/hfq-jcgzllyp4s8mhdfgsdavn6a.png)

Or build and upload to the cloud if nobody used this combination of framework & Xcode versions previously.

### Static linking
Almost the same. Run in your Cartfile.resolved location:
```bash
bb_carthage_cache -static
```
and for osx:
```bash
bb_carthage_cache -platform mac -static
```

**Warning!**
- Be sure to remove copy-frameworks step from your build steps when using static linking. Otherwise it will double your dependencies size.
- Static linking drops all resources from frameworks, keep that in mind. Dynamic & static linking together is not supported (yet)

# Support
There're a few limitations still:
- Cartfile.resolved won't be updated when you add a framework to your Cartfile. You must do it manually (yet)

If you find yourself stuck upon an obstacle feel free to open an issue in any comfortable format. I check it almost every day. 

# Uninstall
Maybe you would open an issue first? 
But if you strongly decided to purge that evil from your property... then:
```bash
sudo rm -f /usr/local/bin/bb_carthage_cache
```
That's it. Bye-bye.

# Motivation
Leave a star if you find it useful. It hugely motivates to go on with the development. 
