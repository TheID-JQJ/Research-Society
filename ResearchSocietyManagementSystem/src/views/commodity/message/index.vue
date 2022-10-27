<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="序号" align="center" width="100" />
      <el-table-column prop="picture" label="图片" align="center" />
      <el-table-column prop="name" label="名称" align="center" />
      <el-table-column prop="price" label="价格/(元)" align="center" />
      <el-table-column prop="integralPrice" label="积分价格" align="center" />
      <el-table-column prop="monthlySales" label="月销量" align="center" />
      <el-table-column prop="inventory" label="库存数量" align="center" />
      <el-table-column prop="kind" label="种类" align="center" />
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
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
        },
        {
          id: 2,
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
        },
        {
          id: 3,
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
        },
        {
          id: 4,
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
        },
        {
          id: 5,
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
        },
        {
          id: 6,
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
        },
        {
          id: 7,
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
        },
        {
          id: 8,
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
        },
        {
          id: 9,
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
        },
        {
          id: 10,
          picture: '空',
          name: '蜜雪冰城',
          price: 10,
          integralPrice: 1000,
          monthlySales: 100,
          inventory: 100,
          kind: '饮料'
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
