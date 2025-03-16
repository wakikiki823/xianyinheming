<template>
  <view class="container">
    <!-- 顶部导航栏 -->
    <view class="nav-bar">
      <view class="nav-tabs">
        <text class="nav-item" :class="{ active: currentTab === 'friends' }" @click="switchTab('friends')">朋友圈</text>
        <text class="nav-item" :class="{ active: currentTab === 'square' }" @click="switchTab('square')">广场</text>
        <text class="nav-item" :class="{ active: currentTab === 'follow' }" @click="switchTab('follow')">关注</text>
      </view>
      <view class="search-post">
        <image class="search-icon" src="/static/sousuo-3 1.png" @click="goToSearch"></image>
        <view class="post-btn" @click="goToPost">发布</view>
      </view>
    </view>

    <!-- 分类标签 -->
    <scroll-view class="category-scroll" scroll-x="true" show-scrollbar="false">
      <view class="category-list">
        <view v-for="(item, index) in categories" 
              :key="index" 
              class="category-item" 
              :class="{ active: currentCategory === item }"
              @click="switchCategory(item)">
          {{ item }}
        </view>
      </view>
    </scroll-view>

    <!-- 内容列表 -->
    <scroll-view class="content-scroll" scroll-y="true">
      <!-- 内容卡片1 -->
      <view class="content-card" v-for="post in posts" :key="post.id">
        <view class="user-info">
          <image class="avatar" :src="post.avatar" mode="aspectFill"></image>
          <view class="user-detail">
            <view class="username">{{ post.username }}</view>
            <view class="user-tag">{{ post.tag }}</view>
          </view>
          <view class="follow-btn" @click="handleFollow(post)">去看看</view>
          <view 
            class="delete-btn" 
            @click="deletePost(post.id)"
            style="right: 20rpx; top: 20rpx;"
          >
            <image src="/static/shanchu.png" class="delete-icon"></image>
          </view>
        </view>
        <view class="post-content">
          <text class="post-text">{{ post.content }}</text>
          <image class="post-image" :src="post.image" mode="aspectFill"></image>
        </view>
        <view class="interaction-bar">
          <view class="interaction-item" @click="handleLike(post)">
            <image class="icon" :src="post.hasLiked ? '/static/dianzan (1).png' : '/static/dianzan.png'"></image>
            <text>{{ post.likes }}</text>
          </view>
          <view class="interaction-item" @click="handleComment(post)">
            <image class="icon" src="/static/pinglun.png"></image>
          </view>
          <view class="interaction-item" @click="handleShare(post)">
            <image class="icon" src="/static/fenxiang.png"></image>
          </view>
        </view>
      </view>
    </scroll-view>
  </view>
</template>

<script>
import DeleteIcon from '@/components/DeleteIcon.vue';

export default {
  components: {
    DeleteIcon
  },
  data() {
    return {
      currentTab: 'square',
      currentCategory: '全部',
      categories: ['全部', '直播', '吉他', '二胡'],
      posts: [
        {
          id: 1,
          avatar: '/static/th.jpg',
          username: 'CC 🔥',
          tag: '诚信认证五年',
          content: '一曲古韵的魅力，穿越千年',
          image: '/static/th.jpg',
          likes: 27,
          hasLiked: false
        },
        {
          id: 2,
          avatar: '/static/th.jpg',
          username: '雪见乐的mmm',
          tag: '创作达人',
          content: '奇怪的民族乐器大赏',
          image: '/static/2024-06-01 1.13 1.png',
          likes: 47,
          hasLiked: false
        }
      ],
      userInfo: uni.getStorageSync('userInfo') || {}
    }
  },
  methods: {
    // 切换顶部主导航
    switchTab(tab) {
      this.currentTab = tab;
      // 这里可以根据不同tab加载不同的数据
      uni.showToast({
        title: '切换到' + this.getTabName(tab),
        icon: 'none'
      });
    },
    
    // 获取标签名称
    getTabName(tab) {
      const tabNames = {
        'friends': '朋友圈',
        'square': '广场',
        'follow': '关注'
      };
      return tabNames[tab] || '';
    },

    // 切换分类
    switchCategory(category) {
      this.currentCategory = category;
      // 这里可以根据分类加载对应的数据
      uni.showToast({
        title: '切换到' + category,
        icon: 'none'
      });
    },

    // 点击搜索
    goToSearch() {
      uni.navigateTo({
        url: '/uni_modules/yt-searchPage/pages/HM-search/HM-search'
      });
    },

    // 点击发布
    goToPost() {
      uni.navigateTo({
        url: '/pages/post/post'
      });
    },

    // 点击关注/去看看
    handleFollow(post) {
      uni.showToast({
        title: '操作成功',
        icon: 'none'
      });
    },

    // 点赞
    handleLike(post) {
      post.hasLiked = !post.hasLiked;
      post.likes += post.hasLiked ? 1 : -1;
      uni.showToast({
        title: post.hasLiked ? '点赞成功' : '取消点赞',
        icon: 'none'
      });
    },

    // 评论
    handleComment(post) {
      uni.navigateTo({
        url: `/pages/comment/comment?id=${post.id}`
      });
    },

    // 分享
    handleShare(post) {
      uni.showShareMenu({
        withShareTicket: true,
        menus: ['shareAppMessage', 'shareTimeline']
      });
    },

    deletePost(postId) {
      uni.showModal({
        title: '确认删除',
        content: '确定要删除这条内容吗？',
        success: (res) => {
          if (res.confirm) {
            this.posts = this.posts.filter(post => post.id !== postId);
            // 实际项目这里需要调用API删除服务器数据
          }
        }
      });
    }
  }
}
</script>

<style>
.container {
  min-height: 100vh;
  background-color: #F8F8F8;
  padding-bottom: 120rpx;
}

/* 顶部导航栏样式 */
.nav-bar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background: #FFFFFF;
  padding: 20rpx 30rpx;
  display: flex;
  justify-content: space-between;
  align-items: center;
  z-index: 100;
  box-shadow: 0 2rpx 10rpx rgba(0, 0, 0, 0.05);
}

.nav-tabs {
  display: flex;
  gap: 60rpx;
}

.nav-item {
  font-size: 34rpx;
  color: #999999;
  position: relative;
  padding: 10rpx 0;
  transition: all 0.3s ease;
}

.nav-item.active {
  color: #333333;
  font-weight: bold;
  font-size: 36rpx;
}

.nav-item.active::after {
  content: '';
  position: absolute;
  bottom: -6rpx;
  left: 50%;
  transform: translateX(-50%);
  width: 48rpx;
  height: 6rpx;
  background: #FFB800;
  border-radius: 3rpx;
}

.search-post {
  display: flex;
  align-items: center;
  gap: 20rpx;
}

.search-icon {
  width: 44rpx;
  height: 44rpx;
  opacity: 0.7;
  cursor: pointer;
  transition: all 0.3s ease;
}

.search-icon:active {
  opacity: 0.5;
  transform: scale(0.95);
}

.post-btn {
  padding: 12rpx 36rpx;
  background: linear-gradient(135deg, #FFB800 0%, #FF9500 100%);
  border-radius: 30rpx;
  color: #FFFFFF;
  font-size: 28rpx;
  font-weight: 500;
  box-shadow: 0 4rpx 8rpx rgba(255, 184, 0, 0.2);
}

/* 分类标签样式 */
.category-scroll {
  position: fixed;
  top: 100rpx;
  left: 0;
  right: 0;
  background: #FFFFFF;
  padding: 24rpx 0;
  white-space: nowrap;
  z-index: 99;
  box-shadow: 0 2rpx 10rpx rgba(0, 0, 0, 0.03);
}

.category-list {
  display: inline-flex;
  padding: 0 30rpx;
  gap: 20rpx;
}

.category-item {
  padding: 12rpx 36rpx;
  background: #F5F5F5;
  border-radius: 30rpx;
  font-size: 28rpx;
  color: #666666;
  transition: all 0.3s ease;
}

.category-item.active {
  background: linear-gradient(135deg, #FFB800 0%, #FF9500 100%);
  color: #FFFFFF;
  box-shadow: 0 4rpx 8rpx rgba(255, 184, 0, 0.2);
}

/* 内容列表样式 */
.content-scroll {
  margin-top: 180rpx;
  padding: 20rpx 30rpx;
}

.content-card {
  background: #FFFFFF;
  border-radius: 24rpx;
  padding: 36rpx;
  margin-bottom: 30rpx;
  box-shadow: 0 4rpx 16rpx rgba(0, 0, 0, 0.04);
}

.user-info {
  position: relative;
  display: flex;
  align-items: center;
  margin-bottom: 10rpx;
}

.avatar {
  width: 88rpx;
  height: 88rpx;
  border-radius: 44rpx;
  margin-right: 24rpx;
  border: 4rpx solid #FFFFFF;
  box-shadow: 0 4rpx 8rpx rgba(0, 0, 0, 0.1);
}

.user-detail {
  flex: 1;
}

.username {
  font-size: 32rpx;
  font-weight: bold;
  color: #333333;
  margin-bottom: 8rpx;
}

.user-tag {
  font-size: 24rpx;
  color: #999999;
  background: #F8F8F8;
  padding: 4rpx 16rpx;
  border-radius: 20rpx;
  display: inline-block;
}

.follow-btn {
  padding: 12rpx 28rpx;
  background: #F5F5F5;
  border-radius: 30rpx;
  font-size: 26rpx;
  margin-bottom: -78rpx;
  color: #666666;
  transition: all 0.3s ease;
}

.follow-btn:active {
  background: #EBEBEB;
}

.post-content {
  margin-bottom: 20rpx;
}

.post-text {
  font-size: 30rpx;
  color: #333333;
  margin: 24rpx 0;
  line-height: 1.6;
}

.post-image {
  width: 100%;
  height: 420rpx;
  border-radius: 16rpx;
  box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.06);
}

.interaction-bar {
  display: flex;
  gap: 50rpx;
  margin-top: 24rpx;
  padding-top: 24rpx;
  border-top: 2rpx solid #F5F5F5;
}

.interaction-item {
  display: flex;
  align-items: center;
  gap: 10rpx;
  font-size: 28rpx;
  color: #999999;
  transition: all 0.3s ease;
  cursor: pointer;
}

.interaction-item:active {
  opacity: 0.8;
}

.icon {
  width: 40rpx;
  height: 40rpx;
  opacity: 0.8;
}

.nav-item, .category-item {
  cursor: pointer;
}

.nav-item:active, .category-item:active {
  opacity: 0.8;
}

.post-btn {
  cursor: pointer;
}

.post-btn:active {
  transform: scale(0.98);
}

.delete-btn {
  position: absolute;
  right: 20rpx;
  top: 20rpx;
  z-index: 2;
  padding: 8rpx;
  background: rgba(255, 255, 255, 0.5);
  border-radius: 50%;
}

.delete-icon {
  width: 40rpx;
  height: 40rpx;
  transition: transform 0.2s ease;
}

.delete-btn:active .delete-icon {
  transform: scale(0.9);
  opacity: 0.8;
}
</style>