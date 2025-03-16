<template>
  <view class="collection-container">
    <view class="header">
      <text class="title">我的收藏</text>
      <text class="count">共{{ collectedSongs.length }}首</text>
      <image 
        src="/static/music-note.png" 
        class="header-icon"
      />
    </view>
    
    <scroll-view 
      class="song-list"
      scroll-y
      :show-scrollbar="false"
    >
      <view 
        class="song-card"
        v-for="(song, index) in collectedSongs"
        :key="song.id"
        @click="playSong(song)"
        :style="{ 
          background: index % 2 === 0 ? '#FFFFFF' : '#F8F9FA',
          transform: 'scale(' + (isCardHovered(index) ? 0.98 : 1) + ')'
        }"
        @touchstart="hoverIndex = index"
        @touchend="hoverIndex = -1"
      >
        <image 
          class="cover"
          :src="song.cover"
          mode="aspectFill"
          style="border-radius: 12rpx;"
        />
        <view class="info">
          <text class="song-title">{{ song.title }}</text>
          <text class="artist">{{ song.artist }}</text>
          <view class="song-tags">
            <text 
              v-for="(tag, i) in song.tags" 
              :key="i"
              class="tag"
            >{{ tag }}</text>
          </view>
        </view>
        <image 
          class="play-icon" 
          src="/static/play-circle.png"
          :style="{ filter: isCardHovered(index) ? 'brightness(0.9)' : 'none' }"
        />
      </view>
      
      <view v-if="collectedSongs.length === 0" class="empty">
        <image class="empty-icon" src="/static/music-note.png"/>
        <text class="empty-text">暂无收藏歌曲</text>
      </view>
    </scroll-view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      currentIndex: -1,
      audioContext: null,
      collectedSongs: [],
      hoverIndex: -1
    }
  },
  methods: {
    goBack() {
      uni.navigateBack();
    },
    formatDuration(seconds) {
      const mins = Math.floor(seconds / 60)
      const secs = seconds % 60
      return `${mins}:${secs.toString().padStart(2, '0')}`
    },
    goToPlayer(song) {
      uni.setStorageSync('currentPlaylist', this.collectedSongs)
      uni.navigateTo({
        url: `/pages/player/player?song=${encodeURIComponent(JSON.stringify(song))}`
      });
    },
    async getAudioDuration(url) {
      return new Promise((resolve) => {
        const audio = uni.createInnerAudioContext();
        audio.src = url;
        audio.onCanplay(() => {
          resolve(Math.round(audio.duration));
          audio.destroy();
        });
        audio.onError(() => {
          resolve(0); // 获取失败时返回0
          audio.destroy();
        });
      });
    },
    async scanMusicFiles() {
      const res = await uni.request({
        url: '/static/music/music-list.json',
        method: 'GET'
      })
      
      this.collectedSongs = res.data.map(item => ({
        ...item,
        cover: `/static/${item.cover}`,
        url: `/static/${item.file}`
      }))
    },
    formatFileName(name) {
      // 将文件名转换为友好格式（例如：song_name -> Song Name）
      return name.replace(/_/g, ' ')
                 .replace(/(^|\s)\S/g, match => match.toUpperCase());
    },
    async validateFile(path) {
      try {
        const res = await uni.getFileSystemManager().stat({ path })
        console.log(`文件信息: ${path}`, res)
        return res.size > 1024 // 文件大小需大于1KB
      } catch {
        return false
      }
    },
    testPlay(url) {
      const testAudio = uni.createInnerAudioContext()
      testAudio.src = url
      testAudio.play()
      testAudio.onPlay(() => {
        uni.showToast({ title: '可正常播放', icon: 'none' })
        testAudio.stop()
        testAudio.destroy()
      })
      testAudio.onError((err) => {
        console.error('播放测试失败:', err)
        uni.showToast({ title: `播放失败: ${err.errMsg}`, icon: 'none' })
        testAudio.destroy()
      })
    },
    playSong(song) {
      this.goToPlayer(song);
    },
    isCardHovered(index) {
      return this.hoverIndex === index;
    }
  },
  mounted() {
    this.scanMusicFiles();
  },
  onUnload() {
    if (this.audioContext) {
      this.audioContext.destroy();
    }
  }
}
</script>

<style>
.collection-container {
  padding: 40rpx;
}

/* 新增头部样式 */
.header {
  position: relative;
  padding: 60rpx 40rpx 40rpx;
  background: linear-gradient(135deg, #FF7675 0%, #FFB800 100%);
  border-radius: 0 0 40rpx 40rpx;
  margin-bottom: 40rpx;
  color: #FFFFFF;
}

.header-icon {
  position: absolute;
  right: 40rpx;
  top: 50%;
  transform: translateY(-50%);
  width: 120rpx;
  height: 120rpx;
  opacity: 0.1;
}

.title {
  font-size: 48rpx;
  font-weight: 700;
  letter-spacing: 1rpx;
}

.count {
  font-size: 28rpx;
  opacity: 0.9;
  margin-top: 12rpx;
}

/* 歌曲列表优化 */
.song-card {
  display: flex;
  align-items: center;
  padding: 32rpx;
  margin: 0 40rpx 24rpx;
  border-radius: 24rpx;
  transition: all 0.3s ease;
  box-shadow: 0 4rpx 16rpx rgba(0,0,0,0.05);
}

.song-card:active {
  transform: scale(0.98);
}

.cover {
  width: 120rpx;
  height: 120rpx;
  border-radius: 16rpx;
  margin-right: 32rpx;
  flex-shrink: 0;
}

.info {
  flex: 1;
  min-width: 0;
}
.song-title {
  font-size: 34rpx;
  color: #2D3436;
  margin-bottom: 8rpx;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.artist {
  font-size: 28rpx;
  color: #636E72;
  margin-bottom: 12rpx;
}

/* 新增标签样式 */
.song-tags {
  display: flex;
  gap: 12rpx;
}

.tag {
  padding: 4rpx 16rpx;
  background: rgba(255,118,117,0.1);
  border-radius: 20rpx;
  font-size: 24rpx;
  color: #FF7675;
}

.play-icon {
  width: 48rpx;
  height: 48rpx;
  transition: all 0.3s ease;
  flex-shrink: 0;
}

/* 空状态样式 */
.empty {
  text-align: center;
  padding: 120rpx 0;
}
.empty-icon {
  width: 200rpx;
  height: 200rpx;
  opacity: 0.2;
  margin-bottom: 24rpx;
}
.empty-text {
  display: block;
  margin-top: 24rpx;
  color: #636E72;
  font-size: 32rpx;
}
</style> 