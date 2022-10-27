<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="序号" align="center" width="100" />
      <el-table-column prop="name" label="种类名称" align="center" />
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
          name: '饮料'
        },
        {
          id: 2,
          name: '饮料'
        },
        {
          id: 3,
          name: '饮料'
        },
        {
          id: 4,
          name: '饮料'
        },
        {
          id: 5,
          name: '饮料'
        },
        {
          id: 6,
          name: '饮料'
        },
        {
          id: 7,
          name: '饮料'
        },
        {
          id: 8,
          name: '饮料'
        },
        {
          id: 9,
          name: '饮料'
        },
        {
          id: 10,
          name: '饮料'
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
