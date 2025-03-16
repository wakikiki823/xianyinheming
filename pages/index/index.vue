<template>
  <view class="container">
    <image class="logo" src="@/static/logo.png" mode="aspectFit"></image>
    <text class="xuangui-text">拾音和鸣</text>
    <text class="subtitle">
      普及智能乐器教学
      <br />
      让中国民族乐器走向世界
    </text>
  </view>
</template>

<script>
export default {
  data() {
    return {};
  },
  mounted() {
    console.log("App 加载成功！");
    // 延迟1秒后跳转到login.vue页面
    setTimeout(() => {
      console.log("正在跳转到登录页面...");
      uni.redirectTo({
        url: '/pages/login/login' // 确保路径正确
      });
    }, 1000); // 延迟1秒
  }
}
</script>


<style>
.container {
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* logo稍微偏上 */
  align-items: center;
  width: 100%; /* 容器宽度 */
  height: 100vh; /* 容器高度 */
  background: linear-gradient( 360deg, #FAE599 0%, #F0FEEC 48%, #FFF2D8 100%); /* 上下两端明亮黄色，中间渐变为白色 */
  border-radius: 0px 0px 0px 0px; /* 容器圆角 */
  margin: auto;
  padding-top: 60px; /* logo上方留出空白 */
}

.logo {
  width: 166px; /* logo宽度 */
  height: 158px; /* logo高度 */
  margin-bottom: 20px; /* logo与文字的间距 */
  border-radius: 0px 0px 0px 0px; /* logo圆角 */
}

.xuangui-text {
  width: 190px; /* 设置宽度 */
  height: 52px; /* 设置高度 */
  font-family: FZDeShengTiS, FZDeShengTiS; /* 字体 */
  font-weight: 800; /* 字体加粗 */
  font-size: 40px; /* 字号 */
  color: #F8B91D; /* 字体颜色 */
  line-height: 47px; /* 行高 */
  letter-spacing: 4px; /* 字符间距 */
  text-align: center; /* 居中文本 */
  font-style: normal; /* 正常字体样式 */
  text-transform: none; /* 无大写转换 */
  margin-bottom: 250px; /* 与下方文字的间距 */
}

.subtitle {
 width: 260px;
 height: 36px;
 font-family: FZTiHeiS, FZTiHeiS;
 font-weight: 400;
 font-size: 16px;
 color: #A3710E;
 line-height: 19px;
 letter-spacing: 1px;
 text-align: center;
 font-style: normal;
 text-transform: none;
}

.subtitle br {
  display: block;
  margin-top: 10px; /* 在换行符后增加间距 */
}
</style>
