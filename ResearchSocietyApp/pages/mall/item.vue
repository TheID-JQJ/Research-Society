<template>
	<view>
		{{message}}
	</view>
</template>

<script>
	export default {
		data() {
			return {
				message: {}
			}
		},
		onLoad(message) {
			this.message = message
		},
		methods: {
			
		}
	}
</script>

<style>

</style>
