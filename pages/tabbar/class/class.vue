<template>
  <view class="container">
    <!-- 顶部问候语 -->
    <view class="greeting-section">
      <view class="greeting-text">
        <text class="greeting">早上好</text>
        <text class="emoji">☀️</text>
      </view>
      <text class="sub-greeting">让我们找到您的灵感！</text>
    </view>

    <!-- 顶部工具栏 -->
    <view class="tools-bar">
      <view class="search-btn" @click="goToSearch">
        <image class="search-icon" src="/static/sousuo-3 1.png"></image>
      </view>
      <view class="notification-btn">
        <image class="notification-icon" src="/static/jingqing 1.png"></image>
      </view>
    </view>

    <!-- 今日课程 -->
    <view class="today-class">
      <text class="section-title">今日课程</text>
      <view class="class-card">
        <image class="class-image" src="/static/th.jpg" mode="aspectFill"></image>
        <view class="class-overlay">
          <text class="class-title">如何从零开始</text>
          <text class="class-subtitle">琵琶学习</text>
          <view class="start-btn">
            <text>点击观看</text>
          </view>
        </view>
      </view>
    </view>

    <!-- 请继续课程 -->
    <view class="continue-class">
      <text class="section-title">请继续课程</text>
      <view class="continue-card">
        <view class="progress-section">
          <view class="title-row">
            <text class="continue-title">零基础二胡课程</text>
            <text class="progress-text">12/36</text>
          </view>
          <text class="continue-progress">实课章节：第50课 完整教学</text>
          <view class="progress-bar">
            <view class="progress-inner" :style="{ width: courseProgress + '%' }"></view>
          </view>
        </view>
        <view class="continue-study-btn">
          <text>继续学习</text>
        </view>
      </view>
    </view>

    <!-- 推荐课程 -->
    <view class="recommended-courses">
      <text class="section-title">推荐课程</text>
      <scroll-view class="course-scroll" scroll-x="true">
        <view class="course-list">
          <view class="course-item" v-for="(course, index) in recommendedCourses" :key="index">
            <image :src="course.image" class="course-image" mode="aspectFill"></image>
            <text class="course-name">{{ course.name }}</text>
          </view>
        </view>
      </scroll-view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      recommendedCourses: [
        { image: '/static/Card-2.png', name: '零基础二胡入门' },
        { image: '/static/Card-1.png', name: '零基础古筝教学' },
        { image: '/static/logoo.png', name: '指弹零基础' },
        { image: '/static/Card-1.png', name: '零基础古筝教学' }
      ],
      courseProgress: 61 // 课程进度百分比
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
  padding: 40rpx;
  background: linear-gradient(180deg, #ffffff 0%, #ffffff 15%, #FFFAE8 100%);
  min-height: 100vh;
}

/* 顶部问候语样式 */
.greeting-section {
  margin-bottom: 20rpx;
}

.greeting-text {
  display: flex;
  align-items: center;
  margin-bottom: 8rpx;
}

.greeting {
  font-size: 48rpx;
  font-weight: bold;
  color: #4C2605;
}

.emoji {
  font-size: 48rpx;
  margin-left: 10rpx;
}

.sub-greeting {
  font-size: 28rpx;
  color: #4C2605;
  opacity: 0.8;
}

/* 工具栏样式 */
.tools-bar {
  position: absolute;
  top: 30rpx;
  right: 30rpx;
  display: flex;
  align-items: center;
  gap: 20rpx;
}

.search-btn, .notification-btn {
  width: 80rpx;
  height: 80rpx;
  background: #FFFFFF;
  border-radius: 40rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.06);
}

.search-icon, .notification-icon {
  width: 44rpx;
  height: 44rpx;
  opacity: 0.7;
}

/* 今日课程样式 */
.section-title {
  font-size: 36rpx;
  font-weight: bold;
  color: #4C2605;
  margin-bottom: 20rpx;
}

.class-card {
  position: relative;
  width: 100%;
  height: 360rpx;
  border-radius: 30rpx;
  overflow: hidden;
  margin-bottom: 40rpx;
}

.class-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.class-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 30rpx;
  background: linear-gradient(transparent, rgba(0, 0, 0, 0.7));
}

.class-title {
  font-size: 40rpx;
  color: #FFFFFF;
  font-weight: bold;
  margin-bottom: 10rpx;
  display: block;
}

.class-subtitle {
  font-size: 28rpx;
  color: rgba(255, 255, 255, 0.9);
  display: block;
  margin-bottom: 20rpx;
}

.start-btn {
  display: inline-block;
  padding: 12rpx 40rpx;
  background: #FFB800;
  border-radius: 30rpx;
  color: #FFFFFF;
  font-size: 28rpx;
  width: 160rpx;
  text-align: center;
}

/* 继续课程样式 */
.continue-card {
  background: #FFB800;
  border-radius: 30rpx;
  padding: 30rpx;
  display: flex;
  flex-direction: column;
  margin-bottom: 40rpx;
  position: relative;
  overflow: hidden;
}

.title-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12rpx;
  padding: 0 20rpx;
}

.continue-title {
  font-size: 34rpx;
  color: #FFFFFF;
  font-weight: bold;
  display: block;
}

.progress-text {
  font-size: 26rpx;
  color: #FFFFFF;
  opacity: 0.9;
}

.continue-progress {
  font-size: 26rpx;
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 16rpx;
  display: block;
  padding: 0 20rpx;
}

.progress-section {
  margin-bottom: 16rpx;
}

.progress-bar {
  width: calc(100% - 40rpx);
  height: 8rpx;
  background: rgba(255, 255, 255, 0.3);
  border-radius: 4rpx;
  overflow: hidden;
  margin: 0 20rpx;
}

.progress-inner {
  height: 100%;
  background: #FFFFFF;
  border-radius: 5rpx;
  transition: width 0.3s ease;
}

.continue-study-btn {
  display: inline-block;
  padding: 10rpx 36rpx;
  background: #A36911;
  border-radius: 30rpx;
  color: #FFFFFF;
  font-size: 30rpx;
  margin-left: 20rpx;
  margin-top: 16rpx;
  width: 160rpx;
  text-align: center;
}

/* 推荐课程样式 */
.course-scroll {
  width: 100%;
}

.course-list {
  display: flex;
  padding: 10rpx 0;
}

.course-item {
  margin-right: 20rpx;
  width: 200rpx;
}

.course-image {
  width: 200rpx;
  height: 200rpx;
  border-radius: 20rpx;
  margin-bottom: 10rpx;
}

.course-name {
  font-size: 26rpx;
  color: #4C2605;
  display: block;
}
</style>