<template>
	<view class="content">
		<view class="find" :class="findClass">
			找回密码
			<button type="default" @click="loginClick('findToLogin')">back</button>
		</view>
		
		<view class="login" :class="loginClass">
			登录
			
			<view style="width:100%;display: flex; justify-content: space-around;">
				<button type="default" @click="loginClick('find')">find</button>
				<button type="default" @click="loginClick('register')">register</button>
			</view>
		</view>
		
		<view class="register" :class="registerClass">
			注册
			<button type="default" @click="loginClick('registerToLogin')">back</button>
		</view>
		
		<view class="test">
			
		</view>
	</view>
</template>

<script>
	export default {
		onLoad(params) {
			console.log(params)
		},
		data() {
			return {
				findClass: '',
				loginClass: '',
				registerClass: ''
			};
		},
		methods: {
			loginClick(mes) {
				if(mes == 'register') {
					this.loginClass = 'login-animation-left'
					this.registerClass = 'register-animation-left'
				}
				if(mes == 'registerToLogin') {
					this.loginClass = 'find-animation-right'
					this.registerClass = 'register-animation-right'
				}
				if(mes == 'find') {
					this.loginClass = 'login-animation-right'
					this.findClass = 'find-animation-right'
				}
				if(mes == 'findToLogin') {
					this.loginClass = 'register-animation-left'
					this.findClass = 'find-animation-left'
				}
			}
		}
	}
</script>

<style lang="scss" scoped>
.content {
	height: 100vh; 
	display: flex; 
	flex-direction: column; 
	align-items: center; 
	
	transform-style: preserve-3d;
	perspective: 2000rpx;
}

.find,
.login,
.register{
	width: 500rpx;
	height: 500rpx;
	background-color: rgba($color: #000000, $alpha: .5);
	
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	align-items: center;
	
	position: fixed;
	top: 200rpx;
}

.find {
	transform: translateX(-400rpx) rotateY(-90deg);
}

.login {
	transform: translateZ(400rpx);
}

.register {
	transform: translateX(400rpx) rotateY(90deg);
}
</style>

<style scoped>
@keyframes loginRight {
	from {
		transform: translateX(0) translateZ(400rpx) rotateY(0deg);
	}
	to {
		transform: translateX(400rpx) translateZ(0) rotateY(90deg);
	}
}
@keyframes loginLeft {
	from {
		transform: translateX(0) translateZ(400rpx) rotateY(0deg);
		
	}
	to {
		transform: translateX(-400rpx) translateZ(0) rotateY(-90deg);
	}
}
.login-animation-right {
	animation-name: loginRight;
	animation-duration: 2s;
	transform: translateX(400rpx) rotateY(90deg);
}
.login-animation-left {
	animation-name: loginLeft;
	animation-duration: 2s;
	transform: translateX(-400rpx) rotateY(-90deg);
}

@keyframes registerRight {
	from {
		transform: translateX(0) translateZ(400rpx) rotateY(0deg);
	}
	to {
		transform: translateX(400rpx) translateZ(0) rotateY(90deg);
	}
}
@keyframes registerLeft {
	from {
		transform: translateX(400rpx) translateZ(0) rotateY(90deg);
	}
	to {
		transform: translateX(0) translateZ(400rpx) rotateY(0deg);
	}
}
.register-animation-right {
	animation-name: registerRight;
	animation-duration: 2s;
	transform: translateX(400rpx) rotateY(90deg);
}
.register-animation-left {
	animation-name: registerLeft;
	animation-duration: 2s;
	transform: translateZ(400rpx) rotateY(0deg);
}
	
@keyframes findRight {
	from {
		transform: translateX(-400rpx) translateZ(0) rotateY(-90deg);
	}
	to {
		transform: translateX(0) translateZ(400rpx) rotateY(0deg);
	}
}
@keyframes findLeft {
	from {
		transform: translateX(0) translateZ(400rpx) rotateY(0deg);
	}
	to {
		transform: translateX(-400rpx) translateZ(0) rotateY(-90deg);
	}
}
.find-animation-right {
	animation-name: findRight;
	animation-duration: 2s;
	transform: translateZ(400rpx) rotateY(0deg);
}
.find-animation-left {
	animation-name: findLeft;
	animation-duration: 2s;
	transform: translateX(-400rpx) rotateY(-90deg);
}
</style>
