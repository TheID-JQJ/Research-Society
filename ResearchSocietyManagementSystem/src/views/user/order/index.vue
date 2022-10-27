<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="序号" align="center" width="100" />
      <el-table-column prop="userNumber" label="账号" align="center" />
      <el-table-column prop="goodsId" label="商品号" align="center" />
      <el-table-column prop="orderNumber" label="订单号" align="center" />
      <el-table-column prop="quantity" label="数量" align="center" />
      <el-table-column prop="realPrice" label="实际价格/(元)" align="center" />
      <el-table-column prop="realIntegralPrice" label="实际积分价格" align="center" />
      <el-table-column prop="validity" label="有效期" align="center" width="150" />
      <el-table-column prop="createTime" label="创建时间" align="center" width="150" />
      <el-table-column prop="payTime" label="支付时间" align="center" width="150" />
      <el-table-column prop="consumeTime" label="消费时间" align="center" width="150" />
      <el-table-column prop="status" label="状态" align="center" />
      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button v-if="isEdit!==$index" type="primary" size="mini" @click="edit($index)">
            编辑
          </el-button>
          <el-button v-if="isEdit===$index" type="danger" size="mini" @click="handleDelete(row,$index)">
            删除
          </el-button>
          <el-button v-if="isEdit===$index" type="primary" size="mini" @click="edit(-1)">
            完成
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      style="padding: 30px; text-align: center;"
      :hide-on-single-page="true"
      background
      layout="prev, pager, next"
      :total="100"
    />
  </div>
</template>

<script>
export default {

  data() {
    return {
      tableData: [
        {
          id: 1,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        },
        {
          id: 2,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        },
        {
          id: 3,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        },
        {
          id: 4,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        },
        {
          id: 5,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        },
        {
          id: 6,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        },
        {
          id: 7,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        },
        {
          id: 8,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        },
        {
          id: 9,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        },
        {
          id: 10,
          userNumber: '123456',
          goodsId: '12',
          orderNumber: '20221027245612',
          quantity: 5,
          realPrice: 15,
          realIntegralPrice: 0,
          validity: '2022.10.27 12:00:00',
          createTime: '2022.9.27 12:00:00',
          payTime: '2022.9.27 12:05:00',
          consumeTime: '空',
          status: '未消费'
        }
      ],
      isEdit: -1
    }
  },
  computed: {

  },

  methods: {
    edit(index) {
      this.isEdit = index
    }
  }
}
</script>
