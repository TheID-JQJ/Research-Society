<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="sid" label="座位号" align="center" width="200" />
      <el-table-column prop="status" label="状态" align="center" width="200">
        <template slot-scope="{row}">
          <el-tag :type="row.status===1?'success':'danger'">
            {{ row.status===1?'使用中':'未使用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="详情" align="center">
        <el-table-column prop="user" label="使用者" align="center" />
        <el-table-column prop="start" label="开始使用时间" align="center" />
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
          sid: 1,
          status: 1,
          user: 'hkc',
          start: '2022.9.24 12:00'
        },
        {
          sid: 2,
          status: 0,
          user: '无',
          start: '无'
        },
        {
          sid: 3,
          status: 0,
          user: '无',
          start: '无'
        },
        {
          sid: 4,
          status: 0,
          user: '无',
          start: '无'
        },
        {
          sid: 5,
          status: 0,
          user: '无',
          start: '无'
        },
        {
          sid: 6,
          status: 0,
          user: '无',
          start: '无'
        },
        {
          sid: 7,
          status: 0,
          user: '无',
          start: '无'
        },
        {
          sid: 8,
          status: 0,
          user: '无',
          start: '无'
        },
        {
          sid: 9,
          status: 0,
          user: '无',
          start: '无'
        },
        {
          sid: 10,
          status: 0,
          user: '无',
          start: '无'
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
