// +build !ubports

package main

/*
#cgo CFLAGS: -pipe -O2 -Wall -Wextra -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -O2 -Wall -Wextra -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../GoProject -I. -I/home/duoduo/Qt/5.15.0/gcc_64/include -I/home/duoduo/Qt/5.15.0/gcc_64/include/QtGui -I/home/duoduo/Qt/5.15.0/gcc_64/include/QtCore -I. -isystem /usr/include/libdrm -I/home/duoduo/Qt/5.15.0/gcc_64/mkspecs/linux-g++
#cgo LDFLAGS: -O1 -Wl,-rpath,/home/duoduo/Qt/5.15.0/gcc_64/lib
#cgo LDFLAGS:  /home/duoduo/Qt/5.15.0/gcc_64/lib/libQt5Gui.so /home/duoduo/Qt/5.15.0/gcc_64/lib/libQt5Core.so -lGL -lpthread
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
