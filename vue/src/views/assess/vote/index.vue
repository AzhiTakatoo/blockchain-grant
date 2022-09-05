<template>
  <div class="container">
    <el-alert type="success" center show-icon>
      <p>学号: {{ wyuUserId }}</p>
        <p>姓名: {{ wyuUserName }}</p>
          <p>以下数据为在校学生给获助学金学生的助后评分</p>
          <p>下面列表将持续更新全部申请助学金的名单</p>
    </el-alert>
    <div v-if="voteList.length==0" style="text-align: center;">
      <el-alert title="查询不到数据" type="warning" />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in voteList" :key="index" :span="6" :offset="1">
        <el-card class="vote-card">
          <div slot="header" class="clearfix">
            获奖学金学号:
            <span style="color: rgb(255, 0, 0);">{{ val.stipendId }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">评分人数: </el-tag>
            <span>{{ val.stuNum }} 人</span>
          </div>
          <div class="item">
            <el-tag type="danger">助后评分: </el-tag>
            <span>{{ val.averageVote }} 分</span>
          </div>
          <el-button type="text" @click="handleVote(val.stipendId)">查看详情</el-button>

          <el-dialog title="评分详情" :visible.sync="dialogTableVisible">
            <el-table :data="voteDataOnly">
              <el-table-column property="stipendId" label="获助学金学生" width="150"></el-table-column>
              <el-table-column property="voteId" label="评分人" width="150"></el-table-column>
              <el-table-column property="vote" label="具体评分"></el-table-column>
            </el-table>
          </el-dialog>

        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import { queryVote, queryVoteOnly } from '@/api/vote'

  export default {
    name: 'Vote',
    data() {
      var checkArea = (rule, value, callback) => {
        if (value <= 0) {
          callback(new Error('必须大于0'))
        } else {
          callback()
        }
      }
      return {
        loading: true,
        loadingDialog: false,
        voteList: [],
        dialogTableVisible: false,
        voteDataOnly: [],
        valItem: {}
      }
    },
    computed: {
      ...mapGetters([
        'wyuUserId',
        'roles',
        'wyuUserName',
      ])
    },
    created() {
      if (this.roles[0] === 'admin') {
        queryVote().then(response => {

          if (response !== null) {
            this.voteList = response
          }
          this.loading = false
        }).catch(_ => {
          this.loading = false
        })
      } else {
        queryVote().then(response => {

          if (response !== null) {
            this.voteList = response
          }
          this.loading = false
        }).catch(_ => {
          this.loading = false
        })
      }
    },
    methods: {
      handleVote(bevoteId) {
        this.dialogTableVisible = true
        queryVoteOnly({ stipendId: bevoteId }).then(response => {
          this.voteDataOnly = response

        }).catch(() => {
          this.loading = false
        })
      },
    }
  }
</script>

<style>
  .container {
    width: 100%;
    /* text-align: center; */
    min-height: 100%;
    overflow: hidden;
  }

  .tag {
    float: left;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
  }

  .clearfix:after {
    clear: both
  }

  .vote-card {
    width: 280px;
    height: 250px;
    margin: 18px;
  }
</style>