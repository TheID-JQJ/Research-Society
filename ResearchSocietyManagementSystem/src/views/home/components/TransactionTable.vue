<template>
  <div style="background-color: #fff;">
    <el-table :data="list" style="width: 100%;padding-top: 15px;">
      <el-table-column label="内容" min-width="200">
        <template slot-scope="scope">
          {{ scope.row.content | orderNoFilter }}
        </template>
      </el-table-column>
      <el-table-column label="时间" width="195" align="center">
        <template slot-scope="scope">
          {{ scope.row.price | toThousandFilter }}
        </template>
      </el-table-column>
      <el-table-column label="状态" width="100" align="center">
        <template slot-scope="{row}">
          <el-tag :type="row.status | statusFilter">
            {{ row.status=='success'?'成功':'失败' }}
          </el-tag>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      style="padding: 30px; text-align: center;"
      :hide-on-single-page="true"
      background
      layout="prev, pager, next"
      :total="30"
      @current-change="currentChange"
    />
  </div>
</template>

<script>
import { transactionList } from '@/api/remote-search'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        success: 'success',
        fail: 'danger'
      }

      return statusMap[status]
    },
    orderNoFilter(str) {
      return str.substring(0, 30)
    },
    toThousandFilter() {
      return '2022.9.24 12:00'
    }
  },
  data() {
    return {
      list: []
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      transactionList().then(response => {
        for (let index = 0; index < 10; index++) {
          this.list.push({
            content: (index + 1) + '号用户进行了XXX操作',
            status: (index % 2) === 0 ? 'fail' : 'success'
          })
        }
      })
    },
    currentChange(page) {
      console.log('this is ' + page)
    }
  }
}
</script>
