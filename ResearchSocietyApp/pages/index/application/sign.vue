<template>
	<view class="content">
		<view class="back" @click="back()">
			< 返回
		</view>
		
		<button class="sign-button" @click="click">签到</button>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				result: ''
			}
		},
		methods: {
			click() {
				let that = this
				uni.scanCode({
					success: res => {
						that.result = res.result
					}
				})
			},
			jump(url, params, type) {
				type = type==null?'navigateTo':type
				// 页面的跳转
				this.$u.route({
					url,
					type,
					params
				})
			},
			back() {
				this.jump('', {}, 'back')
			}
		}
	}
</script>

<style lang="scss" scoped>
.content{
	width: 100vw;
	height: 100vh;
	background-image: linear-gradient(#aaaaff, #00ffff);
	position: relative;
	display: flex;
	justify-content: center;
	align-items: center;
	
	.back {
		position: absolute;
		top: 80rpx;
		left: 50rpx;
		color: #fff;
	}
	
	.sign-button {
		width: 50vw;
		height: 50vw;
		background-color: rgba($color: #fff, $alpha: .5);
		// border: 1rpx solid #000000;
		border-radius: 50%;
		box-shadow: 5px 8px 10px #888888;
		
		text-align: center;
		line-height: 50vw;
		font-size: larger;
		font-weight: bold;
	}
}
</style>
