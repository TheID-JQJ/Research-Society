<template>
	<view class="content">
		<view class="swiper">
			<u-swiper
				:list="swiperList"
				height="150"
				radius="20"
				bgColor="#f7edf5"
				@change="change"
				@click="click"
			>
			</u-swiper>
		</view>
		
		<view class="gradient">
			<view class="order">
				<view class="appoint">
					预约
				</view>
				
				<view class="grid">
					<u-grid
						:border="false"
						@click="click"
						col="4"
					>
						<u-grid-item
							v-for="(baseListItem,baseListIndex) in baseList"
							:key="baseListIndex"
						>
							<u-icon
								:customStyle="{paddingTop:20+'rpx'}"
								:name="baseListItem.name"
								:size="40"
								:color="baseListItem.color"
							></u-icon>
							<text class="grid-text">{{baseListItem.title}}</text>
						</u-grid-item>
					</u-grid>
					<u-toast ref="uToast" />
				</view>
			</view>
		</view>
		
		<h3>学习信息</h3> 
		
		<view class="message">
			<view class="size message1">
				message1
			</view>
			<view class="size message2">
				message2
			</view>
			<view class="size message3">
				message3
			</view>
			<view class="size message4">
				message4
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				title: 'index',
				swiperList: [
					'https://img.ruanwenpu.com/Public/upfiles/60291/202104/20210409113856123009.jpg',
					'https://pic1.zhimg.com/v2-30effef829f9e1647e3644d18b9e3eea_r.jpg?source=172ae18b'
				],
				baseList: [
					{
						name: "calendar-fill",
						color: "#ff557f",
						title: "签到"
					},
					{
						name: "clock-fill",
						color: "#ffff00",
						title: "剩余时长"
					},
					{
						name: "grid-fill",
						color: "#55ff7f",
						title: "排行榜"
					},
					{
						name: "gift-fill",
						color: "#aa55ff",
						title: "积分兑换"
					}
				]
			}
		},
		onLoad() {
			
		},
		methods: {
			change() {
				
			},
			click(e) {
				console.log(e)
			}
		}
	}
</script>

<style lang="scss">
.content {
	// width: 100vw;
	// height: 100vh;
	// background-color: #e2e2e2;
	// background-image: linear-gradient(#e2e2e2, #ffffff);
	
	// padding-top: 50rpx;
	
	// .text-area {
	// 	background-color: #e2e2e2;
	// 	height: 100rpx;
	// }
	.swiper {
		background-color: #f7edf5;
		padding: 40rpx;
	}
	
	.gradient {
		background-image: linear-gradient(#f7edf5, #ffffff);
		padding: 40rpx 40rpx 0 40rpx;
		
		.order {
			background-color: #fff;
		
			padding: 40rpx;
			
			border-radius: 40rpx 40rpx 0 0;
			
			.appoint {
				width: 200rpx;
				height: 200rpx;
				background-color: #007AFF;
				margin: 0 auto;
				border-radius: 50%;
				
				color: #fff;
				text-align: center;
				line-height: 200rpx;
			}
			
			.grid {
				// background-color: #007AFF;
				margin: 40rpx 0 0 0;
			}
		}
	}
	
	.message {
		padding: 40rpx;
		
		display: flex;
		
		.size {
			width: 40%;
			height: 150rpx;
		}
		
		.message1 {
			background-color: #007AFF;
		}
		
		.message2 {
			
		}
		
		.message3 {
			
		}
		
		.message4 {
			background-color: #007AFF;
		}
	}
}
</style>
