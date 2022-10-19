<template>
	<view class="content">
		<view class="back" @click="back()">
			< 返回
		</view>
		
		<view class="kinds">
			<view class="kinds-item" :class="selected==item.id?'selected':''" v-for="item, index in kindsList" @click="kindsClick(item.id)">
				{{item.name}}
				<u-icon name="play-right-fill" v-if="selected==item.id"></u-icon>
			</view>
		</view>
		
		<view class="goods">			
			<view class="goods-item" v-for="item, index in goodsList" @click="toItem(item)">
				<view class="goods-img" :style="item.img">
				</view>
				<view class="goods-message">
					<h3 class="title">{{item.title}}</h3>
					<text class="text">月售{{item.sales}}</text>
					<view class="price">
						<text class="price-int">￥{{item.priceInt}}</text>
						<text class="price-float">.{{item.priceFloat}}</text>/
						<text class="price-integral">{{item.integral}}积分</text>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				kindsList: [{id: 1, name: '奶茶'}, {id: 2, name: '饮料'}, {id: 3, name: '零食'}, {id: 4, name: '资料'}],
				selected: 1,
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
					}
				]
			}
		},
		onLoad() {
			this.selected = this.kindsList[0].id
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
			kindsClick(id) {
				this.selected = id
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
	
	.kinds {
		width: 20vw;
		height: 88vh;
		background-color: rgba($color: #fff, $alpha: .3);
		position: fixed;
		left: 0;
		border-radius: 0 20rpx 20rpx 0;
		display: flex;
		flex-direction: column;
		align-items: center;
		overflow: scroll;
		
		.kinds-item {
			width: 20vw;
			height: 10vw;
			margin: 20rpx 0 0 0;
			line-height: 10vw;
			display: flex;
			justify-content: center;
		}
		
		.selected {
			font-size: larger;
		}
	}
	
	.goods {
		width: 75vw;
		height: 88vh;
		background-color: rgba($color: #fff, $alpha: .3);
		padding-bottom: 20rpx;
		position: fixed;
		right: 0;
		border-radius: 20rpx 0 0 20rpx;
		display: flex;
		flex-direction: column;
		align-items: center;
		overflow: scroll;
		
		.goods-item {
			width: 70vw;
			height: 30vw;
			background-color: rgba($color: #fff, $alpha: .5);
			margin-top: 20rpx;
			border-radius: 20rpx;
			display: flex;
			justify-content: space-between;
			align-items: center;
			
			.goods-img {
				width: 30vw;
				height: 30vw;
				background-size: 100% 100%;
				border-radius: 20rpx;
			}
			
			.goods-message {
				width: 35vw;
				height: 100%;
				display: flex;
				flex-direction: column;
				justify-content: space-around;
				
				.title {
					font-size: large;
					font-weight: normal;
				}
				
				.text {
					color: #888;
					font-size: small;
				}
				
				.price {
					color: #f00;
					font-size: small;
					
					.price-int {
						font-size: medium;
					}
				}
			}
		}
	}
}
</style>
