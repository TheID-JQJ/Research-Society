<template>
	<view class="my-grid">
		<u-grid
			:border="false"
			@click="gridClick"
			:col="myGridCol"
		>
			<u-grid-item
				v-for="(item,index) in myGridList"
				:key="index"
			>
				<slot class="slot" :item="item"></slot>
			</u-grid-item>
		</u-grid>
	</view>
</template>

<script>
	export default {
		name:"myGrid",
		props: ['myGridList', 'myGridCol'],
		data() {
			return {
				
			};
		},
		methods: {
			gridClick(index) {
				// 页面的跳转
				this.$u.route({
					url: this.myGridList[index].url
				})
			}
		}
	}
</script>

