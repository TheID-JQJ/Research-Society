<template>
	<view class="my-form">
		<view class="img-background" :class="backgroundClass"/>
		
		<view class="my-form-body">
			<view class="find" :class="findClass">
				<slot name="find"></slot>
				<view class="back" @click="loginClick('findToLogin')">
					返回
				</view>
			</view>
			
			<view class="login" :class="loginClass">
				<slot name="login"></slot>
				<view class="alternative">
					<text @click="loginClick('find')">找回密码</text>
					<text @click="loginClick('register')">注册</text>
				</view>
			</view>
			
			<view class="register" :class="registerClass">
				<slot name="register"></slot>
				<view class="back" @click="loginClick('registerToLogin')">
					返回
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		name:"myForm",
		data() {
			return {
				backgroundClass: '',
				findClass: '',
				loginClass: '',
				registerClass: '',
			};
		},
		methods: {
			loginClick(mes) {
				if(mes == 'register') {
					this.loginClass = 'login-animation-left'
					this.registerClass = 'register-animation-left'
					this.backgroundClass = 'background-right'
				}
				if(mes == 'registerToLogin') {
					this.loginClass = 'find-animation-right'
					this.registerClass = 'register-animation-right'
					this.backgroundClass = 'background-right-back'
				}
				if(mes == 'find') {
					this.loginClass = 'login-animation-right'
					this.findClass = 'find-animation-right'
					this.backgroundClass = 'background-left'
				}
				if(mes == 'findToLogin') {
					this.loginClass = 'register-animation-left'
					this.findClass = 'find-animation-left'
					this.backgroundClass = 'background-left-back'
				}
			},
		}
	}
</script>

<style lang="scss" scoped>
.my-form-body {
	height: 100vh;
	width: 100vw;
	
	display: flex; 
	flex-direction: column; 
	justify-content: center;
	align-items: center; 
	
	transform-style: preserve-3d;
	perspective: 5000rpx;
}

.img-background {
	height: 100vh;
	width: 200vw;
	background-image: url(../../static/img/v2-30effef829f9e1647e3644d18b9e3eea_r.jpg);
	background-position: center;
	background-repeat: no-repeat; 
	background-size: auto 100%;
	
	position: fixed;
	top: 0;
	z-index: -1;
	
	transform: translateX(-50vw);
}

.find,
.login,
.register{
	width: 600rpx;
	height: 600rpx;
	background-color: rgba($color: #000000, $alpha: .4);
	border-radius: 40rpx;
	
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	align-items: center;
	
	position: fixed;
	color: #fff
}

.find {
	transform: rotateY(-90deg) translateZ(500rpx);
	
	.back {
		margin-bottom: 30rpx;
		text-align: center;
		font-size: small;
	}
}

.login {
	transform: translateZ(500rpx);
	
	.alternative {
		width: 500rpx;
		margin-bottom: 30rpx;
		
		display: flex;
		justify-content: space-between;
		
		font-size: small;
	}
}

.register {
	transform: rotateY(90deg) translateZ(500rpx);
	
	.back {
		margin-bottom: 30rpx;
		text-align: center;
		font-size: small;
	}
}
</style>

<style scoped>
@keyframes loginRight {
	from {
		transform: rotateY(0deg) translateZ(500rpx);
	}
	to {
		transform: rotateY(90deg) translateZ(500rpx);
	}
}
@keyframes loginLeft {
	from {
		transform: rotateY(0deg) translateZ(500rpx);
		
	}
	to {
		transform: rotateY(-90deg) translateZ(500rpx);
	}
}
.login-animation-right {
	animation-name: loginRight;
	animation-duration: 1s;
	transform: rotateY(90deg) translateZ(500rpx);
}
.login-animation-left {
	animation-name: loginLeft;
	animation-duration: 1s;
	transform: rotateY(-90deg) translateZ(500rpx);
}

@keyframes registerRight {
	from {
		transform: rotateY(0deg) translateZ(500rpx);
	}
	to {
		transform: rotateY(90deg) translateZ(500rpx);
	}
}
@keyframes registerLeft {
	from {
		transform: rotateY(90deg) translateZ(500rpx);
	}
	to {
		transform: rotateY(0deg) translateZ(500rpx);
	}
}
.register-animation-right {
	animation-name: registerRight;
	animation-duration: 1s;
	transform: rotateY(90deg) translateZ(500rpx);
}
.register-animation-left {
	animation-name: registerLeft;
	animation-duration: 1s;
	transform: rotateY(0deg) translateZ(500rpx);
}
	
@keyframes findRight {
	from {
		transform: rotateY(-90deg) translateZ(500rpx);
	}
	to {
		transform: rotateY(0deg) translateZ(500rpx);
	}
}
@keyframes findLeft {
	from {
		transform: rotateY(0deg) translateZ(500rpx);
	}
	to {
		transform: rotateY(-90deg) translateZ(500rpx);
	}
}
.find-animation-right {
	animation-name: findRight;
	animation-duration: 1s;
	transform: rotateY(0deg) translateZ(500rpx);
}
.find-animation-left {
	animation-name: findLeft;
	animation-duration: 1s;
	transform: rotateY(-90deg) translateZ(500rpx);
}

@keyframes backgroundRight {
	from {
		transform: translateX(-50vw);
	}
	to {
		transform: translateX(-100vw);
	}
}
@keyframes backgroundRightBack {
	from {
		transform: translateX(-100vw);
	}
	to {
		transform: translateX(-50vw);
	}
}
@keyframes backgroundLeft {
	from {
		transform: translateX(-50vw);
	}
	to {
		transform: translateX(0);
	}
}
@keyframes backgroundLeftBack {
	from {
		transform: translateX(0);
	}
	to {
		transform: translateX(-50vw);
	}
}

.background-right {
	animation-name: backgroundRight;
	animation-duration: 1.5s;
	transform: translateX(-100vw);
}
.background-right-back {
	animation-name: backgroundRightBack;
	animation-duration: 1.5s;
	transform: translateX(-50vw);
}
.background-left {
	animation-name: backgroundLeft;
	animation-duration: 1.5s;
	transform: translateX(0);
}
.background-left-back {
	animation-name: backgroundLeftBack;
	animation-duration: 1.5s;
	transform: translateX(-50vw);
}
</style>

