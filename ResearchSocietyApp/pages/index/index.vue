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
				<u-count-to :startVal="0" :endVal="myLayoutData.item.message" :bold="true"></u-count-to>
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
					'../../static/img/20210409113856123009.jpg',
					'../../static/img/v2-30effef829f9e1647e3644d18b9e3eea_r.jpg'
				],
				applicationList: [
					{
						name: 'calendar-fill',
						color: '#ff557f',
						title: '签到',
						url: '/pages/index/application/sign'
					},
					{
						name: 'clock-fill',
						color: '#ffff00',
						title: '续时',
						url: "/pages/mall/application/account"
					},
					{
						name: 'grid-fill',
						color: '#55ff7f',
						title: '排行榜',
						url: '/pages/index/application/sort'
					},
					{
						name: 'gift-fill',
						color: '#aa55ff',
						title: '积分兑换',
						url: '/pages/mall/application/sort'
					},
					{
						name: 'man-add-fill',
						color: '#00ffff',
						title: '学友匹配',
						url: '/pages/index/application/friend'
					},
					{
						name: 'pushpin-fill',
						color: '#aaff7f',
						title: '公告',
						url: '/pages/index/application/notice'
					},
					{
						name: 'error-circle-fill',
						color: '#ffaaff',
						title: '反馈',
						url: '/pages/index/application/feedback'
					},
					{
						name: 'setting-fill',
						color: '#a2a2a2',
						title: '编辑',
						url: '/pages/index/application/editor'
					}
				],
				messageList: [
					{
						title: '学习总时长(/分钟)',
						message: 600,
						url: 'pages/index/application/study'
					},
					{
						title: '今日学习时长(/分钟)',
						message: 100,
						url: 'pages/index/application/study'
					},
					{
						title: '剩余时长(/分钟)',
						message: 80,
						url: 'pages/index/application/study'
					},
					{
						title: '排名',
						message: 1,
						url: 'pages/index/application/study'
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
			change() {
				
			},
			click(e) {
				console.log(e)
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
