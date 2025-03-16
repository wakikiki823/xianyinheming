<template>
  <view class="container">
    <!-- 顶部导航栏 -->
    <view class="nav-bar">
      <view class="back-btn" @click="goBack">
        <image src="/static/qianwang 3.png" mode="aspectFit" class="back-icon"></image>
      </view>
      <text class="title">扫描二维码</text>
      <!-- 添加闪光灯控制 -->
      <view class="flash-btn" @click="toggleFlash">
        <image :src="flashIcon" mode="aspectFit" class="flash-icon"></image>
      </view>
    </view>
    
    <!-- 扫描区域 -->
    <view class="scan-area">
      <camera 
        device-position="back" 
        :flash="flashMode" 
        class="camera"
        @error="handleError"
      >
        <cover-view class="scan-box">
          <cover-view class="scan-border"></cover-view>
          <cover-view class="scan-line"></cover-view>
          <!-- 添加四个角 -->
          <cover-view class="corner-box">
            <cover-view class="corner lt"></cover-view>
            <cover-view class="corner rt"></cover-view>
            <cover-view class="corner lb"></cover-view>
            <cover-view class="corner rb"></cover-view>
          </cover-view>
        </cover-view>
      </camera>
    </view>
    
    <!-- 提示文本 -->
    <view class="tips">
      <text>将二维码放入框内，即可自动扫描</text>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      isScanning: false,
      flashMode: 'off',
      scanSound: null
    }
  },
  computed: {
    flashIcon() {
      return this.flashMode === 'on' ? '/static/flash-on.png' : '/static/flash-off.png'
    }
  },
  onLoad() {
    // 初始化扫描音效
    this.scanSound = uni.createInnerAudioContext()
    this.scanSound.src = '/static/scan.mp3'
    
    // 检查相机权限
    uni.authorize({
      scope: 'scope.camera',
      success: () => {
        this.startScan()
      },
      fail: () => {
        uni.showToast({
          title: '请授权相机权限',
          icon: 'none'
        })
      }
    })
  },
  onUnload() {
    // 销毁音效实例
    if (this.scanSound) {
      this.scanSound.destroy()
    }
  },
  methods: {
    toggleFlash() {
      this.flashMode = this.flashMode === 'on' ? 'off' : 'on'
    },
    startScan() {
      if (this.isScanning) return
      this.isScanning = true
      
      uni.scanCode({
        onlyFromCamera: true,
        scanType: ['qrCode'],
        success: (res) => {
          // 播放扫描音效
          this.scanSound.play()
          
          console.log('扫码结果：', res)
          uni.showToast({
            title: '扫描成功',
            icon: 'success'
          })
          
          // 震动反馈
          uni.vibrateShort()
          
          // 延迟返回，让用户看到成功提示
          setTimeout(() => {
            uni.navigateBack({
              delta: 1
            })
          }, 1500)
        },
        fail: (err) => {
          console.error('扫码失败：', err)
          uni.showToast({
            title: '扫描失败',
            icon: 'none'
          })
        },
        complete: () => {
          this.isScanning = false
        }
      })
    },
    goBack() {
      uni.navigateBack()
    },
    handleError(e) {
      console.error('相机错误：', e)
      uni.showToast({
        title: '相机出现错误',
        icon: 'none'
      })
    }
  }
}
</script>

<style>
.container {
  width: 100vw;
  height: 100vh;
  background-color: #000;
  position: relative;
}

.nav-bar {
  height: 44px;
  width: 100%;
  background-color: #000;
  display: flex;
  align-items: center;
  padding: 0 15px;
  position: relative;
}

.back-btn {
  width: 44px;
  height: 44px;
  display: flex;
  align-items: center;
}

.back-icon {
  width: 20px;
  height: 20px;
}

.title {
  color: #fff;
  font-size: 18px;
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}

.scan-area {
  width: 100%;
  height: 70vh;
  position: relative;
}

.camera {
  width: 100%;
  height: 100%;
}

.scan-box {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 200px;
  height: 200px;
  background-color: transparent;
}

.scan-border {
  position: absolute;
  width: 100%;
  height: 100%;
  border: none;
}

.scan-line {
  position: absolute;
  width: 100%;
  height: 6px;
  background: linear-gradient(to bottom, transparent, #07c160, transparent);
  top: 0;
  border-radius: 3px;
}

.tips {
  position: absolute;
  bottom: 15%;
  left: 50%;
  transform: translateX(-50%);
  color: #fff;
  font-size: 14px;
  text-align: center;
}

.flash-btn {
  position: absolute;
  right: 15px;
  width: 44px;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.flash-icon {
  width: 24px;
  height: 24px;
}

.corner-box {
  position: absolute;
  width: 100%;
  height: 100%;
}

.corner {
  position: absolute;
  width: 20px;
  height: 20px;
  border: 2px solid #07c160;
}

.lt {
  left: -2px;
  top: -2px;
  border-right: none;
  border-bottom: none;
}

.rt {
  right: -2px;
  top: -2px;
  border-left: none;
  border-bottom: none;
}

.lb {
  left: -2px;
  bottom: -2px;
  border-right: none;
  border-top: none;
}

.rb {
  right: -2px;
  bottom: -2px;
  border-left: none;
  border-top: none;
}
</style> 