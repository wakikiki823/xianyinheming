<template>
  <view class="container">
    <!-- logo部分 -->
    <image class="logo" src="@/static/Group 10.png" mode="aspectFit"></image>
    <text class="title"></text>

    <!-- 手机号登录按钮 -->
    <button class="login-button" @click="loginWithPhone">手机号登录</button>
    
    <!-- 立即体验按钮 -->
    <button class="experience-button" @click="experience">立即体验</button>
    
    <!-- 第三方登录按钮 -->
    <view class="third-party">
      <image class="third-party-button" src="@/static/weixin 1.png" @click="loginWithWeChat" mode="aspectFit"></image>
      <image class="third-party-button" src="@/static/QQ 1.png" @click="loginWithQQ" mode="aspectFit"></image>
      <image class="third-party-button" src="@/static/weibo 1.png" @click="loginWithWeibo" mode="aspectFit"></image>
    </view>
    
    <!-- 其他登录方式 -->
    <text class="other-login">其他登录方式</text>
  </view>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    // 手机号登录
    loginWithPhone() {
      console.log("触发手机号登录跳转");
      uni.navigateTo({
        url: '/pages/phone/phone',
        success: function() {
          console.log("跳转成功");
        },
        fail: function(err) {
          console.error("跳转失败", err);
        }
      });
    }
,
    // 立即体验
    experience() {
      uni.navigateTo({
        url: '/pages/chose/chose' // 进入应用主页
      });
    },
    // 微信登录
    loginWithWeChat() {
      uni.login({
        provider: 'weixin',
        success: res => {
          console.log('微信登录成功', res);
        },
        fail: err => {
          console.error('微信登录失败', err);
        }
      });
    },
    // QQ登录
    loginWithQQ() {
      uni.login({
        provider: 'qq',
        success: res => {
          console.log('QQ登录成功', res);
        },
        fail: err => {
          console.error('QQ登录失败', err);
        }
      });
    },
    // 微博登录
    loginWithWeibo() {
      uni.login({
        provider: 'weibo',
        success: res => {
          console.log('微博登录成功', res);
        },
        fail: err => {
          console.error('微博登录失败', err);
        }
      });
    }
  }
};
</script>

<style>
.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: linear-gradient(180deg, #FFF2D9 0%, #FFFFFF 30%);
  padding-top: 80px;
}

.logo {
  width: 220px;
  height: 220px;
}

.title {
  font-size: 24px;
  font-weight: bold;
  color: #F8B91D;
  margin-top: 100px;
  margin-bottom: 40px;
}

.login-button,
.experience-button {
  width: 288px;
  height: 52px;
  background: #FFEE93;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.25);
  border-radius: 20px;
  font-size: 18px;
  font-weight: bold;
  color: #A3710E;
  text-align: center;
  line-height: 52px;
  margin-bottom: 20px;
  transition: all 0.3s ease; /* 为按钮添加平滑过渡效果 */
}

.login-button:active,
.experience-button:active {
  transform: scale(0.95); /* 按钮点击时缩小 */
  box-shadow: 0px 6px 10px rgba(0, 0, 0, 0.3); /* 点击时的阴影效果 */
}

.third-party {
  display: flex;
  justify-content: space-between;
  width: 50%;
  margin-top: 70px;
}

.third-party-button {
  width: 34px;
  height: 34px;
  background: #ffffff;
  border-radius: 0px 0px 0px 0px;
  transition: all 0.3s ease; /* 为第三方按钮添加平滑过渡效果 */
}

.third-party-button:active {
  transform: scale(0.9); /* 第三方按钮点击时缩小 */
}

.other-login {
  font-size: 14px;
  color: #A3710E;
  margin-top: 10px;
  cursor: pointer;
}
</style>
