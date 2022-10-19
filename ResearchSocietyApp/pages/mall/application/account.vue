<template>
	<view class="content">
		<view class="back" @click="back()">
			< 返回
		</view>
		
		<view class="account">
			<view class="money" @click="addMoney()">
				<h3 class="account-title">余额</h3>
				<text class="account-content">100.00</text>
			</view>
			<view class="time-and-integral">
				<view class="integral" @click="addTime()">
					<h3 class="account-title">积分</h3>
					<text class="account-content">1000</text>
				</view>
				<view class="time" @click="addTime()">
					<h3 class="account-title">剩余时长(/分钟)</h3>
					<text class="account-content">1000</text>
				</view>
			</view>
		</view>
		
		<view class="add add-money" v-if="isAddMoney">
			<view class="triangle" />
			<h3 class="add-title">
				<text>充值</text>
				<text>提现</text>
			</h3>
			<view class="add-list">
				<view class="add-item" :class="index==moneySelected?'selected':''" 
					:id="index" v-for="item, index in moneyList"
					@click="selectMoney(index, item.value)"
				>
					{{item.text}}
				</view>
				<view class="add-item custom-item" :class="5==moneySelected?'selected':''">
					<text v-if="!isCustomMoney" @click="customMoney()">自定义</text>
					<view class="" v-else style="width: 50%;">
						<u--input
							type="number"
						    border="none"
							:focus="true"
							style="border: 1rpx solid #fff;"
							@blur="moneyBlur"
						  ></u--input>
					</view>
					<text v-if="isCustomMoney">元</text>
				</view>
			</view>
			<view class="add-botton">
				<u-button type="error" @click="buyMoney()">充值</u-button>
			</view>
		</view>
		<view class="add add-time" v-else>
			<view class="triangle" />
			<h3 class="add-title">
				<text>购买时长</text>
				<text></text>
			</h3>
			<view class="add-list">
				<view class="add-item" :class="index==timeSelected?'selected':''" 
					:id="index" v-for="item, index in timeList"
					@click="selectTime(index, item.value)"
				>
					{{item.text}}
				</view>
				<view class="add-item custom-item" :class="5==timeSelected?'selected':''">
					<text v-if="!isCustomTime" @click="customTime()">自定义</text>
					<view class="" v-else style="width: 50%;">
						<u-input
							type="number"
						    border="none"
							:focus="true"
							style="border: 1rpx solid #fff;"
							@blur="timeBlur"
						  ></u-input>
					</view>
					<text v-if="isCustomTime">分钟</text>
				</view>
			</view>
			<view class="add-botton">
				<u-button type="error" @click="buyTime()">购买</u-button>
			</view>
		</view>
		
		<view class="tips" style="margin: 20rpx;" v-if="!isAddMoney">
			<text style="color: #666;">
				注：(1积分=0.01元)
				购买时长时可用积分直接购买或抵扣相应金额。
			</text>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				isAddMoney: false,
				moneyList:[{text: '10元', value: 10}, {text: '20元', value: 20}, {text: '30元', value: 30}, {text: '50元', value: 50}, {text: '100元', value: 100}],
				timeList:[{text: '60分钟', value: 60}, {text: '2小时', value: 120}, {text: '24小时', value: 1440}, {text: '7天', value: 10080}, {text: '30天', value: 43200}],
				moneySelected: 0,
				timeSelected: 0,
				isCustomMoney: false,
				isCustomTime: false,
				moneyValue: 10,
				timeValue: 60,
			}
		},
		onLoad(value) {
			if (value=='money') {
				this.isAddMoney = true
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
			addMoney() {
				this.isAddMoney = true
			},
			addTime() {
				this.isAddMoney = false
			},
			selectMoney(i, value) {
				this.isCustomMoney = false
				this.moneySelected = i
				this.moneyValue = value
			},
			customMoney() {
				this.isCustomMoney = true
				this.moneySelected = 5
			},
			selectTime(i, value) {
				this.isCustomTime = false
				this.timeSelected = i
				this.timeValue = value
			},
			customTime() {
				this.isCustomTime = true
				this.timeSelected = 5
			},
			moneyBlur(value) {
				this.moneyValue = value
			},
			timeBlur(value) {
				this.timeValue = value
			},
			buyMoney() {
				console.log(this.moneyValue)
			},
			buyTime() {
				console.log(this.timeValue)
			},
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
	
	.account {
		display: flex;
		justify-content: space-around;
		margin: 20rpx 0;
		
		.money {
			background-color: rgba($color: #fff, $alpha: .3);
			width: 26vw;
			height: 26vw;
			padding: 2vw;
			border-radius: 20rpx;
			display: flex;
			flex-direction: column;
			justify-content: center;
		}
		
		.time-and-integral {
			background-color: rgba($color: #fff, $alpha: .3);
			width: 56vw;
			height: 26vw;
			padding: 2vw;
			border-radius: 20rpx;
			display: flex;
			
			.time,
			.integral {
				width: 28vw;
				height: 26vw;
				display: flex;
				flex-direction: column;
				justify-content: center;
			}
		}
		
		.account-title {
			font-size: small;
			font-weight: normal;
		}
		
		.account-content {
			font-size: x-large;
		}
	}
	
	.add {
		width: 95vw;
		// height: 50vh;
		background-color: rgba($color: #fff, $alpha: .3);
		margin: 20rpx auto;
		border-radius: 20rpx;
	}
	
	.triangle {
		width: 0;
		height: 0;
		border: 10px solid;
		border-color: transparent transparent rgba($color: #fff, $alpha: .3) transparent;
		position: absolute;
		top: -20px;
	}
	
	.add-title {
		width: 90%;
		display: flex;
		justify-content: space-between;
		padding: 20rpx 40rpx;
		font-size: medium;
		font-weight: normal;
	}
	
	.add-list {
		padding: 20rpx;
		display: flex;
		flex-wrap: wrap;
		justify-content: space-around;
		
		.add-item {
			width:25vw;
			height: 15vw;
			background-color: rgba($color: #fff, $alpha: .5);
			margin-bottom: 20rpx;
			border-radius: 20rpx;
			
			text-align: center;
			line-height: 15vw;
		}
		
		.selected {
			background-color: rgba($color: #aaff7f, $alpha: .5);
		}
	}
	
	.custom-item {
		display: flex;
		justify-content: center;
		align-items: center;
	}
	
	.add-botton {
		padding: 40rpx 25vw;
		
	}
	
	.add-money {
		position: relative;
		
		.triangle {
			left: 10vw;
		}
	}
	
	.add-time {
		position: relative;
		
		.triangle {
			left: 75vw;
		}
	}
}
</style>
