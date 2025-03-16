<template>
  <view class="container">
    <!-- 搜索栏部分 -->
    <view class="search-bar">
      <input 
        type="text" 
        v-model="searchQuery" 
        placeholder="搜索" 
        class="search-input"
      />
      <view 
        class="search-btn-wrapper"
        @click.stop="openSearchPage"
      >
        <image 
          class="search-btn" 
          :src="searchBtnIcon"
        ></image>
      </view>
    </view>

    <!-- 个性推荐部分 -->
    <view class="personal-recommendation">
      <view class="title">个性推荐</view>
      <!-- 使用 scroll-view 实现左右滑动 -->
      <scroll-view class="recommendation-scroll" scroll-x="true">
        <view class="recommendation-items">
          <view class="item" v-for="(playlist, index) in playlists" :key="index">
            <image :src="playlist.image" class="playlist-image" />
            <view class="playlist-title">{{ playlist.title }}</view>
            <view class="playlist-description">{{ playlist.description }}</view>
          </view>
        </view>
      </scroll-view>
    </view>

    <!-- 在个性推荐后添加 -->
    <view class="hot-recommend">
      <view class="section-title">
        <text>热门推荐</text>
        <view class="more" @click="goToHotList">更多 ></view>
      </view>
      <scroll-view class="hot-scroll" scroll-x>
        <view 
          v-for="(song, index) in hotSongs" 
          :key="index" 
          class="hot-item"
          @click="playHotSong(song)"
        >
          <image :src="song.cover" class="hot-cover"></image>
          <view class="hot-info">
            <text class="hot-title">{{ song.title }}</text>
            <text class="hot-artist">{{ song.artist }}</text>
          </view>
        </view>
      </scroll-view>
    </view>

    <!-- 新增功能部分 -->
    <view class="extra-column">
      <view class="extra-title"></view>
      <view class="extra-content">
        <!-- 训练入口 -->
        <view class="extra-item" @click="navigateTo('/pages/training-entry')">
          <image :src="'/static/logo.png'" class="training-icon" />
          <view class="extra-text">训练入口</view>
        </view>

        <!-- 每项功能 -->
        <view class="extra-item" v-for="(item, index) in extraItems" :key="index" @click="navigateTo(item.link)">
          <image :src="item.icon" class="extra-icon" />
          <view class="extra-text">{{ item.title }}</view>
        </view>
      </view>
    </view>

    <!-- 推荐课程部分 -->
    <view class="recommended-courses">
      <view class="title">推荐课程</view>
      <!-- 使用 scroll-view 实现左右滑动 -->
      <scroll-view class="courses-scroll" scroll-x="true">
        <view class="courses-items">
          <view class="course" v-for="(course, index) in recommendedCourses" :key="index">
            <image :src="course.image" class="course-image" />
            <view class="course-title">{{ course.title }}</view>
          </view>
        </view>
      </scroll-view>
    </view>

    <!-- 训练建议部分 -->
    <view class="training-suggestions">
      <view class="title">训练建议</view>
      <view class="suggestion-card">
        <view class="suggestion-level">{{ trainingSuggestion.level }}</view>
        <view class="suggestion-description">
          {{ trainingSuggestion.description }}
        </view>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      searchQuery: "",  // 搜索框中的内容
      searchBtnIcon: '/static/huaban13 1.png',  // 搜索按钮图标路径
      trainingIcon: '/static/logo.png', // 训练入口图标

      // 新增功能栏的内容
      extraItems: [
        { icon: '/static/paihangbang 1.png', title: '排行榜', link: '/pages/ranking' },
        { icon: '/static/gedan 1.png', title: '歌单', link: '/pages/playlist' },
        { icon: '/static/shipinbofang 1.png', title: '视频播放', link: '/pages/video' },
        { icon: '/static/xunlianjilu 1.png', title: '训练记录', link: '/pages/training-record' },
      ],

      // 初始 play list 和新增的列表一起显示
      playlists: [
        { image: '/static/Thumbnail.png', title: '每日推荐', description: '符合你口味的新鲜好歌' },
        { image: '/static/Card.png', title: '私人漫游', description: '多种模式随心播放' },
        { image: '/static/playlist3.jpg', title: 'Playlist 3', description: '其他推荐内容' },
        { image: '/static/Card.png', title: '私人漫游', description: '多种模式随心播放' },  // 新增项
      ],
      recommendedCourses: [
        { image: '/static/Card-2.png', title: '零基础二胡入门' },
        { image: '/static/Card-1.png', title: '零基础古筝教学' },
        { image: '/static/C.png', title: '指弹零基础' },
		{ image: '/static/Card-1.png', title: '零基础古筝教学' },
      ],
      // 训练建议改为一项
      trainingSuggestion: { level: '1级', title: '训练建议 1', description: '学习基础知识，逐步掌握节奏感，练习一些简单的曲目。' },
    };
  },
  methods: {
    // 跳转到搜索页面，并传递搜索框内容
    openSearchPage() {
      // 确保searchQuery有值，才进行跳转
      if (this.searchQuery.trim()) {
        uni.navigateTo({
          url: `/uni_modules/yt-searchPage/pages/HM-search/HM-search?searchQuery=${encodeURIComponent(this.searchQuery)}`
        });
      } else {
        uni.showToast({
          title: '请输入搜索内容',
          icon: 'none'
        });
      }
    },
    // 跳转到目标页面
    navigateTo(link) {
      uni.navigateTo({
        url: link,
      });
    },
    playHotSong(song) {
      uni.setStorageSync('currentPlaylist', this.hotSongs)
      uni.navigateTo({
        url: `/pages/player/player?song=${encodeURIComponent(JSON.stringify(song))}`
      })
    }
  },
};
</script>

<style scoped>
.container {
  padding: 20px;
  background: linear-gradient(180deg, #FCD159 0%, #FFEE93 15%, #FFFAE8 100%);
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

/* 标题样式 */
.title {
  width: 79px;
  height: 24px;
  font-family: FZLanTingHeiS-B-GB, FZLanTingHeiS-B-GB;
  font-weight: 400;
  font-size: 20px;
  color: #4C2605;
  line-height: 24px;
  text-align: left;
  font-style: normal;
  text-transform: none;
  white-space: nowrap; /* 不换行 */
}

/* 个性推荐部分 */
.personal-recommendation .playlist-image {
  width: 156px; /* 固定宽度 */
  height: 156px;
  background: #B5DBFF;
  border-radius: 8px;
}

.playlist-title, .playlist-description {
  margin-top: 10px;
  font-size: 14px;
}

.playlist-description {
  font-size: 12px;
  color: #888;
}

/* 使用 scroll-view 实现横向滑动 */
.recommendation-scroll, .courses-scroll {
  display: flex;
  overflow-x: scroll;
  white-space: nowrap;
  width: 100%;
}

.recommendation-items, .courses-items {
  display: flex;
}

.item, .course {
  display: inline-block;
  margin-right: 10px; /* 图片之间的间距 */
  text-align: center;
  width: 180px; /* 固定宽度，确保每个图片的宽度不变 */
}

.item .playlist-image {
  width: 156px;
  height: 156px;
}

.course .course-image {
  width: 100px; /* 固定宽度 */
  height: 100px;
}

.playlist-title, .course-title {
  margin-top: 8px;
  font-size: 14px;
}

.playlist-description {
  font-size: 12px;
  color: #888;
}

/* 新增功能栏 */
.extra-column {
  margin-bottom: 20px;
}

.extra-title {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 10px;
}

.extra-content {
  display: flex;
  justify-content: space-between;
}

.extra-item {
  text-align: center;
  width: 22%;
}

.extra-icon {
  width: 40px;
  height: 40px;
  margin-bottom: 5px;
}

.extra-text {
  font-size: 12px;
}

/* 训练建议部分 */
.suggestion-card {
  background: #FFFAE9;
  padding: 15px;
  border-radius: 8px;
  margin-top: 10px;
}

.suggestion-level {
  font-size: 16px;
  font-weight: bold;
  color: #4C2605;
}

.suggestion-description {
  margin-top: 10px;
  font-size: 14px;
  color: #4C2605;
}

/* 训练入口图标 */
.training-icon {
  width: 58px; /* 保证logo大小不变 */
  height: 58px;
  border-radius: 0;
  margin-bottom: 5px;
}

.hot-recommend {
  margin: 40rpx 0;
}

.hot-scroll {
  white-space: nowrap;
  padding: 20rpx 0;
}

.hot-item {
  display: inline-block;
  width: 280rpx;
  margin-right: 30rpx;
  background: rgba(255,255,255,0.9);
  border-radius: 16rpx;
  overflow: hidden;
}

.hot-cover {
  width: 100%;
  height: 280rpx;
}

.hot-info {
  padding: 20rpx;
}

.hot-title {
  font-size: 28rpx;
  color: #333;
}

.hot-artist {
  font-size: 24rpx;
  color: #666;
}
</style>
