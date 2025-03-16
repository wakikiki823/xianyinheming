<template>
	<view class="content">
		<view class="search-box">
			<view class="input-box">
				<input 
					type="text" 
					:adjust-position="true" 
					:placeholder="defaultKeyword" 
					@input="inputChange" 
					v-model="keyword" 
					@confirm="doSearch(false)"
					placeholder-class="placeholder-class" 
					confirm-type="search"
					class="search-input"
				>
				<view class="search-icon">
					<image src="/static/search.png" mode="aspectFit"></image>
				</view>
			</view>
			<view class="search-btn" @tap="doSearch(false)">搜索</view>
		</view>
		
		<view class="search-keyword">
			<scroll-view class="keyword-list-box" v-show="isShowKeywordList" scroll-y>
				<block v-for="(row,index) in keywordList" :key="index">
					<view class="keyword-entry" hover-class="keyword-entry-tap">
						<view class="keyword-text" @tap.stop="doSearch(keywordList[index].keyword)">
							<view class="search-icon-small">
								<image src="/static/search-small.png" mode="aspectFit"></image>
							</view>
							<rich-text :nodes="row.htmlStr"></rich-text>
						</view>
						<view class="keyword-img" @tap.stop="setKeyword(keywordList[index].keyword)">
							<image src="/static/arrow-up.png" mode="aspectFit"></image>
						</view>
					</view>
				</block>
			</scroll-view>
			
			<scroll-view class="keyword-box" v-show="!isShowKeywordList" scroll-y>
				<view class="keyword-block" v-if="oldKeywordList.length>0">
					<view class="keyword-list-header">
						<view class="header-title">
							<image src="/static/his.png" mode="aspectFit" class="header-icon"></image>
							<text>历史搜索</text>
						</view>
						<view class="header-action" @tap="oldDelete">
							<image src="/static/shanchu.png" mode="aspectFit"></image>
						</view>
					</view>
					<view class="keyword">
						<view 
							v-for="(keyword,index) in oldKeywordList" 
							@tap="doSearch(keyword)" 
							:key="index"
							class="keyword-item"
						>{{keyword}}</view>
					</view>
				</view>
				
				<view class="keyword-block">
					<view class="keyword-list-header">
						<view class="header-title">
							<image src="/static/hot.png" mode="aspectFit" class="header-icon"></image>
							<text>热门搜索</text>
						</view>
						<view class="header-action" @tap="hotToggle">
							<image :src="'/static/eye'+ forbid +'.png'" mode="aspectFit"></image>
						</view>
					</view>
					<view class="keyword" v-if="forbid==''">
						<view 
							v-for="(keyword,index) in hotKeywordList" 
							@tap="doSearch(keyword)" 
							:key="index"
							class="keyword-item hot-item"
							:class="index < 3 ? 'hot-'+index : ''"
						>
							<text class="hot-index" v-if="index < 3">{{index + 1}}</text>
							{{keyword}}
						</view>
					</view>
					<view class="hide-hot-tis" v-else>
						<text>当前热门搜索已隐藏</text>
					</view>
				</view>
			</scroll-view>
		</view>
	</view>
</template>

<script>
	//引用mSearch组件，如不需要删除即可
	/*import mSearch from '@/components/mehaotian-search-revision/mehaotian-search-revision.vue';*/
	export default {
		data() {
			return {
				defaultKeyword: "",
				keyword: "",
				oldKeywordList: [],
				hotKeywordList: [],
				keywordList: [],
				forbid: '',
				isShowKeywordList: false
			}
		},
		onLoad(options) {
			this.init();
			// 如果有传入搜索参数，则设置到搜索框中
			if (options.searchQuery) {
				this.keyword = decodeURIComponent(options.searchQuery);
				// 自动执行搜索
				this.doSearch(this.keyword);
			}
		},
		/*components: {
			//引用mSearch组件，如不需要删除即可
			mSearch
		},*/
		methods: {
			init() {
				this.loadDefaultKeyword();
				this.loadOldKeyword();
				this.loadHotKeyword();

			},
			blur(){
				uni.hideKeyboard()
			},
			//加载默认搜索关键字
			loadDefaultKeyword() {
				//定义默认搜索关键字，可以自己实现ajax请求数据再赋值,用户未输入时，以水印方式显示在输入框，直接不输入内容搜索会搜索默认关键字
				this.defaultKeyword = "默认的搜索内容";
			},
			//加载历史搜索,自动读取本地Storage
			loadOldKeyword() {
				uni.getStorage({
					key: 'OldKeys',
					success: (res) => {
						var OldKeys = JSON.parse(res.data);
						this.oldKeywordList = OldKeys;
					}
				});
			},
			//加载热门搜索
			loadHotKeyword() {
				//定义热门搜索关键字，可以自己实现ajax请求数据再赋值
				this.hotKeywordList = ['二胡️', '吉他', '周杰伦', '方文山','邓紫棋', '伍佰', '考级', '张杰', '林檎', 'linggo'];
			}, 
			//监听输入
			inputChange(event) {
				//兼容引入组件时传入参数情况
				var keyword = event.detail?event.detail.value:event;
				if (!keyword) {
					this.keywordList = [];
					this.isShowKeywordList = false;
					return;
				}
				this.isShowKeywordList = true;
				//以下示例截取淘宝的关键字，请替换成你的接口
				uni.request({
					url: 'https://suggest.taobao.com/sug?code=utf-8&q=' + keyword, //仅为示例
					success: (res) => {
						this.keywordList = [];
						this.keywordList = this.drawCorrelativeKeyword(res.data.result, keyword);
						
					}
				});
			},
			//高亮关键字
			drawCorrelativeKeyword(keywords, keyword) {
				var len = keywords.length,
					keywordArr = [];
				for (var i = 0; i < len; i++) {
					var row = keywords[i];
					//定义高亮#9f9f9f
					var html = row[0].replace(keyword, "<span style='color: #9f9f9f;'>" + keyword + "</span>");
					html = '<div>' + html + '</div>';
					var tmpObj = {
						keyword: row[0],
						htmlStr: html
					};
					keywordArr.push(tmpObj)
				}
				return keywordArr;
			},
			//顶置关键字
			setKeyword(index) {
				this.keyword = this.keywordList[index].keyword;
			},
			//清除历史搜索
			oldDelete() {
				uni.showModal({
					content: '确定清除历史搜索记录？',
					success: (res) => {
						if (res.confirm) {
							console.log('用户点击确定');
							this.oldKeywordList = [];
							uni.removeStorage({
								key: 'OldKeys'
							});
						} else if (res.cancel) {
							console.log('用户点击取消');
						}
					}
				});
			},
			//热门搜索开关
			hotToggle() {
				this.forbid = this.forbid ? '' : '_forbid';
			},
			//执行搜索
			doSearch(keyword) {
				keyword = keyword===false?this.keyword:keyword;
				this.keyword = keyword;
				this.saveKeyword(keyword); //保存为历史 
				uni.showToast({
					title: keyword,
					icon: 'none',
					duration: 2000
				});
				//以下是示例跳转淘宝搜索，可自己实现搜索逻辑
				
				// //#ifdef APP-PLUS
				// plus.runtime.openURL(encodeURI('taobao://s.taobao.com/search?q=' + keyword));
				// //#endif
				// //#ifdef H5
				// window.location.href = 'taobao://s.taobao.com/search?q=' + keyword
				// //#endif
				
			},
			//保存关键字到历史记录
			saveKeyword(keyword) {
				uni.getStorage({
					key: 'OldKeys',
					success: (res) => {
						var OldKeys = JSON.parse(res.data);
						var findIndex = OldKeys.indexOf(keyword);
						if (findIndex == -1) {
							OldKeys.unshift(keyword);
						} else {
							OldKeys.splice(findIndex, 1);
							OldKeys.unshift(keyword);
						}
						//最多10个纪录
						OldKeys.length > 10 && OldKeys.pop();
						uni.setStorage({
							key: 'OldKeys',
							data: JSON.stringify(OldKeys)
						});
						this.oldKeywordList = OldKeys; //更新历史搜索
					},
					fail: (e) => {
						var OldKeys = [keyword];
						uni.setStorage({
							key: 'OldKeys',
							data: JSON.stringify(OldKeys)
						});
						this.oldKeywordList = OldKeys; //更新历史搜索
					}
				});
			}
		}
	}
</script>
<style>
.content {
	background-color: #f8f8f8;
	min-height: 100vh;
}

.search-box {
	width: 100%;
	background-color: #ffffff;
	padding: 30rpx 30rpx;
	display: flex;
	justify-content: space-between;
	align-items: center;
	position: sticky;
	top: 0;
	z-index: 100;
	box-shadow: 0 2rpx 10rpx rgba(0,0,0,0.05);
}

.input-box {
	flex: 1;
	height: 72rpx;
	background: #f5f5f5;
	border-radius: 36rpx;
	padding: 0 30rpx;
	display: flex;
	align-items: center;
	margin-right: 20rpx;
}

.search-input {
	flex: 1;
	height: 100%;
	font-size: 28rpx;
	color: #333;
}

.search-icon {
	width: 40rpx;
	height: 40rpx;
	margin-right: 20rpx;
}

.search-icon image {
	width: 100%;
	height: 100%;
}

.search-btn {
	width: 120rpx;
	height: 72rpx;
	background: linear-gradient(to right, #FFD700, #FFA500);
	border-radius: 36rpx;
	display: flex;
	justify-content: center;
	align-items: center;
	color: #fff;
	font-size: 28rpx;
	font-weight: 500;
}

.keyword-list-box {
	max-height: calc(100vh - 132rpx);
	background-color: #fff;
}

.keyword-entry {
	padding: 20rpx 30rpx;
	display: flex;
	align-items: center;
	justify-content: space-between;
	border-bottom: 2rpx solid #f5f5f5;
	transition: all 0.3s;
}

.keyword-entry-tap {
	background-color: #f8f8f8;
}

.keyword-text {
	flex: 1;
	display: flex;
	align-items: center;
	color: #333;
	font-size: 28rpx;
}

.search-icon-small {
	width: 30rpx;
	height: 30rpx;
	margin-right: 15rpx;
}

.keyword-img {
	width: 40rpx;
	height: 40rpx;
	margin-left: 20rpx;
}

.keyword-img image {
	width: 100%;
	height: 100%;
}

.keyword-block {
	padding: 30rpx;
	background-color: #fff;
	margin-bottom: 20rpx;
}

.keyword-list-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: 30rpx;
}

.header-title {
	display: flex;
	align-items: center;
}

.header-icon {
	width: 40rpx;
	height: 40rpx;
	margin-right: 10rpx;
}

.header-title text {
	font-size: 30rpx;
	color: #333;
	font-weight: 500;
}

.header-action {
	width: 40rpx;
	height: 40rpx;
}

.header-action image {
	width: 100%;
	height: 100%;
}

.keyword {
	display: flex;
	flex-wrap: wrap;
	margin: 0 -10rpx;
}

.keyword-item {
	padding: 12rpx 30rpx;
	background-color: #f5f5f5;
	border-radius: 30rpx;
	margin: 10rpx;
	font-size: 26rpx;
	color: #666;
	position: relative;
	transition: all 0.3s;
}

.keyword-item:active {
	transform: scale(0.95);
	opacity: 0.8;
}

.hot-item {
	padding-left: 50rpx;
}

.hot-index {
	position: absolute;
	left: 20rpx;
	font-weight: bold;
}

.hot-0 {
	background: #FFE4E1;
	color: #FF4500;
}

.hot-1 {
	background: #E6E6FA;
	color: #9370DB;
}

.hot-2 {
	background: #E0FFFF;
	color: #20B2AA;
}

.hide-hot-tis {
	text-align: center;
	padding: 30rpx;
	color: #999;
	font-size: 28rpx;
}

.placeholder-class {
	color: #999;
}
</style>