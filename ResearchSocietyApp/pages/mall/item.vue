<template>
	<view class="content">
		<u--image :showLoading="true" :src="message.imgUrl" width="100%" height="400"></u--image>
		
		<view class="head">
			<view class="price">
				￥{{message.priceInt}}.{{message.priceFloat}}
			</view>
			<view class="description">
				{{message.title}}
			</view>
		</view>
		
<!-- 		<view class="recommendation">
			<text v-if="vuex_language=='Chinese'">推荐：</text>
			<text v-else>recommend：</text>
		</view> -->
		
		<view class="bottom">
			<view class="left">
				<u-icon name="star" size="25" label="收藏" @click="toStar(item.gid)" v-if="!isStar"></u-icon>
				<u-icon name="star-fill" color="#f00" size="25" label="已收藏" @click="disStar(item.gid)" v-else></u-icon>
			</view>
			
			<view class="right">
				<u-button shape="circle" color="#007AFF" @click="addCart(item.gid)" style="width: 45%;">加入购物车</u-button>
				<u-button shape="circle" color="#f00" @click="toCreateOrder(item.gid)" style="width: 45%;">立即购买</u-button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				message: {},
				item: {
					gid: 1
				},
				isStar: false
			}
		},
		onLoad(message) {
			this.message = message
		},
		methods: {
			
		}
	}
</script>

<style lang="scss" scoped>
.content {
	
	.head {
		margin-top: 40rpx;
		margin-left: 20rpx;
		margin-right: 20rpx;
		
		border-radius: 20rpx;
		box-shadow: 5rpx 5rpx 10rpx 0rpx #a8a8a8;
		
		.price {
			padding-top: 20rpx;
			padding-left: 20rpx;
			
			font-size: larger;
			font-weight: bold;
			color: #f00;
		}
		
		.description {
			padding: 20rpx;
		}
	}
	
	// .recommendation {
	// 	margin-top: 20rpx;
	// 	margin-left: 5%;
	// 	margin-right: 5%;
		
	// 	font-size: large;
	// 	font-weight: bold;
	// }
	
	.bottom {
		width: 100vw;
		padding: 40rpx 0;
		position: fixed;
		bottom: 0;
		display: flex;
		justify-content: space-around;
		align-items: center;
		
		.right {
			width: 55vw;
			padding-left: 100rpx;
			display: flex;
		}
	}
}
</style>
