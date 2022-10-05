<template>
	<view class="content">
		<!-- 轮播图 -->
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
		
		<!-- 信息统计 -->
		<myLayout class="message" :myLayoutList="messageList" :myLayoutGutter="20" v-slot="myLayoutData">
			<view class="size">
				<h3 class="title-size">{{myLayoutData.item.title}}</h3>
				<text class="text-size">{{myLayoutData.item.message}}</text>
			</view>
		</myLayout>
		
		<!-- 应用 -->
		<view class="application-background">
			<view class="application">
				<view class="appoint" @click="toAppoint">
					预约
				</view>
				
				<myGrid class="grid" :myGridList="applicationList" :myGridCol="4" v-slot="myGridData">
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
				applicationList: [
					{
						name: 'calendar-fill',
						color: '#ff557f',
						title: '签到'
					},
					{
						name: 'clock-fill',
						color: '#ffff00',
						title: '续时'
					},
					{
						name: 'grid-fill',
						color: '#55ff7f',
						title: '排行榜'
					},
					{
						name: 'gift-fill',
						color: '#aa55ff',
						title: '积分兑换'
					},
					{
						name: 'man-add-fill',
						color: '#00ffff',
						title: '学友匹配'
					},
					{
						name: 'pushpin-fill',
						color: '#aaff7f',
						title: '公告'
					},
					{
						name: 'error-circle-fill',
						color: '#ffaaff',
						title: '反馈'
					},
					{
						name: 'setting-fill',
						color: '#a2a2a2',
						title: '编辑'
					}
				],
				messageList: [
					{
						title: '学习总时长',
						message: '48时30分'
					},
					{
						title: '今日学习时长',
						message: '2时50分'
					},
					{
						title: '剩余时长',
						message: '1时40分'
					},
					{
						title: '排名',
						message: '1'
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
			},
			jump(path, params) {
				// 页面的跳转
				this.$u.route({
					url: path,
					params
				})
			},
			toAppoint() {
				this.jump('pages/index/appoint', {})
			}
		}
	}
</script>

<style lang="scss">
.content {
	.swiper {
		background-color: #f7edf5;
		padding: 40rpx;
	}
	
	.message {
		padding: 40rpx;
		padding-top: 0;
		padding-bottom: 0;
		background-color: #f7edf5;
		
		.size {
			height: 150rpx;
			margin-bottom: 40rpx;
			background-color: #fff;
			
			border-radius: 20rpx;
			
			text-align: center;
			line-height: 70rpx;
		}
		
		.title-size {
			font-size: small;
			font-weight: normal;
		}
		
		.text-size {
			font-size: large;
			font-weight: bold;
		}
	}
	
	.application-background {
		background-image: linear-gradient(#f7edf5, #ffffff);
		padding: 0rpx 40rpx 0 40rpx;
		
		.application {
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
				margin-top: 40rpx;
				
				font-size: small;
				
				.grid-item {
					margin-bottom: 20rpx;
					
					display: flex;
					flex-direction: column;
					align-items: center;
				}
			}
		}
	}

}
</style>
