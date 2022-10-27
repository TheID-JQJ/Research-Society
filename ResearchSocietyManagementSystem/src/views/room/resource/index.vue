<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="序号" align="center" width="100" />
      <el-table-column prop="picture" label="图片" align="center" width="100" />
      <el-table-column prop="name" label="名称" align="center" width="100" />
      <el-table-column prop="price" label="租价/天" align="center" width="100" />
      <el-table-column prop="integralPrice" label="积分租价/天" align="center" width="100" />
      <el-table-column prop="rentAmount" label="月租量" align="center" width="100" />
      <el-table-column prop="kind" label="种类" align="center" width="100" />
      <el-table-column prop="status" label="状态" align="center" width="100">
        <template slot-scope="{row}">
          <el-tag :type="row.status===1?'success':'danger'">
            {{ row.status===1?'已借出':'未借出' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="详情" align="center">
        <el-table-column prop="user" label="租借者" align="center" />
        <el-table-column prop="startTime" label="租借时间" align="center" />
        <el-table-column prop="expirationTime" label="到期时间" align="center" />
      </el-table-column>
      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button v-if="isEdit!==$index" type="primary" size="mini" @click="edit($index)">
            编辑
          </el-button>
          <el-button v-if="isEdit!==$index && row.status===1" type="success" size="mini">
            归还
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
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 1,
          user: 'hkc',
          startTime: '2022.9.24',
          expirationTime: '2022.10.1'
        },
        {
          id: 2,
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 1,
          user: 'hkc',
          startTime: '2022.9.24',
          expirationTime: '2022.10.1'
        },
        {
          id: 3,
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 1,
          user: 'hkc',
          startTime: '2022.9.24',
          expirationTime: '2022.10.1'
        },
        {
          id: 4,
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 1,
          user: 'hkc',
          startTime: '2022.9.24',
          expirationTime: '2022.10.1'
        },
        {
          id: 5,
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 1,
          user: 'hkc',
          startTime: '2022.9.24',
          expirationTime: '2022.10.1'
        },
        {
          id: 6,
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 1,
          user: 'hkc',
          startTime: '2022.9.24',
          expirationTime: '2022.10.1'
        },
        {
          id: 7,
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 1,
          user: 'hkc',
          startTime: '2022.9.24',
          expirationTime: '2022.10.1'
        },
        {
          id: 8,
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 1,
          user: 'hkc',
          startTime: '2022.9.24',
          expirationTime: '2022.10.1'
        },
        {
          id: 9,
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 1,
          user: 'hkc',
          startTime: '2022.9.24',
          expirationTime: '2022.10.1'
        },
        {
          id: 10,
          name: '安徒生童话',
          price: 0.5,
          integralPrice: 50,
          rentAmount: 5,
          kind: '书',
          status: 0,
          user: '空',
          startTime: '空',
          expirationTime: '空'
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
