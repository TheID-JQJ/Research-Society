<template>
	<view>
		<u-list
			class="item-list"
			@scrolltolower="scrolltolower"
		>
			<u-list-item
				v-for="(item, index) in indexList"
				:key="index"
			>
				<u-cell
					:title="`公告-${index + 1}`"
				>
				</u-cell>
			</u-list-item>
		</u-list>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				indexList: [1,2,3,4,5,6,7],
			}
		},
		methods: {
			scrolltolower() {
				
			}
		}
	}
</script>

<style>

</style>
