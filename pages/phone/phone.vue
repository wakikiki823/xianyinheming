<template>
  <view class="container">
    <!-- 上半部分，logo部分 -->
    <view class="logo-container">
      <image class="ellipse" src="@/static/Ellipse 2062.png" mode="aspectFit"></image>
      <image class="logo" src="@/static/litlogo.png" mode="aspectFit"></image>
      <!-- 添加的文字 -->
      <text class="hello-text"></text>
      <text class="welcome-text">欢迎来到<br>拾音和鸣</text>
    </view>

    <!-- 欢迎消息部分 -->
    <image class="welcome-image" src="@/static/huanying 1.png" mode="aspectFit"></image>

    <!-- 下半部分输入框部分 -->
    <view class="bottom-container">
      <!-- 手机号输入框 -->
      <view class="input-container">
        <input class="input" type="number" placeholder="请输入手机号" />
        <button class="send-code-button" @click="sendVerificationCode">发送验证码</button>
      </view>

      <!-- 验证码输入框 -->
      <input class="input" type="number" placeholder="请输入验证码" />

      <!-- 登录按钮 -->
      <button class="login-button" @click="loginWithPhone">登录</button>

      <!-- 忘记密码和注册链接 -->
      <view class="link-container">
        <text class="link-text">忘记密码？</text>
        <text class="link-text">新用户注册</text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    loginWithPhone() {
      uni.navigateTo({
        url: '/pages/chose/chose' // 跳转到主页或其他页面
      });
    },
    sendVerificationCode() {
      console.log('发送验证码');
    }
  }
};
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  width: 100%;
  height: 50%;
  background: #FFE68B; /* 渐变背景 */
  padding-top: 80px; /* 调整上方间距，留出空间 */
  position: relative; /* 使内部绝对定位元素相对此容器定位 */
}

.logo-container {
  position: absolute;
  top: 20px; /* 离屏幕上方20px */
  left: 20px; /* 离屏幕左侧20px */
  width: 100px; /* 缩小Logo尺寸 */
  height: auto;
  z-index: 10;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.ellipse {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 50%; /* 设置圆形背景 */
}

.logo {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 90%; /* 缩小Logo尺寸 */
  height: 90%;
  transform: translate(-50%, -45%);
}

.hello-text {
  width: 94px;
  height: 4px;
  font-family: Impact, Impact;
  font-weight: 400;
  font-size: 26px;
  color: #4C2605;
  line-height: 28px;
  letter-spacing: 1px;
  text-align: center;
  font-style: normal;
  text-transform: none;
  margin-top: 110px;
}

.welcome-text {
  width: 94px;
  height: 4px;
  font-family: Impact, Impact;
  font-weight: 400;
  font-size: 18px;
  color: #4C2605;
  line-height: 28px;
  letter-spacing: 1px;
  text-align: center;
  font-style: normal;
  text-transform: none;
  margin-top: 0px;
}

.welcome-image {
  width: 300px; /* 设置欢迎图片宽度 */
  height: 280px; /* 设置欢迎图片高度 */
  position: absolute;
  top: 80px; /* 将图片稍微下移，避免被白色块遮挡 */
  left: 74%;
  transform: translateX(-50%); /* 将图片水平居中 */
  z-index: 1; /* 设置较低的层级 */
}

.bottom-container {
  width: 393px;
  height: 393px; /* 增大白色部分的高度，调整为 11/17 屏幕高度 */
  background: #FFFFFF;
  border-radius: 100px 100px 20px 20px; /* 设置下半部分的圆角 */
  margin-top: 200px; /* 向下移动白色部分 */
  padding: 20px; /* 内边距 */
  display: flex;
  flex-direction: column;
  justify-content: center; /* 垂直居中 */
  align-items: center; /* 水平居中 */
  z-index: 2; /* 设置较高的层级，使其覆盖 welcome-image */
}

.input-container {
  display: flex;
  justify-content: space-between;
  width: 80%;
}

.input {
  width: 70%; /* 设置输入框宽度为80% */
  height: 50px;
  background-color: #FFEE93; /* 设置背景色 */
  border-radius: 20px;
  padding-left: 15px;
  font-size: 16px;
  border: 1px solid #ccc;
  margin-bottom: 20px; /* 输入框间隔 */
}

.send-code-button {
  width: 30%;
  height: 50px;
  background-color: #FFEE93;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.25);
  border-radius: 20px;
  border: none;
  font-size: 2px;
  font-weight: bold;
  color: #A3710E;
  text-align: center;
  line-height: 50px;
}

.login-button {
  width: 80%;
  height: 50px;
  background-color: #FFEE93;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.25);
  border-radius: 20px;
  border: none;
  font-size: 18px;
  font-weight: bold;
  color: #A3710E;
  text-align: center;
  line-height: 50px;
  margin-bottom: 20px;
}

.link-container {
  display: flex;
  justify-content: space-between;
  width: 80%;
}

.link-text {
  font-size: 14px;
  color: #A3710E;
  cursor: pointer;
}
</style>