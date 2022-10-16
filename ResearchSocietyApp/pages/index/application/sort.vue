<template>
	<view class="content">
		<view class="top">
			<view class="search">
				<u-search></u-search>
			</view>
			
			<view class="tabs">
				<u-tabs 
					:list="tabs" 
					@change="tabsChange" 
					lineColor="#ff814f"
					:activeStyle="{
						fontWeight: 'bold',
						transform: 'scale(1.05)'
					}"
					 :inactiveStyle="{
						transform: 'scale(1)'
					}"
					itemStyle="padding-left: 15vw; padding-right: 15vw; height: 34px;"
				></u-tabs>
			</view>
		</view>
		
		<u-list
			class="item-list"
			@scrolltolower="scrolltolower"
		>
			<u-list-item
				v-for="(item, index) in indexList"
				:key="index"
			>
				<u-cell
					:title="`列表长度-${index + 1}`"
				>
					<u-avatar
						slot="icon"
						shape="square"
						size="35"
						:src="urls[index]"
						customStyle="margin: -3px 5px -3px 0"
					></u-avatar>
				</u-cell>
			</u-list-item>
		</u-list>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				tabs: [{name: "用户榜"}, {name: "好友榜"}],
				indexList: [1,2,3,4,5,6,7],
				urls: [
					'https://cdn.uviewui.com/uview/album/1.jpg',
					'https://cdn.uviewui.com/uview/album/2.jpg',
					'https://cdn.uviewui.com/uview/album/3.jpg',
					'https://cdn.uviewui.com/uview/album/4.jpg',
					'https://cdn.uviewui.com/uview/album/5.jpg',
					'https://cdn.uviewui.com/uview/album/6.jpg',
					'https://cdn.uviewui.com/uview/album/7.jpg'
				]
			}
		},
		methods: {
			tabsChange() {
				console.log("change")
			},
			scrolltolower() {
				
			}
		}
	}
</script>

<style lang="scss" scoped>
.content {
	.top {
		background-color: #fff;
		position: fixed;
		// top: 0;
		z-index: 99;
		
		display: flex;
		flex-direction: column;
		align-items: center;
		
		.search {
			width: 90vw;
			background-color: #fff;
			padding: 20rpx;
		}
	}
	
	.item-list {
		padding-top: 200rpx;
	}
}
</style>
