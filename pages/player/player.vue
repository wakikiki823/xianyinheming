<template>
  <view class="player-container" @click="handleUserInteraction">
    <!-- 歌曲信息 -->
    <view class="song-header">
      <text class="song-title">{{ currentSong.title }}</text>
      <text class="song-artist">{{ currentSong.artist }}</text>
    </view>

    <!-- 封面容器 -->
    <view class="cover-wrapper">
      <image 
        class="album-cover"
        :src="currentSong.cover"
        mode="aspectFill"
        style="border-radius: 40rpx 40rpx 250rpx 250rpx; height: 600rpx;"
      />
      <!-- 线性进度条 -->
      <view class="linear-progress-container">
        <view class="time-display">
          <text>{{ formatTime(currentTime) }}</text>
          <text>{{ formatTime(totalDuration) }}</text>
        </view>
        <view 
          class="linear-progress"
          @touchstart="onSliderTouchStart"
          @touchmove.prevent="onSliderMove"
          @touchend="onSliderEnd"
          @click="onTrackClick"
        >
          <view class="progress-fill" :style="{width: progress + '%'}"></view>
          <view class="progress-thumb" :style="{left: progress + '%'}"></view>
        </view>
      </view>
    </view>

    <!-- 歌词显示 -->
    <view class="lyrics">
      <text class="current-lyric">{{ currentLyric }}</text>
      <text class="next-lyric" v-for="(lyric, index) in nextLyrics" :key="index">{{ lyric }}</text>
    </view>

    <!-- 控制按钮 -->
    <view class="control-bar">
      <button class="control-btn prev-btn" @click="playPrevious">
        <image 
          class="control-icon" 
          src="/static/prev.png" 
          style="width: 60rpx; height: 60rpx;"
        ></image>
      </button>
      <button 
        class="play-btn" 
        @click="togglePlay"
        :disabled="!isReady"
      >
        <image 
          class="play-icon" 
          :src="isPlaying ? '/static/zanting.png' : '/static/bofang.png'"
          style="width: 80rpx; height: 80rpx;"
        ></image>
      </button>
      <button class="control-btn next-btn" @click="playNext">
        <image 
          class="control-icon" 
          src="/static/next.png" 
          style="width: 60rpx; height: 60rpx;"
        ></image>
      </button>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      currentTime: 0,
      isPlaying: false,
      currentSong: {},
      audioContext: null,
      playList: [],
      currentIndex: 0,
      currentLyric: "",
      nextLyrics: [],
      controlIcons: {
        prev: '/static/prev.png',
        next: '/static/next.png',
        play: '/static/zanting.png',
        pause: '/static/bofang.png'
      },
      isReady: false,
      autoPlay: false,
      totalDuration: 0,
      isDragging: false,
      progress: 0,
      bufferProgress: 0,
      rotation: 0
    };
  },
  computed: {
    progressWidth() {
      return this.totalDuration > 0 
        ? (this.currentTime / this.totalDuration) * 100
        : 0
    }
  },
  watch: {
    isDragging(newVal) {
      if (!newVal) {
        this.currentTime = this.audioContext.currentTime
      }
    },
    currentTime: {
      handler(newVal) {
        if (!this.isDragging && this.totalDuration > 0) {
          this.progress = (newVal / this.totalDuration) * 100
        }
      },
      immediate: true
    },
    progress(newVal) {
      this.$forceUpdate()
    }
  },
  methods: {
    async initAudio() {
      if (this.audioContext) {
        this.audioContext.destroy()
      }
      
      this.audioContext = uni.createInnerAudioContext()
      this.audioContext.obeyMuteSwitch = false
      this.isReady = false
      
      this.audioContext.onPlay(() => {
        this.isPlaying = true
        this.startProgressUpdate()
      })
      
      this.audioContext.onPause(() => {
        this.isPlaying = false
        this.stopProgressUpdate()
      })
      
      this.audioContext.onEnded(() => {
        this.playNext()
      })

      this.audioContext.onTimeUpdate(() => {
        if (!this.isDragging) {
          this.currentTime = this.audioContext.currentTime
        }
      })
      
      this.audioContext.onCanplay(() => {
        this.totalDuration = this.audioContext.duration
        this.isReady = true
      })

      this.audioContext.onProgressUpdate = () => {
        if (this.audioContext.buffered) {
          this.bufferProgress = (this.audioContext.buffered / this.totalDuration) * 100
        }
      }

      try {
        const valid = await this.validateFile(this.currentSong.url)
        if (!valid) {
          uni.showToast({ title: '音频文件不存在', icon: 'none' })
          return
        }
        
        this.audioContext.src = this.currentSong.url
        await new Promise(resolve => this.audioContext.onCanplay(resolve))
        
        if (this.autoPlay) {
          this.audioContext.play()
        }
      } catch (err) {
        console.error('音频初始化失败:', err)
        uni.showToast({ title: '音频加载失败', icon: 'none' })
      }
    },
    
    startProgressUpdate() {
      this.timer = setInterval(() => {
        this.currentTime = this.audioContext.currentTime
      }, 1000)
    },
    
    stopProgressUpdate() {
      clearInterval(this.timer)
    },
    
    togglePlay() {
      if (!this.audioContext) {
        uni.showToast({ title: '播放器未初始化', icon: 'none' })
        return
      }
      
      if (!this.isReady) {
        uni.showToast({ title: '音频尚未准备好', icon: 'none' })
        return
      }

      if (this.isPlaying) {
        this.audioContext.pause()
      } else {
        try {
          this.audioContext.play()
          this.isPlaying = true
        } catch (err) {
          console.error('播放失败:', err)
          uni.showToast({ title: '播放启动失败', icon: 'none' })
        }
      }
    },
    
    playPrevious() {
      this.currentIndex = (this.currentIndex - 1 + this.playList.length) % this.playList.length
      this.switchSong(this.playList[this.currentIndex])
    },
    
    playNext() {
      this.currentIndex = (this.currentIndex + 1) % this.playList.length
      this.switchSong(this.playList[this.currentIndex])
    },
    
    switchSong(song) {
      this.currentSong = song
      this.currentTime = 0
      this.initAudio()
      this.loadLyrics()
    },
    
    loadLyrics() {
      // 歌词加载逻辑
    },
    formatTime(seconds) {
      const totalSeconds = Number(seconds) || 0;
      const mins = Math.floor(totalSeconds / 60);
      const secs = (totalSeconds % 60).toFixed(1);
      return `${mins}:${secs.padStart(4, '0')}`;
    },
    getErrorText(code) {
      const errors = {
        10001: '系统错误',
        10002: '网络错误',
        10003: '文件错误',
        10004: '格式错误'
      }
      return errors[code] || '未知错误'
    },
    async validateFile(path) {
      // 小程序/APP端验证
      if (typeof uni.getFileSystemManager === 'function') {
        return new Promise(resolve => {
          uni.getFileSystemManager().access({
            path: path,
            success: () => resolve(true),
            fail: () => resolve(false)
          })
        })
      }
      
      // H5端验证
      if (process.env.VUE_APP_PLATFORM === 'h5') {
        try {
          const response = await fetch(path, { method: 'HEAD' })
          return response.ok
        } catch {
          return false
        }
      }

      // 其他平台默认通过验证
      return true
    },
    handleUserInteraction() {
      if (!this.autoPlay) {
        uni.$emit('userInteraction')
      }
    },
    onSliderTouchStart() {
      this.isDragging = true
    },
    onSliderMove(e) {
      const rect = this.$refs.track.getBoundingClientRect()
      const percent = Math.min(Math.max((e.touches[0].clientX - rect.left) / rect.width, 0), 1) * 100
      this.progress = percent
    },
    onSliderEnd() {
      this.seekTo(this.progress)
      this.isDragging = false
    },
    seekTo(percent) {
      const time = (percent / 100) * this.totalDuration
      this.audioContext.seek(time)
      this.currentTime = time
    },
    updateBuffered() {
      if (this.audioContext.buffered.length > 0) {
        this.bufferProgress = (this.audioContext.buffered / this.totalDuration) * 100
      }
    },
    onTrackClick(e) {
      const rect = this.$refs.track.getBoundingClientRect()
      const percent = Math.min(Math.max((e.clientX - rect.left) / rect.width, 0), 1) * 100
      this.seekTo(percent)
    },
    deletePost(postId) {
      uni.showModal({
        title: '确认删除',
        content: '确定要删除这条内容吗？',
        success: (res) => {
          if (res.confirm) {
            this.posts = this.posts.filter(post => post.id !== postId);
            // 实际项目需调用云函数
          }
        }
      });
    }
  },
  onLoad(options) {
    this.playList = uni.getStorageSync('currentPlaylist') || []
    this.currentIndex = this.playList.findIndex(item => item.id === JSON.parse(options.song).id)
    this.currentSong = {
      ...JSON.parse(options.song),
      url: `/static/${JSON.parse(options.song).file}`
    }
    
    uni.$on('userInteraction', () => {
      this.autoPlay = true
      if (!this.isPlaying) {
        this.togglePlay()
      }
    })
  },
  onUnload() {
    this.audioContext && this.audioContext.destroy()
  }
};
</script>

<style>
.player-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 80rpx 40rpx 0;
  position: relative;
  overflow: hidden;
  min-height: 100vh;
}

/* 调整背景分割线 */
.player-container::before {
  content: '';
  position: absolute;
  top: 0rpx;
  left: -20%;
  right: -20%;
  height: 280rpx;
  background: #FFB800;
  border-radius:0 0 80rpx 80rpx; /* 增大底部圆角 */
  box-shadow: 0 -4rpx 16rpx rgba(0, 0, 0, 0.05);
  z-index: 0;
}

/* 新增白色背景层 */
.player-container::after {
  content: '';
  position: absolute;
  top: 160rpx; /* 在橙色区域下方 */
  left: 0;
  right: 0;
  bottom: 0;
  background: #fff;
  border-radius: 80rpx 80rpx 0 0; /* 顶部圆角 */
  z-index: 1;
}

/* 调整歌曲信息位置 */
.song-header {
  margin: 0rpx 0 60rpx; /* 原为100rpx */
  text-align: center;
  z-index: 1;
  position: relative;
  color: #fff;
}

/* 封面容器 */
.cover-wrapper {
  position: relative;
  width: 600rpx;
  height: 600rpx;
  margin: 20rpx 0 40rpx;
  z-index: 2;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding-bottom: 100rpx; /* 为进度条预留空间 */
}

/* 专辑封面 */
.album-cover {
  width: 80%;
  height: 520rpx;
  margin: 0 auto;
  border-radius: 40rpx 40rpx 250rpx 250rpx;
  box-shadow: 0 20rpx 40rpx rgba(0,0,0,0.2);
  position: relative;
  z-index: 2;
  object-fit: cover;
  display: block;
}

/* 线性进度条容器 */
.linear-progress-container {
  width: 90%;
  margin: 0 auto;
  position: absolute;
  bottom: 40rpx; /* 调整到封面底部 */
}

/* 时间显示 */
.time-display {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20rpx;
}

.time-display text {
  font-size: 24rpx;
  color: #666;
  font-variant-numeric: tabular-nums;
}

/* 进度条轨道 */
.linear-progress {
  position: relative;
  height: 16rpx;
  background: rgba(0,0,0,0.1);
  border-radius: 4rpx;
}

/* 进度填充 */
.progress-fill {
  position: absolute;
  height: 100%;
  background: #FFD700;
  border-radius: 13rpx;
  transition: width 0.3s ease;
}

/* 进度指示点 */
.progress-thumb {
  position: absolute;
  width: 24rpx;
  height: 24rpx;
  background: #FFD700;
  border: 4rpx solid #fff;
  border-radius: 80%;
  top: 80%;
  transform: translate(-50%, -50%);
  box-shadow: 0 2rpx 8rpx rgba(0,0,0,0.2);
}

/* 歌曲信息 */
.song-title {
  font-size: 48rpx;
  font-weight: 600;
  color: inherit;
  margin-bottom: 20rpx;
  line-height: 1.2;
}

.song-artist {
  font-size: 32rpx;
  color: rgba(255,255,255,0.9);
  letter-spacing: 1rpx;
}

/* 控制按钮 */
.control-bar {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 60rpx 40rpx 40rpx;
  gap: 40rpx;
  position: relative;
  z-index: 2;
  background: #fff;
  margin-top: -40rpx;
  box-shadow: none;
}

.control-btn {
  width: 120rpx;
  height: 60rpx;
  padding: 10rpx;
  background: transparent;
  border: none;
}

.play-btn {
  width: 160rpx;
  height: 160rpx;
  margin: 0 20rpx;
  background: #FFD700 ;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 8rpx 24rpx rgba(255,215,0,0.3) ;
}

.play-icon {
  width: 80rpx;
  height: 80rpx;
  transition: transform 0.2s ease;
}

.play-btn:active .play-icon {
  transform: scale(0.95);
}

.control-icon {
  width: 40rpx;
  height: 40rpx;
}

.lyrics {
  text-align: center;
  margin: 20rpx 0 60rpx;
  padding: 40rpx 40rpx 60rpx;
  background: #fff;
  border-radius: 0;
  margin-top: -60rpx;
  z-index: 2; /* 提升层级 */
}

.current-lyric {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 5px;
}

.next-lyric {
  font-size: 16px;
  color: #666;
}
</style>
