<script>
	export default {
		onLaunch: function() {
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style>
	/*每个页面公共css */
	/* 新增全局样式 */
	page {
		--primary-color: #FF7675;
		--text-primary: #2D3436;
		--text-secondary: #636E72;
		--background-light: #F8F9FA;
		
		font-family: -apple-system, BlinkMacSystemFont, 'Helvetica Neue', sans-serif;
		background-color: #FFFFFF;
	}

	/* 统一按钮样式 */
	button::after {
		border: none !important;
	}
	button[disabled] {
		opacity: 0.5 !important;
	}

	/* 加载动画 */
	@keyframes spin {
		to { transform: rotate(360deg); }
	}
	.loading-spinner {
		width: 48rpx;
		height: 48rpx;
		border: 6rpx solid var(--primary-color);
		border-top-color: transparent;
		border-radius: 50%;
		animation: spin 0.8s linear infinite;
	}
</style>
