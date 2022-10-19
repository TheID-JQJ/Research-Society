<template>
	<view class="content">
		<!-- 账号信息 -->
		<view class="account-background">
			<view class="account">
				<view class="appoint">
					<u-avatar
						shape="circle"
						size="100%"
						src="../../static/img/20210409113856123009.jpg"
					></u-avatar>
				</view>
				
				<view class="user">
					<h3 class="username">用户名</h3>
					<view class="user-message">
						<text>个性签名</text>
						<u-icon name="edit-pen"></u-icon>
					</view>
				</view>
				
				<myGrid class="grid" :myGridList="accountList" :myGridCol="3" v-slot="myGridData">
					<view class="grid-item">
						<text class="large">{{myGridData.item.number}}</text>
						<text class="small">{{myGridData.item.title}}</text>
					</view>
				</myGrid>
			</view>
		</view>
		
		<!-- 工具 -->
		<view class="util">
			<myGrid class="grid" :myGridList="utilList" :myGridCol="4" v-slot="myGridData">
				<view class="grid-item">
					<u-icon
						:name="myGridData.item.name"
						:size="30"
					></u-icon>
					<text class="grid-text">{{myGridData.item.title}}</text>
				</view>
			</myGrid>
		</view>
		
		<!-- 订单 -->
		<view class="order">
			<view class="title">
				<text>我的订单</text>
				<text style="color: #848484;" @click="jump('/pages/mall/application/order', {})">全部订单></text>
			</view>

			<myGrid class="grid" :myGridList="orderList" :myGridCol="4" v-slot="myGridData">
				<view class="grid-item">
					<u-icon
						:name="myGridData.item.name"
						:size="30"
					></u-icon>
					<text class="grid-text">{{myGridData.item.title}}</text>
				</view>
			</myGrid>
		</view>
		
		<!-- 账户 -->
		<view class="wallet">
			<myGrid class="grid" :myGridList="walletList" :myGridCol="4" v-slot="myGridData">
				<view class="grid-item" v-if="myGridData.item.title!='wallet'">
					<text style=";margin-bottom: 10rpx;">{{myGridData.item.number}}</text>
					<text style="font-size: small;">{{myGridData.item.title}}</text>
				</view>
				<view class="grid-item grid-item-list" v-else>
					<u-icon
						:name="myGridData.item.name"
						:size="30"
					></u-icon>
					<text style="font-size: small;">账户</text>
				</view>
			</myGrid>
		</view>
		
		<!-- 推荐商品 -->
		<myLayout class="recommended" :myLayoutList="recommendedList" :myLayoutGutter="10" v-slot="myLayoutData">
			<view class="recommended-content">
				<view class="recommended-img" :style="myLayoutData.item.img" />
				<h3 class="recommended-title">{{myLayoutData.item.title}}</h3>
				<text class="recommended-text">{{myLayoutData.item.score}}分 | 月售{{myLayoutData.item.sales}}</text>
				<view class="recommended-price">
					<text class="recommended-price-int">￥{{myLayoutData.item.priceInt}}</text>
					<text class="recommended-price-float">.{{myLayoutData.item.priceFloat}}</text>/
					<text class="recommended-price-integral">{{myLayoutData.item.integral}}积分</text>
				</view>
			</view>
		</myLayout>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				accountList: [
					{
						number: 0,
						title: '兑换券',
						url: '/pages/mall/application/voucher'
					},
					{
						number: 0,
						title: '优惠券',
						url: '/pages/mall/application/coupons'
					},
					{
						number: 255,
						title: '积分',
						url: '/pages/mall/application/earnPoints'
					}
				],
				utilList: [
					{
						name: 'shopping-cart',
						title: '购物车',
						url: '/pages/mall/application/shoppingCart'
					},
					{
						name: 'star',
						title: '收藏',
						url: '/pages/profile/application/star'
					},
					{
						name: 'clock',
						title: '足迹',
						url: '/pages/profile/application/history'
					},
					{
						name: 'setting',
						title: '设置',
						url: '/pages/profile/application/setting'
					},
				],
				orderList: [
					{
						name: 'order',
						title: '待付款',
						url: '/pages/mall/application/order'
					},
					{
						name: 'hourglass',
						title: '待使用',
						url: '/pages/mall/application/order'
					},
					{
						name: 'checkmark-circle',
						title: '已完成',
						url: '/pages/mall/application/order'
					},
					{
						name: 'kefu-ermai',
						title: '退款',
						url: '/pages/mall/application/order'
					},
				],
				walletList: [
					{
						number: '0.00元',
						title: '余额',
						url: '/pages/mall/application/account'
					},
					{
						number: '0时0分',
						title: '剩余时长',
						url: '/pages/mall/application/account'
					},
					{
						number: 255,
						title: '积分',
						url: '/pages/mall/application/account'
					},
					{
						name: 'rmb-circle',
						title: 'wallet',
						url: '/pages/mall/application/account'
					}
				],
				recommendedList: [
					{
						img: 'background-image: url("https://img.zcool.cn/community/01206b5e7884aaa80120a89582caf9.jpg@1280w_1l_2o_100sh.jpg");',
						title: '珍珠奶茶',
						score: 5.0,
						sales: 100,
						priceInt: 6,
						priceFloat: 5,
						integral: 650
					},
					{
						img: 'background-image: url("https://ts1.cn.mm.bing.net/th/id/R-C.de9b84d1464e862e1519425432d715a4?rik=Q8C1pNtrZEs40A&riu=http%3a%2f%2fb2bimgcdn.nbdeli.com%2fStorage%2fItemImages%2f18%2f1802%2f1802002%2f100043267%2fViewBig%2f100043267.jpg&ehk=gHvS0kcPS3rWUyyc5PYxV2yWWzgb5%2bw113qg8ZpaGzs%3d&risl=&pid=ImgRaw&r=0");',
						title: '可口可乐',
						score: 5.0,
						sales: 100,
						priceInt: 3,
						priceFloat: 0,
						integral: 300
					},
					{
						img: 'background-image: url("https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1");',
						title: '干脆面',
						score: 5.0,
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					},
					{
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						title: '笔记本',
						score: 5.0,
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					}
				]
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
			}
		}
	}
</script>

<style lang="scss">
.content {
	background-image: linear-gradient(#f7edf5, #ffffff);
	
	.account-background {
		padding: 150rpx 40rpx 0 40rpx;
		
		.account {
			background-color: #fff;
			padding: 40rpx;
			padding-top: 150rpx;
			padding-bottom: 0;
			
			border: 1rpx solid #f7edf5;
			border-radius: 40rpx;
			
			position: relative;
			
			.appoint {
				width: 200rpx;
				height: 200rpx;
				
				position: absolute;
				top: -100rpx;
				left: 50%;
				transform: translateX(-50%);
			}
			
			.user {
				height: 100rpx;
				
				text-align: center;
				margin-bottom: 20rpx;
				
				.username {
					margin-bottom: 10rpx;
					
					font-size: larger;
				}
				.user-message {
					font-size: small;
					
					display: flex;
					justify-content: center;
				}
			}
			
			.grid {
				.grid-item {
					margin-bottom: 20rpx;
					
					display: flex;
					flex-direction: column;
					align-items: center;
					
					.large {
						display: block;
						font-size: large;
						font-weight: 600;
					}
					
					.small {
						display: block;
						font-size: small;
					}
				}
			}
		}
	}
	
	.util {
		background-color: #fff;
		margin: 20rpx;
		
		border: 1rpx solid #f7edf5;
		border-radius: 20rpx;
		
		font-size: small;
		
		.grid {
			margin: 20rpx 0;
			
			.grid-item {
				display: flex;
				flex-direction: column;
				align-items: center;
			}
		}
	}
	
	.order {
		background-color: #fff;
		margin: 20rpx;
		
		border: 1rpx solid #f7edf5;
		border-radius: 20rpx;
		
		font-size: small;
		
		.title {
			padding: 20rpx 40rpx;
			
			display: flex;
			justify-content: space-between;
			
			border-bottom: 2rpx solid #f7edf5;
		}
		
		.grid {
			margin: 20rpx 0;
			
			.grid-item {
				display: flex;
				flex-direction: column;
				align-items: center;
			}
		}
	}
	
	.wallet {
		background-color: #fff;
		margin: 20rpx;
		
		border: 1rpx solid #f7edf5;
		border-radius: 20rpx;
		
		.grid {
			margin: 20rpx 0;
			
			.grid-item {
				display: flex;
				flex-direction: column;
				align-items: center;
			}
			
			.grid-item-list {
				width: 100%;
				
				border-left: 2rpx solid #f7edf5;
			}
		}
	}
	
	.recommended {
		margin: 20rpx;
		.recommended-content {
			background-color: #fff;
			margin-bottom: 20rpx;
			
			border: 2rpx solid #f7edf5;
			border-radius: 20rpx;
			
			.recommended-img {
				height: 350rpx;
				background-size:100% 100%;
				background-repeat:no-repeat;
				
				border-radius: 20rpx 20rpx 0 0 ;
			}
			
			.recommended-title {
				margin-left: 20rpx;
				font-weight: normal;
			}
			
			.recommended-text {
				margin-left: 20rpx;
				
				color: #848484;
				font-size: small;
			}
			
			.recommended-price {
				margin: 20rpx 0 20rpx 20rpx;
				display: flex;
				
				color: #f00;
				line-height: 50rpx;
				
				.recommended-price-float,
				.recommended-price-integral {
					font-size: small;
					line-height: 55rpx;
				}
			}
		}
	}
}
</style>
