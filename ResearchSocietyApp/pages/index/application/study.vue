<template>
	<view class="content">
		<view class="first">
			
		</view>
		
		<view class="second">			<view class="second-item ranking">
				<h2>1</h2>
				<text class="title">排名</text>
			</view>
			<view class="second-item total-time">
				<h2 class="val">5h30m</h2>
				<text class="title">共学习</text>
			</view>
			<view class="second-item todat-time">
				<h2 class="val">1h5m</h2>
				<text class="title">今日学习</text>
			</view>
		</view>
		
		<view class="third">
			
		</view>
		
		<view class="fourth">
			<view class="to-do-list">
				
			</view>
			
			<view class="application">
				<view class="end" @click="end">
					结束
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				
			}
		},
		onLoad() {
			// #ifdef APP-PLUS
			setTimeout(() => {
				plus.navigator.setFullscreen(true);
				plus.navigator.hideSystemNavigation()
				
				plus.screen.unlockOrientation()
				plus.screen.lockOrientation('landscape-primary')
			}, 500)
			// #endif
		},
		onUnload() {
			// #ifdef APP-PLUS
			plus.navigator.setFullscreen(false);
			plus.navigator.showSystemNavigation()
			plus.screen.unlockOrientation()
			plus.screen.lockOrientation('portrait-primary')
			// #endif
		},
		methods: {
			jump(url, params, type) {
				type = type==null?'navigateTo':type
				// 页面的跳转
				this.$u.route({
					url,
					type,
					params
				})
			},
			end() {
				this.jump('/pages/index/index', {}, 'tab')
			}
		}
	}
</script>

<style lang="scss" scoped>
.content {
	width: 100vw;
	height: 100vh;
	// background-color: #007AFF;
	background-image: linear-gradient(to bottom right ,#aaaaff, #00ffff);
	
	display: flex;
	justify-content: space-around;
	align-items: center;
	
	.first {
		width: 10vw;
		height: 80vh;
		background-color: rgba($color: #fff, $alpha: .5);
		
		border-radius: 40rpx;
	}
	
	.second {
		width: 10vw;
		height: 80vh;
		
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		
		.second-item {
			width: 10vw;
			height: 10vw;
			background-color: rgba($color: #fff, $alpha: .5);
			border-radius: 40rpx;
			display: flex;
			flex-direction: column;
			justify-content: space-around;
			align-items: center;
			
			color: #fff;
			
			.val {
				font-size: medium;
			}
			
			.title {
				font-size: small;
			}
		}
	}
	
	.third {
		width: 30vw;
		height: 80vh;
		background-color: rgba($color: #fff, $alpha: .5);
		
		border-radius: 40rpx;
	}
	
	.fourth {
		width: 30vw;
		height: 80vh;
		
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		
		.to-do-list {
			width: 100%;
			height: 70%;
			
			background-color: rgba($color: #fff, $alpha: .5);
			
			border-radius: 40rpx;
		}
		
		.application {
			width: 100%;
			height: 20%;
			background-color: rgba($color: #fff, $alpha: .5);
			border-radius: 40rpx;
			display: flex;
			justify-content: space-around;
			align-items: center;
			
			.end {
				width: 100rpx;
				height: 100rpx;
				background-color: rgba($color: #fa3534, $alpha: .5);
				border: 1rpx solid #ffffff;
				border-radius: 50%;
				
				line-height: 100rpx;
				text-align: center;
				color: #fff;
			}
		}
	}
}
</style>
