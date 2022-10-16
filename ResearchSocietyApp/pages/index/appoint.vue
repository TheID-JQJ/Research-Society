<template>
	<view class="content">
		<view class="sit-img">
			<view class="item" :style="'transform: translateX('+index*53+'vw);'" v-for="item, index in sitImgList">
			</view>
		</view>
		
		<view class="sit-botton">
			<view class="left-botton">
				<
			</view>
			<view class="center-botton">
				预约
			</view>
			<view class="right-botton">
				>
			</view>
		</view>
		
		<view class="sit-list">
			<view class="item" v-for="item, index in sitList">
				{{item+1}}
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				sitImgList: [
					1,2,3,4,5
				],
				sitList: [
					
				],
				select: 0
			};
		},
		onLoad(params) {
			for (var i = 0; i < 20; i++) {
				this.sitList.push(i)
			}
		},
		methods: {

		}
	}
</script>

<style lang="scss" scoped>
.content {
	width: 100vw;
	height: 100vh;
	background-image: linear-gradient(#aaaaff, #00ffff);
	position: fixed;
	overflow: scroll;
	transform-style: preserve-3d;
	perspective: 5000rpx;
	
	.sit-img {
		width: 100vw;
		height: 90vw;
		position: relative;
		
		.item {
			width: 50vw;
			height: 70vw;
			background-color: rgba($color: #fff, $alpha: .3);
			border-radius: 40rpx;
			position: absolute;
			top: 80rpx;
			left: 25vw;
		}
	}
	
	.sit-botton {
		display: flex;
		justify-content: center;
		
		.left-botton {
			width: 20vw;
			height: 20vw;
			// background-color: rgba($color: #fff, $alpha: .3);
			
			text-align: center;
			line-height: 20vw;
			color: #fff;
			font-size: x-large;
		}
		
		.center-botton {
			width: 20vw;
			height: 20vw;
			background-color: rgba($color: #fff, $alpha: .3);
			border-radius: 50%;
			
			text-align: center;
			line-height: 20vw;
			color: #fff;
		}
		
		.right-botton {
			width: 20vw;
			height: 20vw;
			// background-color: rgba($color: #fff, $alpha: .3);
			
			text-align: center;
			line-height: 20vw;
			color: #fff;
			font-size: x-large;
		}
	}
	
	.sit-list{
		margin: 40rpx;
		padding: 20rpx 0;
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
		background-color: rgba($color: #fff, $alpha: .5);
		border-radius: 40rpx;
		
		.item {
			width: 10vw;
			height: 10vw;
			background-color: #fff;
			margin: 20rpx;
			
			text-align: center;
			line-height: 10vw;
		}
	}
}
</style>