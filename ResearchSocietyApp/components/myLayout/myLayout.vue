<template>
	<view class="my-layout">
		<u-row 
			:gutter="myLayoutGutter==null?0:myLayoutGutter"
			v-for="i in myLayoutList.length%2==0?myLayoutList.length/2:(myLayoutList.length+1)/2"
			:key="i"
		>
			<u-col span="6" v-for="j in 2" v-if="myLayoutList[(i-1)*2+j-1]!=null">
				<slot :item="myLayoutList[(i-1)*2+j-1]"></slot>
			</u-col>
		</u-row>
	</view>
</template>

<script>
	export default {
		name:'myLayout',
		props: ['myLayoutList', 'myLayoutGutter'],
		data() {
			return {
			};
		},
	}
</script>
