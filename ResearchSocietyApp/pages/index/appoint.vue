<template>
	<view class="content">
		<view class="back" @click="back()">
			< 返回
		</view>
		
		<view class="sit-img">
			<view class="item sit-img-item" :id="sitImgStyle" :style="sitImgList[select%2]">
			</view>
			
			<view class="item first" :id="firstStyle">
			</view>
			<view class="item second" :id="secondStyle">
			</view>
			<view class="item third" :id="thirdStyle">
			</view>
			<view class="item forth" :id="forthStyle">
			</view>
		</view>
		
		<view class="sit-botton">
			<view class="left-botton" @click="toRight()">
				<u-icon name="rewind-left" color="#fff" size="32"></u-icon>
			</view>
			<view class="center-botton" @click="appoint()">
				<image src="../../static/icon/appoint.png" style="width: 100%; height: 100%;"></image>
			</view>
			<view class="right-botton" @click="toLeft()">
				<u-icon name="rewind-right" color="#fff" size="32"></u-icon>
			</view>
		</view>
		
		<view class="sit-list">
			<view class="item" :class="select==index?'select':''" v-for="item, index in sitList" @click="sitClick(index)">
				{{item+1}}
			</view>
		</view>
		
		<view class="select-time">
			<u-button class="button" type="default" shape="circle" @click="isStartTimeShow()">
				<text v-if="startTime==''">请选择时间段</text>
				<text v-else style="color: #ff1313;">{{startTime+' - '+endTime}}</text>
			</u-button>
		</view>
		<u-datetime-picker
			:show="startTimeShow"
			v-model="startTime"
			mode="time"
			:minHour="7"
			:maxHour="22"
			title="请选择开始时间"
			confirmText="下一步"
			:closeOnClickOverlay="true"
			@cancel="isStartTimeShow()"
			@close="isStartTimeShow()"
			@confirm="isEndTimeShow()"
		></u-datetime-picker>
		<u-datetime-picker
			:show="endTimeShow"
			v-model="endTime"
			mode="time"
			:minHour="7"
			:maxHour="22"
			title="请选择结束时间"
			confirmText="完成"
			:closeOnClickOverlay="true"
			@cancel="cancelEndTime()"
			@close="cancelEndTime()"
			@confirm="finishTime()"
		></u-datetime-picker>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				sitImgList: [
					'background-image: url("../../static/img/20210409113856123009.jpg")',
					'background-image: url("../../static/img/v2-30effef829f9e1647e3644d18b9e3eea_r.jpg")'
				],
				sitList: [
				],
				sitImgStyle: '',
				firstStyle: '',
				secondStyle: '',
				thirdStyle: '',
				forthStyle: '',
				styleLeftList: ['forth-left', 'third-left', 'second-left', 'first-left'],
				styleRightList: ['first-right', 'second-right', 'third-right', 'forth-right'],
				imgStyleList: ['img-in', 'img-in-two'],
				clickNum: 0,
				select: 0,
				startTime: '',
				endTime: '',
				startTimeShow: false,
				endTimeShow: false
			};
		},
		onLoad(params) {
			for (var i = 0; i < 20; i++) {
				this.sitList.push(i)
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
			async toLeft(i) {
				this.sitImgStyle = this.imgStyleList[this.clickNum%2]
				this.forthStyle = this.styleLeftList[this.clickNum]
				this.thirdStyle = this.styleLeftList[(this.clickNum+1)%4]
				this.secondStyle = this.styleLeftList[(this.clickNum+2)%4]
				this.firstStyle = this.styleLeftList[(this.clickNum+3)%4]
				this.clickNum = (this.clickNum+1)==4?0:this.clickNum+1
				if (i) {
					this.select = i
				} else {
					this.select = (this.select+1)==this.sitList.length?0:this.select+1
				}
				// console.log(this.select)
			},
			toRight(i) {
				this.sitImgStyle = this.imgStyleList[this.clickNum%2]
				this.firstStyle = this.styleRightList[this.clickNum]
				this.secondStyle = this.styleRightList[(this.clickNum+1)%4]
				this.thirdStyle = this.styleRightList[(this.clickNum+2)%4]
				this.forthStyle = this.styleRightList[(this.clickNum+3)%4]
				this.clickNum = (this.clickNum+1)==4?0:this.clickNum+1
				if (i) {
					this.select = i
				} else {
					this.select = (this.select-1)==-1?this.sitList.length-1:this.select-1
				}
				// console.log(this.select)
			},
			appoint() {
				console.log(this.endTime)
			},
			sitClick(index) {
				if (index>this.select) {
					this.toLeft(index)
				} else if (index<this.select) {
					this.toRight(index)
				}
			},
			isStartTimeShow() {
				this.startTimeShow = !this.startTimeShow
			},
			isEndTimeShow() {
				this.endTimeShow = !this.endTimeShow
			},
			cancelEndTime() {
				this.isEndTimeShow()
				this.startTime = ''
			},
			finishTime() {
				this.isEndTimeShow()
				this.isStartTimeShow()
			}
		}
	}
</script>

<style lang="scss" scoped>
.content {
	width: 100vw;
	height: 100vh;
	background-image: linear-gradient(#aaaaff, #00ffff);
	position: fixed;
	overflow: scroll;
	transform-style: preserve-3d;
	perspective: 1000rpx;
	
	.back {
		z-index: 99;
		position: absolute;
		top: 80rpx;
		left: 50rpx;
		color: #fff;
	}
	
	.sit-img {
		width: 100vw;
		height: 100vw;
		position: relative;
		overflow: hidden;
		
		.item {
			width: 50vw;
			height: 70vw;
			background-color: rgba($color: #fff, $alpha: .3);
			border-radius: 40rpx;
			position: absolute;
			top: 150rpx;
			left: 25vw;
		}
		.sit-img-item {
			z-index: 99;
			background-size:100% 100%;
		}
		
		.first {
			transform: translateX(-70vw) translateZ(-100rpx);
		}
		
		.third {
			transform: translateX(70vw) translateZ(-100rpx);
		}
		
		.forth {
			transform: translateX(100vw) translateZ(-100rpx);
		}
	}
	
	.sit-botton {
		display: flex;
		justify-content: center;
		margin-bottom: 60rpx;
		
		.left-botton,
		.right-botton {
			width: 20vw;
			height: 20vw;
			
			display: flex;
			justify-content: center;
			align-items: center;
		}
		
		.center-botton {
			width: 20vw;
			height: 20vw;
		}
	}
	
	.sit-list{
		width: 90vw;
		height: 60vw;
		margin: 40rpx;
		padding: 20rpx 0;
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
		background-color: rgba($color: #fff, $alpha: .5);
		border-radius: 40rpx;
		overflow: scroll;
		
		.item {
			z-index: 99;
			width: 10vw;
			height: 10vw;
			background-color: #fff;
			margin: 20rpx;
			border-radius: 20rpx;
			
			text-align: center;
			line-height: 10vw;
		}
		
		.select {
			background-color: #aaff7f;
		}
	}
	
	.select-time {
		width: 50vw;
		margin: 0 auto;
		
		.button {
			background-color: rgba($color: #fff, $alpha: .5);
			color: #888;
		}
	}
}
</style>

<style scoped>
@keyframes ftr {
	from {
		transform: translateX(-70vw) translateZ(-100rpx);
	}
	to {
		transform: translateX(0) translateZ(0);
	}
}
@keyframes ftl {
	from {
		transform: translateX(-70vw) translateZ(-100rpx);
	}
	to {
		transform: translateX(-100vw) translateZ(-100rpx);
	}
}
#first-right {
	animation-name: ftr;
	animation-duration: 1s;
	transform: translateX(0) translateZ(0);
}
#first-left {
	animation-name: ftl;
	animation-duration: 1s;
	transform: translateX(100vw) translateZ(-100rpx);
}

@keyframes str {
	from {
		transform: translateX(0) translateZ(0);
	}
	to {
		transform: translateX(70vw) translateZ(-100rpx);
	}
}
@keyframes stl {
	from {
		transform: translateX(0) translateZ(0);
		
	}
	to {
		transform: translateX(-70vw) translateZ(-100rpx);
	}
}
#second-right {
	animation-name: str;
	animation-duration: 1s;
	transform: translateX(70vw) translateZ(-100rpx);
}
#second-left {
	animation-name: stl;
	animation-duration: 1s;
	transform: translateX(-70vw) translateZ(-100rpx);
}
@keyframes ttr {
	from {
		transform: translateX(70vw) translateZ(-100rpx);
	}
	to {
		transform: translateX(100vw) translateZ(-100rpx);
	}
}
@keyframes ttl {
	from {
		transform: translateX(70vw) translateZ(-100rpx);
		
	}
	to {
		transform: translateX(0) translateZ(0);
	}
}
#third-right {
	animation-name: ttr;
	animation-duration: 1s;
	transform: translateX(-100vw) translateZ(-100rpx);
}
#third-left {
	animation-name: ttl;
	animation-duration: 1s;
	transform: translateX(0) translateZ(0);
}
@keyframes fotr {
	from {
		transform: translateX(-100vw) translateZ(-100rpx);
	}
	to {
		transform: translateX(-70vw) translateZ(-100rpx);
	}
}
@keyframes fotl {
	from {
		transform: translateX(100vw) translateZ(-100rpx);
		
	}
	to {
		transform: translateX(70vw) translateZ(-100rpx);
	}
}
#forth-right {
	animation-name: fotr;
	animation-duration: 1s;
	transform: translateX(-70vw) translateZ(-100rpx);
}
#forth-left {
	animation-name: fotl;
	animation-duration: 1s;
	transform: translateX(70vw) translateZ(-100rpx);
}
@keyframes ii {
	0% {
		opacity: 0;
	}
	
	25% {
		opacity: 0;
	}
	
	100% {
		opacity: 1;
	}
}
@keyframes iit {
	0% {
		opacity: 0;
	}
	
	50% {
		opacity: 0;
	}
	
	100% {
		opacity: 1;
	}
}
#img-in {
	animation-name: ii;
	animation-duration: 3s;
	opacity: 1;
}
#img-in-two {
	animation-name: iit;
	animation-duration: 3s;
	opacity: 1;
}
</style>