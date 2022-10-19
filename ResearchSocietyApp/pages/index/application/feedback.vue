<template>
	<view class="content">
		<view class="question">
			<h3 class="title">
				问题和意见
			</h3>
			
			<u--textarea v-model="text" placeholder="请输入内容" ></u--textarea>
		</view>
		
		<view class="button">
			<u-button type="success">提交</u-button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				text: ''
			}
		},
		methods: {
			
		}
	}
</script>

<style lang="scss" scoped>
.content {
	display: flex;
	flex-direction: column;
	align-items: center;
	
	.question {
		width: 90vw;
		padding: 20rpx;
		
		.title {
			margin: 20rpx;
			font-size: medium;
			font-weight: normal;
			color: #888;
		}
	}
	
	.button {
		width: 50vw;
	}
}
</style>
