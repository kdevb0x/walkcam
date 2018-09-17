LOCAL_PATH := $(call my-dir)

include $(CLEAR_VARS)

LOCAL_MODULE    := walkcam
LOCAL_SRC_FILES := lib/libwalkcamactivity.so
LOCAL_LDLIBS    := -llog -landroid -lcamera

include $(PREBUILT_SHARED_LIBRARY)
