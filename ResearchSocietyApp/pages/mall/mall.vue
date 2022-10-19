<template>
	<view class="content">
		<!-- 搜索框 -->
		<view class="search">
			<u-search></u-search>
		</view>
		
		<!-- 应用 -->
		<view class="application">
			<myGrid class="grid" :myGridList="gridList" :myGridCol="4" v-slot="myGridData">
				<view class="grid-item">
					<u-icon
						:customStyle="{paddingTop:20+'rpx'}"
						:name="myGridData.item.name"
						:size="30"
						:color="myGridData.item.color"
					></u-icon>
					<text class="grid-text">{{myGridData.item.title}}</text>
				</view>
			</myGrid>
		</view>
		
		<!-- 推荐分组 -->
		<myLayout class="sort" :myLayoutList="sortList" v-slot="myLayoutData">
			<view  class="sort-content">
				<view class="sort-name">
					<h3 class="sort-name-title" style="margin-left: 20rpx;">{{myLayoutData.item.title}}</h3>
					<u-icon :name="myLayoutData.item.ico" size="50" style="margin-left: 20rpx;"></u-icon>
				</view>
				<view class="sort-img" :style="myLayoutData.item.img" />
			</view>
		</myLayout>
		
		<!-- 推荐商品 -->
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
</template>

<script>
	export default {
		data() {
			return {
				gridList: [
					{
						name: 'calendar-fill',
						color: '#ff557f',
						title: '分类',
						url: '/pages/mall/application/sort'
					},
					{
						name: 'clock-fill',
						color: '#ffff00',
						title: '购物车',
						url: '/pages/mall/application/shoppingCart'
					},
					{
						name: 'grid-fill',
						color: '#55ff7f',
						title: '订单',
						url: '/pages/mall/application/order'
					},
					{
						name: 'gift-fill',
						color: '#aa55ff',
						title: '账户',
						url: '/pages/mall/application/account'
					},
					{
						name: 'man-add-fill',
						color: '#00ffff',
						title: '兑换券',
						url: '/pages/mall/application/voucher'
					},
					{
						name: 'gift-fill',
						color: '#aaff7f',
						title: '优惠券',
						url: '/pages/mall/application/coupons'
					},
					{
						name: 'integral-fill',
						color: '#ffaaff',
						title: '赚积分',
						url: '/pages/mall/application/earnPoints'
					},
					{
						name: 'setting-fill',
						color: '#a2a2a2',
						title: '编辑',
						url: '/pages/mall/application/editor'
					}
				],
				sortList: [
					{
						title: '奶茶铺子',
						ico: '../../static/svg/milkTea.svg',
						img: 'background-image: url("https://tse1-mm.cn.bing.net/th/id/OIP-C.ZJUMQOtHPnWLkVuPxjVSSwHaIv?pid=ImgDet&rs=1");',
						url: '/pages/mall/shop'
					},
					{
						title: '零食铺子',
						ico: '../../static/svg/snacks.svg',
						img: 'background-image: url("https://tse4-mm.cn.bing.net/th/id/OIP-C.s3FA_4X7BSfTiKurjfq5AQHaHa?pid=ImgDet&rs=1");',
						url: '/pages/mall/shop'
					},
					{
						title: '学习资料',
						ico: '../../static/svg/study.svg',
						img: 'background-image: url("https://www.amphasisdesign.com/image/cache/catalog/corporategifts/Customised%20Products/customized-pu-leather-notebook-800x800.jpg");',
						url: '/pages/mall/shop'
					},
					{
						title: '兑换时长',
						ico: '../../static/svg/time.svg',
						img: 'background-image: url("https://ts1.cn.mm.bing.net/th/id/R-C.96db53825942c485e2a1d31f1679fff3?rik=coDhrKg1RnfYrg&riu=http%3a%2f%2fpicapi.zhituad.com%2fphoto%2f69%2f35%2f46EAB.jpg&ehk=LcGVQNUtSfKbEVtiHu68jp%2bKwDpSJrvnVdQkHKCzHNM%3d&risl=&pid=ImgRaw&r=0&sres=1&sresct=1");',
						url: 'pages/mall/application/account'
					}
				],
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
			},
			itemClick(index) {
				this.jump('/pages/mall/item', this.recommendedList[index])
			}
		}
	}
</script>

<style lang="scss">
.content {
	background-image: linear-gradient(#f7edf5, #ffffff);
	
	.search {
		background-color: #fff;
		padding: 20rpx;
	}
	
	.application {
		background-color: #fff;
		margin: 20rpx;
		
		border-radius: 20rpx;
		
		font-size: small;
		
		.grid {
			.grid-item {
				margin-bottom: 20rpx;
				
				display: flex;
				flex-direction: column;
				align-items: center;
			}
		}
	}
	
	.sort {
		background-color: #fff;
		margin: 20rpx;
		padding: 20rpx;
		
		border-radius: 20rpx;
		
		.sort-content {
			display: flex;
			justify-content: space-around;
			
			.sort-name {
				.sort-name-title {
					font-size: small;
					font-weight: normal;
				}
			}
			.sort-img {
				height: 120rpx;
				width: 120rpx;
				background-size:100% 100%;
				background-repeat:no-repeat;
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
