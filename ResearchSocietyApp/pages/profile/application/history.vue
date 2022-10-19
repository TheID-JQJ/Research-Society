<template>
	<view class="content">
		<view class="back" @click="back()">
			< 返回
		</view>
		
		<view class="title">
			该功能维护中...
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				
			}
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
			back() {
				this.jump('', {}, 'back')
			},
		}
	}
</script>

<style lang="scss" scoped>
.content {
	background-image: linear-gradient(#aaaaff, #00ffff);
	width: 100vw;
	height: 100vh;
	padding-top: 10vh;
	position: fixed;
	
	.back {
		z-index: 99;
		position: absolute;
		top: 5vh;
		left: 2vh;
		color: #fff;
	}
	
	.title {
		position: absolute;
		top: 50vw;
		left: 50vw;
		transform: translateX(-50%);
	}
}
</style>
