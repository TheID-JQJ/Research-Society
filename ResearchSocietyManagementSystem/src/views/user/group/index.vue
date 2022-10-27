<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="序号" align="center" width="100" />
      <el-table-column prop="groupNumber" label="群号" align="center" />
      <el-table-column prop="avatar" label="群头像" align="center" />
      <el-table-column prop="name" label="群名称" align="center" />
      <el-table-column prop="ownerNumber" label="群主账号" align="center" />
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
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
        },
        {
          id: 2,
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
        },
        {
          id: 3,
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
        },
        {
          id: 4,
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
        },
        {
          id: 5,
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
        },
        {
          id: 6,
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
        },
        {
          id: 7,
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
        },
        {
          id: 8,
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
        },
        {
          id: 9,
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
        },
        {
          id: 10,
          groupNumber: '22334',
          avatar: '空',
          name: '好友之家',
          ownerNumber: '123456'
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
