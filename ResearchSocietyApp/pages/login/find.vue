<template>
	<view class="content">
		<view class="img-background"/>
		
		<view class="find-content">
			<h1 class="title">找回密码</h1>
			
			<view class="find-form">
				<u--form>
					<u-form-item class="username" borderBottom>
						<u--input
								v-model="username"
								border="none"
								maxlength="10"
								placeholder="请输入账号/邮箱"
						></u--input>
					</u-form-item>
				</u--form>
			</view>
			
			<view class="next-button">
				<u-button class="next" type="primary" shape="circle" @click="next">下一步</u-button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				username: ''
			}
		},
		onLoad() {
			
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
			next() {
				console.log('next')
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
	
	.find-content {
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
		
		.find-form {
			padding: 100rpx;
			
			// .username {
			// 	margin-bottom: 40rpx;
			// }
		}
		
		.next-button {
			padding: 0 100rpx;
			
			.next {
				width: 100%;
				height: 80rpx;
				margin-bottom: 40rpx;
			}
		}
	}
}
</style>
