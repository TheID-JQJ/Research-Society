<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="序号" align="center" width="100" />
      <el-table-column prop="title" label="标题" align="center" width="200" />
      <el-table-column prop="content" label="内容" align="center" />
      <el-table-column prop="time" label="时间" align="center" width="200" />
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
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
        },
        {
          id: 2,
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
        },
        {
          id: 3,
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
        },
        {
          id: 4,
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
        },
        {
          id: 5,
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
        },
        {
          id: 6,
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
        },
        {
          id: 7,
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
        },
        {
          id: 8,
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
        },
        {
          id: 9,
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
        },
        {
          id: 10,
          title: '公告标题',
          content: '公告内容',
          time: '2022.10.27 12:00:00'
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
