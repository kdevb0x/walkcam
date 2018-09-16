all: project build

build:
	ndk-build
	ant debug

install:
	ndk-build
	ant debug install

project:
	# (required) go get github.com/xlab/android-go/cmd/android-project
	android-project update --target android-24 --name WalkCamActivity --path .


clean:
	rm -rf libs obj bin gen jni/lib
	rm -f proguard-project.txt project.properties
	rm -f local.properties build.xml
