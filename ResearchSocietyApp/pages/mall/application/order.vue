<template>
	<view class="content">
		<view class="back" @click="back()">
			< 返回
		</view>
		
		<view class="head">
			<text class="title-text" :class="titleSelected==index?'selected':''" v-for="item, index in titleList" @click="titleClick(index)">
				{{item}}
			</text>
		</view>
		
		<view class="body">
			<view class="body-item" v-for="item, index in goodsList">
				<view class="goods-img" :style="item.img" @click="toItem(item)">
				</view>
				<view class="goods-message">
					<h3 class="title">{{item.title}}</h3>
					<view class="price">
						<text class="price-int">￥{{item.priceInt}}</text>
						<text class="price-float">.{{item.priceFloat}}</text>
					</view>
					<text class="text">数量：{{item.sales}}</text>
					<text class="message" style="color: #2979FF;">详情 ></text>
				</view>
				<view class="goods-botton" v-if="type=='finesh'">
					<u-button shape="circle" style="height: 8vw;">退款</u-button>
					<u-button type="error" shape="circle" style="height: 8vw;">再次购买</u-button>
				</view>
				<view class="goods-botton" v-else-if="type=='wait'">
					<u-button shape="circle" style="height: 8vw;">取消</u-button>
					<u-button type="error" shape="circle" style="height: 8vw;">立即付款</u-button>
				</view>
				<view class="goods-botton" v-else>
					<u-button shape="circle" style="height: 8vw;">退款</u-button>
					<u-button type="error" shape="circle" style="height: 8vw;">去使用</u-button>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				titleList: ['全部', '待付款', '待使用', '已完成', '退款'],
				titleSelected: 0,
				type: 'finish',
				goodsList: [
					{
						img: 'background-image: url("https://img.zcool.cn/community/01206b5e7884aaa80120a89582caf9.jpg@1280w_1l_2o_100sh.jpg");',
						imgUrl: 'https://img.zcool.cn/community/01206b5e7884aaa80120a89582caf9.jpg@1280w_1l_2o_100sh.jpg',
						title: '珍珠奶茶',
						sales: 100,
						priceInt: 6,
						priceFloat: 5,
						integral: 650,
						kid: 1
					},
					{
						img: 'background-image: url("https://ossmuyuanec.oss-cn-beijing.aliyuncs.com/Storage/Shop/1221/Products/17977/1.png");',
						imgUrl: 'https://ossmuyuanec.oss-cn-beijing.aliyuncs.com/Storage/Shop/1221/Products/17977/1.png',
						title: '可口可乐',
						sales: 100,
						priceInt: 3,
						priceFloat: 0,
						integral: 300,
						kid: 2
					},
					{
						img: 'background-image: url("https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1");',
						imgUrl: 'https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1',
						title: '干脆面',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100,
						kid: 3
					},
					{
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						imgUrl: 'https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg',
						title: '笔记本',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100,
						kid: 4
					},
					{
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						imgUrl: 'https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg',
						title: '笔记本',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100,
						kid: 4
					},
					{
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						imgUrl: 'https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg',
						title: '笔记本',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100,
						kid: 4
					},
					{
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						imgUrl: 'https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg',
						title: '笔记本',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100,
						kid: 4
					}
				],
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
			toItem(item) {
				this.jump('/pages/mall/item', item)
			},
			titleClick(index) {
				this.titleSelected = index
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
		height: 5vh;
		background-color: rgba($color: #fff, $alpha: .3);
		margin: 0 20rpx;
		border-radius: 20rpx;
		display: flex;
		justify-content: space-around;
		align-items: center;
		
		.selected {
			font-size: larger;
		}
	}
	
	.body {
		height: 80vh;
		background-color: rgba($color: #fff, $alpha: .3);
		margin: 20rpx;
		padding: 20rpx;
		padding-top: 0;
		border-radius: 20rpx;
		overflow: scroll;
		
		.body-item {			
			width: 100%;
			height: 30vw;
			background-color: rgba($color: #fff, $alpha: .5);
			margin-top: 20rpx;
			border-radius: 20rpx;
			display: flex;
			justify-content: space-around;
			align-items: center;
			
			.goods-img {
				width: 20vw;
				height: 20vw;
				background-size: 100% 100%;
				border-radius: 20rpx;
			}
			
			.goods-message {
				width: 30vw;
				height: 100%;
				display: flex;
				flex-direction: column;
				justify-content: space-around;
				
				.title {
					font-size: medium;
					font-weight: normal;
				}
				
				.price,
				.text,
				.message {
					font-size: small;
				}
			}
			
			.goods-botton {
				height: 100%;
				display: flex;
				flex-direction: column;
				justify-content: space-around;
			}
		}
	}
}
</style>
