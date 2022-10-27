<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="序号" align="center" width="100" />
      <el-table-column prop="type" label="类型" align="center" width="200" />
      <el-table-column prop="userNumber" label="对象" align="center" width="200" />
      <el-table-column prop="time" label="时间" align="center" width="200" />
      <el-table-column prop="money" label="金额" align="center" width="100" />
      <el-table-column prop="details" label="详情" align="center" />
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
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
        },
        {
          id: 2,
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
        },
        {
          id: 3,
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
        },
        {
          id: 4,
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
        },
        {
          id: 5,
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
        },
        {
          id: 6,
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
        },
        {
          id: 7,
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
        },
        {
          id: 8,
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
        },
        {
          id: 9,
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
        },
        {
          id: 10,
          type: '购买商品',
          userNumber: '123456',
          time: '2022.10.26 12:00:00',
          money: 10,
          details: '购买了1号商品'
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
