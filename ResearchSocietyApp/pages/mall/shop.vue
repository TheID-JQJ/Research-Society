<template>
	<view class="content">
		<view class="back" @click="back()">
			< 返回
		</view>
		
		<!-- 搜索框 -->
		<view class="search">
			<u-search></u-search>
		</view>
		
		<view class="body">
			<myLayout class="recommended" :myLayoutList="recommendedList" :myLayoutGutter="10" v-slot="myLayoutData">
				<view class="recommended-content">
					<view class="recommended-img" :style="myLayoutData.item.img" />
					<h3 class="recommended-title">{{myLayoutData.item.title}}</h3>
					<text class="recommended-text">月售{{myLayoutData.item.sales}}</text>
					<view class="recommended-price">
						<text class="recommended-price-int">￥{{myLayoutData.item.priceInt}}</text>
						<text class="recommended-price-float">.{{myLayoutData.item.priceFloat}}</text>/
						<text class="recommended-price-integral">{{myLayoutData.item.integral}}积分</text>
					</view>
				</view>
			</myLayout>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				recommendedList: [
					{
						img: 'background-image: url("https://img.zcool.cn/community/01206b5e7884aaa80120a89582caf9.jpg@1280w_1l_2o_100sh.jpg");',
						imgUrl: 'https://img.zcool.cn/community/01206b5e7884aaa80120a89582caf9.jpg@1280w_1l_2o_100sh.jpg',
						title: '珍珠奶茶',
						sales: 100,
						priceInt: 6,
						priceFloat: 5,
						integral: 650
					},
					{
						img: 'background-image: url("https://ossmuyuanec.oss-cn-beijing.aliyuncs.com/Storage/Shop/1221/Products/17977/1.png");',
						imgUrl: 'https://ossmuyuanec.oss-cn-beijing.aliyuncs.com/Storage/Shop/1221/Products/17977/1.png',
						title: '可口可乐',
						sales: 100,
						priceInt: 3,
						priceFloat: 0,
						integral: 300
					},
					{
						img: 'background-image: url("https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1");',
						imgUrl: 'https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1',
						title: '干脆面',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					},
					{
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						imgUrl: 'https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg',
						title: '笔记本',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					},
					{
						img: 'background-image: url("https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1");',
						imgUrl: 'https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1',
						title: '干脆面',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					},
					{
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						imgUrl: 'https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg',
						title: '笔记本',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					},
					{
						img: 'background-image: url("https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1");',
						imgUrl: 'https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1',
						title: '干脆面',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					},
					{
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						imgUrl: 'https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg',
						title: '笔记本',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					},
					{
						img: 'background-image: url("https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1");',
						imgUrl: 'https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1',
						title: '干脆面',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					},
					{
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						imgUrl: 'https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg',
						title: '笔记本',
						sales: 100,
						priceInt: 1,
						priceFloat: 0,
						integral: 100
					}
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
			itemClick(index) {
				this.jump('/pages/mall/item', this.recommendedList[index])
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
		position: fixed;
		top: 5vh;
		left: 2vh;
		color: #fff;
	}
	
	.search {
		padding: 20rpx;
		padding-top: 0;
	}
	
	.body {
		height: 83vh;
		overflow: scroll;
		
		.recommended {
			margin: 20rpx;
			.recommended-content {
				background-color: rgba($color: #fff, $alpha: .5);
				margin-bottom: 20rpx;
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
}
</style>
