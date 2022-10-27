<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="座位号" align="center" width="100" />
      <el-table-column prop="picture" label="图片" align="center" width="150" />
      <el-table-column prop="kind" label="种类" align="center" width="150" />
      <el-table-column prop="status" label="状态" align="center" width="150">
        <template slot-scope="{row}">
          <el-tag :type="row.status===1?'danger':'success'">
            {{ row.status===1?'使用中':'空闲' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="详情" align="center">
        <el-table-column prop="userNumber" label="使用者账号" align="center" />
        <el-table-column prop="startTime" label="开始使用时间" align="center" />
      </el-table-column>
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
          kind: '普通座位',
          status: 1,
          userNumber: 123456789,
          startTime: '2022.10.26'
        },
        {
          id: 2,
          kind: '普通座位',
          status: 1,
          userNumber: 123456789,
          startTime: '2022.10.26'
        },
        {
          id: 3,
          kind: '普通座位',
          status: 1,
          userNumber: 123456789,
          startTime: '2022.10.26'
        },
        {
          id: 4,
          kind: '普通座位',
          status: 1,
          userNumber: 123456789,
          startTime: '2022.10.26'
        },
        {
          id: 5,
          kind: '普通座位',
          status: 1,
          userNumber: 123456789,
          startTime: '2022.10.26'
        },
        {
          id: 6,
          kind: '普通座位',
          status: 1,
          userNumber: 123456789,
          startTime: '2022.10.26'
        },
        {
          id: 7,
          kind: '普通座位',
          status: 1,
          userNumber: 123456789,
          startTime: '2022.10.26'
        },
        {
          id: 8,
          kind: '普通座位',
          status: 1,
          userNumber: 123456789,
          startTime: '2022.10.26'
        },
        {
          id: 9,
          kind: '普通座位',
          status: 1,
          userNumber: 123456789,
          startTime: '2022.10.26'
        },
        {
          id: 10,
          kind: '普通座位',
          status: 0,
          userNumber: '空',
          startTime: '空'
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
