<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="序号" align="center" width="100" />
      <el-table-column prop="sub" label="角色" align="center" />
      <el-table-column prop="obj" label="资源" align="center" />
      <el-table-column prop="act" label="方法" align="center" />
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
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
        },
        {
          id: 2,
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
        },
        {
          id: 3,
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
        },
        {
          id: 4,
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
        },
        {
          id: 5,
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
        },
        {
          id: 6,
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
        },
        {
          id: 7,
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
        },
        {
          id: 8,
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
        },
        {
          id: 9,
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
        },
        {
          id: 10,
          sub: 'user',
          obj: '/api/test/*',
          act: 'get'
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
