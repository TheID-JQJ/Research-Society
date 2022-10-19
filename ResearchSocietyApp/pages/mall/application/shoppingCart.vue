<template>
	<view class="content">
		<view class="back" @click="back()">
			< 返回
		</view>
		
		<view class="cart-title">
			<view class="select-all">
				<u-checkbox-group @change="selectAllChange">
					<u-checkbox :checked="allSelected"></u-checkbox>
				</u-checkbox-group>
				<text>全选</text>
			</view>
			<view class="total">
				共{{itemList.length}}个商品
			</view>
			<view class="edit" @click="editChange" v-if="!isEdit">
				编辑
			</view>
			<view class="edit" @click="editChange" v-else>
				取消
			</view>
		</view>
		
		<view class="goods">
			<view class="goods-list" v-for="item, index in itemList">
				<view class="goods-img" :style="item.img" @click="toItem(item)">
				</view>
				<view class="goods-message" @click="toItem(item)">
					<h3 class="title">{{item.title}}</h3>
					<text class="text">月售{{item.sales}}</text>
					<view class="price">
						<text class="price-int">￥{{item.priceInt}}</text>
						<text class="price-float">.{{item.priceFloat}}</text>/
						<text class="price-integral">{{item.integral}}积分</text>
					</view>
				</view>
				<view class="checkbox">
					<u-checkbox-group @change="checkedChange(index)">
						<u-checkbox :checked="isChecked[index]"></u-checkbox>
					</u-checkbox-group>
				</view>
			</view>
		</view>
		
		<view class="cart-button">
			<view class="total-selected">
				已选择{{selectedGoods.length}}个商品
			</view>
			<u-button class="button" type="primary" shape="circle" v-if="!isEdit">
				立即购买
			</u-button>
			<u-button class="button" type="error" shape="circle" v-else>
				删除
			</u-button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				itemList: [
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
				],
				allSelected: false,
				isEdit: false,
				isChecked: []
			}
		},
		computed: {
			selectedGoods() {
				let selectedGoodsList = []
				
				for (var i = 0; i < this.isChecked.length; i++) {
					if (this.isChecked[i]) {
						// selectedGoodsList.push(this.itemList[i].id)
						selectedGoodsList.push(this.itemList[i].kid) //这儿要改
					}
				}
				
				return selectedGoodsList
			}
		},
		onLoad() {
			for (var i = 0; i < this.itemList.length; i++) {
				this.isChecked.push(false)
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
			selectAllChange() {
				this.allSelected = !this.allSelected
				for (var i = 0; i < this.isChecked.length; i++) {
					this.$set(this.isChecked, i, this.allSelected)
				}
			},
			editChange() {
				this.isEdit = !this.isEdit
			},
			checkedChange(index) {
				this.$set(this.isChecked, index, !this.isChecked[index])
				for (var i = 0; i < this.isChecked.length; i++) {
					if (!this.isChecked[i]) {
						this.allSelected = false
						return
					}
				}
				this.allSelected = true
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
	
	.cart-title {
		height: 5vh;
		margin: 0 20rpx;
		background-color: rgba($color: #fff, $alpha: .3);
		border-radius: 20rpx;
		display: flex;
		justify-content: space-between;
		line-height: 5vh;
		text-align: center;
		
		.select-all {
			width: 25%;
			display: flex;
			justify-content: center;
		}
		
		.total {
			width: 25%;
		}
		
		.edit {
			width: 25%;
		}
	}
	
	.goods {
		height: 70vh;
		margin: 20rpx;
		padding: 20rpx;
		padding-top: 0;
		background-color: rgba($color: #fff, $alpha: .3);
		border-radius: 20rpx;
		overflow: scroll;
		
		.goods-list {
			width: 100%;
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
				width: 40vw;
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
	
	.cart-button {
		height: 9vh;
		margin: 0 20rpx;
		background-color: rgba($color: #fff, $alpha: .3);
		border-radius: 20rpx;
		display: flex;
		// flex-direction: column;
		justify-content: space-around;
		align-items: center;
		
		.total-selected {
			width: 50%;
			text-align: center;
		}
		
		.button {
			width: 40%;
		}
	}
}
</style>
