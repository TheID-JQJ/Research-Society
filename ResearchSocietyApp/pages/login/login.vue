<template>
	<view class="content">
		<myForm>
			<view class="find" slot="find">
				<h2 class="title">找回密码</h2>
				
				<view class="username" borderBottom>
					<u--input
							v-model="username"
							border="bottom"
							maxlength="10"
							placeholder="请输入账号/邮箱"
							color="#fff"
					></u--input>
				</view>
				
				<view class="submit-button">
					<u-button class="submit" type="primary" shape="circle" @click="next()">下一步</u-button>
				</view>
			</view>
			
			<view class="login" slot="login">
				<h1 class="title">研学社</h1>
				
				<view class="username">
					<u--input
							v-model="username"
							border="bottom"
							maxlength="10"
							placeholder="请输入账号/邮箱"
							color="#fff"
					></u--input>
				</view>
				
				<view class="password">
					<u--input
							v-model="password"
							border="bottom"
							type="password"
							maxlength="16"
							placeholder="请输入密码"
							color="#fff"
					></u--input>
				</view>
				
				<view class="submit-button">
					<u-button class="submit" type="primary" shape="circle" @click="submit">登录</u-button>
				</view>
			</view>
			
			<view class="register" slot="register">
				<h2 class="title">注册</h2>
				
				<view class="username">
					<u--input
							v-model="username"
							border="bottom"
							maxlength="10"
							placeholder="请输入邮箱"
							color="#fff"
					></u--input>
				</view>
				
				<view class="password">
					<u--input
							v-model="password"
							border="bottom"
							type="password"
							maxlength="16"
							placeholder="请输入密码"
							color="#fff"
					></u--input>
				</view>
				
				<view class="submit-button">
					<u-button class="submit" type="primary" shape="circle" @click="register()">注册</u-button>
				</view>
			</view>
		</myForm>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				username: '',
				password: ''
			}
		},
		onLoad() {
			
		},
		methods: {
			login() {
				uni.switchTab({
					url: '/pages/index/index'
				})
			},
			submit() {
				this.login()
			},
			
			jump(url, params, type) {
				type = type==null?'navigateTo':type
				// 页面的跳转
				this.$u.route({
					url,
					type,
					params
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
.title {
	text-align: center;
}

.find,
.login,
.register {
	width: 500rpx;
	height: 550rpx;
	
	display: flex;
	flex-direction: column;
	justify-content: space-around;
}
</style>