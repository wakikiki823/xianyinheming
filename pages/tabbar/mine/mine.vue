<template>
  <view class="container">
    <!-- 右上角的图标部分 -->
    <view class="top-icons">
      <image 
        class="top-icon" 
        src="/static/iconfontscan 1.png"
        @click="goToScanPage"
      ></image>
      <image 
        class="top-icon" 
        src="/static/shezhi 1.png"
        @click="goToSettingsPage"
      ></image>
    </view>

    <!-- 用户信息部分 -->
    <view class="profile-section">
      <!-- 头像部分，点击后选择头像 -->
      <image 
        class="profile-avatar" 
        :src="profile.avatar" 
        @click="chooseAvatar"
      ></image>

      <!-- 右侧的用户信息 -->
      <view class="user-info">
        <!-- 显示用户名称，点击输入框修改 -->
        <input 
          v-model="profile.name" 
          class="user-name" 
          placeholder="点击编辑名称"
          @blur="updateName"
        />
        <text class="vip-level">VIP等级：{{ profile.vipLevel }}</text>
        <text class="account-id">账户ID：{{ profile.accountId }}</text>
      </view>

       <view class="coins">
        <image class="coin-icon" src="/static/huiyuanjifenfangan 1.png"></image>
        <view class="coin-frame">
          <text class="coin-count">{{ profile.coins }}</text>
        </view>
      </view>
    </view>

    <!-- 菜单按钮 -->
    <view class="main-section">
      <view class="training-record">
        <image 
          class="training-icon"  
          src="/static/tuanjie 2.png"
        ></image>  
        <text>训练记录</text>
      </view>
      <view class="right-section">
        <view class="course">
          <image class="menu-icon" src="/static/kecheng-3 1 (2).png"></image>
          <text>我的课程</text>
        </view>
        <view class="collection" @click="navigateToCollection">
          <image class="menu-icon" src="/static/shoucang 1.png"></image>  
          <text>我的收藏</text>
        </view>
      </view>
    </view>

    <!-- 额外选项 -->
    <view class="settings-options">
      <view class="settings-item">
        <image class="settings-icon" src="/static/dongtai 1.png"></image>  
        <text>我的收藏</text>
        <image class="arrow-icon" src="/static/qianwang 3.png"></image>
      </view>
      <view class="settings-item">
        <image class="settings-icon" src="/static/jianchagengxin 1.png"></image>  
        <text>检查更新</text>
        <image class="arrow-icon" src="/static/qianwang 3.png"></image>
      </view>
      <view class="settings-item">
        <image class="settings-icon" src="/static/yijianfankui 1.png"></image>  
        <text>意见反馈</text>
        <image class="arrow-icon" src="/static/qianwang 3.png"></image>
      </view>
      <view class="settings-item">
        <image class="settings-icon" src="/static/yonghuxieyi 1.png"></image>  
        <text>用户协议</text>
        <image class="arrow-icon" src="/static/qianwang 3.png"></image>
      </view>
      <view class="settings-item">
        <image class="settings-icon" src="/static/yinsizhengce 1.png"></image>  
        <text>隐私政策</text>
        <image class="arrow-icon" src="/static/qianwang 3.png"></image>
      </view>
      <view class="settings-item">
        <image class="settings-icon" src="/static/guanyuwomen 1.png"></image>  
        <text>关于我们</text>
        <image class="arrow-icon" src="/static/qianwang 3.png"></image>
      </view>
      <!-- 空白块用于美化界面 -->
      <view class="settings-item empty">
        <!-- 空白部分，保持不变 -->
      </view>
    </view>

  </view>
</template>

<script>
export default {
  data() {
    return {
      profile: {
        avatar: "/static/logoo.png", // 默认头像
        name: "",
        vipLevel: 0, // VIP等级初始值为0
        accountId: Math.floor(100000 + Math.random() * 900000), // 生成6位随机数作为账号ID
        coins: 0 // 金币数量初始值为0
      },
      likedSongs: []
    };
  },
  created() {
    this.loadLikedSongs()
    uni.$on('update-liked-songs', this.loadLikedSongs)
  },
  methods: {
    chooseAvatar() {
      uni.chooseImage({
        count: 1,
        success: (res) => {
          const tempFilePath = res.tempFilePaths[0];
          this.profile.avatar = tempFilePath;
        },
        fail: (err) => {
          console.log("选择图片失败", err);
        },
      });
    },

    updateName() {
      console.log("新名称是：" + this.profile.name);
    },

    goToScanPage() {
      uni.navigateTo({
        url: '/pages/scan/scan'
      });
    },

    goToSettingsPage() {
      uni.navigateTo({
        url: '/pages/settings/settings'
      });
    },

    navigateToCollection() {
      uni.navigateTo({
        url: '/pages/collected-songs/collected-songs'
      });
    },

    loadLikedSongs() {
      this.likedSongs = uni.getStorageSync('likedSongs') || []
    }
  }
};
</script>


<style scoped>
	.container {
		  --round-radius: 10px;
		}
	
	
.coin-frame {
	  width: 100rpx;
	  height: 38rpx;
	  border-radius: 8px; /* 替换图片为CSS圆角 */
	  background: linear-gradient(90deg, #FFD700, #FFB300); /* 新增渐变背景 */
	  display: flex;
	  justify-content: center;
	  align-items: center;
	}
		
		.training-record,
			.course, 
			.collection {
			  border-radius: var(--round-radius);
			}
			
	.settings-item {
		  border-radius: var(--round-radius);
		  margin-bottom: 8px; /* 增加间距避免圆角重叠 */
		  border-bottom: none; /* 移除分割线 */
		}

	.settings-options {
		  gap: 8px; /* 使用gap控制间距 */
		}
	
.settings-item.empty {
	  border-radius: var(--round-radius);
	  background: transparent; /* 改为透明背景 */
	  box-shadow: none;
	}
	

	
	
	.training-icon {
	  width: 120px;  /* 放大图标尺寸 */
	  height: 120px;
	  object-fit: contain; /* 保持比例 */
	  margin-bottom: 10px; /* 增加与文字的间距 */
	   border-radius: var(--round-radius); /* 新增圆角 */
	}
	
	/* 保持其他图标原有样式不变 */
	
	
.container {
  padding: 20px;
  background: linear-gradient(1deg, #FFFFFF 42%, #FFEE93 91%, #FFD561 100%);
}

.top-icons {
  display: flex;
  justify-content: flex-end;
  margin-top: 15px;
  position: absolute;
  right: 20px;
  top: 20px;
}

.top-icon {
  width: 23px;
  height: 23px;
  margin-left: 10px;
  cursor: pointer;
}

.profile-section {
  display: flex;
  align-items: flex-start;
  margin-top: 25px;
  margin-bottom: 20px;
  position: relative;
}

.profile-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  cursor: pointer;
  margin-right: 15px;
}

.user-info {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.user-name {
  font-size: 18px;
  font-weight: bold;
  text-align: left;
  border: none;
  background: transparent;
  margin-bottom: 5px;
}

.vip-level,
.account-id {
  font-size: 14px;
  color: #888;
}

.coins {
  position: absolute;
  right: 20px;
  bottom: 0;
  display: flex;
  align-items: center;
}

.coin-icon {
  width: 65rpx;
  height: 65rpx;
}

.coin-frame {
  width: 100rpx;
  height: 38rpx;
  position: relative;
  background-image: url('/static/Rectangle 1601.png');
  background-size: cover;
  display: flex;
  justify-content: center;
  align-items: center;
}

.coin-count {
  font-size: 18px;
  color: #fff;
  text-align: center;
}

/* 菜单按钮样式 */
.main-section {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

.training-record {
  width: 40%;
  height: 150px;
  background-color: #ffffff;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* 阴影效果 */
}

.right-section {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 125px;
  width: 40%; /* 调整宽度，减少占用空间 */
  padding-right: 20px; /* 添加右侧间距 */
}

.course, .collection {
  width: 100%;
  height: 48%;
  background-color: #ffffff;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 10px;
  margin-bottom: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* 阴影效果 */
}

.menu-icon {
  width: 40px;
  height: 40px;
  object-fit: contain;
  margin-bottom: 5px;
}

/* 额外选项部分 */
.settings-options {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  margin-top: 20px;
}

.settings-item {
  width: 88%;
  background-color: white;
  border: none;
  border-radius: 5px;
  margin-bottom: 0px;
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #8b8b8b; /* 分割线 */
  justify-content: space-between; /* 让图标、文字和箭头图片两端对齐 */
}

.settings-item.empty {
  height: 25px;
  width: 88%;
  background-color: #ffffff;  /* 空白部分，增加界面的美化 */
}

.settings-icon {
  width: 30px;
  height: 30px;
  object-fit: contain;
  margin-right: 10px; /* 图标和文字的间隔 */
}

.arrow-icon {
  width: 20px; /* 调整箭头图片的宽度 */
  height: 20px; /* 调整箭头图片的高度 */
  object-fit: contain;
}

.settings-item text {
  flex: 1; /* 让文字部分占据剩余空间 */
}
</style>