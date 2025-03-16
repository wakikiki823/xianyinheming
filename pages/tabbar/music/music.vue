<template>
  <view class="container">
    <!-- 搜索栏 -->
    <view class="search-box">
      <view class="search-bar">
        <input 
          type="text" 
          v-model="searchText" 
          placeholder="搜索" 
          class="search-input"
        />
        <view 
          class="search-btn-wrapper"
          @click.stop="goToSearch"
        >
          <image 
            class="search-btn" 
            src="/static/huaban13 1.png"
          ></image>
        </view>
      </view>
    </view>

    <!-- 为您推荐 -->
    <view class="recommend-section">
      <view class="section-title">为您推荐</view>
      <view class="recommend-list">
        <!-- 推荐卡片1 -->
        <view class="recommend-card">
          <image class="recommend-image" src="/static/jhk-1717082281905 1.png" mode="aspectFill"></image>
        </view>
        <!-- 推荐卡片2 -->
        <view class="recommend-card">
          <image class="recommend-image" src="/static/11.44 1.png" mode="aspectFill"></image>
        </view>
        <!-- 推荐卡片3 -->
        <view class="recommend-card">
          <image class="recommend-image" src="/static/11.45 1.png" mode="aspectFill"></image>
        </view>
      </view>
    </view>

    <!-- 曲风分类 -->
    <view class="genre-section">
      <view class="section-title">曲风分类</view>
      <view class="genre-grid">
        <!-- 摇滚 -->
        <view class="genre-item">
          <image class="genre-image" src="/static/2024-05-30 11.53 1.png" mode="aspectFill"></image>
          <text class="genre-name">摇滚</text>
          <text class="genre-name">摇滚音乐</text>
        </view>
        <!-- 乡村 -->
        <view class="genre-item">
          <image class="genre-image" src="/static/2024-05-30 11.54 1.png" mode="aspectFill"></image>
          <text class="genre-name">乡村</text>
          <text class="genre-name">乡村音乐</text>
        </view>
        <!-- 世界音乐 -->
        <view class="genre-item">
          <image class="genre-image" src="/static/2024-05-30 11.55 1.png" mode="aspectFill"></image>
          <text class="genre-name">世界音乐</text>
          <text class="genre-name">世界音乐</text>
        </view>
        <!-- 民谣 -->
        <view class="genre-item">
          <image class="genre-image" src="/static/2024-05-30 11.56 1.png" mode="aspectFill"></image>
          <text class="genre-name">民谣</text>
          <text class="genre-name">民谣音乐</text>
        </view>
        <!-- 嘻哈 -->
        <view class="genre-item">
          <image class="genre-image" src="/static/2024-05-30 11.57 1.png" mode="aspectFill"></image>
          <text class="genre-name">嘻哈</text>
          <text class="genre-name">嘻哈音乐</text>
        </view>
        <!-- 说唱 -->
        <view class="genre-item">
          <image class="genre-image" src="/static/2024-05-30 11.59 1.png" mode="aspectFill"></image>
          <text class="genre-name">说唱</text>
          <text class="genre-name">说唱音乐</text>
        </view>
        <!-- 流行 -->
        <view class="genre-item">
          <image class="genre-image" src="/static/2024-05-31 12.00 1.png" mode="aspectFill"></image>
          <text class="genre-name">流行</text>
          <text class="genre-name">流行音乐</text>
        </view>
        <!-- 古典 -->
        <view class="genre-item">
          <image class="genre-image" src="/static/2024-05-31 12.01 1.png" mode="aspectFill"></image>
          <text class="genre-name">古典</text>
          <text class="genre-name">古典音乐</text>
          
        </view>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      searchText: ''
    }
  },
  methods: {
    goToSearch() {
      uni.navigateTo({
        url: '/uni_modules/yt-searchPage/pages/HM-search/HM-search'
      })
    }
  }
}
</script>

<style>
.container {
  padding: 30rpx 20rpx;
  background: linear-gradient(180deg, #FFF6E5 0%, #FFFFFF 100%);
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  padding-bottom: 0;
}

/* 搜索栏样式 */
.search-box {
  padding: 20rpx 30rpx;
}

.search-bar {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  background: #FFFFFF;
  height: 90rpx;
  border-radius: 45rpx;
  padding: 0 20rpx 0 30rpx;
  box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.06);
}

.search-btn-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 80rpx;
  height: 80rpx;
  cursor: pointer;
  transition: all 0.3s ease;
}

.search-btn-wrapper:active {
  transform: scale(0.9);
}

.search-btn {
  width: 48rpx;
  height: 48rpx;
  opacity: 0.7;
  transition: opacity 0.3s ease;
}

.search-btn-wrapper:active .search-btn {
  opacity: 1;
}

.search-input {
  flex: 1;
  height: 100%;
  font-size: 28rpx;
  background-color: transparent;
  border: none;
  padding: 0 20rpx;
  color: #333;
  &::placeholder {
    color: #999;
  }
}

/* 推荐区域样式 */
.recommend-section {
  margin-top: 30rpx;
  padding: 0 20rpx;
}

.section-title {
  font-family: 'PingFang SC', sans-serif;
  font-size: 36rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 25rpx;
  position: relative;
  padding-left: 24rpx;
}

.section-title::before {
  content: '';
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 8rpx;
  height: 32rpx;
  background: #FFB800;
  border-radius: 4rpx;
}

.recommend-list {
  white-space: nowrap;
  margin: 0 -20rpx;
  padding: 10rpx 20rpx;
}

.recommend-card {
  display: inline-block;
  width: 260rpx;
  height: 340rpx;
  margin-right: 24rpx;
  border-radius: 16rpx;
  overflow: hidden;
  box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
}

.recommend-card:active {
  transform: scale(0.98);
}

.recommend-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* 曲风分类样式 */
.genre-section {
  margin-top: 40rpx;
  padding: 0;
  margin-bottom: 0;
  margin-left: -20rpx;
  margin-right: -20rpx;
  flex: 1;
}

.genre-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  background: linear-gradient(180deg, #FFD028 0%, #FFB800 100%);
  border-radius: 30rpx 30rpx 0 0;
  padding: 30rpx 40rpx 20rpx;
  min-height: 500rpx;
  box-shadow: 0 -4rpx 16rpx rgba(0, 0, 0, 0.05);
}

.genre-item {
  width: 23%;
  margin-bottom: 30rpx;
  text-align: center;
  transition: all 0.3s ease;
}

.genre-item:active {
  transform: scale(0.95);
}

.genre-image {
  width: 110rpx;
  height: 110rpx;
  border-radius: 20rpx;
  margin-bottom: 12rpx;
  box-shadow: 0 4rpx 8rpx rgba(0, 0, 0, 0.1);
}

.genre-name {
  font-size: 24rpx;
  color: #333;
  display: block;
  line-height: 1.3;
  margin-bottom: 4rpx;
}

.genre-name:last-child {
  color: #666;
  font-size: 22rpx;
}

/* 隐藏滚动条但保持滚动功能 */
.recommend-list {
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none;
}

.recommend-list::-webkit-scrollbar {
  display: none;
}
</style>