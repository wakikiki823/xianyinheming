<template>
  <view class="container">
    <view class="title-container">
      <text class="title" >选择你感兴趣的音乐标签</text>
    </view>
    <view class="subtitle-container">
      <text class="subtitle">你的选择将帮助我们更精准地推荐音乐和课程</text>
    </view>

    <view class="tags">
      <!-- 使用v-for动态生成标签 -->
      <view 
        class="tag" 
        v-for="(tag, index) in tags" 
        :key="index"
        :class="{'selected': selectedTags.includes(tag)}" 
        :style="getTagStyle(index)" 
        @click="toggleTag(tag)"
      >
        <text class="tag-text" :style="getTextStyle(index)">{{ tag }}</text>
      </view>
    </view>

    <button 
      class="next-button" 
      :disabled="selectedTags.length === 0"
      @click="goToNextStep"
    >
      下一步 (已选{{ selectedTags.length }}个)
    </button>
  </view>
</template>

<script>
export default {
  data() {
    return {
      tags: ['流行', '摇滚', '经典', '说唱', '爵士', '欧美', '新世纪', '古典'], // 标签数据
      selectedTags: [] // 已选择的标签
    };
  },
  methods: {
    // 切换标签的选择状态
    toggleTag(tag) {
      if (this.selectedTags.includes(tag)) {
        // 取消选择
        this.selectedTags = this.selectedTags.filter(item => item!== tag);
      } else {
        // 添加到已选
        this.selectedTags.push(tag);
      }
    },
    // 获取标签的动态样式（宽高不同）
    getTagStyle(index) {
      const tagSizes = [
        { width: '175px', height: '162px' },
        { width: '72px', height: '67px' },
        { width: '136px', height: '126px' },
        { width: '61px', height: '59px' },
        { width: '82px', height: '76px' },
        { width: '109px', height: '110px' },
        { width: '205px', height: '196px' },
        { width: '150px', height: '139px' }
      ];
      const size = tagSizes[index];
      return {
        width: size.width,
        height: size.height,
        background: 'radial-gradient(circle, #FEBB01 0%, #FFE08B 50%, #ffffff 100%)',
        borderRadius: '50%',
      };
    },
    // 获取标签文字的动态样式（字体大小不同）
    getTextStyle(index) {
      const tagSizes = [
        { width: '175px', height: '162px' },
        { width: '72px', height: '67px' },
        { width: '136px', height: '126px' },
        { width: '61px', height: '59px' },
        { width: '82px', height: '76px' },
        { width: '109px', height: '110px' },
        { width: '205px', height: '196px' },
        { width: '150px', height: '139px' }
      ];
      const size = tagSizes[index];
      // 根据标签的宽度计算字体大小
      const fontSize = Math.min(parseInt(size.width) * 0.3, 30);
      return {
        fontSize: `${fontSize}px`
      };
    },
   // 下一步按钮点击时的跳转逻辑
   goToNextStep() {
     if (this.selectedTags.length > 0) {
       // 如果是跳转到 tabBar 页面，使用 uni.switchTab
       uni.switchTab({
         url: '/pages/tabbar/home/home', // 跳转到 tabBar 页面
         success: () => {
           console.log('跳转成功');
         },
         fail: (err) => {
           console.error('跳转失败', err);
         }
       });
     } else {
       uni.showToast({
         title: '请先选择至少一个标签',
         icon: 'none'
       });
     }
   }
  }
};
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  /* justify-content: center; 移除原有的水平居中 */
  align-items: center;
  width: 100%;
  height: 100%;
  background: #FFE68B; 
  padding-top: 40px;
}

.title-container {
  width: 100%;
  text-align: left;
  margin-bottom: 5px; /* 增加与下方元素的距离 */
}

.title {
  font-size: 24px;
  font-weight: bold;
  color: #4C2605;
  margin-bottom: 30px;
}

.subtitle-container {
  width: 100%;
  text-align: left;
  margin-bottom: 40px; /* 增加与下方元素的距离 */
}

.subtitle {
  font-size: 16px;
  color: #3E3E3E;
  margin-bottom: 70px;
}

.tags {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 10px; /* 缩小标签之间的间距 */
}

.tag {
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease; /* 平滑过渡效果 */
  border-radius: 50%;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1); /* 添加阴影 */
}

.tag.selected {
  background: #FEBB01; /* 选中状态背景色 */
  transform: scale(1.2); /* 放大效果 */
  box-shadow: 0px 8px 15px rgba(0, 0, 0, 0.3); /* 放大时的阴影效果 */
}

.tag-text {
  color: #643812;
  font-weight: bold;
}

.next-button {
  position: fixed;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  width: 288px;
  height: 52px;
  background: #ffaa00;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.25);
  border-radius: 20px;
  font-size: 18px;
  font-weight: bold;
  color: #000000;
  line-height: 52px;
  z-index: 1000;
}

@supports (bottom: env(safe-area-inset-bottom)) {
  .next-button {
    bottom: calc(20px + env(safe-area-inset-bottom));
  }
}

.next-button:disabled {
  background-color: #000000;
  color: #000000;
  cursor: not-allowed;
}
</style>