<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="序号" align="center" width="100" />
      <el-table-column prop="userNumber" label="账号" align="center" />
      <el-table-column prop="password" label="密码" align="center" />
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
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
        },
        {
          id: 2,
          userNumber: '123456',
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
        },
        {
          id: 3,
          userNumber: '123456',
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
        },
        {
          id: 4,
          userNumber: '123456',
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
        },
        {
          id: 5,
          userNumber: '123456',
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
        },
        {
          id: 6,
          userNumber: '123456',
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
        },
        {
          id: 7,
          userNumber: '123456',
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
        },
        {
          id: 8,
          userNumber: '123456',
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
        },
        {
          id: 9,
          userNumber: '123456',
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
        },
        {
          id: 10,
          userNumber: '123456',
          password: 'E10ADC3949BA59ABBE56E057F20F883E'
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
