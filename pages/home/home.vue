<!-- 搜索框容器 -->
<view class="search-container">
  <view class="search-box">
    <input 
      class="search-input"
      placeholder="搜索歌曲"
      placeholder-class="placeholder-style"
    />
    <button 
      class="search-btn"
      style="background: #fff; width: 100rpx; height: 100rpx;"
    >
      <image 
        src="/static/huaban13 1.png" 
        class="search-icon"
        style="width: 50rpx; height: 50rpx;"
      />
    </button>
  </view>
</view>

<style>
.search-container {
  padding: 40rpx;
  background: #fff;
  border-radius: 24rpx;
  box-shadow: 0 8rpx 24rpx rgba(0,0,0,0.06);
}

.search-box {
  display: flex;
  align-items: center;
  background: #f8f9fa;
  border-radius: 50rpx;
  padding: 20rpx 30rpx;
}

.search-btn {
  transition: transform 0.2s ease;
}

.search-btn:active {
  transform: scale(0.95);
}

.search-icon {
  width: 40rpx;
  height: 40rpx;
}
</style> 