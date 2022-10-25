<template>
  <div style="background-color: #fff;">
    <el-table :data="list" style="width: 100%;padding-top: 15px;">
      <el-table-column label="内容" min-width="200">
        <template slot-scope="scope">
          {{ scope.row.order_no | orderNoFilter }}
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
            {{ row.status=='pending'?'失败':'成功' }}
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
        pending: 'danger'
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
      list: null
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      transactionList().then(response => {
        this.list = response.data.items.slice(0, 10)
      })
    },
    currentChange(page) {
      console.log('this is ' + page)
    }
  }
}
</script>
