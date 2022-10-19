<template>
	<view class="content">
		<view class="back" @click="back()">
			< 返回
		</view>
		
		<view class="head">			
			<u-tabs 
				:list="tabs" 
				@change="tabsChange" 
				lineColor="#ff814f"
				:activeStyle="{
					fontWeight: 'bold',
					transform: 'scale(1.05)'
				}"
				 :inactiveStyle="{
					transform: 'scale(1)'
				}"
				itemStyle="padding-left: 15vw; padding-right: 15vw; height: 34px;"
			></u-tabs>
		</view>
		
		<view class="body">
			<view class="user-item" v-for="item, index in indexList">
				<u-avatar
					shape="square"
					size="35"
					:src="urls[index]"
					customStyle="margin: -3px 5px -3px 0"
				></u-avatar>
				<text class="username">用户{{item}}</text>
				<text class="no" :class="index<3?'no-fst':''">No:{{item}}</text>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				tabs: [{name: "用户榜"}, {name: "好友榜"}],
				indexList: [1,2,3,4,5,6,7],
				urls: [
					'https://cdn.uviewui.com/uview/album/1.jpg',
					'https://cdn.uviewui.com/uview/album/2.jpg',
					'https://cdn.uviewui.com/uview/album/3.jpg',
					'https://cdn.uviewui.com/uview/album/4.jpg',
					'https://cdn.uviewui.com/uview/album/5.jpg',
					'https://cdn.uviewui.com/uview/album/6.jpg',
					'https://cdn.uviewui.com/uview/album/7.jpg'
				]
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
			tabsChange() {
				console.log("change")
			},
			scrolltolower() {
				
			}
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
	
	.head {
		background-color: rgba($color: #fff, $alpha: .3);
		padding: 10rpx 0;
		margin: 20rpx;
		margin-top: 0;
		border-radius: 20rpx;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	
	.body {
		height: 80vh;
		background-color: rgba($color: #fff, $alpha: .3);
		margin: 0 20rpx;
		padding: 20rpx;
		border-radius: 20rpx;
		overflow: scroll;
		
		.user-item {
			height: 8vh;
			background-color: rgba($color: #fff, $alpha: .5);
			margin-bottom: 20rpx;
			border-radius: 20rpx;
			display: flex;
			justify-content: space-around;
			align-items: center;
			
			.username {
				width: 60%;
			}
			
			.no {
				font-size: small;
			}
			
			.no-fst {
				color: #f00;
				font-size: large;
			}
		}
	}
}
</style>
