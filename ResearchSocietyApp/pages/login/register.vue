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
		
		<view class="register-content">
			<h1 class="title">注册</h1>
			
			<view class="register-form">
				<u--form>
					<u-form-item class="email" borderBottom>
						<u--input
								v-model="email"
								border="none"
								type="email"
								placeholder="请输入邮箱"
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
				<u-button class="submit" type="primary" shape="circle" @click="submit">注册</u-button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				email: '',
				password: ''
			}
		},
		onLoad() {
			
		},
		methods: {
			register() {
				uni.switchTab({
					url: '/pages/index/index'
				})
			},
			submit() {
				this.register()
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
	
	.register-content {
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
		
		.register-form {
			padding: 100rpx;
			
			.email {
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
		}
	}
}
</style>
