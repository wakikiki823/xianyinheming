<template>
  <view class="container">
    <!-- 顶部导航栏 -->
    <view class="nav-bar">
      <text class="cancel" @click="goBack">取消</text>
      <text class="title">发布内容</text>
      <text class="publish" @click="publishContent" :class="{ active: canPublish }">发布</text>
    </view>

    <!-- 内容编辑区 -->
    <view class="content-edit">
      <textarea class="text-input" 
                v-model="content" 
                placeholder="分享你的音乐故事..." 
                maxlength="500"
                :adjust-position="false"
      />
      
      <!-- 媒体内容预览 -->
      <view class="media-preview" v-if="mediaList.length > 0">
        <view class="media-item" v-for="(item, index) in mediaList" :key="index">
          <image v-if="item.type === 'image'" 
                 :src="item.url" 
                 mode="aspectFill" 
                 class="preview-image"
          />
          <video v-if="item.type === 'video'" 
                 :src="item.url" 
                 class="preview-video"
          />
          <text class="delete-btn" @click="deleteMedia(index)">×</text>
        </view>
      </view>

      <!-- 添加媒体按钮 -->
      <view class="media-tools">
        <view class="tool-item" @click="chooseImage">
          <image src="/static/image.png" class="tool-icon"></image>
          <text>图片</text>
        </view>
        <view class="tool-item" @click="chooseVideo">
          <image src="/static/video.png" class="tool-icon"></image>
          <text>视频</text>
        </view>
      </view>

      <!-- 音乐选择区域 -->
      <view class="music-card" v-if="selectedMusic" @click="showMusicPicker = true">
        <image :src="selectedMusic.cover" class="music-cover"></image>
        <view class="music-info">
          <text class="music-title">{{ selectedMusic.title }}</text>
          <text class="music-artist">{{ selectedMusic.artist }}</text>
        </view>
        <text class="music-delete" @click.stop="selectedMusic = null">×</text>
      </view>
      
      <view class="add-music" v-else @click="showMusicPicker = true">
        <image src="/static/music-note.png" class="add-icon"></image>
        <text>关联音乐</text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      content: '', // 文字内容
      mediaList: [], // 媒体列表
      isSubmitting: false, // 是否正在提交
      selectedMusic: null, // 选中的音乐信息
      showMusicPicker: false, // 是否显示音乐选择器
      userInfo: uni.getStorageSync('userInfo') || {
        avatar: '/static/default-avatar.png',
        username: '未登录用户'
      }
    }
  },
  computed: {
    canPublish() {
      return !this.isSubmitting;
    }
  },
  onLoad() {
    // 监听返回按钮
    const pages = getCurrentPages();
    const page = pages[pages.length - 1];
    if (typeof page.onBackPress === 'function') {
      page.onBackPress = () => {
        this.checkBeforeLeave();
        return true;
      }
    }
  },
  methods: {
    // 检查是否有未保存内容
    checkBeforeLeave() {
      if (this.content.trim() !== '' || this.mediaList.length > 0) {
        uni.showModal({
          title: '提示',
          content: '是否放弃当前编辑的内容？',
          success: (res) => {
            if (res.confirm) {
              uni.navigateBack();
            }
          }
        });
        return true;
      }
      return false;
    },

    // 返回上一页
    goBack() {
      if (!this.checkBeforeLeave()) {
        uni.navigateBack();
      }
    },

    // 删除媒体
    deleteMedia(index) {
      this.mediaList.splice(index, 1);
    },

    // 新增上传方法
    async uploadFile(file) {
      return new Promise((resolve, reject) => {
        uni.showLoading({ title: '上传中...' });
        const uploadTask = uni.uploadFile({
          url: '/api/upload',
          filePath: file.url,
          name: 'file',
          formData: {
            'user': 'test'
          },
          success: (uploadFileRes) => {
            console.log('上传成功', uploadFileRes.data);
            resolve(uploadFileRes.data.url);
          },
          fail: (err) => {
            console.error('上传失败详情:', {
              errCode: err.errCode,
              errMsg: err.errMsg,
              statusCode: err.statusCode
            });
            uni.showModal({
              content: `上传失败: ${JSON.stringify(err)}`,
              showCancel: false
            });
            reject(err);
          },
          complete: () => {
            uni.hideLoading();
          }
        });
        
        uploadTask.onProgressUpdate = (res) => {
          console.log('上传进度:', res.progress);
          uni.showLoading({ 
            title: `上传中 ${res.progress}%`,
            mask: true
          });
        };
      });
    },

    // 发布内容
    async publishContent() {
      if (!this.canPublish || this.isSubmitting) return;

      this.isSubmitting = true;
      uni.showLoading({
        title: '发布中...',
        mask: true
      });

      try {
        const uploadedMedia = [];
        for (const media of this.mediaList) {
          try {
            const remoteUrl = await this.uploadFile(media); // 替换为真实上传
            uploadedMedia.push(remoteUrl); 
          } catch (err) {
            console.error('上传失败:', err);
            throw err; // 终止发布流程
          }
        }

        // 构建发布数据
        const newPost = {
          id: Date.now(),
          avatar: this.userInfo.avatar,
          username: this.userInfo.username,
          tag: '音乐爱好者',
          content: this.content,
          mediaList: uploadedMedia.map(url => ({
            type: url.includes('.mp4') ? 'video' : 'image',
            url
          })),
          likes: 0,
          hasLiked: false,
          userId: this.userInfo.id,
          createTime: new Date().toISOString()
        };

        // 更新广场页面数据
        const pages = getCurrentPages();
        const communityPage = pages.find(page => page.route === 'pages/tabbar/community/community');
        if (communityPage) {
          communityPage.$vm.posts.unshift(newPost);
        }

        uni.hideLoading();
        uni.showToast({
          title: '发布成功',
          icon: 'success'
        });

        // 返回上一页
        setTimeout(() => {
          uni.navigateBack();
        }, 1500);
      } catch (err) {
        console.error('Publish failed:', err);
        uni.hideLoading();
        uni.showToast({
          title: '发布失败，请重试',
          icon: 'none'
        });
      } finally {
        this.isSubmitting = false;
      }
    },

    // 从音乐库选择
    async selectFromMusicLib() {
      const res = await uni.navigateTo({
        url: '/pages/tabbar/music/music?mode=select'
      });
      // 通过事件通道接收音乐数据
      const eventChannel = this.getOpenerEventChannel();
      eventChannel.on('selectMusic', (music) => {
        this.selectedMusic = music;
      });
    },

    async chooseImage() {
      const res = await uni.chooseImage({
        count: 99,
        sizeType: ['original'],
        sourceType: ['album', 'camera']
      });

      res.tempFilePaths.forEach(path => {
        this.mediaList.push({
          type: 'image',
          url: path
        });
      });
    },

    async chooseVideo() {
      const res = await uni.chooseVideo({
        sourceType: ['album', 'camera'],
        maxDuration: 0,
        compressed: false
      });

      this.mediaList.push({
        type: 'video',
        url: res.tempFilePath
      });
    }
  }
}
</script>

<style>
.container {
  min-height: 100vh;
  background: #FFFFFF;
}

.nav-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20rpx 30rpx;
  border-bottom: 1rpx solid #F5F5F5;
}

.cancel {
  font-size: 32rpx;
  color: #666666;
}

.title {
  font-size: 34rpx;
  font-weight: bold;
  color: #333333;
}

.publish {
  font-size: 32rpx;
  color: #CCCCCC;
}

.publish.active {
  color: #FFB800;
  font-weight: 500;
}

.content-edit {
  padding: 30rpx;
}

.text-input {
  width: 100%;
  height: 240rpx;
  font-size: 30rpx;
  line-height: 1.6;
  color: #333333;
  padding: 0;
}

.media-preview {
  display: flex;
  flex-wrap: wrap;
  gap: 20rpx;
  margin-top: 30rpx;
}

.media-item {
  position: relative;
  width: 220rpx;
  height: 220rpx;
  border-radius: 12rpx;
  overflow: hidden;
}

.preview-image, .preview-video {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.delete-btn {
  position: absolute;
  top: 10rpx;
  right: 10rpx;
  width: 40rpx;
  height: 40rpx;
  background: rgba(0, 0, 0, 0.5);
  color: #FFFFFF;
  border-radius: 20rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
}

.media-tools {
  display: flex;
  gap: 40rpx;
  margin-top: 40rpx;
  padding-top: 20rpx;
  border-top: 1rpx solid #F5F5F5;
}

.tool-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8rpx;
  position: relative;
}

.tool-item:active {
  opacity: 0.7;
}

.tool-icon {
  width: 48rpx;
  height: 48rpx;
  opacity: 0.8;
  transition: all 0.3s ease;
}

.tool-item:active .tool-icon {
  transform: scale(0.95);
}

.preview-video {
  width: 100%;
  height: 100%;
  object-fit: cover;
  background: #000;
}

.delete-btn {
  position: absolute;
  top: 10rpx;
  right: 10rpx;
  width: 40rpx;
  height: 40rpx;
  background: rgba(0, 0, 0, 0.5);
  color: #FFFFFF;
  border-radius: 20rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
  transition: all 0.3s ease;
}

.delete-btn:active {
  background: rgba(0, 0, 0, 0.7);
  transform: scale(0.9);
}

.publish {
  font-size: 32rpx;
  color: #CCCCCC;
  transition: all 0.3s ease;
}

.publish.active {
  color: #FFB800;
  font-weight: 500;
}

.publish.active:active {
  opacity: 0.8;
  transform: scale(0.98);
}

.music-card {
  display: flex;
  align-items: center;
  padding: 20rpx;
  background: #f8f8f8;
  border-radius: 12rpx;
  margin: 20rpx 0;
  position: relative;
}

.music-cover {
  width: 80rpx;
  height: 80rpx;
  border-radius: 8rpx;
  margin-right: 20rpx;
}

.music-info {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.music-title {
  font-size: 28rpx;
  color: #333;
}

.music-artist {
  font-size: 24rpx;
  color: #666;
}

.add-music {
  display: flex;
  align-items: center;
  padding: 20rpx;
  color: #FFB800;
  border: 1rpx dashed #FFB800;
  border-radius: 12rpx;
  margin: 20rpx 0;
}

.add-icon {
  width: 40rpx;
  height: 40rpx;
  margin-right: 15rpx;
}

.music-delete {
  position: absolute;
  right: 20rpx;
  top: 50%;
  transform: translateY(-50%);
  font-size: 40rpx;
  color: #999;
  padding: 10rpx;
}

.search-input {
  user-select: auto;
  pointer-events: auto;
}
</style> 