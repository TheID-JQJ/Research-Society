<template>
	<view class="content">
<!-- 		<view class="video-background">
			<video
				class="video"
				src="../../static/VID_20221009_152517.mp4"
				:loop="true" :autoplay="true"
				:muted="true" :controls="false"
				:enable-progress-gesture="false"
				object-fit="cover"
			/>
		</view> -->
		
		<view class="img-background"/>
		
		<view class="login-content">
			<h1 class="title">研学社</h1>
			
			<view class="login-form">
				<u--form>
					<u-form-item class="username" borderBottom>
						<u--input
								v-model="username"
								border="none"
								maxlength="10"
								placeholder="请输入账号/邮箱"
						></u--input>
					</u-form-item>
					
					<u-form-item class="password" borderBottom>
						<u--input
								v-model="password"
								border="none"
								type="password"
								maxlength="16"
								placeholder="请输入密码"
						></u--input>
					</u-form-item>
						
				</u--form>
			</view>
			
			<view class="submit-button">
				<u-button class="submit" type="primary" shape="circle" @click="submit">登录</u-button>
				<view class="alternative">
					<text @click="toFind">找回密码</text>
					<text @click="toRegister">注册</text>
				</view>
			</view>
		</view>
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
			},
			toFind() {
				this.jump('pages/login/find', {})
			},
			toRegister() {
				this.jump('pages/login/register', {})
			}
			
		}
	}
</script>

<style lang="scss">
.content {
	.video-background {
		height: 100vh;
		width: 100vw;
		
		overflow: hidden;
		
		position: fixed;
		top: 0;
		z-index: -1;
		
		.video {
			width: 100%;
			height: 100%;
			
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
		}
	}
	
	.img-background {
		height: 100vh;
		width: 100vw;
		background-image: url(../../static/img/v2-30effef829f9e1647e3644d18b9e3eea_r.jpg);
		background-position: center;
		background-repeat: no-repeat; 
		background-size: auto 100%;
		
		position: fixed;
		top: 0;
		z-index: -1;
	}
	
	.login-content {
		height: 100vh;
		width: 100vw;
		background-color: rgba($color: #000000, $alpha: .5);
		
		display: flex;
		flex-direction: column;
		justify-content: center;
		
		.title {
			height: 20rpx;
			text-align: center;
			padding: 40rpx;
			color: #fff;
		}
		
		.login-form {
			padding: 100rpx;
			
			.username {
				margin-bottom: 40rpx;
			}
			
			.password {
				// margin-bottom: 40rpx;
			}
		}
		
		.submit-button {
			padding: 0 100rpx;
			
			.submit {
				width: 100%;
				height: 80rpx;
				margin-bottom: 40rpx;
			}
			
			.alternative {
				display: flex;
				justify-content: space-between;
				
				text-align: center;
				// color: #848484;
				color: #fff;
			}
		}
	}
}
</style>
