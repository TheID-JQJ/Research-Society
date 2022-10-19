<template>
	<view class="content">
		<view class="logout">
			<u-button type="error" @click="toLogin">退出登录</u-button>
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
			toLogin() {
				this.jump('/pages/login/login', {}, 'reLaunch')
			}
		}
	}
</script>

<style lang="scss" scoped>
.content {

	.logout{
		margin: 20rpx;
	}
}
</style>
